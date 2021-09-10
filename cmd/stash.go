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
	"errors"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/elijahjpassmore/gohoard/pkg"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

// stashCmd represents the stash command
var stashCmd = &cobra.Command{
	Use:        "stash",
	Short:      "Put a password in the password hoard",
	Long:       "Put a password in the password hoard.",
	Args:       cobra.MinimumNArgs(1),
	SuggestFor: []string{"add", "new", "create"},
	Run: func(cmd *cobra.Command, args []string) {
		minLength, _ := cmd.Flags().GetInt("min-length")
		digits, _ := cmd.Flags().GetBool("no-digits")
		symbols, _ := cmd.Flags().GetBool("no-symbols")
		capitals, _ := cmd.Flags().GetBool("no-capitals")

		for _, path := range args {
			password := pkg.NewPassword(minLength, !digits, !symbols, !capitals)
			stashPassword([]byte(password), path)
		}
	},
}

//init adds the new command to rootCmd
func init() {
	rootCmd.AddCommand(stashCmd)

	stashCmd.Flags().IntP("min-length", "l", 20, "minimum password length")
	stashCmd.Flags().BoolP("no-digits", "d", false, "omit digits")
	stashCmd.Flags().BoolP("no-symbols", "s", false, "omit symbols")
	stashCmd.Flags().BoolP("no-capitals", "c", false, "omit capitals")
}

//stashPassword stores a given password to a given file path
func stashPassword(password []byte, hoardPath string) error {
	dir := fmt.Sprintf("%s%s", userConfig.HoardPath, filepath.Dir(hoardPath))
	file := filepath.Base(hoardPath) // file name without dir
	fullPath := fmt.Sprintf("%s/%s", dir, file)

	os.MkdirAll(dir, os.ModePerm)
	os.WriteFile(fullPath, password, 0644)

	// Copy the password to the clipboard.
	err := clipboard.WriteAll(string(password))
	if err != nil {
		return errors.New("missing CLI clipboard")
	}

	encryptFile(fullPath)
	os.Remove(fullPath)

	return nil
}

//encryptFile encrypts a given file
func encryptFile(filePath string) {
	cmd := exec.Command("gpg", "-r", userConfig.KeyId, "-e", filePath)
	cmd.Run()
}
