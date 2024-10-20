package tests

import (
	"strings"
	"testing"

	gotemir "github.com/blablatdinov/gotemir/src"
)

func TestCompare(t *testing.T) {
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
		for i := 0; i < len(got); i++ {
			if testCase.expected[i] != got[i] {
				t.Errorf("Expected: %s, actual: %s\n", testCase.expected, got[i])
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
					i,
					testCase.srcDir,
					testCase.testsDir,
					got[i],
					testCase.expected[i],
				)
			}
		}
	}
}
