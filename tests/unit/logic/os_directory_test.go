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

package logic_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	gotemir "github.com/blablatdinov/gotemir/src/logic"
)

func prepareFiles(t *testing.T) string {
	t.Helper()
	tempDir, err := os.MkdirTemp(os.TempDir(), "gotemir_test")
	if err != nil {
		t.Fatalf("Fail on create temp dir: %s", err)
	}
	srcDirPath := tempDir + "/src"
	handlersDirPath := srcDirPath + "/handlers"
	entryPath := srcDirPath + "/entry.py"
	filePath := handlersDirPath + "/file.py"
	if err = os.Mkdir(srcDirPath, 0o770); err != nil {
		t.Fatalf("Fail on create dir %s: %s", srcDirPath, err)
	}
	if err = os.Mkdir(handlersDirPath, 0o770); err != nil {
		t.Fatalf("Fail on create dir %s: %s", handlersDirPath, err)
	}
	if err = os.WriteFile(entryPath, []byte(""), 0o600); err != nil {
		t.Fatalf("Fail on create file %s: %s", entryPath, err)
	}
	if err = os.WriteFile(filePath, []byte(""), 0o600); err != nil {
		t.Fatalf("Fail on create file %s: %s", filePath, err)
	}
	return tempDir
}

func TestOsDirectory(t *testing.T) {
	t.Parallel()
	tempDir := prepareFiles(t)
	osDir := gotemir.OsDirectoryCtor(tempDir)
	expected := []string{"src/entry.py", "src/handlers/file.py"}
	localizedExpected := make([]string, len(expected))
	for idx, expectedFile := range expected {
		localized, err := filepath.Localize(expectedFile)
		if err != nil {
			t.Fatalf("Fail on localized path: %s", expectedFile)
		}
		localizedExpected[idx] = localized
	}

	got, err := osDir.Structure()
	if err != nil {
		t.Fatalf("Fail on parse dir: %s", err)
	}
	if len(got) != 2 {
		t.Errorf(
			strings.Join(
				[]string{
					"Actual must contains 2 elements",
					"Actual: %v",
					"Expected: %v",
					"\n",
				},
				"\n",
			),
			got,
			localizedExpected,
		)
	}
	for idx, actual := range got {
		if actual != localizedExpected[idx] {
			t.Errorf(
				strings.Join(
					[]string{
						"Incompare actual and expected at index=%d",
						"Actual: %v != Expected: %v",
						"\n",
					},
					"\n",
				),
				idx,
				got,
				localizedExpected,
			)
		}
	}
}
