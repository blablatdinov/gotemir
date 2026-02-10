// SPDX-FileCopyrightText: Copyright (c) 2024-2026 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
// SPDX-License-Identifier: MIT

package logic

type FkComparedStructures struct {
	FilesWthtTests    []string
	TestsWthtSrcFiles []string
}

func FkComparedStructuresCtor(FilesWthtTests []string, TestsWthtSrcFiles []string) ComparedStructures {
	return FkComparedStructures{
		FilesWthtTests,
		TestsWthtSrcFiles,
	}
}

func (fkComparedStructures FkComparedStructures) FilesWithoutTests() ([]string, error) {
	return fkComparedStructures.FilesWthtTests, nil
}

func (fkComparedStructures FkComparedStructures) TestsWithoutSrcFiles() ([]string, error) {
	return fkComparedStructures.TestsWthtSrcFiles, nil
}
