package pihole

import (
	"github.com/cryog0at/wtf/cfg"
	"github.com/olebedev/config"
)

const (
	defaultFocusable = true
	defaultTitle     = "Pi-hole"
)

type Settings struct {
	*cfg.Common

	wrapText       bool
	apiUrl         string
	token          string
	showTopItems   int
	showTopClients int
	maxClientWidth int
	maxDomainWidth int
	showSummary    bool
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiUrl:         ymlConfig.UString("apiUrl"),
		token:          ymlConfig.UString("token"),
		showSummary:    ymlConfig.UBool("showSummary", true),
		showTopItems:   ymlConfig.UInt("showTopItems", 5),
		showTopClients: ymlConfig.UInt("showTopClients", 5),
		maxClientWidth: ymlConfig.UInt("maxClientWidth", 20),
		maxDomainWidth: ymlConfig.UInt("maxDomainWidth", 20),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.token).
		Service(settings.apiUrl).Load()

	return &settings
}
