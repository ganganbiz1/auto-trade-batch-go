package env

import "os"

type Env interface {
	GetAPIKey() string
	GetAPISecret() string
}

type EnvImpl struct {
	apiKey    string
	apiSecret string
}

func NewEnv() *EnvImpl {
	return &EnvImpl{
		apiKey:    os.Getenv("BITBANK_API_KEY"),
		apiSecret: os.Getenv("BITBANK_API_SECRET"),
	}
}

func (e *EnvImpl) GetAPIKey() string {
	return e.apiKey
}

func (e *EnvImpl) GetAPISecret() string {
	return e.apiSecret
}
