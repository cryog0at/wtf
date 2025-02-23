package twitch

import (
	"os"

	"github.com/cryog0at/wtf/cfg"
	"github.com/cryog0at/wtf/utils"
	"github.com/olebedev/config"
)

const (
	defaultFocusable = true
)

type Settings struct {
	*cfg.Common

	numberOfResults int      `help:"Number of results to show. Default is 10." optional:"true"`
	clientId        string   `help:"Client Id (default is env var TWITCH_CLIENT_ID)"`
	languages       []string `help:"Stream languages" optional:"true"`
	gameIds         []string `help:"Twitch Game IDs" optional:"true"`
	streamType      string   `help:"Type of stream 'live' (default), 'all', 'vodcast'" optional:"true"`
	userIds         []string `help:"Twitch user ids" optional:"true"`
	userLogins      []string `help:"Twitch user names" optional:"true"`
}

func defaultLanguage() []interface{} {
	var defaults []interface{}
	defaults = append(defaults, "en")
	return defaults
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	twitch := ymlConfig.UString("twitch")
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, twitch, defaultFocusable, ymlConfig, globalConfig),

		numberOfResults: ymlConfig.UInt("numberOfResults", 10),
		clientId:        ymlConfig.UString("clientId", os.Getenv("TWITCH_CLIENT_ID")),
		languages:       utils.ToStrs(ymlConfig.UList("languages", defaultLanguage())),
		streamType:      ymlConfig.UString("streamType", "live"),
		gameIds:         utils.ToStrs(ymlConfig.UList("gameIds", make([]interface{}, 0))),
		userIds:         utils.ToStrs(ymlConfig.UList("userIds", make([]interface{}, 0))),
		userLogins:      utils.ToStrs(ymlConfig.UList("userLogins", make([]interface{}, 0))),
	}
	return &settings
}
