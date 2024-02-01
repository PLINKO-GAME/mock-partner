package main

import (
	"github.com/caarlos0/env/v6"
)

type config struct {
	CoreURL    string `env:"CONF_CORE_URL" envDefault:"http://plinko-core:8080"`
	HTTPPort   string `env:"CONF_HTTP_PORT" envDefault:":8080"`
	PrivateKey string
	PublicKey  string
}

func newConfig() (*config, error) {
	conf := new(config)
	err := env.Parse(conf)
	if err != nil {
		return nil, err
	}

	// plinko-core public key
	conf.PublicKey = "-----BEGIN PUBLIC KEY-----\nMIIBITANBgkqhkiG9w0BAQEFAAOCAQ4AMIIBCQKCAQBWCDdaPGjJb64kKonq0ipx\n+MTj4lkEMCRvh60QHO7qO2JwZUGiu5rTrPgVrFjA8BzKq20+7++50LFfSKNY0VZh\nzph08uZl4jjFSM6SM4uTXaI2mCDgT0T1zjdygPj0xuMPFE8UdR0/e12DtpWOqhfA\nKWT4QGO0xv/uO9Mwr/srnvY9BRe0WELbdqIYPCJIPhCTw4iD+k2wHNWff04xEOJE\nTMrWyQILeWQQAvMw/9ROZaqzHTmYR1+V6Gubk2Ywg7FJ05dKNRcxy/vT1/F3no7F\nD5jbxUqQT21GlR/7bFDHcydo541X6c0UL13V4EHpZkkHnvMy4qBWIgfrKYkeg0hb\nAgMBAAE=\n-----END PUBLIC KEY-----"
	// mock-partner private key
	conf.PrivateKey = "-----BEGIN RSA PRIVATE KEY-----\nMIIEogIBAAKCAQEAhcnimjhkkGue02/EdYkV47tR0L36Wcunn/1WW4noRLD8M04Y\n9G36UT2uzt20u5KzKwOk1NdW0/0WYqQCMJ4DfKVdaLiXpZFsLsirT8r8cC0mpuNA\nAUHdl/bGDJa9kFrbwT6RyTkIMAxLnG16yeiuam/799ve9V04kDvC2F/YFCX1Sumr\nJCVgbyw9K/xpYWp7Fppt/pB3KXlvZf0FyIDkI9j767tMfuTlOVLPY/n5x+ZaVmF7\nvxhKJknoboHY+9v5KaU7XeTrThZfSl0LjboBBpFl1gWpKNglGQbOvjOdijV51xgi\n8IeIZ90u2W3LtTbL+k+c16qMVOaFA3sJVQwZCwIDAQABAoIBAAJ/JPppyP6k837Q\nnCLxXvYz/a/ei7h3Q3aJ3L2ykiIOB3bRo0eUcdJoS0XS/1dswmkwFThfmGA2Xd+T\nXfMYT8pYr6iPoUzWrOUmm4POru1M+mas4PnlB8SZN1Lu0TTLbURq7X+Kz+tNn2+Y\n32y7Kd4UnugeM0fy6GZQpy8wgrDFiWKphSB7pz/czovAR7BQLbCIhpGgQc4LF22D\n1EAx5i/wSU5b8KgGeiOsdpT9LzOaFosaeq1tDGlwpN4JKuj6sSgvhST96f/XihD5\nnwBNIDiM5IFMr4MplJThr2oFTQH7TyK6/bswvvURReyXiKpVaaZ/QTyY2hQ3yrjV\nFZGPX2ECgYEA9StqKNYmqg8W+X6MUE0EyAlHUHey30dgz1nn3uHmynhQMWEWc2X0\ni7r17h1Yi37DN0veIJRW3e1/DDypezIebQkaoyGmuy2Ogh3WCrH/nL/WVMgcAIXu\nLRS111M4LtJFOc5vbCw8U/34OFH0T6k5/8gjtrcx12ExZzezoSfXiBECgYEAi7Lh\nTmRwOHcYvMH84RETb4BoBq3V97CCO91ZpinDQa8z9v3tsEl5NhY0N9cceMvnB/+f\nqQ3tNHCrZcktXLXwLLPDCtYU9BpY/sRZy39kCsEo8t8xtGPjiI+rH0AhheTgUTCc\n5TK3LwKmlkUGGDpww9eOIFljbIMeoH2+HOU2C1sCgYA8SfTNHfxcDWHk8I2ooYfv\nePikfQrrhS31T3KJiJusZnGx8uIGdqfwRIV9jJHdm8p9qpZxBIloAaMgazpyJRz+\nSyLVwsyxcr58mMGt15+3+CTIrHzWVBkB1PnyfXBvcx263VzhCO+859NGZkDh5gdx\nMtI1eE81W50+eKAfnSCPQQKBgE7oWW9IOEMMsoJcKJSQaqP+qcOsCUIBB279FphO\n2qWNaxLGV63NspOkcxZfgQuSUQspjmuVHDkUsxupSOAnPGRjnXXPesJu53nwOrBB\nYqbYeGLHQ3IbQfhu/j+Gn+jbYQE7LkQgI2yAWMxkbI7e47cbWIJZO1mdrn0EyY/U\nwHQlAoGAMqsSsgQM4IKzkK3kW+WmCEtpfPu+A6gaDAPVWdY1UU/r5IfGSV47r+bE\nj2BQaPQFiLzqA4SRDuvZHDRyI3MsqT9t1Mx6fBJ+x2kb3Muctc/SR2JpsIgbB1xv\nKk1vBRzeosnDbvbnZVkFih/lWLSEuLyE2iqy+kwScwIhiBcN/a8=\n-----END RSA PRIVATE KEY-----\n"

	return conf, nil
}
