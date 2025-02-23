package circleci

import (
	"os"

	"github.com/cryog0at/wtf/cfg"
	"github.com/olebedev/config"
)

const (
	defaultFocusable = false
	defaultTitle     = "CircleCI"
)

type Settings struct {
	*cfg.Common

	apiKey         string `help:"Your CircleCI API token."`
	numberOfBuilds int    `help:"The number of build, 10 by default"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:         ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_CIRCLE_API_KEY"))),
		numberOfBuilds: ymlConfig.UInt("numberOfBuilds", 10),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.apiKey).Load()

	return &settings
}
