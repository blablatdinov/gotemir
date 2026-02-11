// SPDX-FileCopyrightText: Copyright (c) 2024-2026 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
// SPDX-License-Identifier: MIT

package logic

import (
	"fmt"
	"regexp"
)

type WarnOnUnusedConfig struct {
	origin ComparedStructures
	config Config
	srcDir Directory
}

// UnusedPattern описывает паттерн из конфига, который совпадает с файлами, для которых уже есть тесты.
type UnusedPattern struct {
	Pattern      string
	MatchedPaths []string
}

func WarnOnUnusedConfigCtor(origin ComparedStructures, config Config, srcDir Directory) ComparedStructures {
	return WarnOnUnusedConfig{origin, config, srcDir}
}

func (warnOnUnusedConfig WarnOnUnusedConfig) FilesWithoutTests() ([]string, error) {
	return warnOnUnusedConfig.origin.FilesWithoutTests()
}

func (warnOnUnusedConfig WarnOnUnusedConfig) TestsWithoutSrcFiles() ([]string, error) {
	return warnOnUnusedConfig.origin.TestsWithoutSrcFiles()
}

// UnusedTestFreeFilesPatterns возвращает паттерны test-free-files, которые совпадают с файлами, для которых уже есть тесты.
func (warnOnUnusedConfig WarnOnUnusedConfig) UnusedTestFreeFilesPatterns() ([]UnusedPattern, error) {
	srcFiles, err := warnOnUnusedConfig.srcDir.Structure()
	if err != nil {
		return nil, fmt.Errorf("src dir structure: %w", err)
	}
	filesWithoutTests, err := warnOnUnusedConfig.origin.FilesWithoutTests()
	if err != nil {
		return nil, err
	}
	withoutTestsSet := make(map[string]struct{}, len(filesWithoutTests))
	for _, p := range filesWithoutTests {
		withoutTestsSet[p] = struct{}{}
	}
	var filesWithTests []string
	for _, f := range srcFiles {
		abs, err := f.Absolute()
		if err != nil {
			continue
		}
		if _, without := withoutTestsSet[abs]; !without {
			filesWithTests = append(filesWithTests, abs)
		}
	}
	var result []UnusedPattern
	for _, pattern := range warnOnUnusedConfig.config.GotemirConfig.TestFreeFiles {
		re, err := regexp.Compile(pattern)
		if err != nil {
			return nil, fmt.Errorf("fail parsing regex %q in .gotemir.yaml: %w", pattern, err)
		}
		var matched []string
		for _, path := range filesWithTests {
			found := re.FindString(path)
			if len(found) == len(path) {
				matched = append(matched, path)
			}
		}
		if len(matched) > 0 {
			result = append(result, UnusedPattern{Pattern: pattern, MatchedPaths: matched})
		}
	}
	return result, nil
}
