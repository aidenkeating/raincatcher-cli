package main

import (
	"os"

	"./commands"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "wfm"
	app.Usage = "A cli for the WFM demo application"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		commands.NewCloneCommand(os.Stdin, os.Stdout),
		commands.NewCleanCommand(os.Stdin, os.Stdout),
		commands.NewInstallCommand(os.Stdin, os.Stdout),
	}

	app.Run(os.Args)
}
