package i18n

import (
	"embed"
	"sync"

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
	lock                 = &sync.Mutex{}
	singleton *Localizer = nil
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
