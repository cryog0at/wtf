package twitter

import (
	"os"

	"github.com/cryog0at/wtf/cfg"
	"github.com/olebedev/config"
)

const (
	defaultFocusable = true
	defaultTitle     = "Twitter"
)

type Settings struct {
	*cfg.Common

	bearerToken    string
	consumerKey    string
	consumerSecret string
	count          int
	screenNames    []interface{}
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		bearerToken:    ymlConfig.UString("bearerToken", os.Getenv("WTF_TWITTER_BEARER_TOKEN")),
		consumerKey:    ymlConfig.UString("consumerKey", os.Getenv("WTF_TWITTER_CONSUMER_KEY")),
		consumerSecret: ymlConfig.UString("consumerSecret", os.Getenv("WTF_TWITTER_CONSUMER_SECRET")),
		count:          ymlConfig.UInt("count", 5),
		screenNames:    ymlConfig.UList("screenName"),
	}

	settings.SetDocumentationPath("twitter/tweets")

	return &settings
}
