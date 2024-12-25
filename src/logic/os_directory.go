// The MIT License (MIT)
//
// Copyright (c) 2024-2025 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
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
	splittedPath := strings.Split(osDirectory.path, ",")
	for _, sPath := range splittedPath {
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
