package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/0xzer0x/go-pray/cmd/calendar"
	"github.com/0xzer0x/go-pray/cmd/next"
	"github.com/0xzer0x/go-pray/cmd/notify"
	"github.com/0xzer0x/go-pray/internal/config"
	"github.com/0xzer0x/go-pray/internal/util"
)

var rootCmd = &cobra.Command{
	Use:   "go-pray",
	Short: "Prayer times CLI application",
	Long:  `Prayer times CLI application`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func bindStringFlag(name string, usage string) {
	rootCmd.PersistentFlags().String(name, "", usage)
	err := viper.BindPFlag(name, rootCmd.PersistentFlags().Lookup(name))
	if err != nil {
		util.ErrExit("error binding " + name + " flag")
	}
}

func registerFlags() {
	bindStringFlag("config", "config file")
	bindStringFlag("timezone", "calculation timezone")
	bindStringFlag("method", "calculation method")

	for name, usage := range map[string]string{"location.lat": "calculation latitude", "location.long": "calculation longitude"} {
		rootCmd.PersistentFlags().Float64(name, 0, usage)
		err := viper.BindPFlag(name, rootCmd.PersistentFlags().Lookup(name))
		if err != nil {
			util.ErrExit("error binding " + name + " flag")
		}
	}
}

func addSubCommands() {
	rootCmd.AddCommand(notify.NotifyCmd)
	rootCmd.AddCommand(calendar.CalendarCmd)
	rootCmd.AddCommand(next.NextCommand)
}

func init() {
	cobra.OnInitialize(config.Initialize)
	addSubCommands()
	registerFlags()
}
