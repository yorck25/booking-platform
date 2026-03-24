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
		return nil, errors.New("no secret key")
	}
	config.JwtSecretKey = []byte(key)

	ConnectionStr := os.Getenv("CONNECTION_STR")
	if ConnectionStr == "" {
		return nil, errors.New("no CONNECTION_STR")
	}
	config.ConnectionStr = ConnectionStr

	return config, nil
}
