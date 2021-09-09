/*
Copyright © 2021 Elijah J. Passmore <elijah@elijahjpassmore.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	FilePath string
	KeyId    string
	HoardPath string
}

var cfgFile string
var userConfig Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gohoard",
	Short: "gohoard is a super simple password manager",
	Long:  "gohoard is a super simple prescriptive password manager written in Go utilizing GPG encryption.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&userConfig.FilePath, "config", "", "config file (default is $HOME/.config/gohoard/gohoard.toml)")
}

// initConfig reads in config file and ENV variables if set
func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	viper.SetDefault("hoardpath", fmt.Sprintf("%s/.gohoard/", home))

	if userConfig.FilePath != "" {
		// Use config file from the flag.
		viper.SetConfigFile(userConfig.FilePath)
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("missing config")
		} else {
			readConfig()
		}
	} else {
		userConfig.FilePath = fmt.Sprintf("%s/.config/gohoard/gohoard.toml", home)
		viper.SetConfigFile(userConfig.FilePath)
		viper.AutomaticEnv() // read in environment variables that match

		if err := viper.ReadInConfig(); err != nil {
			os.Create(userConfig.FilePath) // create empty config file
			viper.WriteConfig()
			writeNewConfig()
		} else {
			readConfig()
		}
	}
}

//readConfig reads config keys to an instance of Config
func readConfig() {
	keyId := viper.Get("keyid")
	if keyId == nil {
		fmt.Println("no key specified in config file")
		os.Exit(1)
	}
	hoardPath := viper.Get("hoardpath")
	userConfig.KeyId = fmt.Sprintf("%s", keyId)
	userConfig.HoardPath = fmt.Sprintf("%s", hoardPath)
}

//writeNewConfig writes a new config file to userConfig.FilePath
func writeNewConfig() {
	// Get user input through terminal.
	fmt.Print("gpg key ID (gpg --list-keys): ")
	var keyId string
	fmt.Scanln(&keyId)

	// Create the user config.
	userConfig.KeyId = keyId

	// Set the variables.
	viper.Set("keyid", keyId)

	// Write the new config.
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println("test")
		fmt.Println(err.Error())
	}
}
