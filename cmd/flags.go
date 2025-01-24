package cmd

import (
	"github.com/spf13/viper"

	"github.com/0xzer0x/go-pray/internal/util"
)

func registerGlobalFlags() {
	// NOTE: register string flags
	for name, usage := range map[string]string{
		"config":             "config file",
		"format":             "output format",
		"adhan":              "path to adhan mp3",
		"timezone":           "prayer times timezone",
		"calculation.method": "calculation method",
	} {
		rootCmd.PersistentFlags().String(name, "", usage)
		err := viper.BindPFlag(name, rootCmd.PersistentFlags().Lookup(name))
		if err != nil {
			util.ErrExit("failed to bind %s flag", name)
		}
	}

	// NOTE: register float flags
	for name, usage := range map[string]string{
		"location.lat":  "calculation latitude",
		"location.long": "calculation longitude",
	} {
		rootCmd.PersistentFlags().Float64(name, 0, usage)
		err := viper.BindPFlag(name, rootCmd.PersistentFlags().Lookup(name))
		if err != nil {
			util.ErrExit("failed to bind %s flag", name)
		}
	}
}
