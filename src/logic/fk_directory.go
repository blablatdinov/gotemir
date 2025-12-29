// SPDX-FileCopyrightText: Copyright (c) 2024-2026 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
// SPDX-License-Identifier: MIT

package logic

type FkDirectory struct {
	origin []Path
}

func FkDirectoryCtor(origin []Path) Directory {
	return FkDirectory{
		origin,
	}
}

func (fkDirectory FkDirectory) Structure() ([]Path, error) {
	return fkDirectory.origin, nil
}
