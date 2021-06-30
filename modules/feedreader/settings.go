package feedreader

import (
	"github.com/cryog0at/wtf/cfg"
	"github.com/cryog0at/wtf/utils"
	"github.com/olebedev/config"
)

const (
	defaultFocusable = true
	defaultTitle     = "Feed Reader"
)

// Settings defines the configuration properties for this module
type Settings struct {
	*cfg.Common

	feeds     []string `help:"An array of RSS and Atom feed URLs"`
	feedLimit int      `help:"The maximum number of stories to display for each feed"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := &Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		feeds:     utils.ToStrs(ymlConfig.UList("feeds")),
		feedLimit: ymlConfig.UInt("feedLimit", -1),
	}

	return settings
}
