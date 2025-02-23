package bamboohr

import (
	"os"

	"github.com/cryog0at/wtf/cfg"
	"github.com/olebedev/config"
)

const (
	defaultFocusable = false
	defaultTitle     = "BambooHR"
)

type Settings struct {
	*cfg.Common

	apiKey    string `help:"Your BambooHR API token."`
	subdomain string `help:"Your BambooHR API subdomain name."`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:    ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_BAMBOO_HR_TOKEN"))),
		subdomain: ymlConfig.UString("subdomain", os.Getenv("WTF_BAMBOO_HR_SUBDOMAIN")),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.apiKey).Load()

	return &settings
}
