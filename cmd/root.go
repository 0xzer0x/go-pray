package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/0xzer0x/go-pray/cmd/calendar"
	"github.com/0xzer0x/go-pray/cmd/daemon"
	"github.com/0xzer0x/go-pray/cmd/next"
	"github.com/0xzer0x/go-pray/internal/config"
	"github.com/0xzer0x/go-pray/internal/util"
)

var rootCmd = &cobra.Command{
	Use:   "go-pray",
	Short: "Prayer times CLI application",
	Long:  `Prayer times CLI application`,
}

func registerGlobalFlags() {
	// NOTE: register string flags
	for name, usage := range map[string]string{
		"config":   "config file",
		"timezone": "calculation timezone",
		"method":   "calculation method",
		"format":   "output format",
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

func init() {
	cobra.OnInitialize(config.Initialize)
	rootCmd.AddCommand(daemon.DaemonCmd)
	rootCmd.AddCommand(calendar.CalendarCmd)
	rootCmd.AddCommand(next.NextCommand)
	registerGlobalFlags()
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
