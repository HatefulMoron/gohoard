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
	FilePath  string
	KeyId     string
	HoardPath string
}

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
		// If the user provides their own config path, use config file from the flag.
		viper.SetConfigFile(userConfig.FilePath)
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println(fmt.Sprintf("cannot find config at: %s", userConfig.FilePath))
			os.Exit(1)
		} else {
			readConfig()
		}
	} else {
		// If the user does not provide their own config path, try and read the default.
		userConfig.FilePath = fmt.Sprintf("%s/.config/gohoard/gohoard.toml", home) // default config path
		viper.SetConfigFile(userConfig.FilePath)
		viper.AutomaticEnv() // read in environment variables that match

		if err := viper.ReadInConfig(); err != nil {
			_, err := os.Create(userConfig.FilePath)
			if err != nil {
				fmt.Println(fmt.Sprintf("cannot write to file: %s", userConfig.FilePath))
			}
			writeNewConfig()
			readConfig()
		} else {
			readConfig()
		}
	}
}

//readConfig reads config keys to an instance of Config
func readConfig() {
	keyId := viper.Get("keyid")
	if keyId == nil {
		fmt.Println(fmt.Sprintf(`GPG key must be specified in %s

E.g:
	keyid: "YOUR_KEY_HERE"`, userConfig.FilePath))
		os.Exit(1)
	}
	hoardPath := viper.Get("hoardpath")
	userConfig.KeyId = fmt.Sprintf("%s", keyId)
	userConfig.HoardPath = fmt.Sprintf("%s", hoardPath)
}

//writeNewConfig writes a new config file to userConfig.FilePath
func writeNewConfig() {
	// Get the GPG key ID.
	fmt.Print("GPG key ID (gpg --list-keys): ")
	var keyId string
	_, err := fmt.Scanln(&keyId)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Get the gohoard directory.
	fmt.Print("gohoard directory [$HOME/.gohoard/]: ")
	var hoardPath string
	_, err = fmt.Scanln(&hoardPath)
	if err == nil {
	}

	// Create the user config.
	userConfig.KeyId = keyId

	// Set the variables.
	viper.Set("keyid", keyId)
	if hoardPath != "" {
		viper.Set("hoardpath", hoardPath)
	}

	// Write the new config.
	err = viper.WriteConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
