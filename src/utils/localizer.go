package utils

import (
	"encoding/json"
	"log"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type LocalizerUtil struct {
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
}

// NewLocalizerUtil creates a new instance of LocalizerUtil
func NewLocalizerUtil() *LocalizerUtil {
	util := &LocalizerUtil{}
	util.initI18n()
	return util
}

func (l *LocalizerUtil) initI18n() {
	l.bundle = i18n.NewBundle(language.English)
	l.bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	_, err := l.bundle.LoadMessageFile("./locales/pt-BR.json")
	if err != nil {
		log.Printf("Error loading pt-BR.json: %v", err)
	}

	_, err = l.bundle.LoadMessageFile("./locales/en.json")
	if err != nil {
		log.Printf("Error loading en.json: %v", err)
	}

	l.localizer = i18n.NewLocalizer(l.bundle, language.English.String(), language.BrazilianPortuguese.String())
}

func (l *LocalizerUtil) SetLocalizer(lang string, accept string) {
	l.localizer = i18n.NewLocalizer(l.bundle, lang, accept)
}

func (l *LocalizerUtil) Localize(config i18n.LocalizeConfig) string {
	return l.localizer.MustLocalize(&config)
}

func (l *LocalizerUtil) GetLocalizer() *i18n.Localizer {
	return l.localizer
}
