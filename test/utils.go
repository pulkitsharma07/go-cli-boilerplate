package test

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// Reset should be called after each test case to clear any state which can
// persist between test runs.
func Reset() {
	// Resets Viper
	viper.Reset()

	// Reset ENV
	// To ensure that ENV doesn't leak for the next test case
	pathBackup := os.Getenv("PATH")
	os.Clearenv()
	os.Setenv("PATH", pathBackup)
}

// RunCLI is used to launch the cli.
// commandLine can be passed any string which user can pass from the command line
func RunCLI(cmd *cobra.Command, commandLine string) (string, error) {
	// Buffer to capture output generated from the CLI
	buf := new(bytes.Buffer)

	// Capture both stdout and stderr
	cmd.SetOutput(buf)

	// Passing in the commandLine arguments
	cmd.SetArgs(strings.Split(commandLine, " "))

	// Execute cli
	err := cmd.Execute()

	// Merge everything into a single line for matching Regex
	// TODO: make this optional ?
	output := strings.Replace(buf.String(), "\n", "", -1)

	if err != nil {
		return output, err
	}

	return output, nil
}
