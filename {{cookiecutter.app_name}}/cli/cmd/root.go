package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/config"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var Conf *config.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "{{cookiecutter.app_name}}",
	Short: "graphql server",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/{{cookiecutter.app_config_file}}.yaml)")
	pFlags := rootCmd.PersistentFlags()

	// api server config
	pFlags.String("api.hostip", "", "api server ip")
	pFlags.Int("api.port", 8080, "api server port")

	// database config
	pFlags.String("database.type", "mysql", "database type [sqlite|mysql|postgres]")
	pFlags.String("database.user", "", "database user")
	pFlags.String("database.password", "", "database password")
	pFlags.String("database.name", "", "database name")
	pFlags.String("database.host", "", "database hostname")
	pFlags.Int("database.port", 3306, "database port")
	pFlags.Bool("database.debug", false, "debug database flag")

	if err := viper.BindPFlags(pFlags); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name "{{cookiecutter.app_config_file}}" (without extension).
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc")
		viper.SetConfigName("{{cookiecutter.app_config_file}}")
	}

	viper.AutomaticEnv() // read in environment variables that match

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	Conf = config.NewDefaultConfig()
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		err := viper.Unmarshal(&Conf)
		if err != nil {
			fmt.Printf("unable to decode into config struct: %v\n", err)
		}
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		return
	}

	err := viper.Unmarshal(&Conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct: %v\n", err)
	}
}
