package i18n

import (
	"io/ioutil"
	"strings"

	"golang.org/x/text/language"
)

// Load apply the configuration to load language messages
func Load(c Config) error {
	files, err := ioutil.ReadDir(c.PathLangFiles)
	if err != nil {
		return err
	}
	defaultExists := false
	currentExists := false
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".json") {
			lang, err := language.Parse(strings.TrimRight(f.Name(), ".json"))
			if err != nil {
				return err
			}

			c.langs = append(c.langs, lang)

			if lang.String() == c.DefaultLang {
				defaultExists = true
			}
			if lang.String() == c.CurrentLang {
				currentExists = true
			}
		}
	}

	if len(c.langs) == 0 {
		return &NotFoundLanguagesError{}
	}

	if !defaultExists {
		c.DefaultLang = c.langs[0].String()
	}
	if !currentExists {
		c.CurrentLang = c.DefaultLang
	}

	config = c

	return nil
}
