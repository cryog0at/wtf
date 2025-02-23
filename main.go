package main

import (
	"fmt"
	"log"
	"os"

	// Blank import of tzdata embeds the timezone database to allow Windows hosts to find timezone
	// information even if the timezone database is not available on the local system. See release
	// notes at https://golang.org/doc/go1.15#time/tzdata for details. This prevents "no timezone
	// data available" errors in clocks module.
	_ "time/tzdata"

	"github.com/logrusorgru/aurora"
	"github.com/pkg/profile"

	"github.com/cryog0at/wtf/app"
	"github.com/cryog0at/wtf/cfg"
	"github.com/cryog0at/wtf/flags"
	"github.com/cryog0at/wtf/utils"
	"github.com/cryog0at/wtf/wtf"
)

var (
	date    = "dev"
	version = "dev"
)

/* -------------------- Main -------------------- */

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Parse and handle flags
	flags := flags.NewFlags()
	flags.Parse()

	// Load the configuration file
	cfg.Initialize(flags.HasCustomConfig())
	config := cfg.LoadWtfConfigFile(flags.ConfigFilePath())

	wtf.SetTerminal(config)

	flags.RenderIf(version, date, config)

	if flags.Profile {
		defer profile.Start(profile.MemProfile).Stop()
	}

	openFileUtil := config.UString("wtf.openFileUtil", "open")
	openURLUtil := utils.ToStrs(config.UList("wtf.openUrlUtil", []interface{}{}))
	utils.Init(openFileUtil, openURLUtil)

	/* Initialize the App Manager */
	appMan := app.NewAppManager()
	appMan.MakeNewWtfApp(config, flags.Config)

	currentApp, err := appMan.Current()
	if err != nil {
		fmt.Printf("\n%s %v\n", aurora.Red("ERROR"), err)
		os.Exit(1)
	}

	currentApp.Run()
}
