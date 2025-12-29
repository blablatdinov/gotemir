// SPDX-FileCopyrightText: Copyright (c) 2024-2026 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
// SPDX-License-Identifier: MIT

package logic

import "slices"

type CmprdStructures struct {
	srcDir   Directory
	testsDir Directory
}

func CmprdStructuresCtor(srcDir, testsDir Directory) ComparedStructures {
	return CmprdStructures{
		srcDir,
		testsDir,
	}
}

func (cmprdStructures CmprdStructures) FilesWithoutTests() ([]string, error) {
	filesWithoutTests := make([]string, 0)
	testFiles, _ := cmprdStructures.testsDir.Structure()
	srcFiles, _ := cmprdStructures.srcDir.Structure()
	for _, srcFile := range srcFiles {
		relativePath, _ := srcFile.Relative()
		testFileVariants := TestFileNameVariantsCtor(relativePath).AsList()
		testFileFound := false
	out:
		for _, testFile := range testFiles {
			relativePath, _ := testFile.Relative()
			for _, testFileVariant := range testFileVariants {
				if relativePath == testFileVariant {
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
	return filesWithoutTests, nil
}

func (cmprdStructures CmprdStructures) TestsWithoutSrcFiles() ([]string, error) {
	testsWithoutSrcFiles := make([]string, 0)
	testFiles, _ := cmprdStructures.testsDir.Structure()
	srcFiles, _ := cmprdStructures.srcDir.Structure()
	for _, testFile := range testFiles {
		relativeTestPath, _ := testFile.Relative()
		srcFileVariants := SourceFileNameVariantCtor(relativeTestPath).AsList()
		srcFileFound := false
		for _, srcFile := range srcFiles {
			relativeSrcPath, _ := srcFile.Relative()
			if slices.Contains(srcFileVariants, relativeSrcPath) {
				srcFileFound = true
			}
		}
		if !srcFileFound {
			val, _ := testFile.Absolute()
			testsWithoutSrcFiles = append(testsWithoutSrcFiles, val)
		}
	}
	return testsWithoutSrcFiles, nil
}
