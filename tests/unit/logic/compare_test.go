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
	"strings"
	"testing"

	gotemir "github.com/blablatdinov/gotemir/src/logic"
)

func TestCompare(t *testing.T) {
	t.Parallel()
	cases := []struct {
		srcDir   gotemir.Directory
		testsDir gotemir.Directory
	}{
		{
			srcDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{gotemir.FkPathCtor("src/logic.go", ".")},
			),
			testsDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{gotemir.FkPathCtor("tests/logic_test.go", ".")},
			),
		},
		{
			srcDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{gotemir.FkPathCtor("src/logic.py", ".")},
			),
			testsDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{gotemir.FkPathCtor("tests/logic_test.py", ".")},
			),
		},
		{
			srcDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{gotemir.FkPathCtor("src/logic.py", ".")},
			),
			testsDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{gotemir.FkPathCtor("tests/test_logic.py", ".")},
			),
		},
		{
			srcDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{
					gotemir.FkPathCtor("src/handlers/users.py", "."),
					gotemir.FkPathCtor("src/logic/auth.py", "."),
				},
			),
			testsDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{
					gotemir.FkPathCtor("tests/handlers/test_users.py", "."),
					gotemir.FkPathCtor("tests/logic/test_auth.py", "."),
				},
			),
		},
		// TODO: case when source code and test in one directory //nolint:godox
		// {
		// 	srcDir:   gotemir.FkDirectoryCtor([]string{"logic.go"}),
		// 	testsDir: gotemir.FkDirectoryCtor([]string{"logic_test.go"}),
		// },
	}
	for _, testCase := range cases {
		got := gotemir.Compare(testCase.srcDir, testCase.testsDir)
		if len(got) > 0 {
			t.Errorf(
				strings.Join(
					[]string{
						"Found files without tests",
						"src directory content: %v",
						"tests directory content: %v",
						"Actual: %v",
						"\n",
					},
					"\n",
				),
				testCase.srcDir,
				testCase.testsDir,
				got,
			)
		}
	}
}

func TestFileWithoutTest(t *testing.T) {
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
				[]gotemir.Path{gotemir.FkPathCtor("src/logic.go", ".")},
			),
			testsDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{gotemir.FkPathCtor("tests/unbounded_test.go", ".")},
			),
			expected: []string{"src/logic.go"},
		},
		{
			name: "Test not exist",
			srcDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{gotemir.FkPathCtor("src/logic.go", ".")},
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
					gotemir.FkPathCtor("src/handlers/users.py", "."),
					gotemir.FkPathCtor("src/logic/auth.py", "."),
				},
			),
			testsDir: gotemir.FkDirectoryCtor(
				[]gotemir.Path{
					gotemir.FkPathCtor("tests/handlers/test_users.py", "."),
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
