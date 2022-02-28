package i18n

import (
	"github.com/NaujOyamat/infinity/logs"
	"github.com/NaujOyamat/infinity/times"
	"github.com/allegro/bigcache/v3"
)

const (
	NotInitialized = "..::-@#@-::.."
)

var (
	config Config
	logger logs.Logger
)

func init() {
	config = Config{
		PathLangFiles: NotInitialized,
	}
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(200 * times.Years))
	if err != nil {
		panic(err)
	}
	config.cacheMap = cache
	logger = logs.NewLogrusLogger()
}
