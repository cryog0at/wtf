package app

import (
	"github.com/cryog0at/wtf/modules/asana"
	"github.com/cryog0at/wtf/modules/azuredevops"
	"github.com/cryog0at/wtf/modules/bamboohr"
	"github.com/cryog0at/wtf/modules/bargraph"
	"github.com/cryog0at/wtf/modules/buildkite"
	cdsfavorites "github.com/cryog0at/wtf/modules/cds/favorites"
	cdsqueue "github.com/cryog0at/wtf/modules/cds/queue"
	cdsstatus "github.com/cryog0at/wtf/modules/cds/status"
	"github.com/cryog0at/wtf/modules/circleci"
	"github.com/cryog0at/wtf/modules/clocks"
	"github.com/cryog0at/wtf/modules/cmdrunner"
	"github.com/cryog0at/wtf/modules/covid"
	"github.com/cryog0at/wtf/modules/cryptoexchanges/bittrex"
	"github.com/cryog0at/wtf/modules/cryptoexchanges/blockfolio"
	"github.com/cryog0at/wtf/modules/cryptoexchanges/cryptolive"
	"github.com/cryog0at/wtf/modules/datadog"
	"github.com/cryog0at/wtf/modules/devto"
	"github.com/cryog0at/wtf/modules/digitalclock"
	"github.com/cryog0at/wtf/modules/digitalocean"
	"github.com/cryog0at/wtf/modules/docker"
	"github.com/cryog0at/wtf/modules/exchangerates"
	"github.com/cryog0at/wtf/modules/feedreader"
	"github.com/cryog0at/wtf/modules/football"
	"github.com/cryog0at/wtf/modules/gcal"
	"github.com/cryog0at/wtf/modules/gerrit"
	"github.com/cryog0at/wtf/modules/git"
	"github.com/cryog0at/wtf/modules/github"
	"github.com/cryog0at/wtf/modules/gitlab"
	"github.com/cryog0at/wtf/modules/gitlabtodo"
	"github.com/cryog0at/wtf/modules/gitter"
	"github.com/cryog0at/wtf/modules/googleanalytics"
	"github.com/cryog0at/wtf/modules/grafana"
	"github.com/cryog0at/wtf/modules/gspreadsheets"
	"github.com/cryog0at/wtf/modules/hackernews"
	"github.com/cryog0at/wtf/modules/healthchecks"
	"github.com/cryog0at/wtf/modules/hibp"
	"github.com/cryog0at/wtf/modules/ipaddresses/ipapi"
	"github.com/cryog0at/wtf/modules/ipaddresses/ipinfo"
	"github.com/cryog0at/wtf/modules/jenkins"
	"github.com/cryog0at/wtf/modules/jira"
	"github.com/cryog0at/wtf/modules/krisinformation"
	"github.com/cryog0at/wtf/modules/kubernetes"
	"github.com/cryog0at/wtf/modules/logger"
	"github.com/cryog0at/wtf/modules/mercurial"
	"github.com/cryog0at/wtf/modules/nbascore"
	"github.com/cryog0at/wtf/modules/newrelic"
	"github.com/cryog0at/wtf/modules/opsgenie"
	"github.com/cryog0at/wtf/modules/pagerduty"
	"github.com/cryog0at/wtf/modules/pihole"
	"github.com/cryog0at/wtf/modules/pocket"
	"github.com/cryog0at/wtf/modules/power"
	"github.com/cryog0at/wtf/modules/resourceusage"
	"github.com/cryog0at/wtf/modules/rollbar"
	"github.com/cryog0at/wtf/modules/security"
	"github.com/cryog0at/wtf/modules/spacex"
	"github.com/cryog0at/wtf/modules/spotify"
	"github.com/cryog0at/wtf/modules/spotifyweb"
	"github.com/cryog0at/wtf/modules/status"
	"github.com/cryog0at/wtf/modules/stocks/finnhub"
	"github.com/cryog0at/wtf/modules/stocks/yfinance"
	"github.com/cryog0at/wtf/modules/subreddit"
	"github.com/cryog0at/wtf/modules/textfile"
	"github.com/cryog0at/wtf/modules/todo"
	"github.com/cryog0at/wtf/modules/todo_plus"
	"github.com/cryog0at/wtf/modules/transmission"
	"github.com/cryog0at/wtf/modules/travisci"
	"github.com/cryog0at/wtf/modules/twitch"
	"github.com/cryog0at/wtf/modules/twitter"
	"github.com/cryog0at/wtf/modules/twitterstats"
	"github.com/cryog0at/wtf/modules/unknown"
	"github.com/cryog0at/wtf/modules/uptimerobot"
	"github.com/cryog0at/wtf/modules/victorops"
	"github.com/cryog0at/wtf/modules/weatherservices/arpansagovau"
	"github.com/cryog0at/wtf/modules/weatherservices/prettyweather"
	"github.com/cryog0at/wtf/modules/weatherservices/weather"
	"github.com/cryog0at/wtf/modules/zendesk"
	"github.com/cryog0at/wtf/wtf"
	"github.com/olebedev/config"
	"github.com/rivo/tview"
)

