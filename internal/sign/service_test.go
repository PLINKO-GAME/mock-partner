package sign

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_signAndVerify(t *testing.T) {
	privateKeyPem := "-----BEGIN RSA PRIVATE KEY-----\nMIIEogIBAAKCAQEAhcnimjhkkGue02/EdYkV47tR0L36Wcunn/1WW4noRLD8M04Y\n9G36UT2uzt20u5KzKwOk1NdW0/0WYqQCMJ4DfKVdaLiXpZFsLsirT8r8cC0mpuNA\nAUHdl/bGDJa9kFrbwT6RyTkIMAxLnG16yeiuam/799ve9V04kDvC2F/YFCX1Sumr\nJCVgbyw9K/xpYWp7Fppt/pB3KXlvZf0FyIDkI9j767tMfuTlOVLPY/n5x+ZaVmF7\nvxhKJknoboHY+9v5KaU7XeTrThZfSl0LjboBBpFl1gWpKNglGQbOvjOdijV51xgi\n8IeIZ90u2W3LtTbL+k+c16qMVOaFA3sJVQwZCwIDAQABAoIBAAJ/JPppyP6k837Q\nnCLxXvYz/a/ei7h3Q3aJ3L2ykiIOB3bRo0eUcdJoS0XS/1dswmkwFThfmGA2Xd+T\nXfMYT8pYr6iPoUzWrOUmm4POru1M+mas4PnlB8SZN1Lu0TTLbURq7X+Kz+tNn2+Y\n32y7Kd4UnugeM0fy6GZQpy8wgrDFiWKphSB7pz/czovAR7BQLbCIhpGgQc4LF22D\n1EAx5i/wSU5b8KgGeiOsdpT9LzOaFosaeq1tDGlwpN4JKuj6sSgvhST96f/XihD5\nnwBNIDiM5IFMr4MplJThr2oFTQH7TyK6/bswvvURReyXiKpVaaZ/QTyY2hQ3yrjV\nFZGPX2ECgYEA9StqKNYmqg8W+X6MUE0EyAlHUHey30dgz1nn3uHmynhQMWEWc2X0\ni7r17h1Yi37DN0veIJRW3e1/DDypezIebQkaoyGmuy2Ogh3WCrH/nL/WVMgcAIXu\nLRS111M4LtJFOc5vbCw8U/34OFH0T6k5/8gjtrcx12ExZzezoSfXiBECgYEAi7Lh\nTmRwOHcYvMH84RETb4BoBq3V97CCO91ZpinDQa8z9v3tsEl5NhY0N9cceMvnB/+f\nqQ3tNHCrZcktXLXwLLPDCtYU9BpY/sRZy39kCsEo8t8xtGPjiI+rH0AhheTgUTCc\n5TK3LwKmlkUGGDpww9eOIFljbIMeoH2+HOU2C1sCgYA8SfTNHfxcDWHk8I2ooYfv\nePikfQrrhS31T3KJiJusZnGx8uIGdqfwRIV9jJHdm8p9qpZxBIloAaMgazpyJRz+\nSyLVwsyxcr58mMGt15+3+CTIrHzWVBkB1PnyfXBvcx263VzhCO+859NGZkDh5gdx\nMtI1eE81W50+eKAfnSCPQQKBgE7oWW9IOEMMsoJcKJSQaqP+qcOsCUIBB279FphO\n2qWNaxLGV63NspOkcxZfgQuSUQspjmuVHDkUsxupSOAnPGRjnXXPesJu53nwOrBB\nYqbYeGLHQ3IbQfhu/j+Gn+jbYQE7LkQgI2yAWMxkbI7e47cbWIJZO1mdrn0EyY/U\nwHQlAoGAMqsSsgQM4IKzkK3kW+WmCEtpfPu+A6gaDAPVWdY1UU/r5IfGSV47r+bE\nj2BQaPQFiLzqA4SRDuvZHDRyI3MsqT9t1Mx6fBJ+x2kb3Muctc/SR2JpsIgbB1xv\nKk1vBRzeosnDbvbnZVkFih/lWLSEuLyE2iqy+kwScwIhiBcN/a8=\n-----END RSA PRIVATE KEY-----\n"
	publicKey := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAhcnimjhkkGue02/EdYkV\n47tR0L36Wcunn/1WW4noRLD8M04Y9G36UT2uzt20u5KzKwOk1NdW0/0WYqQCMJ4D\nfKVdaLiXpZFsLsirT8r8cC0mpuNAAUHdl/bGDJa9kFrbwT6RyTkIMAxLnG16yeiu\nam/799ve9V04kDvC2F/YFCX1SumrJCVgbyw9K/xpYWp7Fppt/pB3KXlvZf0FyIDk\nI9j767tMfuTlOVLPY/n5x+ZaVmF7vxhKJknoboHY+9v5KaU7XeTrThZfSl0LjboB\nBpFl1gWpKNglGQbOvjOdijV51xgi8IeIZ90u2W3LtTbL+k+c16qMVOaFA3sJVQwZ\nCwIDAQAB\n-----END PUBLIC KEY-----\n"
	expected := "GeSW3VGWUlhxf47j5HCCrUNXkc1N1Y6OYDeixygE5JBYXmZubagSXLDlVSAalqu6EPda003AQAIg/lRUAa04gExy8+vmAToMZAeUmE+yj2vBLZd5LfpWHDW468PEvvxTY3KXfEd0mxBScWJV6prMN9H4NlmQL1u4NTeru6lPklxtfcDrqORCU18sJZ9vb4HsrYJpcqGfRfms/FhWyxiCW4Pdg2rJRERyyDN4f/NTIepWmXb0eMmEufuR/Mv7xujSSqLZyLlgUkmMeDMe6CBlRZHJ0x5Yhsm8xAuQWYoQ6mhO+JM4RaDNi3c8nfkbLxaDFhydfdAB7CjuA++nNQW7Fw=="

	service := New(privateKeyPem, publicKey)
	body := "{\"user\":\"3nYTOSjdlF6UTz9Ir\",\"country\":\"XX\",\"currency\":\"EUR\",\"operator_id\":1,\"token\":\"cd6bd8560f3bb8f84325152101adeb45\",\"platform\":\"GPL_DESKTOP\",\"game_code\":\"clt_dragonrising\",\"lang\":\"en\",\"lobby_url\":\"https://examplecasino.io\",\"ip\":\"::ffff:10.0.0.39\"}"

	signed, err := service.Sign([]byte(body))
	assert.Nil(t, err)
	assert.Equal(t, expected, signed)

	verify, err := service.verify(signed, []byte(body))
	assert.Nil(t, err)
	assert.True(t, verify)
}
