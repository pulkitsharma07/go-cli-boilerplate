package main

import (
	"github.com/pulkitsharma07/go-cli-boilerplate/cmds"
	"github.com/pulkitsharma07/go-cli-boilerplate/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test: ./basecommand
func TestBaseCommand(t *testing.T) {
	test.Reset()

	output, err := test.RunCLI(cmd.Command(), "")

	assert.Nil(t, err, "there should be no error for the base command")

	// Modify the regex, add more assertions as expected.
	assert.Regexp(t, "basecommand", output, "Should display whatever is expected")
}
