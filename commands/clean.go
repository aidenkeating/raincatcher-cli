package commands

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/urfave/cli"
)

type cleanCommand struct {
	in  io.Reader
	out io.Writer
}

func (cmd *cleanCommand) Clean() cli.Command {
	return cli.Command{
		Name:        "clean",
		Description: "remove the required repositories for the demo application",
		Action:      cmd.cleanAction,
	}
}

func (cmd *cleanCommand) cleanAction(ctx *cli.Context) error {

	appRepos, moduleRepos := getRepoLists()

	for _, appRepo := range appRepos {

		_, repoPath := getRepoInfo("../%s", appRepo)

		fmt.Printf("Removing %s\n", repoPath)

		if err := cleanRepo(repoPath); err != nil {
			return cli.NewExitError("Failed to clone repository", 1)
		}
	}

	for _, moduleRepo := range moduleRepos {

		_, repoPath := getRepoInfo("../modules/%s", moduleRepo)

		if err := cleanRepo(repoPath); err != nil {
			return cli.NewExitError("Failed to clone repository", 1)
		}
	}

	return cli.NewExitError("Completed cleaning all repositories.", 0)
}

func cleanRepo(repoPath string) error {

	nodeModulesFolder := fmt.Sprintf("%s/node_modules", repoPath)

	fmt.Printf("Deleting %s\n", nodeModulesFolder)

	gitCmd := "rm"
	gitArgs := []string{"-rf", nodeModulesFolder}
	if err := exec.Command(gitCmd, gitArgs...).Run(); err != nil {
		return err
	}
	return nil
}

//NewCleanCommand configures a new clean command
func NewCleanCommand(in io.Reader, out io.Writer) cli.Command {
	cmd := &cleanCommand{
		in:  in,
		out: out,
	}
	return cmd.Clean()
}
