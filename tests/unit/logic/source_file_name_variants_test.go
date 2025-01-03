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

package logic_test

import (
	"path/filepath"
	"sort"
	"strings"
	"testing"

	gotemir "github.com/blablatdinov/gotemir/src/logic"
)

func TestSourceFileVariants(t *testing.T) { //nolint:funlen //Many cases
	t.Parallel()
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Simple",
			input:    "test_user.py",
			expected: []string{"user.py"},
		},
		{
			name:     "Double test in file name",
			input:    "excluded_tests_test.go",
			expected: []string{"excluded_tests.go"},
		},
		{
			name:     "Double test in file name",
			input:    "test_excluded_tests.go",
			expected: []string{"excluded_tests.go", "test_excluded.go"},
		},
		{
			name:     "Double test in file name",
			input:    "excluded_tests_tests.go",
			expected: []string{"excluded_tests.go"},
		},
		{
			name:     "Double test in file name",
			input:    "test_excluded_test.go",
			expected: []string{"excluded_test.go", "test_excluded.go"},
		},
		{
			name:     "Double test in file name",
			input:    "test_test_test_excluded_test.go",
			expected: []string{"test_test_excluded_test.go", "test_test_test_excluded.go"},
		},
		{
			name:     "Pascal case",
			input:    "TestAbc.go",
			expected: []string{"Abc.go"},
		},
	}
	for testIdx, testCase := range cases {
		localizedInput, err := filepath.Localize(testCase.input)
		if err != nil {
			t.Fatalf("Err on localize path: %s. %s", testCase.input, err)
		}
		got := sort.StringSlice(
			gotemir.SourceFileNameVariantCtor(localizedInput).AsList(),
		)
		if len(got) != len(testCase.expected) {
			t.Fatalf(
				"Case %d (%s): len of actual and expected not equal\nActual: %v\nExpected: %v\n",
				testIdx, testCase.name, got, testCase.expected,
			)
		}
		for expIdx, expectedStr := range testCase.expected {
			localizedExpected, err := filepath.Localize(expectedStr)
			if err != nil {
				t.Fatalf("Err on localize path: %s. %s", expectedStr, err)
			}
			testCase.expected[expIdx] = localizedExpected
		}
		testCase.expected = sort.StringSlice(testCase.expected)
		for idx, actualFile := range got {
			localizedActual, err := filepath.Localize(actualFile)
			if err != nil {
				t.Fatalf("Err on localize path: %s. %s", actualFile, err)
			}
			localizedExpected := testCase.expected[idx]
			if localizedExpected != localizedActual {
				t.Errorf(
					strings.Join(
						[]string{
							"Incompare actual and expected at index=%d (%s)",
							"input: %s",
							"test file variants: %v",
							"Actual: '%s' != Expected: %s",
							"\n",
						},
						"\n",
					),
					testIdx+1, testCase.name, localizedInput, got, localizedActual, localizedExpected,
				)
			}
		}
	}
}
