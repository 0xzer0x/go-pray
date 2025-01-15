package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"

	"github.com/0xzer0x/go-pray/internal/format"
)

func MissingKeyError(key string) error {
	return errors.New("missing key: " + key)
}

func ValidateKey(key string) error {
	if !viper.IsSet(key) {
		return fmt.Errorf("%s is not set, use config file or command-line flags to set it.", key)
	}
	return nil
}

func ValidateCalculationParams() error {
	for _, key := range []string{"method", "location.lat", "location.long", "timezone"} {
		if !viper.IsSet(key) {
			return fmt.Errorf(
				"%s is not set, use config file or command-line flags to set it.",
				key,
			)
		}
	}
	return nil
}

func Formatter() format.Formatter {
	value := viper.GetString("format")
	switch value {
	case "json":
		return &format.JsonFormatter{}
	default:
		return &format.TextFormatter{}
	}
}
