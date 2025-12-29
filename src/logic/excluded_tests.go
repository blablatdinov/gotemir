// SPDX-FileCopyrightText: Copyright (c) 2024-2026 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
// SPDX-License-Identifier: MIT

package logic

import (
	"errors"
	"fmt"
)

type ExcludedTestsDirectory struct {
	origin   Directory
	testsDir Directory
}

func ExcludedTestsDirectoryCtor(srcDir, testsPath Directory) Directory {
	return ExcludedTestsDirectory{
		srcDir,
		testsPath,
	}
}

var errWalkingExcludedTestsDirectory = errors.New("fail on walk directory")

func (excludedTestsDirectory ExcludedTestsDirectory) Structure() ([]Path, error) {
	origin, err := excludedTestsDirectory.origin.Structure()
	if err != nil {
		return nil, fmt.Errorf("%w", errWalkingExcludedTestsDirectory)
	}
	updated := make([]Path, 0)
	testsPaths, _ := excludedTestsDirectory.testsDir.Structure()
	for _, srcPath := range origin {
		srcPathStr, _ := srcPath.Absolute()
		testFound := false
		for _, testPath := range testsPaths {
			testPathAbsolute, _ := testPath.Absolute()
			if srcPathStr == testPathAbsolute {
				testFound = true
			}
		}
		if !testFound {
			updated = append(updated, srcPath)
		}
	}
	return updated, nil
}
