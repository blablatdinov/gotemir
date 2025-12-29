// SPDX-FileCopyrightText: Copyright (c) 2024-2026 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
// SPDX-License-Identifier: MIT

package logic

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type OsDirectory struct {
	path      string
	extension string
}

func OsDirectoryCtor(path string, extension string) Directory {
	return OsDirectory{
		path,
		extension,
	}
}

var errWalking = errors.New("fail on walk directory")

func (osDirectory OsDirectory) Structure() ([]Path, error) {
	var files []Path
	splittedPath := strings.SplitSeq(osDirectory.path, ",")
	for sPath := range splittedPath {
		err := filepath.Walk(sPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("%w. Dirname=%s error: %w", errWalking, path, err)
			}
			if !info.IsDir() && strings.HasSuffix(path, osDirectory.extension) {
				fkPath := FkPathCtor(path, sPath)
				files = append(files, fkPath)
			}
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("%w. Dirname=%s error: %w", errWalking, osDirectory.path, err)
		}
	}
	return files, nil
}
