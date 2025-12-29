// SPDX-FileCopyrightText: Copyright (c) 2024-2026 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
// SPDX-License-Identifier: MIT

package logic

import (
	"errors"
	"fmt"
	"path/filepath"
)

type FkPath struct {
	absolute string
	dir      string
}

func FkPathCtor(absolute, dir string) Path {
	return FkPath{
		absolute,
		dir,
	}
}

var errBuildRelative = errors.New("error build relative path")

func (fkPath FkPath) Relative() (string, error) {
	rel, err := filepath.Rel(fkPath.dir, fkPath.absolute)
	if err != nil {
		return "", fmt.Errorf("%w", errBuildRelative)
	}
	return rel, nil
}

func (fkPath FkPath) Absolute() (string, error) {
	return fkPath.absolute, nil
}
