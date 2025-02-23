package {{(Lower .Name)}}

import (
	"github.com/olebedev/config"
	"github.com/cryog0at/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "{{(.Name)}}"
)

// Settings defines the configuration properties for this module
type Settings struct {
	common *cfg.Common

    // Define your settings attributes here
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
        common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

        // Configure your settings attributes here. See http://github.com/olebedev/config for type details
	}

	return &settings
}