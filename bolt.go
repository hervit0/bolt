package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

type Runner interface {
	Execute(command string, args ...string) error
}

type CommandLineRunner struct{}

func (runner CommandLineRunner) Execute(command string, args ...string) error {
	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true

	app.Name = "bolt"
	app.Usage = "Bolt is a lightning fast CLI tool that speeds up development"
	app.Action = func(c *cli.Context) error {
		fmt.Println("This is Bolt, try `bolt help` for more!")
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "git-log",
			Aliases: []string{"gl"},
			Usage:   "See the git logs of your repository in your current directory",
			Action: func(c *cli.Context) error {
				runner := CommandLineRunner{}
				return GitLog(c, runner)
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
