package i18n

import "golang.org/x/text/language"

type Config struct {
	PathLangFiles string
	DefaultLang   string
	CurrentLang   string
	langs         []language.Tag
}
