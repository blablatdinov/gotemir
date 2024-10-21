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
			srcDir:   gotemir.FkDirectoryCtor([]string{"logic.go"}),
			testsDir: gotemir.FkDirectoryCtor([]string{"logic_test.go"}),
		},
		{
			srcDir:   gotemir.FkDirectoryCtor([]string{"logic.py"}),
			testsDir: gotemir.FkDirectoryCtor([]string{"logic_test.py"}),
		},
		{
			srcDir:   gotemir.FkDirectoryCtor([]string{"logic.py"}),
			testsDir: gotemir.FkDirectoryCtor([]string{"test_logic.py"}),
		},
		{
			srcDir: gotemir.FkDirectoryCtor([]string{
				"handlers/users.py",
				"logic/auth.py",
			}),
			testsDir: gotemir.FkDirectoryCtor([]string{
				"handlers/test_users.py",
				"logic/test_auth.py",
			}),
		},
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
		srcDir   gotemir.Directory
		testsDir gotemir.Directory
		expected []string
	}{
		{
			srcDir:   gotemir.FkDirectoryCtor([]string{"logic.go"}),
			testsDir: gotemir.FkDirectoryCtor([]string{"unbounded_test.go"}),
			expected: []string{"logic.go"},
		},
		{
			srcDir:   gotemir.FkDirectoryCtor([]string{"logic.py"}),
			testsDir: gotemir.FkDirectoryCtor([]string{}),
			expected: []string{"logic.py"},
		},
		{
			srcDir: gotemir.FkDirectoryCtor([]string{
				"handlers/users.py",
				"logic/auth.py",
			}),
			testsDir: gotemir.FkDirectoryCtor([]string{
				"handlers/test_users.py",
			}),
			expected: []string{"logic/auth.py"},
		},
	}
	for _, testCase := range cases {
		got := gotemir.Compare(testCase.srcDir, testCase.testsDir)
		if len(got) != len(testCase.expected) {
			t.Fatalf(
				"Len of actual and expected not equal\nActual: %v\nExpected: %v\n",
				got,
				testCase.expected,
			)
		}
		for idx, actualFile := range got {
			if testCase.expected[idx] != actualFile {
				t.Errorf("Expected: %s, actual: %s\n", testCase.expected, actualFile)
				t.Errorf(
					strings.Join(
						[]string{
							"Incompare actual and expected at index=%d",
							"src directory content: %v",
							"tests directory content: %v",
							"Actual: %v != Expected: %v",
							"\n",
						},
						"\n",
					),
					idx,
					testCase.srcDir,
					testCase.testsDir,
					actualFile,
					testCase.expected[idx],
				)
			}
		}
	}
}
