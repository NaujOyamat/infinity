package i18n

import "github.com/allegro/bigcache/v3"

type Config struct {
	PathLangFiles string
	DefaultLang   string
	CurrentLang   string
	langs         []string
	cacheMap      *bigcache.BigCache
}
