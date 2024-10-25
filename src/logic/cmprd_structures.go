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

package logic

type CmprdStructures struct {
	srcDir   Directory
	testsDir Directory
}

func CmprdStructuresCtor(srcDir, testsDir Directory) ComparedStructures {
	return CmprdStructures{
		srcDir,
		testsDir,
	}
}

func (cmprdStructures CmprdStructures) FilesWithoutTests() []string {
	filesWithoutTests := make([]string, 0)
	testFiles, _ := cmprdStructures.testsDir.Structure()
	srcFiles, _ := cmprdStructures.srcDir.Structure()
	for _, srcFile := range srcFiles {
		relativePath, _ := srcFile.Relative()
		testFileVariants := TestFileVariants(relativePath)
		testFileFound := false
	out:
		for _, testFile := range testFiles {
			relativePath, _ := testFile.Relative()
			for _, testFileVariant := range testFileVariants {
				if relativePath == testFileVariant {
					testFileFound = true
					break out
				}
			}
		}
		if !testFileFound {
			val, _ := srcFile.Absolute()
			filesWithoutTests = append(filesWithoutTests, val)
		}
	}
	return filesWithoutTests
}

func (cmprdStructures CmprdStructures) TestsWithoutSrcFiles() []string {
	testsWithoutSrcFiles := make([]string, 0)
	testFiles, _ := cmprdStructures.testsDir.Structure()
	srcFiles, _ := cmprdStructures.srcDir.Structure()
	for _, testFile := range testFiles {
		relativeTestPath, _ := testFile.Relative()
		srcFileVariant := SourceFileVariants(relativeTestPath)
		srcFileFound := false
		for _, srcFile := range srcFiles {
			relativeSrcPath, _ := srcFile.Relative()
			if relativeSrcPath == srcFileVariant {
				srcFileFound = true
				break
			}
		}
		if !srcFileFound {
			val, _ := testFile.Absolute()
			testsWithoutSrcFiles = append(testsWithoutSrcFiles, val)
		}
	}
	return testsWithoutSrcFiles
}
