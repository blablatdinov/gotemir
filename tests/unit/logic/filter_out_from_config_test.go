// SPDX-FileCopyrightText: Copyright (c) 2024-2026 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
// SPDX-License-Identifier: MIT

package logic_test

import (
	"slices"
	"testing"

	"github.com/blablatdinov/gotemir/src/logic"
)

func TestFilterOut(t *testing.T) {
	t.Parallel()
	t.Run("empty_input_returns_empty", func(t *testing.T) {
		t.Parallel()
		filtered := logic.FilterOutFromConfifCtor(
			logic.FkComparedStructuresCtor([]string{}, []string{}),
			logic.Config{
				Version:        0,
				GotemirConfig:  logic.GotemirConfig{TestFreeFiles: []string{}, TestHelpers: []string{}},
			},
		)
		files, err := filtered.FilesWithoutTests()
		if err != nil {
			t.Fatal(err)
		}
		if len(files) != 0 {
			t.Errorf("FilesWithoutTests() = %v, want []", files)
		}
		tests, err := filtered.TestsWithoutSrcFiles()
		if err != nil {
			t.Fatal(err)
		}
		if len(tests) != 0 {
			t.Errorf("TestsWithoutSrcFiles() = %v, want []", tests)
		}
	})
	t.Run("TestFreeFiles_excludes_matching_files", func(t *testing.T) {
		t.Parallel()
		filtered := logic.FilterOutFromConfifCtor(
			logic.FkComparedStructuresCtor(
				[]string{"pkg/foo.go", "pkg/bar.go", "cmd/main.go"},
				[]string{},
			),
			logic.Config{
				Version: 0,
				GotemirConfig: logic.GotemirConfig{
					TestFreeFiles: []string{`pkg/.*\.go`},
					TestHelpers:   []string{},
				},
			},
		)
		got, err := filtered.FilesWithoutTests()
		if err != nil {
			t.Fatal(err)
		}
		want := []string{"cmd/main.go"}
		if !slices.Equal(got, want) {
			t.Errorf("FilesWithoutTests() = %v, want %v", got, want)
		}
	})
	t.Run("TestHelpers_excludes_matching_tests", func(t *testing.T) {
		t.Parallel()
		filtered := logic.FilterOutFromConfifCtor(
			logic.FkComparedStructuresCtor(
				[]string{},
				[]string{"tests/helper.go", "tests/foo_test.go", "tests/util.go"},
			),
			logic.Config{
				Version: 0,
				GotemirConfig: logic.GotemirConfig{
					TestFreeFiles: []string{},
					TestHelpers:   []string{`tests/helper\.go`, `tests/util\.go`},
				},
			},
		)
		got, err := filtered.TestsWithoutSrcFiles()
		if err != nil {
			t.Fatal(err)
		}
		want := []string{"tests/foo_test.go"}
		if !slices.Equal(got, want) {
			t.Errorf("TestsWithoutSrcFiles() = %v, want %v", got, want)
		}
	})
	t.Run("invalid_regex_in_TestFreeFiles_returns_error", func(t *testing.T) {
		t.Parallel()
		filtered := logic.FilterOutFromConfifCtor(
			logic.FkComparedStructuresCtor([]string{"foo.go"}, []string{}),
			logic.Config{
				Version: 0,
				GotemirConfig: logic.GotemirConfig{
					TestFreeFiles: []string{`[invalid`},
					TestHelpers:   []string{},
				},
			},
		)
		_, err := filtered.FilesWithoutTests()
		if err == nil {
			t.Error("FilesWithoutTests() err = nil, want error for invalid regex")
		}
	})
	t.Run("invalid_regex_in_TestHelpers_returns_error", func(t *testing.T) {
		t.Parallel()
		filtered := logic.FilterOutFromConfifCtor(
			logic.FkComparedStructuresCtor([]string{}, []string{"helper.go"}),
			logic.Config{
				Version: 0,
				GotemirConfig: logic.GotemirConfig{
					TestFreeFiles: []string{},
					TestHelpers:   []string{`(?P<name>unclosed`},
				},
			},
		)
		_, err := filtered.TestsWithoutSrcFiles()
		if err == nil {
			t.Error("TestsWithoutSrcFiles() err = nil, want error for invalid regex")
		}
	})
}
