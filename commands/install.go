package commands

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/urfave/cli"
)

type installCommand struct {
	in  io.Reader
	out io.Writer
}

func (cmd *installCommand) Install() cli.Command {
	return cli.Command{
		Name:        "install",
		Description: "install dependencies for all projects in the demo",
		Action:      cmd.installAction,
	}
}

func (cmd *installCommand) installAction(ctx *cli.Context) error {

	// TODO:  Move these to something like a config file
	appRepos, moduleRepos := getRepoLists()

	for _, appRepo := range appRepos {

		_, repoPath := getRepoInfo("../%s", appRepo)

		fmt.Printf("Installing in %s\n", repoPath)

		if err := installNodeModules(repoPath); err != nil {
			return cli.NewExitError("Failed to install dependencies", 1)
		}
	}

	for _, moduleRepo := range moduleRepos {

		_, repoPath := getRepoInfo("../modules/%s", moduleRepo)

		if err := installNodeModules(repoPath); err != nil {
			return cli.NewExitError("Failed to install dependencies", 1)
		}
	}

	return cli.NewExitError("Completed installing all dependencies.", 0)
}

func installNodeModules(repoPath string) error {
	npmCmd := "npm"
	npmArgs := []string{"--prefix", repoPath, "install", repoPath}

	exec.Command(npmCmd, npmArgs...).Run()

	return nil
}

//NewInstallCommand configures a new clean command
func NewInstallCommand(in io.Reader, out io.Writer) cli.Command {
	cmd := &installCommand{
		in:  in,
		out: out,
	}
	return cmd.Install()
}
