// SPDX-FileCopyrightText: Copyright (c) 2024-2026 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
// SPDX-License-Identifier: MIT

package logic

type FkComparedStructures struct {
	FilesWithTests    []string
	TestsWithSrcFiles []string
}

func FkComparedStructuresCtor(filesWithTests []string, testsWithSrcFiles []string) ComparedStructures {
	return FkComparedStructures{
		filesWithTests,
		testsWithSrcFiles,
	}
}

func (fkComparedStructures FkComparedStructures) FilesWithoutTests() ([]string, error) {
	return fkComparedStructures.FilesWithTests, nil
}

func (fkComparedStructures FkComparedStructures) TestsWithoutSrcFiles() ([]string, error) {
	return fkComparedStructures.TestsWithSrcFiles, nil
}
