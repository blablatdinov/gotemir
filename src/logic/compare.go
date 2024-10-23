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

package logic

import (
	"fmt"
	"strings"
)

func Compare(srcDir Directory, testsDir Directory) []string {
	filesWithoutTests := make([]string, 0)
	testFiles, _ := testsDir.Structure()
	srcFiles, _ := srcDir.Structure()
	for _, srcFile := range srcFiles {
		// absolutePath, _ := srcFile.Absolute()
		relativePath, _ := srcFile.Relative()
		// fmt.Printf("src absolute: %s src relative: %s\n", absolutePath, relativePath)
		splittedSrcFile := strings.Split(relativePath, "/")
		var srcFileRelative string
		if len(splittedSrcFile) == 1 {
			srcFileRelative = relativePath
		} else {
			srcFileRelative = strings.Join(splittedSrcFile[1:], "/")
		}
		// fmt.Printf("src relative: %s splitted: %v\n", srcFileRelative, strings.Split(srcFileRelative, "."))
		fmt.Printf("src relative: %s\n", srcFileRelative)
		// os.Exit(0)
		fileExtension := "." + strings.Split(srcFileRelative, ".")[1]
		splittedPath := strings.Split(srcFileRelative, "/")
		fileName := splittedPath[len(splittedPath)-1]
		fileNameWithoutExtension := strings.Split(fileName, ".")[0]
		testFileVariants := []string{
			strings.Replace(
				srcFileRelative,
				fileName,
				fileNameWithoutExtension+"_test"+fileExtension,
				1,
			),
			strings.Replace(
				srcFileRelative,
				fileName,
				"test_"+fileNameWithoutExtension+fileExtension,
				1,
			),
		}
		testFileFound := false
		fmt.Printf("test variants: %v\n", testFileVariants)
	out:
		for _, testFile := range testFiles {
			// absolutePath, _ := testFile.Absolute()
			relativePath, _ := testFile.Relative()
			// fmt.Printf("src absolute: %s src relative: %s\n", absolutePath, relativePath)
			// splittedTestFile := strings.Split(absolutePath, "/")
			splittedTestFile := strings.Split(relativePath, "/")
			testFileRelative := strings.Join(splittedTestFile[1:], "/")
			for _, testFileVariant := range testFileVariants {
				if testFileRelative == testFileVariant {
					testFileFound = true
					break out
				}
			}
		}
		if !testFileFound {
			val, _ := srcFile.Absolute()
			filesWithoutTests = append(filesWithoutTests, val)
		}
	}
	return filesWithoutTests
}
