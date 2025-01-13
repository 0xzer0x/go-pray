package calendar

import (
	"slices"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/0xzer0x/go-pray/internal/config"
	"github.com/0xzer0x/go-pray/internal/util"
)

func validateDate(args []string) {
	for _, date := range args {
		if slices.Contains(util.MapKeys(specialDates), date) {
			continue
		}
		if _, err := time.Parse("2006-01-02", date); err != nil {
			util.ErrExit(
				"invalid date '%s': must be in the format YYYY-MM-DD or one of %s",
				date,
				strings.Join(util.MapKeys(specialDates), ", "),
			)
		}
	}
}

func validateCalendarArgs(cmd *cobra.Command, args []string) {
	validateDate(args)
	err := config.ValidateCalculationParams()
	if err != nil {
		util.ErrExit("%v", err)
	}
}
