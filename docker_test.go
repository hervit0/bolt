package main

import (
	"flag"
	"testing"
)



func TestDockerComposeExec(t *testing.T) {
	containerName := "lead-store"
	containerCmd := "bash"

	testRunner := new(MockRunner)
	testRunner.On("Execute", "docker-compose", []string{"up", "-d", "--build"}).Return(nil).Once()
	testRunner.On("Execute", "docker-compose", []string{"exec", containerName, containerCmd}).Return(nil).Once()

	flags := flag.NewFlagSet("tag", 0)
	flags.String("container-name", containerName, "doc")
	flags.String("container-command", containerCmd, "doc")
	context := GetCliContext(flags)

	DockerComposeExec(context, testRunner)

	testRunner.AssertExpectations(t)
}