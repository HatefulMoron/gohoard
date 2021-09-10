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
)

var version = "0.1.0"
var license = "Apache License, Version 2.0"
var repository = "https://github.com/elijahjpassmore/gohoard"

// aboutCmd represents the about command
var aboutCmd = &cobra.Command{
	Use:        "about",
	Short:      "Print information about gohoard",
	Long:       "Print information about the installed version of gohoard.",
	SuggestFor: []string{"version", "license", "repository"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf(
			`gohoard v%s, licensed under the %s
Read more about the project and report issues at %s`, version, license, repository))
	},
}

//init adds the new command to rootCmd
func init() {
	rootCmd.AddCommand(aboutCmd)
}
