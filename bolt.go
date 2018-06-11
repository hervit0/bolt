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

	cmd.Stdin = os.Stdin
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
		{
			Name:    "docker-compose-exec",
			Aliases: []string{"dex"},
			Usage:   "Build, run and exec on to a local container",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "container-name,n",
					Value: "lead-store",
					Usage: "Name of the container you want to login into. This is defined in the docker-compose.yml",
				},
				cli.StringFlag{
					Name:  "container-command,c",
					Value: "bash",
					Usage: "The command you want to run on the container. Use `bash` or `sh` if you are looking to develop within the container",
				},
			},
			Action: func(c *cli.Context) error {
				runner := CommandLineRunner{}
				return DockerComposeExec(c, runner)
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
