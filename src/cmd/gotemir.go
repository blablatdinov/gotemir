// The MIT License (MIT)
//
// Copyright (c) 2024 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE
// OR OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"errors"
	"log"
	"os"
	"strings"

	gotemir "github.com/blablatdinov/gotemir/src/logic"
	"github.com/urfave/cli/v2"
)

var errOptions = errors.New("you must provide both source and test directories")

func main() {
	app := &cli.App{ //nolint:exhaustruct
		Name: "Gotemir",
		Description: strings.Join(
			[]string{
				"is a tool that verifies if the structure of the test directory",
				"mirrors the structure of the source code directory. It ensures",
				"that for every source file, a corresponding test file exists",
				"in the appropriate directory.",
			},
			" ",
		),
		Action: func(cliCtx *cli.Context) error {
			expectedOptionCount := 2
			if cliCtx.NArg() < expectedOptionCount {
				return errOptions
			}
			filesWithoutTests := gotemir.Compare(
				gotemir.OsDirectoryCtor(
					cliCtx.Args().Get(0),
				),
				gotemir.OsDirectoryCtor(
					cliCtx.Args().Get(1),
				),
			)
			log.Println("Files without tests:")
			for _, fileWithoutTest := range filesWithoutTests {
				log.Printf(" - %s\n", fileWithoutTest)
			}
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
