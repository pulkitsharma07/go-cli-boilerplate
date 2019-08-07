package main

import (
	"github.com/pulkitsharma07/go-cli-boilerplate/cmd"
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

// Test: ./basecommand subcommand1
func TestSubCommand1(t *testing.T) {
	test.Reset()

	output, err := test.RunCLI(cmd.Command(), "subcommand1")

	assert.Nil(t, err, "there should be no error for the base command")

	// Modify the regex, add more assertions as expected.
	assert.Regexp(t, "CATS CATS CATS", output, "Should print CATS...")
}

// Test: ./basecommand nestedsubcommand1 subcommand2
// This command is available inside a nested command
func TestSubCommand2(t *testing.T) {
	test.Reset()

	output, err := test.RunCLI(cmd.Command(), "nestedsubcommand1 subcommand2")

	assert.Nil(t, err, "there should be no error for the base command")

	// Modify the regex, add more assertions as expected.
	assert.Regexp(t, "DOGS DOGS DOGS", output, "Should print DOGS...")
}
