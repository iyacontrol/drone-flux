package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
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
			Usage:  "base URL of the flux controller",
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
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.author",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.StringFlag{
			Name:  "env-file",
			Usage: "source env file",
		},
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Repo: Repo{
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
		},
		Build: Build{
			Number:  c.Int("build.number"),
			Event:   c.String("build.event"),
			Status:  c.String("build.status"),
			Commit:  c.String("commit.sha"),
			Branch:  c.String("commit.branch"),
			Author:  c.String("commit.author"),
			Message: c.String("commit.message"),
			Link:    c.String("build.link"),
		},
		Config: Config{
			URL:         c.String("url"),
			Token:       c.String("token"),
			Namespace:   c.String("namespace"),
			Controller:  c.StringSlice("controller"),
			Exclude:     c.StringSlice("exclude"),
			UpdateImage: c.String("update-image"),
			User:        c.String("user"),
			Message:     c.String("message"),
		},
	}

	return plugin.Exec()
}
