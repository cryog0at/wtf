package victorops

import (
	"os"

	"github.com/cryog0at/wtf/cfg"
	"github.com/olebedev/config"
)

const (
	defaultFocusable = true
	defaultTitle     = "VictorOps"
)

type Settings struct {
	*cfg.Common

	apiID  string
	apiKey string
	team   string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiID:  ymlConfig.UString("apiID", os.Getenv("WTF_VICTOROPS_API_ID")),
		apiKey: ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_VICTOROPS_API_KEY"))),
		team:   ymlConfig.UString("team"),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.apiKey).Load()

	return &settings
}
