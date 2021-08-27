package main

import (
	"os"

	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/chronos"
	"github.com/paketo-buildpacks/packit/pexec"
	"github.com/paketo-buildpacks/packit/scribe"
	noderunscript "github.com/railwayapp/node-run-script"
)

func main() {
	npmExec := pexec.NewExecutable("npm")
	yarnExec := pexec.NewExecutable("yarn")
	scriptManager := noderunscript.NewScriptManager()
	logger := scribe.NewLogger(os.Stdout)

	packit.Run(
		noderunscript.Detect(scriptManager),
		noderunscript.Build(npmExec, yarnExec, scriptManager, chronos.DefaultClock, logger),
	)
}
