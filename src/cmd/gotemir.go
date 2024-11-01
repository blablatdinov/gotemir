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
	"path/filepath"
	"strings"

	gotemir "github.com/blablatdinov/gotemir/src/logic"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

var errOptions = errors.New("you must provide both source and test directories")

func main() {
	app := &cli.App{ //nolint:exhaustruct
		Name:  "Gotemir",
		Usage: "golang tests mirrow",
		Description: strings.Join(
			[]string{
				"is a tool that verifies if the structure of the test directory",
				"mirrors the structure of the source code directory. It ensures",
				"that for every source file, a corresponding test file exists",
				"in the appropriate directory.\n\n",
				"Example of usage\n",
				"We have next project structure:\n",
				"src/\n",
				"├── main.py\n",
				"├── service/\n",
				"│   └── user.py\n",
				"tests/\n",
				"├── main_test.py\n",
				"└── service/\n",
				"    └── user_test.py\n\n",
				"Run gotemir:\n",
				"./gotemir --ext .py src tests",
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
			config, exitStatus := parseConfig()
			if exitStatus != 0 {
				os.Exit(exitStatus)
			}
			testsDir := gotemir.OsDirectoryCtor(
				cliCtx.Args().Get(1),
				cliCtx.String("ext"),
			)
			cmprd := gotemir.CmprdStructuresCtor(
				gotemir.FilteredByConfigFilesCtor(
					gotemir.ExcludedTestsDirectoryCtor(
						gotemir.OsDirectoryCtor(
							cliCtx.Args().Get(0),
							cliCtx.String("ext"),
						),
						testsDir,
					),
					config.TestFreeFiles,
				),
				gotemir.FilteredByConfigFilesCtor(
					testsDir,
					config.TestHelpers,
				),
			)
			filesWithoutTests := cmprd.FilesWithoutTests()
			testsWithoutSourceFiles := cmprd.TestsWithoutSrcFiles()
			exitStatus = writeResult(filesWithoutTests, testsWithoutSourceFiles)
			os.Exit(exitStatus)
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func parseConfig() (gotemir.Config, int) {
	config := gotemir.Config{} //nolint:exhaustruct
	configFileContent, _ := os.ReadFile(".gotemir.yaml")
	err := yaml.Unmarshal(configFileContent, &config)
	if err != nil {
		fmt.Printf("Fail on parse .gotemir.yaml content: %s\n", err) //nolint:forbidigo
		return gotemir.Config{}, 1                                   //nolint:exhaustruct
	}
	localizedPaths := make([]string, 0)
	for _, testFreeFilePath := range config.TestFreeFiles {
		localized, err := filepath.Localize(testFreeFilePath)
		if err != nil {
			fmt.Printf("Fail on localize path: '%s' err: %s\n", testFreeFilePath, err) //nolint:forbidigo
		}
		localizedPaths = append(localizedPaths, localized) //nolint:wsl
	}
	config.TestFreeFiles = localizedPaths
	return config, 0
}

func writeResult(filesWithoutTests, testsWithoutSourceFiles []string) int {
	if len(filesWithoutTests) == 0 && len(testsWithoutSourceFiles) == 0 {
		fmt.Println("Complete!") //nolint:forbidigo
		return 0
	}
	for _, fileWithoutTest := range filesWithoutTests {
		fmt.Printf("%s:0:0 Not found test for file\n", fileWithoutTest) //nolint:forbidigo
	}
	for _, testWithoutSourceFile := range testsWithoutSourceFiles {
		fmt.Printf("%s:0:0 Not found source file for test\n", testWithoutSourceFile) //nolint:forbidigo
	}
	return 1
}
