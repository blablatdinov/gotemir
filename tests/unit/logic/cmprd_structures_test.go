// SPDX-FileCopyrightText: Copyright (c) 2024-2026 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
// SPDX-License-Identifier: MIT

package logic_test

import (
	"errors"
	"path/filepath"
	"strings"
	"testing"

	gotemir "github.com/blablatdinov/gotemir/src/logic"
)

type SeDirectory struct{}

func (seDirectory SeDirectory) Structure() ([]gotemir.Path, error) {
	return []gotemir.Path{}, errors.New("fk error")
}

type SePath struct{}

func (sePath SePath) Relative() (string, error) {
	return "", errors.New("fk error")
}

func (sePath SePath) Absolute() (string, error) {
	return "", errors.New("fk error")
}

type SePathBrokenAbsolute struct {
	relative string
}

func (sePath SePathBrokenAbsolute) Relative() (string, error) {
	return sePath.relative, nil
}

func (sePath SePathBrokenAbsolute) Absolute() (string, error) {
	return "", errors.New("fk error")
}

func TestCompare(t *testing.T) { //nolint:funlen // Many cases
	t.Parallel()
	cases := []struct {
		name     string
		srcDir   gotemir.Directory
		testsDir gotemir.Directory
	}{
		{
			name: "One file case",
			srcDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{gotemir.FkPathCtor(
					filepath.Join("src", "logic.go"),
					"src",
				)},
			),
			testsDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{gotemir.FkPathCtor(
					filepath.Join("tests", "logic_test.go"),
					"tests",
				)},
			),
		},
		{
			name: "Case with test_ prefix",
			srcDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{gotemir.FkPathCtor(
					filepath.Join("src", "logic.py"),
					"src",
				)},
			),
			testsDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{gotemir.FkPathCtor(
					filepath.Join("tests", "test_logic.py"),
					"tests",
				)},
			),
		},
		{
			name: "Case nested deirectories",
			srcDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{
					gotemir.FkPathCtor(
						filepath.Join("src", "handlers", "users.py"),
						"src",
					),
					gotemir.FkPathCtor(
						filepath.Join("src", "logic", "auth.py"),
						"src",
					),
				},
			),
			testsDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{
					gotemir.FkPathCtor(
						filepath.Join("tests", "handlers", "test_users.py"),
						"tests",
					),
					gotemir.FkPathCtor(
						filepath.Join("tests", "logic", "test_auth.py"),
						"tests",
					),
				},
			),
		},
		{
			name:     "Case file and test in one directory",
			srcDir:   gotemir.FkDirectoryCtor([]gotemir.Path{gotemir.FkPathCtor("logic.go", ".")}),
			testsDir: gotemir.FkDirectoryCtor([]gotemir.Path{gotemir.FkPathCtor("logic_test.go", ".")}),
		},
	}
	for _, testCase := range cases {
		got, _ := gotemir.CmprdStructuresCtor(testCase.srcDir, testCase.testsDir).FilesWithoutTests()
		if len(got) > 0 {
			t.Errorf(
				strings.Join(
					[]string{
						"Found files without tests. Case: '%s'",
						"src directory content: %v",
						"tests directory content: %v",
						"Actual: %v",
						"\n",
					},
					"\n",
				),
				testCase.name,
				testCase.srcDir,
				testCase.testsDir,
				got,
			)
		}
	}
}

func TestFileWithoutTest(t *testing.T) { //nolint:funlen // Many cases
	t.Parallel()
	cases := []struct {
		name     string
		srcDir   gotemir.Directory
		testsDir gotemir.Directory
		expected []string
	}{
		{
			name: "Unbounded test",
			srcDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{gotemir.FkPathCtor(
					filepath.Join("src", "logic.go"),
					"src",
				)},
			),
			testsDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{gotemir.FkPathCtor(
					filepath.Join("tests", "unbounded_test.go"),
					"tests",
				)},
			),
			expected: []string{
				filepath.Join("src", "logic.go"),
			},
		},
		{
			name: "Test not exist",
			srcDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{gotemir.FkPathCtor("src/logic.go", "src")},
			),
			testsDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{},
			),
			expected: []string{"src/logic.go"},
		},
		{
			name: "One file without test",
			srcDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{
					gotemir.FkPathCtor("src/handlers/users.py", "src"),
					gotemir.FkPathCtor("src/logic/auth.py", "src"),
				},
			),
			testsDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{
					gotemir.FkPathCtor("tests/handlers/test_users.py", "tests"),
				},
			),
			expected: []string{"src/logic/auth.py"},
		},
	}
	for idx, testCase := range cases {
		got, _ := gotemir.CmprdStructuresCtor(testCase.srcDir, testCase.testsDir).FilesWithoutTests()
		if len(got) != len(testCase.expected) {
			t.Fatalf(
				"Case %d (%s): len of actual and expected not equal\nActual: %v\nExpected: %v\n",
				idx+1, testCase.name, got, testCase.expected,
			)
		}
		for idx, actualFile := range got {
			if testCase.expected[idx] != actualFile {
				t.Errorf(
					strings.Join(
						[]string{
							"Incompare actual and expected at index=%d (%s)",
							"src directory content: %v",
							"tests directory content: %v",
							"Actual: %s != Expected: %s",
							"\n",
						},
						"\n",
					),
					idx+1, testCase.name, testCase.srcDir, testCase.testsDir, actualFile, testCase.expected[idx],
				)
			}
		}
	}
}

func TestErrorHandling(t *testing.T) {
	t.Parallel()
	_, err := gotemir.CmprdStructuresCtor(
		gotemir.FkDirectoryCtor([]gotemir.Path{gotemir.FkPathCtor("", "")}),
		SeDirectory{},
	).FilesWithoutTests()
	if err == nil {
		t.Errorf("Error not handled")
	}
	_, err = gotemir.CmprdStructuresCtor(
		SeDirectory{},
		gotemir.FkDirectoryCtor([]gotemir.Path{gotemir.FkPathCtor("", "")}),
	).FilesWithoutTests()
	if err == nil {
		t.Errorf("Error not handled")
	}
	_, err = gotemir.CmprdStructuresCtor(
		gotemir.FkDirectoryCtor([]gotemir.Path{SePath{}}),
		gotemir.FkDirectoryCtor([]gotemir.Path{gotemir.FkPathCtor("", "")}),
	).FilesWithoutTests()
	if err == nil {
		t.Errorf("Error not handled")
	}
	_, err = gotemir.CmprdStructuresCtor(
		gotemir.FkDirectoryCtor([]gotemir.Path{SePathBrokenAbsolute{"file.go"}}),
		gotemir.FkDirectoryCtor([]gotemir.Path{gotemir.FkPathCtor("file_test.go", "src/file_test.go")}),
	).FilesWithoutTests()
	if err == nil {
		t.Errorf("Error not handled")
	}
	_, err = gotemir.CmprdStructuresCtor(
		gotemir.FkDirectoryCtor([]gotemir.Path{gotemir.FkPathCtor("", "")}),
		SeDirectory{},
	).TestsWithoutSrcFiles()
	if err == nil {
		t.Errorf("Error not handled")
	}
	_, err = gotemir.CmprdStructuresCtor(
		SeDirectory{},
		gotemir.FkDirectoryCtor([]gotemir.Path{gotemir.FkPathCtor("", "")}),
	).TestsWithoutSrcFiles()
	if err == nil {
		t.Errorf("Error not handled")
	}
}
