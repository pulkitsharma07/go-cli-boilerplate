package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"path"

	"github.com/pulkitsharma07/go-cli-boilerplate/constants"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// ConfigFilePath stores the path to the config file to use.
	ConfigFilePath string
	// isVerbose stores whether the verbose flag is passed or not (used for setting the default log level)
	isVerbose bool
)

// baseCmd represents the base command when called without any subcommands
var baseCmd = &cobra.Command{
	Use:   "basecommand",
	Short: "Short description for your base command",
	Long: `
+-+-+-+-+-+-+-+-+-+-+-+-+
|.|.|.|B|a|n|n|e|r|.|.|.|
+-+-+-+-+-+-+-+-+-+-+-+-+
<Some additional Text>
  `,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Help()
		return nil
	},
}

// Command is used for accessing baseCmd in tests (and generating docs, refer docs/doc_generator.go)
func Command() *cobra.Command {
	return baseCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the baseCmd.
func Execute() error {
	err := baseCmd.Execute()
	return err
}

func init() {
	cobra.OnInitialize(initConfig)
	baseCmd.PersistentFlags().StringVar(&ConfigFilePath, "config", "", "config file (default is $HOME/.basecommand.yaml)")
	baseCmd.PersistentFlags().BoolVarP(&isVerbose, "verbose", "v", false, "Run in verbose mode")

	log.SetOutput(os.Stdout)
	initCommandTree()
}

func initCommandTree() {
	//subCommand.Init(baseCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if isVerbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	viper.SetConfigType("yaml")

	if ConfigFilePath != "" {
		// Use config file from the flag.
		viper.SetConfigFile(ConfigFilePath)
	} else {
		initViper()
	}

	// Set the default values in Viper
	setViperDefaults()

	// Set defaults for config
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Debug("Loading config from: ", viper.ConfigFileUsed())
	}
}

func setViperDefaults() {
	viper.SetDefault(constants.ViperBaseURLKey, constants.DefaultBaseURL)
}

func initViper() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		log.Error(err)
		return
	}

	// Search config in home directory with name ".basecommand" (without extension).
	viper.AddConfigPath(home)
	ConfigFilePath = path.Join(home, ".basecommand.yaml")

	// BASECOMMAND_ENV_VAR can be passed into the cli, which can be fetched using: viper.Get("env_var")
	viper.SetEnvPrefix("basecommand")
	viper.SetConfigName(".basecommand")
}
