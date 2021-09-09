/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"github.com/pelletier/go-toml"
	"os"

	"github.com/spf13/cobra"
)

type Config struct {
	KeyId string
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newConfig := Config{
			KeyId: args[0],
		}
		writeConfig(newConfig)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func writeConfig(newConfig Config) {
	dir := fmt.Sprintf("%s/.config/gohoard", os.Getenv("HOME"))
	filePath := fmt.Sprintf("%s/config.toml", dir)
	os.MkdirAll(dir, os.ModePerm)
	file, _ := os.Create(filePath)
	toml.NewEncoder(file).Encode(newConfig)
}

func getConfigField(field string) string {
	dir := fmt.Sprintf("%s/.config/gohoard", os.Getenv("HOME"))
	filePath := fmt.Sprintf("%s/config.toml", dir)
	fileData, _ := os.ReadFile(filePath)
	config, _ := toml.Load(string(fileData))
	return config.Get(field).(string)
}
