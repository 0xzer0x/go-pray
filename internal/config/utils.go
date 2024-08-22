package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"

	"github.com/0xzer0x/go-pray/internal/util"
)

func MissingKeyError(key string) error {
	return errors.New("missing key: " + key)
}

func ValidateSet(key string) {
	if !viper.IsSet(key) {
		util.ErrExit("%s is not set, use config file or command-line flags to set it.", key)
	}
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
