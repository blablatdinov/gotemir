// SPDX-FileCopyrightText: Copyright (c) 2024-2026 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
// SPDX-License-Identifier: MIT

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
	result := make([]string, 0, 2*len(appendixes))
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
