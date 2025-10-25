package i18n

import (
	"embed"
	"strings"
	"sync"
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
)

//go:embed locale.*.toml
var localeFS embed.FS

type Localizer struct {
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
}

var (
	lock                           = &sync.Mutex{}
	singleton           *Localizer = nil
	arabicTimeLocalizer            = strings.NewReplacer(
		"0", "٠",
		"1", "١",
		"2", "٢",
		"3", "٣",
		"4", "٤",
		"5", "٥",
		"6", "٦",
		"7", "٧",
		"8", "٨",
		"9", "٩",
		"h", "س",
		"m", "د",
		"s", "ث",
		"am", "ص",
		"pm", "م",
		"AM", "ص",
		"PM", "م",
	)
)

func GetInstance() (*Localizer, error) {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleton == nil {
			lang := viper.GetString("language")
			loc := Localizer{}

			loc.bundle = i18n.NewBundle(language.English)
			loc.bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
			if _, err := loc.bundle.LoadMessageFileFS(localeFS, "locale.en.toml"); err != nil {
				return nil, err
			}
			if _, err := loc.bundle.LoadMessageFileFS(localeFS, "locale.ar.toml"); err != nil {
				return nil, err
			}
			loc.localizer = i18n.NewLocalizer(loc.bundle, lang)

			singleton = &loc
		}
	}

	return singleton, nil
}

func (l *Localizer) Localize(messageID string, templateData *map[string]any) (string, error) {
	config := &i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: messageID,
		},
	}
	if templateData != nil {
		config.TemplateData = *templateData
	}

	return l.localizer.Localize(config)
}

func (l *Localizer) LocalizeTime(tm time.Time, format string) string {
	formattedTime := tm.Format(format)
	localizedTime := formattedTime

	lang := viper.GetString("language")
	if lang == "ar" {
		localizedTime = arabicTimeLocalizer.Replace(formattedTime)
	}

	return localizedTime
}

func (l *Localizer) LocalizeDuration(d time.Duration) string {
	lang := viper.GetString("language")
	if lang == "ar" {
		return arabicTimeLocalizer.Replace(d.String())
	}
	return d.String()
}

func (l *Localizer) LocalizeTimeString(s string) string {
	lang := viper.GetString("language")
	if lang == "ar" {
		return arabicTimeLocalizer.Replace(s)
	}
	return s
}
