package i18n

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/NaujOyamat/infinity/utils"
	"github.com/allegro/bigcache/v3"
	"golang.org/x/text/language"
)

// FillObject maps each of the object's properties
// with its corresponding translation key specified in the i18n tag.
// Example:
// 	type Person struct {
// 		Name  string `i18n:"Person.Name"`
// 		State string `i18n:"StateLabel"`
// 	}
func FillObject(spec interface{}) error {
	s := reflect.ValueOf(spec)

	if s.Kind() != reflect.Ptr {
		return &InvalidSpecificationError{}
	}
	s = s.Elem()
	if s.Kind() != reflect.Struct {
		return &InvalidSpecificationError{}
	}
	typeOfSpec := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		ftype := typeOfSpec.Field(i)
		if !f.CanSet() {
			continue
		}

		if f.Kind() != reflect.String {
			return &InvalidSpecificationError{}
		}

		key := ftype.Tag.Get("i18n")
		if key == "" {
			key = ftype.Name
		}

		f.SetString(Translate(key))
	}

	return nil
}

// Translate find a entry in the language used
func Translate(key string) string {
	if config.PathLangFiles == NotInitialized {
		panic(&LanguagesNotInitializedError{})
	}

	key = strings.Trim(key, " ")

	entry, err := config.cacheMap.Get(fmt.Sprintf("%s:%s", config.CurrentLang, key))
	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			entry, err = config.cacheMap.Get(fmt.Sprintf("%s:%s", config.DefaultLang, key))
			if err != nil {
				logger.Errorf("entry not fount (%s)", key)
				return ""
			}
			return string(entry)
		}
		logger.Errorf("error %s", err.Error())
		return ""
	}
	return string(entry)
}

// SetDefaultLang set a language loaded in memory as default
func SetDefaultLang(lang string) error {
	for _, l := range config.langs {
		if strings.EqualFold(l, lang) {
			config.DefaultLang = lang
			return nil
		}
	}

	return &LanguageNotFoundError{}
}

// DefaultLang return the default language
func DefaultLang() string {
	return config.DefaultLang
}

// SetCurrentLang set a language loaded in memory to use
func SetCurrentLang(lang string) error {
	for _, l := range config.langs {
		if strings.EqualFold(l, lang) {
			config.CurrentLang = lang
			return nil
		}
	}

	return &LanguageNotFoundError{}
}

// CurrentLang return the current language
func CurrentLang() string {
	return config.CurrentLang
}

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

			err = loadFile(lang.String(), fmt.Sprintf("%s/%s", c.PathLangFiles, f.Name()))
			if err != nil {
				return err
			}

			c.langs = append(c.langs, lang.String())

			if lang.String() == c.DefaultLang {
				defaultExists = true
			}
			if lang.String() == c.CurrentLang {
				currentExists = true
			}
		}
	}

	if len(c.langs) == 0 {
		return &LanguageFileNotFoundError{}
	}

	if !defaultExists {
		c.DefaultLang = c.langs[0]
	}
	if !currentExists {
		c.CurrentLang = c.DefaultLang
	}

	setConfig(c)

	return nil
}

// setConfig set config properties
func setConfig(c Config) {
	c.cacheMap = config.cacheMap
	config = c
}

// loadFile load content file to load keys in memory
func loadFile(tag string, fpath string) error {
	jsonFile, err := os.Open(fpath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	bytesValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	err = json.Unmarshal(bytesValue, &result)
	if err != nil {
		return err
	}

	err = loadKeys(tag, "", result)
	if err != nil {
		return err
	}

	return nil
}

// loadKeys load keys from map to memory
func loadKeys(tag string, prefix string, data map[string]interface{}) error {
	if prefix != "" {
		prefix = fmt.Sprintf("%s.", prefix)
	}

	for k, v := range data {
		key := fmt.Sprintf("%s%s", prefix, k)
		vType := reflect.TypeOf(v)
		if vType.String() == "map[string]interface {}" {
			err := loadKeys(tag, key, v.(map[string]interface{}))
			if err != nil {
				return err
			}
		} else {
			config.cacheMap.Append(fmt.Sprintf("%s:%s", tag, key), []byte(utils.ConvertToString(v)))
		}
	}

	return nil
}
