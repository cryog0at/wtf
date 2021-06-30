package grafana

import (
	"log"
	"os"
	"strings"

	"github.com/cryog0at/wtf/cfg"
	"github.com/olebedev/config"
)

const (
	defaultFocusable = true
	defaultTitle     = "Grafana"
)

type Settings struct {
	*cfg.Common

	apiKey  string `help:"Your Grafana API token."`
	baseURI string `help:"Base url of your grafana instance"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:  ymlConfig.UString("apiKey", os.Getenv("WTF_GRAFANA_API_KEY")),
		baseURI: ymlConfig.UString("baseUri", ""),
	}

	if settings.baseURI == "" {
		log.Fatal("baseUri for grafana is empty, but is required")
	} else {
		settings.baseURI = strings.TrimSuffix(settings.baseURI, "/")
	}

	return &settings
}
