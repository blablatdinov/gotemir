package logic

import (
	"strings"
)

type Directory interface {
	structure() ([]string, error)
}

type FkDirectory struct {
	origin []string
}

func FkDirectoryCtor(origin []string) Directory {
	return FkDirectory{
		origin,
	}
}

func (fkDirectory FkDirectory) structure() ([]string, error) {
	return fkDirectory.origin, nil
}

func Compare(srcDir Directory, testsDir Directory) []string {
	filesWithoutTests := make([]string, 0)
	srcFiles, _ := srcDir.structure()
	testFiles, _ := testsDir.structure()
	for _, srcFile := range srcFiles {
		fileExtension := "." + strings.Split(srcFile, ".")[1]
		splittedPath := strings.Split(srcFile, "/")
		fileName := splittedPath[len(splittedPath)-1]
		fileNameWithoutExtension := strings.Split(fileName, ".")[0]
		testFileVariants := []string{
			strings.Replace(
				srcFile,
				fileName,
				fileNameWithoutExtension+"_test"+fileExtension,
				1,
			),
			strings.Replace(
				srcFile,
				fileName,
				"test_"+fileNameWithoutExtension+fileExtension,
				1,
			),
		}
		testFileFound := false
	out:
		for _, testFile := range testFiles {
			for _, testFileVariant := range testFileVariants {
				if testFile == testFileVariant {
					testFileFound = true
					break out
				}
			}
		}
		if !testFileFound {
			filesWithoutTests = append(filesWithoutTests, srcFile)
		}
	}
	return filesWithoutTests
}
