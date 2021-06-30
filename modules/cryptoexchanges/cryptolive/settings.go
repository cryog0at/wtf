package cryptolive

import (
	"github.com/cryog0at/wtf/cfg"
	"github.com/cryog0at/wtf/modules/cryptoexchanges/cryptolive/price"
	"github.com/cryog0at/wtf/modules/cryptoexchanges/cryptolive/toplist"
	"github.com/olebedev/config"
)

const (
	defaultFocusable = false
	defaultTitle     = "CryptolLive"
)

type colors struct {
	from struct {
		name        string
		displayName string
	}
	to struct {
		name  string
		price string
	}
	top struct {
		from struct {
			name        string
			displayName string
		}
		to struct {
			name  string
			field string
			value string
		}
	}
}

type Settings struct {
	*cfg.Common

	colors

	currencies map[string]interface{}
	top        map[string]interface{}

	priceSettings   *price.Settings
	toplistSettings *toplist.Settings
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	currencies, _ := ymlConfig.Map("currencies")
	top, _ := ymlConfig.Map("top")

	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		currencies: currencies,
		top:        top,

		priceSettings:   price.NewSettingsFromYAML(name, ymlConfig, globalConfig),
		toplistSettings: toplist.NewSettingsFromYAML(name, ymlConfig, globalConfig),
	}

	settings.colors.from.name = ymlConfig.UString("colors.from.name")
	settings.colors.from.displayName = ymlConfig.UString("colors.from.displayName")

	settings.colors.to.name = ymlConfig.UString("colors.to.name")
	settings.colors.to.price = ymlConfig.UString("colors.to.price")

	settings.colors.top.from.name = ymlConfig.UString("colors.top.from.name")
	settings.colors.top.from.displayName = ymlConfig.UString("colors.top.from.displayName")

	settings.colors.top.to.name = ymlConfig.UString("colors.top.to.name")
	settings.colors.top.to.field = ymlConfig.UString("colors.top.to.field")
	settings.colors.top.to.value = ymlConfig.UString("colors.top.to.value")

	settings.SetDocumentationPath("cryptocurrencies/cryptolive")

	return &settings
}
