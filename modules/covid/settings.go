package covid

import (
	"github.com/cryog0at/wtf/cfg"
	"github.com/olebedev/config"
)

const (
	defaultFocusable = false
	defaultTitle     = "Covid tracker"
)

// Settings is the struct for this module's settings
type Settings struct {
	*cfg.Common

	countries []interface{} `help:"Countries (codes) from which to retrieve stats."`
}

// NewSettingsFromYAML returns the settings from the config yaml file
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		// List of countries to retrieve stats from
		countries: ymlConfig.UList("countries"),
	}

	return &settings
}
