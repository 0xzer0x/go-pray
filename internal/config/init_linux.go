package config

import (
	"os"

	"github.com/spf13/viper"

	"github.com/0xzer0x/go-pray/internal/util"
)

func Initialize() {
	if cfgFile := viper.GetString("config"); cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config.yml")
		viper.SetConfigType("yaml")
		if _, defined := os.LookupEnv("XDG_CONFIG_HOME"); defined {
			viper.AddConfigPath("$XDG_CONFIG_HOME/go-pray")
		}
		if _, defined := os.LookupEnv("HOME"); defined {
			viper.AddConfigPath("$HOME/.config/go-pray")
			viper.AddConfigPath("$HOME/.go-pray")
		}
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		util.ErrExit("failed to load configuration: %v", err)
	}
}
