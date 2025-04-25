package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

func MissingKeyError(key string) error {
	return errors.New("missing key: " + key)
}

func ValidateKey(key string) error {
	if !viper.IsSet(key) {
		return fmt.Errorf("%s is not set, use config file or command-line flags to set it", key)
	}
	return nil
}

func ValidateCalculationParams() error {
	for _, key := range []string{"calculation.method", "location.lat", "location.long"} {
		if !viper.IsSet(key) {
			return fmt.Errorf("%s is not set, use config file or command-line flags to set it", key)
		}
	}
	return nil
}
