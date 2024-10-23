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
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	gotemir "github.com/blablatdinov/gotemir/src/logic"
)

func prepareFiles(t *testing.T, paths []string) string {
	t.Helper()
	tempDir, err := os.MkdirTemp(os.TempDir(), "gotemir_test")
	if err != nil {
		t.Fatalf("Fail on create temp dir: %s", err)
	}
	for _, path := range paths {
		dir := filepath.Dir(tempDir + "/" + path)
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			t.Fatalf("failed to create directory %s: %s\n", dir, err)
		}
		file, err := os.Create(tempDir + "/" + path)
		if err != nil {
			t.Fatalf("failed to create file %s: %s\n", tempDir+"/"+path, err)
		}
		err = file.Close()
		if err != nil {
			t.Fatalf("failed to close file %s: %s\n", tempDir+"/"+path, err)
		}
	}
	return tempDir
}

func TestOsDirectory(t *testing.T) {
	t.Parallel()
	tempDir := prepareFiles(t, []string{
		"src/handlers/file.py",
		"src/entry.py",
		"src/README.md",
	})
	osDir := gotemir.OsDirectoryCtor(tempDir+"/src", ".py")
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
		actualVal, _ := actual.Relative()
		if actualVal != localizedExpected[idx] {
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
				got[idx],
				localizedExpected[idx],
			)
		}
	}
}

func TestOsDirectorySeparated(t *testing.T) {
	t.Parallel()
	tempDir := prepareFiles(t, []string{
		"tests/it/test_file.py",
		"tests/unit/test_auth.py",
	})
	osDir := gotemir.OsDirectoryCtor(
		fmt.Sprintf("%s/tests/it/,%s/tests/unit/", tempDir, tempDir),
		".py",
	)
	// TODO: fix this case
	// osDir := gotemir.OsDirectoryCtor(
	// 	fmt.Sprintf("%s/tests/it,%s/tests/unit", tempDir, tempDir),
	// 	".py",
	// )
	// expected := []string{
	localizedExpected := []string{
		tempDir + "/tests/it/test_file.py",
		tempDir + "/tests/unit/test_auth.py",
	}
	relativeExpected := []string{
		"test_file.py",
		"test_auth.py",
	}
	// localizedExpected := make([]string, len(expected))
	// for idx, expectedFile := range expected {
	// 	localized, err := filepath.Localize(expectedFile)
	// 	if err != nil {
	// 		t.Fatalf("Fail on localized path: %s, err: %s", expectedFile, err)
	// 	}
	// 	localizedExpected[idx] = localized
	// }

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
		actualAbsolute, _ := actual.Absolute()
		if actualAbsolute != localizedExpected[idx] {
			t.Errorf(
				strings.Join(
					[]string{
						"Incompare actual absolute and expected at index=%d",
						"Actual: %s != Expected: %s",
						"\n",
					},
					"\n",
				),
				idx,
				actualAbsolute,
				localizedExpected[idx],
			)
		}
		actualRelative, _ := actual.Relative()
		if actualRelative != relativeExpected[idx] {
			t.Errorf(
				strings.Join(
					[]string{
						"Incompare actual relative and expected at index=%d",
						"Actual: %s != Expected: %s",
						"\n",
					},
					"\n",
				),
				idx,
				actualRelative,
				relativeExpected[idx],
			)
		}
	}
}
