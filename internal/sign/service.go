package sign

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const operatorID = "8837f423-6d63-4766-9789-44a75d1d5f22"
const signatureHeader = "X-Signature"
const operatorIdHeader = "X-Operator-ID"

type Service struct {
	privateKeyPem string
	publicKeyPem  string
	privateKey    *rsa.PrivateKey
	publicKey     *rsa.PublicKey
}

func New(privateKeyPem string, publicKeyPem string) *Service {
	return &Service{
		privateKeyPem: privateKeyPem,
		publicKeyPem:  publicKeyPem,
	}
}

func (s *Service) getPrivateKey() *rsa.PrivateKey {
	if s.privateKey == nil {
		block, _ := pem.Decode([]byte(s.privateKeyPem))
		key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			panic(err)
		}
		s.privateKey = key
	}

	return s.privateKey
}

func (s *Service) getPublicKey() *rsa.PublicKey {
	if s.publicKey != nil {
		return s.publicKey
	}

	block, _ := pem.Decode([]byte(s.publicKeyPem))
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	s.publicKey = key.(*rsa.PublicKey)

	return s.publicKey
}

func (s *Service) Sign(body []byte) (string, error) {
	hash := sha256.New()
	hash.Write(body)
	digest := hash.Sum(nil)

	v15, err := rsa.SignPKCS1v15(nil, s.getPrivateKey(), crypto.SHA256, digest)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(v15), nil
}

func (s *Service) verify(signature string, body []byte) (bool, error) {
	key := s.getPublicKey()

	hash := sha256.New()
	hash.Write(body)
	digest := hash.Sum(nil)

	decoded, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		panic(err)
	}

	err = rsa.VerifyPKCS1v15(key, crypto.SHA256, digest, decoded)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *Service) AttachOperatorSignature(r *http.Request, body []byte) {
	signature, err := s.Sign(body)
	if err != nil {
		log.WithError(err).Fatal("failed to sign launch game request")
	}
	r.Header.Add(signatureHeader, signature)
	r.Header.Add(operatorIdHeader, operatorID)
}

func (s *Service) VerifyProviderSignature(c *fiber.Ctx) bool {
	signature, err := extractHeader(c, signatureHeader)
	if err != nil {
		log.WithError(err).Error("no signature supplied")
		return false
	}

	verify, err := s.verify(signature, c.Body())
	if err != nil {
		log.WithError(err).Error("signature verification failed")
		return false
	}

	return verify
}

func extractHeader(c *fiber.Ctx, key string) (string, error) {
	val, ok := c.GetReqHeaders()[key]

	if !ok || len(val) == 0 {
		return "", fmt.Errorf("header [%s] not found or empty", key)
	}

	return val[0], nil
}
