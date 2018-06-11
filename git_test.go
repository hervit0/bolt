package main

import (
	"testing"
)

func TestGitLog(t *testing.T) {
	testRunner := new(MockRunner)
	testRunner.On("Execute", "git", []string{"log", "--graph", "--abbrev-commit", "--decorate"}).Return(nil).Once()

	context := GetCliContext(nil)

	GitLog(context, testRunner)

	testRunner.AssertExpectations(t)
}
