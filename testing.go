package main

import (
	"io/ioutil"
	"flag"

	"github.com/stretchr/testify/mock"
	"github.com/urfave/cli"
)

type MockRunner struct {
	mock.Mock
}

func (m *MockRunner) Execute(command string, args ...string) error {
	m.Called(command, args)
	return nil
}

func GetCliContext(flags *flag.FlagSet) *cli.Context {
	return cli.NewContext(getCliApp(), flags, nil)
}

func getCliApp() *cli.App {
	app := cli.NewApp()
	app.Writer = ioutil.Discard
	return app
}
