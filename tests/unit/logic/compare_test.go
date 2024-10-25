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
	"path/filepath"
	"strings"
	"testing"

	gotemir "github.com/blablatdinov/gotemir/src/logic"
)

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
		got := gotemir.Compare(testCase.srcDir, testCase.testsDir)
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
		got := gotemir.Compare(testCase.srcDir, testCase.testsDir)
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

func TestTestFileVariants(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "File without dir",
			input:    "user.py",
			expected: []string{"user_test.py", "test_user.py"},
		},
		{
			name:     "File in dir",
			input:    "handlers/user.py",
			expected: []string{"handlers/user_test.py", "handlers/test_user.py"},
		},
		{
			name:     "File in nested dir",
			input:    "handlers/auth/user.py",
			expected: []string{"handlers/auth/user_test.py", "handlers/auth/test_user.py"},
		},
	}
	for testIdx, testCase := range cases {
		got := gotemir.TestFileVariants(testCase.input)
		if len(got) != len(testCase.expected) {
			t.Fatalf(
				"Case %d (%s): len of actual and expected not equal\nActual: %v\nExpected: %v\n",
				testIdx, testCase.name, got, testCase.expected,
			)
		}
		for idx, actualFile := range got {
			if testCase.expected[idx] != actualFile {
				t.Errorf(
					strings.Join(
						[]string{
							"Incompare actual and expected at index=%d (%s)",
							"test file variants: %v",
							"Actual: %s != Expected: %s",
							"\n",
						},
						"\n",
					),
					testIdx, testCase.name, got, got[idx], testCase.expected[idx],
				)
			}
		}
	}
}
