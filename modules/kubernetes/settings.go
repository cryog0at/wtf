package kubernetes

import (
	"github.com/cryog0at/wtf/cfg"
	"github.com/cryog0at/wtf/utils"
	"github.com/olebedev/config"
)

const (
	defaultFocusable = false
	defaultTitle     = "Kubernetes"
)

type Settings struct {
	*cfg.Common

	objects    []string `help:"Kubernetes objects to show. Options are: [nodes, pods, deployments]."`
	title      string   `help:"Override the title of widget."`
	kubeconfig string   `help:"Location of a kubeconfig file."`
	namespaces []string `help:"List of namespaces to watch. If blank, defaults to all namespaces."`
	context    string   `help:"Kubernetes context to use. If blank, uses default context"`
}

func NewSettingsFromYAML(name string, moduleConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, moduleConfig, globalConfig),

		objects:    utils.ToStrs(moduleConfig.UList("objects")),
		title:      moduleConfig.UString("title"),
		kubeconfig: moduleConfig.UString("kubeconfig"),
		namespaces: utils.ToStrs(moduleConfig.UList("namespaces")),
		context:    moduleConfig.UString("context"),
	}

	return &settings
}
