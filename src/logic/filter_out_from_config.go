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
	"fmt"
	"regexp"
)

type FilterOutFromConfig struct {
	origin ComparedStructures
	config Config
}

func FilterOutFromConfifCtor(cmprd ComparedStructures, config Config) ComparedStructures {
	return FilterOutFromConfig{cmprd, config}
}

func (filterOutFromConfig FilterOutFromConfig) FilesWithoutTests() ([]string, error) {
	originFiles, err := filterOutFromConfig.origin.FilesWithoutTests()
	if err != nil {
		return []string{}, err //nolint:wrapcheck
	}
	result := make([]string, 0)
	for _, originFile := range originFiles {
		originAbsolute := originFile
		patternFound := false
		for _, pattern := range filterOutFromConfig.config.GotemirConfig.TestFreeFiles {
			regexPattern, err := regexp.Compile(pattern)
			if err != nil {
				return []string{}, fmt.Errorf("fail parsing regex \"%s\" in .gotemir.yaml:\n  %w", pattern, err)
			}
			regexFoundString := regexPattern.FindString(originAbsolute)
			if len(regexFoundString) == len(originAbsolute) {
				patternFound = true
			}
		}
		if !patternFound {
			result = append(result, originFile)
		}
	}
	return result, nil
}

func (filterOutFromConfig FilterOutFromConfig) TestsWithoutSrcFiles() ([]string, error) {
	originFiles, err := filterOutFromConfig.origin.TestsWithoutSrcFiles()
	if err != nil {
		return []string{}, err //nolint:wrapcheck
	}
	result := make([]string, 0)
	for _, originFile := range originFiles {
		originAbsolute := originFile
		patternFound := false
		for _, pattern := range filterOutFromConfig.config.GotemirConfig.TestHelpers {
			regexPattern, err := regexp.Compile(pattern)
			if err != nil {
				return []string{}, fmt.Errorf("fail parsing regex \"%s\" in .gotemir.yaml:\n  %w", pattern, err)
			}
			regexFoundString := regexPattern.FindString(originAbsolute)
			if len(regexFoundString) == len(originAbsolute) {
				patternFound = true
			}
		}
		if !patternFound {
			result = append(result, originFile)
		}
	}
	return result, nil
}
