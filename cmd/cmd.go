package main

import (
	"fmt"

	"github.com/urfave/cli"
)

var commands = []cli.Command{
	{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print version",
		Action: func(c *cli.Context) error {
			if version == "" {
				version = defaultVersion
			}
			fmt.Printf("Version: %s", version)
			return nil
		},
	},
}
