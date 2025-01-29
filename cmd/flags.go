package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/0xzer0x/go-pray/internal/util"
)

func registerGlobalFlags() {
	pflags := rootCmd.PersistentFlags()

	// NOTE: register string flags
	for name, usage := range map[string]string{
		"config":             "config file",
		"format":             "output format",
		"adhan":              "path to adhan mp3",
		"timezone":           "prayer times timezone",
		"calculation.method": "calculation method",
	} {
		pflags.String(name, "", usage)
		err := viper.BindPFlag(name, pflags.Lookup(name))
		if err != nil {
			util.ErrExit("failed to bind %s flag", name)
		}
	}

	// NOTE: register float flags
	for name, usage := range map[string]string{
		"location.lat":  "calculation latitude",
		"location.long": "calculation longitude",
	} {
		pflags.Float64(name, 0, usage)
		err := viper.BindPFlag(name, pflags.Lookup(name))
		if err != nil {
			util.ErrExit("failed to bind %s flag", name)
		}
	}

	err := rootCmd.RegisterFlagCompletionFunc(
		"format",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return []string{"json", "table", "short"}, cobra.ShellCompDirectiveDefault
		},
	)
	if err != nil {
		util.ErrExit("failed to set flag completion: %v", err)
	}

	err = rootCmd.RegisterFlagCompletionFunc(
		"config",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return []string{"yml", "yaml"}, cobra.ShellCompDirectiveFilterFileExt
		},
	)
	if err != nil {
		util.ErrExit("failed to set valid flag file extensions: %v", err)
	}
}
