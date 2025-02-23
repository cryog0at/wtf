package travisci

import (
	"os"

	"github.com/cryog0at/wtf/cfg"
	"github.com/olebedev/config"
)

const (
	defaultFocusable = true
	defaultTitle     = "TravisCI"
)

type Settings struct {
	*cfg.Common

	apiKey  string
	baseURL string `help:"Your TravisCI Enterprise API URL." optional:"true"`
	compact bool
	limit   string
	pro     bool
	sort_by string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:  ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_TRAVIS_API_TOKEN"))),
		baseURL: ymlConfig.UString("baseURL", ymlConfig.UString("baseURL", os.Getenv("WTF_TRAVIS_BASE_URL"))),
		pro:     ymlConfig.UBool("pro", false),
		compact: ymlConfig.UBool("compact", false),
		limit:   ymlConfig.UString("limit", "10"),
		sort_by: ymlConfig.UString("sort_by", "id:desc"),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.apiKey).
		Service(settings.baseURL).Load()

	return &settings
}
