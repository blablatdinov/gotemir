// The MIT License (MIT)
//
// Copyright (c) 2024-2025 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
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
	"path/filepath"
	"strings"
)

type TestFileNameVariants struct {
	path string
}

func TestFileNameVariantsCtor(path string) FileNameVariants {
	return TestFileNameVariants{path}
}

func (testFileNameVariant TestFileNameVariants) AsList() []string {
	dir, file := filepath.Split(testFileNameVariant.path)
	fileExtension := "." + strings.Split(file, ".")[1]
	appendixes := []string{
		// Snake case
		"test_",
		"_test",
		"_tests",
		"tests_",
		// Pascal case
		"Test",
		"Tests",
		// Camel case
		"test",
		"Test",
		"tests",
		"Tests",
	}
	fileNameWithoutExtension := strings.Split(file, ".")[0]
	result := make([]string, 0)
	for _, appendix := range appendixes {
		result = append(
			result,
			filepath.Join(
				dir,
				fileNameWithoutExtension+appendix+fileExtension,
			),
			filepath.Join(
				dir,
				appendix+fileNameWithoutExtension+fileExtension,
			),
		)
	}
	return result
}
