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
	"strings"
	"testing"

	gotemir "github.com/blablatdinov/gotemir/src/logic"
)

func prepareFiles(t *testing.T) string {
	tempDir, err := os.MkdirTemp(os.TempDir(), "gotemir_test")
	if err != nil {
		t.Fatalf("Fail on create temp dir: %s", err)
	}
	srcDirPath := tempDir + "/src"
	handlersDirPath := srcDirPath + "/handlers"
	entryPath := srcDirPath + "/entry.py"
	filePath := handlersDirPath + "/file.py"
	os.Mkdir(srcDirPath, 0770)
	os.Mkdir(handlersDirPath, 0770)
	os.WriteFile(entryPath, []byte(""), 0660)
	os.WriteFile(filePath, []byte(""), 0660)
	return tempDir
}

func TestOsDirectory(t *testing.T) {
	t.Parallel()
	tempDir := prepareFiles(t)
	osDir := gotemir.OsDirectoryCtor(tempDir)
	expected := []string{"src/entry.py", "src/handlers/file.py"}

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
			expected,
		)
	}
	for idx, actual := range got {
		if actual != expected[idx] {
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
				expected,
			)
		}
	}
}
