package config

import (
	"github.com/spf13/viper"

	"github.com/0xzer0x/go-pray/internal/util"
)

func Initialize() {
	if cfgFile := viper.GetString("config"); cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config.yml")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("$XDG_CONFIG_HOME/go-pray")
		viper.AddConfigPath("$HOME/.config/go-pray")
		viper.AddConfigPath("$HOME/.go-pray")
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		util.ErrExit("failed to load configuration: %v", err)
	}
}
