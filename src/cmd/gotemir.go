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
	"fmt"
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
		Flags: []cli.Flag{
			&cli.StringFlag{ //nolint:exhaustruct
				Name:  "ext",
				Value: ".go",
				Usage: "file extension for scan",
			},
		},
		Action: func(cliCtx *cli.Context) error {
			expectedOptionCount := 2
			if cliCtx.NArg() < expectedOptionCount {
				return errOptions
			}
			fmt.Printf("Cli ext=%s\n", cliCtx.String("ext"))
			filesWithoutTests := gotemir.Compare(
				gotemir.OsDirectoryCtor(
					cliCtx.Args().Get(0),
					cliCtx.String("ext"),
				),
				gotemir.OsDirectoryCtor(
					cliCtx.Args().Get(1),
					cliCtx.String("ext"),
				),
			)
			exitStatus := writeResult(filesWithoutTests)
			os.Exit(exitStatus)
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func writeResult(filesWithoutTests []string) int {
	if len(filesWithoutTests) > 0 {
		fmt.Println("Files without tests:") //nolint:forbidigo
	} else {
		fmt.Println("Complete!") //nolint:forbidigo
		return 0
	}
	for _, fileWithoutTest := range filesWithoutTests {
		fmt.Printf(" - %s\n", fileWithoutTest) //nolint:forbidigo
	}
	return 1
}
