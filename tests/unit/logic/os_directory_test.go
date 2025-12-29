// SPDX-FileCopyrightText: Copyright (c) 2024-2026 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
// SPDX-License-Identifier: MIT

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
	tempDir := t.TempDir()
	for _, path := range paths {
		localizedPath, _ := filepath.Localize(path)
		joinedPath := filepath.Join(tempDir, localizedPath)
		dir := filepath.Dir(joinedPath)
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			t.Fatalf("failed to create directory '%s': %s\n", dir, err)
		}
		file, err := os.Create(joinedPath)
		if err != nil {
			t.Fatalf("failed to create file %s: %s\n", joinedPath, err)
		}
		err = file.Close()
		if err != nil {
			t.Fatalf("failed to close file %s: %s\n", joinedPath, err)
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
	osDir := gotemir.OsDirectoryCtor(filepath.Join(tempDir, "src"), ".py")
	expected := []string{"entry.py", "handlers/file.py"}
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
				actualVal,
				localizedExpected[idx],
			)
		}
	}
}

func TestOsDirectorySeparated(t *testing.T) { //nolint:funlen //TODO
	t.Parallel()
	tempDir := prepareFiles(t, []string{
		"tests/it/test_file.py",
		"tests/unit/test_auth.py",
	})
	osDir := gotemir.OsDirectoryCtor(
		fmt.Sprintf(
			"%s,%s",
			filepath.Join(tempDir, "tests", "it"),
			filepath.Join(tempDir, "tests", "unit"),
		),
		".py",
	)
	localizedExpected := []string{
		filepath.Join(tempDir, "tests", "it", "test_file.py"),
		filepath.Join(tempDir, "tests", "unit", "test_auth.py"),
	}
	relativeExpected := []string{
		"test_file.py",
		"test_auth.py",
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
