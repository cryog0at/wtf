package arpansagovau

import (
	"github.com/cryog0at/wtf/cfg"
	"github.com/olebedev/config"
)

const (
	defaultFocusable = false
	defaultTitle     = "ARPANSA UV Data"
)

type Settings struct {
	*cfg.Common

	city string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
		city:   ymlConfig.UString("locationid"),
	}

	settings.SetDocumentationPath("weather_services/arpansagovau")

	return &settings
}
