package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func initApp() *cli.App {
	app := cli.NewApp()
	app.Name = "SwitchHosts!"
	app.Usage = "Switch your hosts!"
	app.Version = "0.0.1"
	app.Authors = []cli.Author{
		{
			Name:  "Yuvv",
			Email: "yuvv_th@outlook.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "global,g",
			Usage: "Global hosts configuration",
		},
	}

	app.Commands = []cli.Command{
		addCommand,
		editCommand,
		delCommand,
		listCommand,
		switchCommand,
		viewCommand,
	}

	return app
}

func main() {
	app := initApp()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
