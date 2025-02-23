package pocket

import (
	"github.com/cryog0at/wtf/cfg"
	"github.com/olebedev/config"
)

const (
	defaultFocusable = true
	defaultTitle     = "Pocket"
)

type Settings struct {
	*cfg.Common

	consumerKey string
	requestKey  *string
	accessToken *string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common:      cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
		consumerKey: ymlConfig.UString("consumerKey"),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.consumerKey).Load()

	return &settings
}
