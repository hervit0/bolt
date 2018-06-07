package main

import (
	"testing"
	"io/ioutil"

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

func TestGitLog(t *testing.T) {
	testRunner := new(MockRunner)
	testRunner.On("Execute", "git", []string{"log", "--graph", "--abbrev-commit", "--decorate"}).Return(nil).Once()

	context := cli.NewContext(getCliApp(), nil, nil)

	GitLog(context, testRunner)

	testRunner.AssertExpectations(t)
}

func getCliApp() *cli.App {
	app := cli.NewApp()
	app.Writer = ioutil.Discard
	return app
}
