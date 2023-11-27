package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "HealthChecker",
		Usage: "A tiny tool that checks whether the website is running or down",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "Domain name",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port number to check",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			port := ctx.String("port")
			if ctx.String("port") == "" {
				port = "80"
			}
			status := Check(ctx.String("domain"), port)
			fmt.Print(status)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
