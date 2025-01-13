package next

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/0xzer0x/go-pray/internal/common"
	"github.com/0xzer0x/go-pray/internal/config"
	"github.com/0xzer0x/go-pray/internal/util"
)

func validateNextArgs(cmd *cobra.Command, args []string) {
	err := config.ValidateCalculationParams()
	if err != nil {
		util.ErrExit("%v", err)
	}

	if len(args) > 0 {
		nextPrayerArg = strings.ToLower(args[0])
		validArg := false
		for name := range common.Prayers {
			if nextPrayerArg == name {
				validArg = true
				break
			}
		}
		if !validArg {
			util.ErrExit(
				"invalid prayer '%s', allowed values are: %s",
				args[0],
				strings.Join(util.MapKeys(common.Prayers), ", "),
			)
		}
	}
}
