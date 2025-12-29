// SPDX-FileCopyrightText: Copyright (c) 2024-2026 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
// SPDX-License-Identifier: MIT

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
