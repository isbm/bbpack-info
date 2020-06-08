package main

import (
	"fmt"
	"os"
	"path"
	"sort"
	"strings"

	bbpak "github.com/isbm/bbpack-info"

	"github.com/urfave/cli/v2"
)

func app(ctx *cli.Context) error {
	m := bbpak.NewBBPakMatcher(ctx.String("path"))

	format := ""
	for _, fmt := range []string{"txt", "csv", "md", "json", ""} {
		if ctx.String("format") == fmt {
			if fmt == "" {
				fmt = "stdout"
			}
			format = fmt
			break
		}
	}
	if format == "" {
		fmt.Printf("Error: format %s not supported\n", ctx.String("format"))
		os.Exit(1)
	}

	if ctx.Bool("list") {
		manifests, err := m.FindManifests()
		if err != nil {
			sort.Strings(manifests)
			fmt.Println("Available manifests:")
			for idx, mfs := range manifests {
				fmt.Printf("  %d. %s\n", idx+1, strings.Split(path.Base(mfs), ".")[0])
			}
		} else {
			fmt.Printf("Error: %s", err.Error())
			os.Exit(1)
		}
	} else {
		// Get manifest
		if ctx.String("manifest") == "" {
			fmt.Println("Error: no manifest has been specified.")
			os.Exit(1)
		} else {
			m.SetTargetManifest(ctx.String("manifest"))
			_, err := m.FindManifests()
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				os.Exit(1)
			}
		}
		m.ParseManifestPackages()

		if ctx.String("package") != "" {
			m.FindRequestedPackage(ctx.String("package"))
			m.Format(format)
		} else {
			m.FindPhysicalPackages()
			m.Format(format)
		}
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
				Usage:   "Name of the manifest to reference package index",
			},
			&cli.StringFlag{
				Aliases: []string{"f"},
				Name:    "format",
				Usage:   "Output in: csv, md, json, txt",
			},
			&cli.BoolFlag{
				Aliases: []string{"l"},
				Name:    "list",
				Usage:   "List available manifests",
			},
			&cli.StringFlag{
				Aliases: []string{"g"},
				Name:    "package",
				Usage:   "Display package information",
			},
			&cli.StringFlag{
				Aliases: []string{"t"},
				Name:    "patch",
				Usage:   "Display package patch info (requires package name)",
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
