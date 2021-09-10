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
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// digCmd represents the dig command
var digCmd = &cobra.Command{
	Use:        "dig",
	Short:      "Copy a password from the password hoard",
	Long:       "Copy a password from the password hoard to the clipboard.",
	Args:       cobra.MinimumNArgs(1),
	SuggestFor: []string{"copy", "get"},
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")

		password, err := getPassword(args[0])
		if err == nil {
			if verbose {
				fmt.Println(fmt.Sprintf("%s copied to clipboard", args[0]))
			}
			err = clipboard.WriteAll(password)
			if err != nil {
				println("missing CLI clipboard (e.g. xclip)")
			}
		} else {
			fmt.Println(err.Error())
		}
	},
}

//init adds the new command to rootCmd
func init() {
	rootCmd.AddCommand(digCmd)
}

//getPassword get the password stored at the given path
func getPassword(filePath string) (string, error) {
	decryptedPath := fmt.Sprintf("%s%s", userConfig.HoardPath, filePath)
	encryptedPath := fmt.Sprintf("%s%s.gpg", userConfig.HoardPath, filePath)

	// Check if the file currently exists.
	_, err := os.OpenFile(encryptedPath, os.O_RDONLY, 0644)
	if os.IsNotExist(err) {
		fmt.Println("test")
		return "", err
	}

	// See if the user is able to decrypt the file.
	err = decryptFile(encryptedPath)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to decrypt: %s, check key ID", encryptedPath))
	}
	password, _ := os.ReadFile(decryptedPath)
	err = os.Remove(decryptedPath)
	if err != nil {
		return "", errors.New(fmt.Sprintf("cannot remove file: %s", decryptedPath))
	}

	return string(password), nil
}

//decryptFile decrypts a file and returns the result
func decryptFile(filePath string) error {
	cmd := exec.Command("gpg", filePath)
	return cmd.Run()
}
