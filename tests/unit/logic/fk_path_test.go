// SPDX-FileCopyrightText: Copyright (c) 2024-2026 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
// SPDX-License-Identifier: MIT

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
