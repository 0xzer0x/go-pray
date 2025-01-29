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
		if val, defined := os.LookupEnv("XDG_CONFIG_HOME"); defined && len(val) > 0 {
			viper.AddConfigPath("$XDG_CONFIG_HOME/go-pray")
		}
		if val, defined := os.LookupEnv("HOME"); defined && len(val) > 0 {
			viper.AddConfigPath("$HOME/.config/go-pray")
			viper.AddConfigPath("$HOME/.go-pray")
		}
	}

	viper.SetEnvPrefix("GO_PRAY")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		util.ErrExit("failed to load configuration: %v", err)
	}
}
