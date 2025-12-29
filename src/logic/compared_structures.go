// SPDX-FileCopyrightText: Copyright (c) 2024-2026 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
// SPDX-License-Identifier: MIT

package logic

type ComparedStructures interface {
	FilesWithoutTests() ([]string, error)
	TestsWithoutSrcFiles() ([]string, error)
}
