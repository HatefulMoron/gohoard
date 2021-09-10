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

package pkg

import (
	"github.com/sethvargo/go-diceware/diceware"
	"math/rand"
	"strings"
	"time"
)

var digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var symbols = []string{"!", "(", ")", "-", ".", "?", "[", "]", "_", "`", "~", ";", ":", "@", "#", "$", "%", "^", "&", "*", "+", "="}

//NewPassword generate a new prescriptive password
func NewPassword(length int, isDigits bool, isSymbols bool, isCapitalize bool) string {
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder

	separator := ""
	if isSymbols {
		separator = symbols[rand.Intn(len(symbols))]
	}

	if isDigits {
		sb.WriteString(digits[rand.Intn(len(digits))])
	}
	for sb.Len() < length {
		newWord, err := diceware.Generate(1)
		if err != nil {
			panic(err)
		}

		if isCapitalize {
			sb.WriteString(strings.Title(newWord[0]))
		} else {
			sb.WriteString(newWord[0])
		}
		sb.WriteString(separator)
	}

	newString := sb.String()
	return newString[:len(newString)-1]
}
