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
	Use:        "stash <name ...>",
	Short:      "Generate a new password and put it in the password hoard",
	Long:       "Generate a new password and put it in the password hoard.",
	Args:       cobra.MinimumNArgs(1),
	SuggestFor: []string{"add", "new", "create"},
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")

		minLength, _ := cmd.Flags().GetInt("min-length")
		overwrite, _ := cmd.Flags().GetBool("overwrite")
		digits, _ := cmd.Flags().GetBool("no-digits")
		symbols, _ := cmd.Flags().GetBool("no-symbols")
		capitals, _ := cmd.Flags().GetBool("no-capitals")

		for _, path := range args {
			if verbose {
				fmt.Println(fmt.Sprintf("stash password path: %s", path))
			}
			password, err := pkg.NewPassword(minLength, !digits, !symbols, !capitals)
			if err != nil {
				println(err.Error())
			}
			err = stashPassword([]byte(password), path, overwrite)
			if err != nil {
				println(err.Error())
			}
		}
	},
}

//init adds the new command to rootCmd
func init() {
	rootCmd.AddCommand(stashCmd)

	stashCmd.Flags().IntP("min-length", "l", 30, "minimum password length")
	stashCmd.Flags().BoolP("overwrite", "o", false, "overwrite existing passwords")
	stashCmd.Flags().BoolP("no-digits", "d", false, "omit digits")
	stashCmd.Flags().BoolP("no-symbols", "s", false, "omit symbols")
	stashCmd.Flags().BoolP("no-capitals", "c", false, "omit capitals")
}

//stashPassword stores a given password to a given file path
func stashPassword(password []byte, hoardPath string, overwrite bool) error {
	dir := fmt.Sprintf("%s%s", userConfig.HoardPath, filepath.Dir(hoardPath))
	file := filepath.Base(hoardPath) // file name without dir
	decryptedPath := fmt.Sprintf("%s/%s", dir, file)
	encryptedPath := fmt.Sprintf("%s.gpg", decryptedPath)

	// Create directories and file.
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	// Either return if the password already exists, or overwrite it.
	isExisting := fileExists(encryptedPath)
	if isExisting {
		if overwrite {
			err = os.Remove(encryptedPath)
			if err != nil {
				return err
			}
		} else {
			return errors.New(fmt.Sprintf("%s already exists in hoard", hoardPath))
		}
	}
	err = os.WriteFile(decryptedPath, password, 0644)
	if err != nil {
		return err
	}

	// Copy the password to the clipboard.
	err = clipboard.WriteAll(string(password))
	if err != nil {
		return errors.New("missing CLI clipboard (e.g xclip)")
	}

	// Encrypt the file.
	err = encryptFile(decryptedPath)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to encrypt: %s", decryptedPath))
	}
	err = os.Remove(decryptedPath)
	if err != nil {
		return err
	}

	return nil
}

//encryptFile encrypts a given file
func encryptFile(filePath string) error {
	cmd := exec.Command("gpg", "-r", userConfig.KeyId, "-e", filePath)

	return cmd.Run()
}
