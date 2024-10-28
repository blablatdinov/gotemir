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
	"errors"
	"fmt"
	"regexp"
)

type FilteredByConfigFiles struct {
	origin Directory
	config Config
}

func FilteredByConfigFilesCtor(origin Directory, config Config) Directory {
	return FilteredByConfigFiles{origin, config}
}

var errFiltering = errors.New("error on filtering")

func (filteredByConfigFiles FilteredByConfigFiles) Structure() ([]Path, error) {
	originFiles, err := filteredByConfigFiles.origin.Structure()
	if err != nil {
		return nil, fmt.Errorf("%w %w", errFiltering, err)
	}
	result := make([]Path, 0)
	for _, originFile := range originFiles {
		originAbsolute, err := originFile.Absolute()
		if err != nil {
			return nil, fmt.Errorf("%w %w", errFiltering, err)
		}
		patternFound := false
		for _, pattern := range filteredByConfigFiles.config.TestFreeFiles {
			patternFound, err = regexp.MatchString(pattern, originAbsolute)
			if err != nil {
				return nil, fmt.Errorf(
					"%w. Fail regexp.Match, pattern: %s. err: %w",
					errFiltering, pattern, err,
				)
			}
		}
		if !patternFound {
			result = append(result, originFile)
		}
	}
	return result, nil
}
