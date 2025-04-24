package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/0xzer0x/go-pray/internal/common"
	"github.com/0xzer0x/go-pray/internal/util"
)

func registerGlobalFlags() {
	var err error
	pflags := rootCmd.PersistentFlags()

	// NOTE: register string flags
	for _, flagInfo := range [][3]string{
		{"language", "en", "output language"},
		{"config", "", "config file"},
		{"format", "", "output format"},
		{"adhan", "", "path to adhan mp3"},
		{"timezone", time.Now().Location().String(), "prayer times timezone"},
		{"calculation.method", "", "calculation method"},
		{"notification.icon", "clock-applet-symbolic", "notification icon name"},
		{"notification.title", "Prayer", "notification title template"},
		{"notification.body", "Time for {{ .CalendarName }} prayer 🕌", "notification body template"},
	} {
		pflags.String(flagInfo[0], flagInfo[1], flagInfo[2])
		err = viper.BindPFlag(flagInfo[0], pflags.Lookup(flagInfo[0]))
		if err != nil {
			util.ErrExit("failed to bind %s flag", flagInfo[0])
		}
	}

	// NOTE: register float flags
	for name, usage := range map[string]string{
		"location.lat":  "calculation latitude",
		"location.long": "calculation longitude",
	} {
		pflags.Float64(name, 0, usage)
		err = viper.BindPFlag(name, pflags.Lookup(name))
		if err != nil {
			util.ErrExit("failed to bind %s flag", name)
		}
	}

	// NOTE: shell completion for flags
	for name, values := range map[string][]string{
		"language":           {"en", "ar"},
		"format":             {"json", "table", "short"},
		"calculation.method": util.MapKeys(common.Methods),
	} {
		err = rootCmd.RegisterFlagCompletionFunc(
			name,
			func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
				return values, cobra.ShellCompDirectiveDefault
			},
		)
		if err != nil {
			util.ErrExit("failed to set flag completion: %v", err)
		}
	}

	// NOTE: shell completion for flags that accept files
	for name, exts := range map[string][]string{
		"config": {"yaml", "yml"},
		"adhan":  {"mp3"},
	} {
		err = rootCmd.RegisterFlagCompletionFunc(
			name,
			func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
				return exts, cobra.ShellCompDirectiveFilterFileExt
			},
		)
		if err != nil {
			util.ErrExit("failed to set valid flag file extensions: %v", err)
		}
	}
}
