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
	"github.com/elijahjpassmore/gohoard/pkg"
	"github.com/spf13/cobra"
)

// stashCmd represents the stash command
var stashCmd = &cobra.Command{
	Use:   "stash",
	Short: "Put a password in the password hoard",
	Long:  "Put a password in the password hoard.",
	Args: cobra.MinimumNArgs(1),
	SuggestFor: []string{"add", "new", "create"},
	Run: func(cmd *cobra.Command, args []string) {
		digits, _ := cmd.Flags().GetBool("no-digits")
		symbols, _ := cmd.Flags().GetBool("no-symbols")
		capitals, _ := cmd.Flags().GetBool("no-capitals")

		password := pkg.NewPassword(20, !digits, !symbols, !capitals)
		fmt.Println(password)
	},
}

func init() {
	rootCmd.AddCommand(stashCmd)

	stashCmd.Flags().BoolP("no-digits", "d", false, "omit digits from new password")
	stashCmd.Flags().BoolP("no-symbols", "s", false, "omit symbols from new password")
	stashCmd.Flags().BoolP("no-capitals", "c", false, "omit capitals from new password")
}
