package runner

import (
	"path"
)

type LanguageSettings struct {
	Path string `json:"path"`
}

type Theme string

const (
	LightTheme Theme = "light"
	DarkTheme        = "dark"
)

type Settings struct {
	Id            string                                  `json:"id"`
	LastProject   *string                                 `json:"lastProject"`
	Languages     map[SupportedLanguages]LanguageSettings `json:"languages"`
	File          string                                  `json:"file"`
	StdoutMaxSize int                                     `json:"stdoutMaxSize"`
	Theme         Theme                                   `json:"theme"`
	CacheSettings *CacheSettings                          `json:"cacheSettings"`
	CaCerts       []struct {
		File string `json:"file"`
	} `json:"caCerts"`
}

type CacheSettings struct {
	CachePresent bool
	Enabled      bool
}

var SettingsFileDefaultLocation = path.Join(CONFIG_FS_BASE, ".settings")

var DefaultSettings = &Settings{
	Id:            newId(),
	File:          SettingsFileDefaultLocation,
	StdoutMaxSize: 5000,
	Theme:         "light",
	CacheSettings: &CacheSettings{
		Enabled:      false,
		CachePresent: false,
	},
}

func LoadSettings(file string) (*Settings, error) {
	var settings Settings
	err := readJSONFileInto(file, &settings)
	return &settings, err
}
