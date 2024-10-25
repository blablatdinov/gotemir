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
	"testing"

	gotemir "github.com/blablatdinov/gotemir/src/logic"
)

func TestFkPath(t *testing.T) {
	t.Parallel()
	cases := []struct {
		file             string
		expectedRelative string
		dir              string
	}{
		{
			"src/tests/test_auth.py",
			"test_auth.py",
			"src/tests",
		},
		{
			"src/tests/test_auth.py",
			"test_auth.py",
			"src/tests/",
		},
		{
			"src/tests/test_auth.py",
			"src/tests/test_auth.py",
			".",
		},
	}
	for idx, testCase := range cases {
		fkPath := gotemir.FkPathCtor(testCase.file, testCase.dir)
		if val, _ := fkPath.Absolute(); val != testCase.file {
			t.Errorf(
				"Case #%d fail\nAbsolute not valid\nActual: %s != Expected %s",
				idx+1,
				val,
				testCase.file,
			)
		}
		localizedExpected, _ := filepath.Localize(testCase.expectedRelative)
		if val, _ := fkPath.Relative(); val != localizedExpected {
			t.Errorf(
				"Case #%d fail\nRelative not valid\nActual: %s != Expected %s",
				idx+1,
				val,
				localizedExpected,
			)
		}
	}
}
