package uptimerobot

import (
	"os"

	"github.com/cryog0at/wtf/cfg"
	"github.com/olebedev/config"
)

const (
	defaultFocusable = true
	defaultTitle     = "Uptime Robot"
)

type Settings struct {
	*cfg.Common

	apiKey        string `help:"An UptimeRobot API key."`
	uptimePeriods string `help:"The periods over which to display uptime (in days, dash-separated)." optional:"true"`
	offlineFirst  bool   `help:"Display offline monitors at the top." optional:"true"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:        ymlConfig.UString("apiKey", os.Getenv("WTF_UPTIMEROBOT_APIKEY")),
		uptimePeriods: ymlConfig.UString("uptimePeriods", "30"),
		offlineFirst:  ymlConfig.UBool("offlineFirst", false),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.apiKey).
		Service("https://api.uptimerobot.com").Load()

	return &settings
}