// MakeWidget creates and returns instances of widgets
func MakeWidget(
	tviewApp *tview.Application,
	pages *tview.Pages,
	moduleName string,
	config *config.Config,
) wtf.Wtfable {
	var widget wtf.Wtfable

	moduleConfig, _ := config.Get("wtf.mods." + moduleName)

	// Don' try to initialize modules that don't exist
	if moduleConfig == nil {
		return nil
	}

	// Don't try to initialize modules that aren't enabled
	if enabled := moduleConfig.UBool("enabled", false); !enabled {
		return nil
	}

	// Always in alphabetical order
	switch moduleConfig.UString("type", moduleName) {
	case "arpansagovau":
		settings := arpansagovau.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = arpansagovau.NewWidget(tviewApp, settings)
	case "asana":
		settings := asana.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = asana.NewWidget(tviewApp, pages, settings)
	case "azuredevops":
		settings := azuredevops.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = azuredevops.NewWidget(tviewApp, pages, settings)
	case "bamboohr":
		settings := bamboohr.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = bamboohr.NewWidget(tviewApp, settings)
	case "bargraph":
		settings := bargraph.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = bargraph.NewWidget(tviewApp, settings)
	case "bittrex":
		settings := bittrex.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = bittrex.NewWidget(tviewApp, settings)
	case "blockfolio":
		settings := blockfolio.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = blockfolio.NewWidget(tviewApp, settings)
	case "buildkite":
		settings := buildkite.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = buildkite.NewWidget(tviewApp, pages, settings)
	case "cdsFavorites":
		settings := cdsfavorites.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = cdsfavorites.NewWidget(tviewApp, pages, settings)
	case "cdsQueue":
		settings := cdsqueue.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = cdsqueue.NewWidget(tviewApp, pages, settings)
	case "cdsStatus":
		settings := cdsstatus.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = cdsstatus.NewWidget(tviewApp, pages, settings)
	case "circleci":
		settings := circleci.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = circleci.NewWidget(tviewApp, settings)
	case "clocks":
		settings := clocks.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = clocks.NewWidget(tviewApp, settings)
	case "covid":
		settings := covid.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = covid.NewWidget(tviewApp, settings)
	case "cmdrunner":
		settings := cmdrunner.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = cmdrunner.NewWidget(tviewApp, settings)
	case "cryptolive":
		settings := cryptolive.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = cryptolive.NewWidget(tviewApp, settings)
	case "datadog":
		settings := datadog.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = datadog.NewWidget(tviewApp, pages, settings)
	case "devto":
		settings := devto.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = devto.NewWidget(tviewApp, pages, settings)
	case "digitalclock":
		settings := digitalclock.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = digitalclock.NewWidget(tviewApp, settings)
	case "digitalocean":
		settings := digitalocean.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = digitalocean.NewWidget(tviewApp, pages, settings)
	case "docker":
		settings := docker.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = docker.NewWidget(tviewApp, pages, settings)
	case "feedreader":
		settings := feedreader.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = feedreader.NewWidget(tviewApp, pages, settings)
	case "football":
		settings := football.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = football.NewWidget(tviewApp, pages, settings)
	case "gcal":
		settings := gcal.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = gcal.NewWidget(tviewApp, settings)
	case "gerrit":
		settings := gerrit.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = gerrit.NewWidget(tviewApp, pages, settings)
	case "git":
		settings := git.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = git.NewWidget(tviewApp, pages, settings)
	case "github":
		settings := github.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = github.NewWidget(tviewApp, pages, settings)
	case "gitlab":
		settings := gitlab.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = gitlab.NewWidget(tviewApp, pages, settings)
	case "gitlabtodo":
		settings := gitlabtodo.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = gitlabtodo.NewWidget(tviewApp, pages, settings)
	case "gitter":
		settings := gitter.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = gitter.NewWidget(tviewApp, pages, settings)
	case "googleanalytics":
		settings := googleanalytics.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = googleanalytics.NewWidget(tviewApp, settings)
	case "gspreadsheets":
		settings := gspreadsheets.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = gspreadsheets.NewWidget(tviewApp, settings)
	case "grafana":
		settings := grafana.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = grafana.NewWidget(tviewApp, pages, settings)
	case "hackernews":
		settings := hackernews.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = hackernews.NewWidget(tviewApp, pages, settings)
	case "healthchecks":
		settings := healthchecks.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = healthchecks.NewWidget(tviewApp, pages, settings)
	case "hibp":
		settings := hibp.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = hibp.NewWidget(tviewApp, settings)
	case "ipapi":
		settings := ipapi.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = ipapi.NewWidget(tviewApp, settings)
	case "ipinfo":
		settings := ipinfo.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = ipinfo.NewWidget(tviewApp, settings)
	case "jenkins":
		settings := jenkins.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = jenkins.NewWidget(tviewApp, pages, settings)
	case "jira":
		settings := jira.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = jira.NewWidget(tviewApp, pages, settings)
	case "kubernetes":
		settings := kubernetes.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = kubernetes.NewWidget(tviewApp, settings)
	case "krisinformation":
		settings := krisinformation.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = krisinformation.NewWidget(tviewApp, settings)
	case "logger":
		settings := logger.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = logger.NewWidget(tviewApp, settings)
	case "mercurial":
		settings := mercurial.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = mercurial.NewWidget(tviewApp, pages, settings)
	case "nbascore":
		settings := nbascore.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = nbascore.NewWidget(tviewApp, pages, settings)
	case "newrelic":
		settings := newrelic.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = newrelic.NewWidget(tviewApp, pages, settings)
	case "opsgenie":
		settings := opsgenie.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = opsgenie.NewWidget(tviewApp, settings)
	case "pagerduty":
		settings := pagerduty.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = pagerduty.NewWidget(tviewApp, settings)
	case "pihole":
		settings := pihole.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = pihole.NewWidget(tviewApp, pages, settings)
	case "power":
		settings := power.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = power.NewWidget(tviewApp, settings)
	case "prettyweather":
		settings := prettyweather.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = prettyweather.NewWidget(tviewApp, settings)
	case "pocket":
		settings := pocket.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = pocket.NewWidget(tviewApp, pages, settings)
	case "resourceusage":
		settings := resourceusage.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = resourceusage.NewWidget(tviewApp, settings)
	case "rollbar":
		settings := rollbar.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = rollbar.NewWidget(tviewApp, pages, settings)
	case "security":
		settings := security.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = security.NewWidget(tviewApp, settings)
	case "spacex":
		settings := spacex.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = spacex.NewWidget(tviewApp, settings)
	case "spotify":
		settings := spotify.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = spotify.NewWidget(tviewApp, pages, settings)
	case "spotifyweb":
		settings := spotifyweb.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = spotifyweb.NewWidget(tviewApp, pages, settings)
	case "status":
		settings := status.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = status.NewWidget(tviewApp, settings)
	case "subreddit":
		settings := subreddit.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = subreddit.NewWidget(tviewApp, pages, settings)
	case "textfile":
		settings := textfile.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = textfile.NewWidget(tviewApp, pages, settings)
	case "todo":
		settings := todo.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = todo.NewWidget(tviewApp, pages, settings)
	case "todo_plus":
		settings := todo_plus.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = todo_plus.NewWidget(tviewApp, pages, settings)
	case "todoist":
		settings := todo_plus.FromTodoist(moduleName, moduleConfig, config)
		widget = todo_plus.NewWidget(tviewApp, pages, settings)
	case "transmission":
		settings := transmission.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = transmission.NewWidget(tviewApp, pages, settings)
	case "travisci":
		settings := travisci.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = travisci.NewWidget(tviewApp, pages, settings)
	case "trello":
		settings := todo_plus.FromTrello(moduleName, moduleConfig, config)
		widget = todo_plus.NewWidget(tviewApp, pages, settings)
	case "twitch":
		settings := twitch.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = twitch.NewWidget(tviewApp, pages, settings)
	case "twitter":
		settings := twitter.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = twitter.NewWidget(tviewApp, pages, settings)
	case "twitterstats":
		settings := twitterstats.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = twitterstats.NewWidget(tviewApp, pages, settings)
	case "uptimerobot":
		settings := uptimerobot.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = uptimerobot.NewWidget(tviewApp, pages, settings)
	case "victorops":
		settings := victorops.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = victorops.NewWidget(tviewApp, settings)
	case "weather":
		settings := weather.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = weather.NewWidget(tviewApp, pages, settings)
	case "zendesk":
		settings := zendesk.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = zendesk.NewWidget(tviewApp, pages, settings)
	case "exchangerates":
		settings := exchangerates.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = exchangerates.NewWidget(tviewApp, pages, settings)
	case "finnhub":
		settings := finnhub.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = finnhub.NewWidget(tviewApp, settings)
	case "yfinance":
		settings := yfinance.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = yfinance.NewWidget(tviewApp, settings)
	default:
		settings := unknown.NewSettingsFromYAML(moduleName, moduleConfig, config)
		widget = unknown.NewWidget(tviewApp, settings)
	}

	return widget
}

// MakeWidgets creates and returns a collection of enabled widgets
func MakeWidgets(tviewApp *tview.Application, pages *tview.Pages, config *config.Config) []wtf.Wtfable {
	widgets := []wtf.Wtfable{}

	moduleNames, _ := config.Map("wtf.mods")

	for moduleName := range moduleNames {
		widget := MakeWidget(tviewApp, pages, moduleName, config)

		if widget != nil {
			widgets = append(widgets, widget)
		}
	}

	return widgets
}
