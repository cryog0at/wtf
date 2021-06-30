package ipapi

import (
	"github.com/cryog0at/wtf/cfg"
	"github.com/olebedev/config"
)

const (
	defaultFocusable = false
	defaultTitle     = "IP API"
)

type colors struct {
	name  string
	value string
}

type Settings struct {
	colors
	*cfg.Common
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
	}

	settings.colors.name = ymlConfig.UString("colors.name", "red")
	settings.colors.value = ymlConfig.UString("colors.value", "white")

	settings.SetDocumentationPath("ipaddress/ipapi")

	return &settings
}
