package main

import (
	"os"

	bbpak "github.com/isbm/bbpack-info"

	"github.com/urfave/cli/v2"
)

func app(ctx *cli.Context) error {
	m := bbpak.NewBBPakMatcher(ctx.String("path"))

	if ctx.Bool("list") {
		m.FindManifests()
	}

	return nil
}

func main() {
	appname := "bbpak"
	app := &cli.App{
		Version: "0.1 Alpha",
		Name:    appname,
		Usage:   "Query installed packages in Yocto's BitBake",
		Action:  app,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Aliases:  []string{"p"},
				Name:     "path",
				Usage:    "Path to the build",
				Required: true,
			},
			&cli.StringFlag{
				Aliases: []string{"m"},
				Name:    "manifest",
				Usage:   "Name of the manifest",
			},
			&cli.BoolFlag{
				Aliases: []string{"l"},
				Name:    "list",
				Usage:   "List found manifests",
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
