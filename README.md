<!---
The MIT License (MIT)

Copyright (c) 2024 <a.ilaletdinov@yandex.ru>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE
OR OTHER DEALINGS IN THE SOFTWARE.
--->
# Gotemir

[![Build status](https://github.com/blablatdinov/gotemir/actions/workflows/pr-check.yaml/badge.svg)](https://github.com/blablatdinov/gotemir/actions/workflows/pr-check.yaml)
[![Lines of code](https://tokei.rs/b1/github/blablatdinov/gotemir)](https://github.com/XAMPPRocky/tokei_rs)
[![Hits-of-Code](https://hitsofcode.com/github/blablatdinov/gotemir)](https://hitsofcode.com/github/blablatdinov/gotemir/view)

Gotemir is a tool that verifies if the structure of the test directory mirrors the structure of the source code directory. It ensures that for every source file, a corresponding test file exists in the appropriate directory.

[On the Layout of Tests](https://www.yegor256.com/2023/01/19/layout-of-tests.html)

## Features

- Recursively scans both source and test directories.
- Verifies that every source file has a corresponding test file.
- Identifies missing test files and outputs results to the console.

## Installation

To use Gotemir, you need to have Go installed. You can download Go here.

Clone the repository or download the code and compile it:

```bash
git clone https://github.com/blablatdinov/gotemir.git
cd gotemir
go build src/cmd/gotemir.go
```

## Usage

Once compiled, you can run Gotemir by specifying the source and test directories.
Example project structure:

```
src/
├── main.go
├── service/
│   └── user.go
tests/
├── main_test.go
└── service/
    └── user_test.go
```

Run `gotemir`

```bash
./gotemir --ext .go src tests
```

## License

[MIT](https://github.com/blablatdinov/gotemir/blob/master/LICENSE)

<!--
TODO

## Examples
-->
