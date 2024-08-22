package calendar

import (
	"github.com/spf13/cobra"
)

var CalendarCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Prayers calendar utilities",
}

func init() {
	CalendarCmd.AddCommand(todayCmd)
	CalendarCmd.AddCommand(dateCmd)
}
