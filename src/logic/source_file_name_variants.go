// SPDX-FileCopyrightText: Copyright (c) 2024-2026 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
// SPDX-License-Identifier: MIT

package logic

import (
	"path/filepath"
	"slices"
	"strings"
)

type SourceFileNameVariant struct {
	path string
}

func SourceFileNameVariantCtor(path string) FileNameVariants {
	return SourceFileNameVariant{path}
}

func (srcFileNameVariant SourceFileNameVariant) AsList() []string {
	testMarkers := []string{
		"test_",
		"_test",
		"tests_",
		"_tests",
		"Tests",
		"Test",
	}
	dir, file := filepath.Split(srcFileNameVariant.path)
	result := make([]string, 0)
	fileNameWithoutExtension := strings.Split(file, ".")[0]
	fileExt := "." + strings.Split(file, ".")[1]
	for _, marker := range testMarkers {
		markerBegin := fileNameWithoutExtension[0:len(marker)] == marker
		fnLen := len(fileNameWithoutExtension)
		markerEnd := fileNameWithoutExtension[fnLen-len(marker):fnLen] == marker
		variant := ""
		if markerBegin {
			variant = fileNameWithoutExtension[len(marker):fnLen] + fileExt
		} else if markerEnd {
			variant = fileNameWithoutExtension[0:fnLen-len(marker)] + fileExt
		}
		if variant != "" && !slices.Contains(result, variant) {
			result = append(
				result,
				filepath.Join(dir, variant),
			)
		}
	}
	return result
}
