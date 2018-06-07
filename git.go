package main

import (
	"github.com/urfave/cli"
)

func GitLog(_ *cli.Context, runner Runner) error {
	err := runner.Execute("git", "log", "--graph", "--abbrev-commit", "--decorate")

	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	return nil
}
