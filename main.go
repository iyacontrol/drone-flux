package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

var build = "0" // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "flux plugin"
	app.Usage = "flux plugin"
	app.Action = run
	app.Version = fmt.Sprintf("1.0.0+%s", build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "url",
			Usage:  "ase URL of the flux controller",
			EnvVar: "FLUX_URL",
		},
		cli.StringFlag{
			Name:   "token",
			Usage:  "Weave Cloud controller token",
			EnvVar: "FLUX_SERVICE_TOKEN",
		},
		cli.StringFlag{
			Name:   "namespace",
			Usage:  "controller namespace",
			EnvVar: "PLUGIN_NAMESPACE,NAMESPACE",
		},
		cli.StringSliceFlag{
			Name:   "controller",
			Usage:  "list of controllers to release <kind>/<name>",
			EnvVar: "CONTROLLER",
		},
		cli.StringSliceFlag{
			Name:   "exclude",
			Usage:  "exclude a controller",
			EnvVar: "EXCLUDE",
		},
		cli.StringFlag{
			Name:   "update-image",
			Usage:  "update a specific image",
			EnvVar: "UPDATE_IMAGE",
		},
		cli.StringFlag{
			Name:   "user",
			Usage:  "override the user reported as initiating the update",
			EnvVar: "USER",
		},
		cli.StringFlag{
			Name:   "message",
			Usage:  "attach a message to the update",
			EnvVar: "MESSAGE",
		},
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	if c.String("env-file") != "" {
		_ = godotenv.Load(c.String("env-file"))
	}
}
