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
		password, err := getPassword(args[0])
		if err == nil {
			clipboard.WriteAll(password)
		} else {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(digCmd)
}

//getPassword get the password stored at the given path
func getPassword(filePath string) (string, error) {
	// TODO: decrypt now encrypted file
	fullPath := fmt.Sprintf("%s/.gohoard/%s", os.Getenv("HOME"), filePath)
	fullPathEncrypted := fmt.Sprintf("%s/.gohoard/%s.gpg", os.Getenv("HOME"), filePath)

	err := decryptFile(fullPathEncrypted)
	if err != nil {
		return "", errors.New(fmt.Sprintf("missing password: %s", fullPathEncrypted))
	}

	password, _ := os.ReadFile(fullPath)
	os.Remove(fullPath)

	return string(password), nil
}
func decryptFile(filePath string) error {
	cmd := exec.Command("gpg", filePath)
	return cmd.Run()
}
