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
	"strings"
	"testing"

	gotemir "github.com/blablatdinov/gotemir/src/logic"
)

func TestWithoutTests(t *testing.T) { //nolint:funlen // TODO
	t.Skip() // TODO
	t.Parallel()
	withoutTests := gotemir.ExcludedTestsDirectoryCtor(
		gotemir.FkDirectoryCtor(
			[]gotemir.Path{
				gotemir.FkPathCtor("src/entry.py", "src"),
				gotemir.FkPathCtor("src/auth.py", "src"),
				gotemir.FkPathCtor("src/tests/entry.py", "src/tests"),
				gotemir.FkPathCtor("src/tests/auth.py", "src/tests"),
			},
		),
		gotemir.FkDirectoryCtor([]gotemir.Path{
			gotemir.FkPathCtor("src/tests", "."),
		}),
	)
	expected := []string{"src/entry.py", "src/auth.py"}
	localizedExpected := make([]string, len(expected))
	for idx, expectedFile := range expected {
		localized, err := filepath.Localize(expectedFile)
		if err != nil {
			t.Fatalf("Fail on localized path: %s", expectedFile)
		}
		localizedExpected[idx] = localized
	}

	got, err := withoutTests.Structure()
	if err != nil {
		t.Fatalf("Fail on parse dir: %s", err)
	}
	if len(got) != 2 {
		t.Fatalf(
			strings.Join(
				[]string{
					"Actual must contains 2 elements",
					"Actual: %v",
					"Expected: %v",
					"\n",
				},
				"\n",
			),
			got, localizedExpected,
		)
	}
	for idx, actual := range got {
		actualVal, _ := actual.Relative()
		localizedActual, err := filepath.Localize(actualVal)
		if err != nil {
			t.Fatalf("Fail on localized path: %s", actual)
		}
		if localizedActual != localizedExpected[idx] {
			t.Errorf(
				strings.Join(
					[]string{
						"Incompare actual and expected at index=%d",
						"Actual: %v != Expected: %v",
						"\n",
					},
					"\n",
				),
				idx, got, localizedExpected,
			)
		}
	}
}
