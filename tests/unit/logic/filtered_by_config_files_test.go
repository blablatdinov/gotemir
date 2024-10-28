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

func TestFilteredbyConfigFiles(t *testing.T) {
	cases := []struct {
		paths    []gotemir.Path
		patterns []string
		expected []string
	}{
		// {
		// 	paths: []gotemir.Path{
		// 		gotemir.FkPathCtor("src/__init__.py", "src"),
		// 	},
		// 	patterns: []string{"__init__.py"},
		// 	expected: []string{"src/__init__.py"},
		// },
		{
			paths: []gotemir.Path{
				gotemir.FkPathCtor("src/__init__.py", "src"),
			},
			patterns: []string{".*__init__.py"},
			expected: []string{},
		},
	}
	for caseIdx, testCase := range cases {
		got, _ := gotemir.FilteredByConfigFilesCtor(
			gotemir.FkDirectoryCtor(testCase.paths),
			testCase.patterns,
		).Structure()
		if len(got) != len(testCase.expected) {
			t.Fatalf(
				"Case %d: len of actual and expected not equal\nActual: %v\nExpected: %v\n",
				caseIdx+1, got, testCase.expected,
			)
		}
		for idx, actual := range got {
			actualAbs, _ := actual.Absolute()
			if actualAbs != testCase.expected[idx] {
				t.Errorf(
					strings.Join(
						[]string{
							"Case %d",
							"Incompare actual and expected at index=%d",
							"Actual: %v != Expected: %v",
							"\n",
						},
						"\n",
					),
					caseIdx, idx, got, testCase.expected[idx],
				)

			}
		}
	}
}
