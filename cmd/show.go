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

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:        "show",
	Short:      "Print a password from the stash",
	Long:       "Print a password from the stash to the terminal.",
	Args:       cobra.MinimumNArgs(1),
	SuggestFor: []string{"print"},
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")

		for _, path := range args {
			if verbose {
				fmt.Println(fmt.Sprintf("Showing password path: %s", path))
			}
			password, err := getPassword(path)
			if err == nil {
				fmt.Println(password)
			} else {
				fmt.Println(err.Error())
			}
		}
	},
}

//init adds the new command to rootCmd
func init() {
	rootCmd.AddCommand(showCmd)
}
