package commands

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/urfave/cli"
)

type cloneCommand struct {
	in  io.Reader
	out io.Writer
}

func (cmd *cloneCommand) Clone() cli.Command {
	return cli.Command{
		Name:        "clone",
		Description: "clone the required repositories for the demo application",
		Action:      cmd.cloneAction,
	}
}

func (cmd *cloneCommand) cloneAction(ctx *cli.Context) error {

	// TODO:  Move these to something like a config file
	appRepos, moduleRepos := getRepoLists()

	for _, appRepo := range appRepos {

		repoURL, repoPath := getRepoInfo("../%s", appRepo)

		fmt.Printf("Cloning %s\n", repoURL)

		if err := cloneRepo(repoURL, repoPath); err != nil {
			return cli.NewExitError("Failed to clone repository", 1)
		}
	}

	for _, moduleRepo := range moduleRepos {

		repoURL, repoPath := getRepoInfo("../modules/%s", moduleRepo)

		fmt.Printf("Cloning %s\n", repoURL)

		if err := cloneRepo(repoURL, repoPath); err != nil {
			return cli.NewExitError("Failed to clone repository", 1)
		}
	}

	return cli.NewExitError("Completed cloning all repositories.", 0)
}

func cloneRepo(repoURL, repoPath string) error {
	gitCmd := "git"
	gitArgs := []string{"clone", repoURL, repoPath}
	if err := exec.Command(gitCmd, gitArgs...).Run(); err != nil {
		return err
	}
	return nil
}

//NewCloneCommand configures a new clone command
func NewCloneCommand(in io.Reader, out io.Writer) cli.Command {
	cmd := &cloneCommand{
		in:  in,
		out: out,
	}
	return cmd.Clone()
}
