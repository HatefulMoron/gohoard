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

// digCmd represents the dig command
var digCmd = &cobra.Command{
	Use:   "dig",
	Short: "Copy a password from the password hoard",
	Long: "Copy a password from the password hoard to the clipboard.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dig called")
	},
}

func init() {
	rootCmd.AddCommand(digCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// digCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// digCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
