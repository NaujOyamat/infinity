package i18n_test

import (
	"testing"

	"github.com/NaujOyamat/infinity/i18n"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	t.Run("ioutil.ReadDir error", func(t *testing.T) {
		expectedErrorMsg := "open -: no such file or directory"
		err := i18n.Load(i18n.Config{
			PathLangFiles: "-",
			DefaultLang:   ".",
			CurrentLang:   ".",
		})

		if assert.Error(t, err) {
			assert.EqualError(t, err, expectedErrorMsg)
		}
	})
}
