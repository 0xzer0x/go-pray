package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/0xzer0x/go-pray/cmd/calendar"
	"github.com/0xzer0x/go-pray/cmd/daemon"
	"github.com/0xzer0x/go-pray/cmd/next"
	"github.com/0xzer0x/go-pray/cmd/version"
	"github.com/0xzer0x/go-pray/internal/config"
)

var rootCmd = &cobra.Command{
	Use:   "go-pray",
	Short: "Prayer times CLI",
	Long:  `Prayer times CLI`,
}

func init() {
	cobra.OnInitialize(config.Initialize)
	rootCmd.AddCommand(daemon.DaemonCmd)
	rootCmd.AddCommand(calendar.CalendarCmd)
	rootCmd.AddCommand(next.NextCommand)
	rootCmd.AddCommand(version.VersionCmd)
	registerGlobalFlags()
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
