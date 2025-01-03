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
