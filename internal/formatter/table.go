package formatter

import (
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/mnadev/adhango/pkg/calc"

	"github.com/0xzer0x/go-pray/internal/common"
	"github.com/0xzer0x/go-pray/internal/version"
)

type TableFormatter struct{}

var (
	headerStyle lipgloss.Style  = lipgloss.NewStyle().Bold(true).Align(lipgloss.Center)
	cellStyle   lipgloss.Style  = lipgloss.NewStyle().Padding(0, 2)
	tableBorder lipgloss.Border = lipgloss.RoundedBorder()
)

func (f *TableFormatter) cellStyle(row, col int) lipgloss.Style {
	switch row {
	case table.HeaderRow:
		return headerStyle
	default:
		return cellStyle
	}
}

func (f *TableFormatter) Calendar(calendar calc.PrayerTimes) (string, error) {
	prayersTable := table.New().
		Border(tableBorder).
		StyleFunc(f.cellStyle).
		Headers("DATE", "PRAYER", "TIME", "REMAINING")

	for _, name := range []string{"fajr", "sunrise", "dhuhr", "asr", "maghrib", "isha"} {
		prayer := common.Prayers[name]
		pt := calendar.TimeForPrayer(prayer)
		prayersTable.Row(
			pt.Format(time.DateOnly),
			common.CalendarName(calendar, prayer),
			pt.Format("03:04PM"),
			time.Until(pt).Truncate(time.Second).String(),
		)
	}

	return prayersTable.String() + "\n", nil
}

func (f *TableFormatter) Prayer(calendar calc.PrayerTimes, prayer calc.Prayer) (string, error) {
	prayerTable := table.New().
		Border(tableBorder).
		StyleFunc(f.cellStyle).
		Headers("DATE", "PRAYER", "TIME", "REMAINING")

	pt := calendar.TimeForPrayer(prayer)
	prayerTable.Row(
		pt.Format(time.DateOnly),
		common.CalendarName(calendar, prayer),
		pt.Format("03:04PM"),
		time.Until(pt).Truncate(time.Second).String(),
	)

	return prayerTable.String() + "\n", nil
}

func (f *TableFormatter) VersionInfo(versionInfo version.VersionInfo) (string, error) {
	versionTable := table.New().
		Border(tableBorder).
		StyleFunc(f.cellStyle).
		Headers("VERSION", "GO VERSION", "BUILD COMMIT", "BUILD TIME", "OS", "ARCH")

	versionTable.Row(
		versionInfo.Version,
		versionInfo.Runtime,
		versionInfo.BuildCommit,
		versionInfo.BuildTime.Format(time.DateTime),
		versionInfo.Os,
		versionInfo.Arch,
	)

	return versionTable.String() + "\n", nil
}
