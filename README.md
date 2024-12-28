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

### From source code

To use Gotemir, you need to have Go installed. You can download Go here.

Clone the repository or download the code and compile it:

```bash
git clone https://github.com/blablatdinov/gotemir.git
cd gotemir
go build src/cmd/gotemir.go
```

### Download from github releases

```
VERSION="0.0.3"
OS="linux"
ARCH="amd64"
curl -O -L -C - https://github.com/blablatdinov/gotemir/releases/download/$VERSION/gotemir-$OS-$ARCH.tar && \
tar xzf gotemir-$OS-$ARCH.tar && \
./gotemir-$OS-$ARCH src tests
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

## Contributing

We welcome contributions to Gotemir! Here's how you can get started:

1. **Fork the repository**: Click the "Fork" button at the top of this page to create a copy of the repository in your GitHub account.

2. **Clone your fork**:

```bash
git clone https://github.com/your-username/gotemir.git
cd gotemir
```

3. **Create a new branch**: Create a branch for your changes to keep them separate from the main branch.

```bash
git checkout -b my-feature-branch
```

4. **Make your changes**: Implement your feature or fix the bug.

5. **Write tests**: Ensure your changes are well-tested and do not break existing functionality.

6. Run linters and formatters:

```bash
task fmt
task lint
```

7. **Commit your changes**: Write a clear and concise commit message.

```bash
git add .
git commit -m "Add feature XYZ"
```

8. **Push to your fork:**

```bash
git push origin my-feature-branch
```

9. **Create a pull request:** Open a pull request to the main repository and provide a detailed description of your changes.

### Development environment

Gotemir uses [Taskfile](https://taskfile.dev/) to simplify common development tasks. To install and use Taskfile, refer to the [official documentation](https://taskfile.dev/installation/). All necessary tasks, such as building, testing, and linting, are encapsulated within the [Taskfile.yml](./Taskfile.yml).

For developers who prefer to avoid installing Go and Python locally, a Docker-based development environment is provided. To use it:

1. Build the Docker image:

```bash
task docker-setup
```

2. Start a container with the project directory mounted:

```bash
task docker-bash
```

Thank you for contributing to Gotemir!

## License

[MIT](https://github.com/blablatdinov/gotemir/blob/master/LICENSE)

<!--
TODO

## Examples
-->
