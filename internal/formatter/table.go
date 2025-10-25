package formatter

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/mnadev/adhango/pkg/calc"

	"github.com/0xzer0x/go-pray/internal/common"
	"github.com/0xzer0x/go-pray/internal/i18n"
	"github.com/0xzer0x/go-pray/internal/version"
)

type TableFormatter struct {
	localizer *i18n.Localizer
}

var (
	headerStyle lipgloss.Style  = lipgloss.NewStyle().Bold(true).Align(lipgloss.Center)
	cellStyle   lipgloss.Style  = lipgloss.NewStyle().Padding(0, 2)
	tableBorder lipgloss.Border = lipgloss.RoundedBorder()
)

func NewTableFormatter() (*TableFormatter, error) {
	var err error
	var localizer *i18n.Localizer
	if localizer, err = i18n.GetInstance(); err != nil {
		return nil, fmt.Errorf("failed to initialize localizer: %w", err)
	}

	tf := &TableFormatter{
		localizer,
	}
	return tf, nil
}

func (f *TableFormatter) cellStyle(row, col int) lipgloss.Style {
	switch row {
	case table.HeaderRow:
		return headerStyle
	default:
		return cellStyle
	}
}

func (f *TableFormatter) newPrayersTable() (*table.Table, error) {
	localizedHeaderMessageIDs := [...]string{"date", "prayer", "time", "remaining"}
	localizedHeaders := make([]string, 0, 4)
	for _, messageID := range localizedHeaderMessageIDs {
		var err error
		var header string
		if header, err = f.localizer.Localize(messageID, nil); err != nil {
			return nil, fmt.Errorf("failed to localize header: %w", err)
		}
		localizedHeaders = append(localizedHeaders, strings.ToUpper(header))
	}

	prayersTable := table.New().
		Border(tableBorder).
		StyleFunc(f.cellStyle).
		Headers(localizedHeaders...)

	return prayersTable, nil
}

func (f *TableFormatter) appendPrayerRow(tbl *table.Table, pname string, ptime time.Time) {
	tbl.Row(
		f.localizer.LocalizeTime(ptime, time.DateOnly),
		pname,
		f.localizer.LocalizeTime(ptime, "03:04 PM"),
		f.localizer.LocalizeDuration(time.Until(ptime).Truncate(time.Second)),
	)
}

func (f *TableFormatter) Calendar(calendar calc.PrayerTimes) (string, error) {
	prayersTable, err := f.newPrayersTable()
	if err != nil {
		return "", err
	}

	for _, name := range []string{"fajr", "sunrise", "dhuhr", "asr", "maghrib", "isha"} {
		prayer := common.Prayers[name]
		pname := common.CalendarName(calendar, prayer)
		ptime := calendar.TimeForPrayer(prayer)
		f.appendPrayerRow(prayersTable, pname, ptime)
	}

	return prayersTable.String() + "\n", nil
}

func (f *TableFormatter) Prayer(calendar calc.PrayerTimes, prayer calc.Prayer) (string, error) {
	prayersTable, err := f.newPrayersTable()
	if err != nil {
		return "", err
	}

	pname := common.CalendarName(calendar, prayer)
	ptime := calendar.TimeForPrayer(prayer)
	f.appendPrayerRow(prayersTable, pname, ptime)

	return prayersTable.String() + "\n", nil
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
