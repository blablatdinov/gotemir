// SPDX-FileCopyrightText: Copyright (c) 2024-2026 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
// SPDX-License-Identifier: MIT

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
	result, err := generic(originFiles, filterOutFromConfig.config.GotemirConfig.TestFreeFiles)
	if err != nil {
		return []string{}, err
	}
	return result, nil
}

func (filterOutFromConfig FilterOutFromConfig) TestsWithoutSrcFiles() ([]string, error) {
	originFiles, err := filterOutFromConfig.origin.TestsWithoutSrcFiles()
	if err != nil {
		return []string{}, err //nolint:wrapcheck
	}
	result, err := generic(originFiles, filterOutFromConfig.config.GotemirConfig.TestHelpers)
	if err != nil {
		return []string{}, err
	}
	return result, nil
}

func generic(files []string, exclude []string) ([]string, error) {
	result := make([]string, 0)
	for _, originFile := range files {
		originAbsolute := originFile
		patternFound := false
		for _, pattern := range exclude {
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
