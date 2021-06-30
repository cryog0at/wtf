// Package exchangerates
package exchangerates

import (
	"github.com/cryog0at/wtf/cfg"
	"github.com/olebedev/config"
)

const (
	defaultFocusable = false
	defaultTitle     = "Exchange rates"
)

// Settings defines the configuration properties for this module
type Settings struct {
	*cfg.Common

	precision int `help:"How many decimal places to display." optional:"true"`

	rates map[string][]string `help:"Defines what currency rates we want to know about"`
	order []string
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		precision: ymlConfig.UInt("precision", 7),

		rates: map[string][]string{},
		order: []string{},
	}

	raw := ymlConfig.UMap("rates", map[string]interface{}{})
	for key, value := range raw {
		settings.order = append(settings.order, key)
		settings.rates[key] = []string{}
		switch value := value.(type) {
		case string:
			settings.rates[key] = []string{value}
		case []interface{}:
			for _, currency := range value {
				str, ok := currency.(string)
				if ok {
					settings.rates[key] = append(settings.rates[key], str)
				}
			}
		}
	}

	return &settings
}
