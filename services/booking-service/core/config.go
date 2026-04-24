package core

import (
	"booking-service/common"
	"errors"
	"os"
)

func LoadConfig() (*common.Config, error) {
	config := &common.Config{}

	key := os.Getenv("JWT_SECRET")
	if key == "" {
		return nil, errors.New("no JWT_SECRET")
	}
	config.JwtSecretKey = []byte(key)

	ConnectionStr := os.Getenv("CONNECTION_STR")
	if ConnectionStr == "" {
		return nil, errors.New("no CONNECTION_STR")
	}
	config.ConnectionStr = ConnectionStr

	TursoToken := os.Getenv("TURSO_TOKEN")
	if TursoToken == "" {
		return nil, errors.New("no TURSO_TOKEN")
	}
	config.TursoToken = TursoToken

	return config, nil
}
