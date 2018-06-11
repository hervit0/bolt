package main

import (
	"github.com/urfave/cli"
)

func DockerComposeExec(context *cli.Context, runner Runner) error {
	err := runner.Execute("docker-compose", "up", "-d", "--build")

	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	name := context.String("container-name")
	command := context.String("container-command")

	err = runner.Execute("docker-compose", "exec", name, command)

	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	return nil
}
