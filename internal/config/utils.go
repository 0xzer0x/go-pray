package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"

	"github.com/0xzer0x/go-pray/internal/pfmt"
	"github.com/0xzer0x/go-pray/internal/util"
)

func MissingKeyError(key string) error {
	return errors.New("missing key: " + key)
}

func ValidateKey(key string) {
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

func FormatStrategy() pfmt.FormatStrategy {
	value := viper.GetString("format")
	switch value {
	case "json":
		return &pfmt.JsonFormatStrategy{}
	default:
		return &pfmt.TextFormatStrategy{}
	}
}
