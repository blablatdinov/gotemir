<!---
The MIT License (MIT).

Copyright (c) 2024-2025 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>

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
-->
# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Installation instructions (#51)
- Dev dockerfile (#56)
- Yamllint (#61)
- Self check (#59)

### Fixes

- Licenses (#58)
- Fix handle invalid regex (#77)

### Changed

- Update `README.md`
- Update license year (#55)
- Update config format (#80)

### Removed

- Isort (#70)

## [0.0.3] - 2024-12-09

### Added

- Check `go mod tidy` applied (#38)
- Build task (#47)

### Fixed

- Bug on ignore file (#49)

### Changed

- Update `README.md`
- `--help` out (#35)
- Update dependencies
- Renovate config
- Replace `Makefile` -> `Taskfile.yml` (#42)
- Update `Taskfile.yml` (#44)

## [0.0.2] - 2024-10-28

### Added

- Compress release archieve
- Filter test free files by regex (#32)
- Ignore test helpers (#33)
- Unit test for `FilteredByConfigFiles` (#34)

### Changed

- Uncomment test for windows (#31)

## [0.0.1] - 2024-10-28

### Added

- Configuration for app (#29)

### Changed

- Update `README.md` (#28)

### Fixed

- Unit test for windows (#30)

## [0.0.1a6] - 2024-10-25

### Fixed

- Read github token

## [0.0.1a5] - 2024-10-25

### Added

- Tar binaries

## [0.0.1a4] - 2024-10-25

### Fixed

- Read github token

## [0.0.1a3] - 2024-10-25

### Removed

- Exclude darwin 386 from release

## [0.0.1a2] - 2024-10-25

### Fixed

- Build command in CI

## [0.0.1a1] - 2024-10-25

### Added

- Init version

[unreleased]: https://github.com/blablatdinov/gotemir/compare/0.0.3...HEAD
[0.0.3]: https://github.com/blablatdinov/gotemir/compare/0.0.2...0.0.3
[0.0.2]: https://github.com/blablatdinov/gotemir/compare/0.0.1...0.0.2
[0.0.1]: https://github.com/blablatdinov/gotemir/compare/0.0.1-a6...0.0.1
[0.0.1a6]: https://github.com/blablatdinov/gotemir/compare/0.0.1-a5...0.0.1-a6
[0.0.1a5]: https://github.com/blablatdinov/gotemir/compare/0.0.1-a4...0.0.1-a5
[0.0.1a4]: https://github.com/blablatdinov/gotemir/compare/0.0.1-a3...0.0.1-a4
[0.0.1a3]: https://github.com/blablatdinov/gotemir/compare/0.0.1-a2...0.0.1-a3
[0.0.1a2]: https://github.com/blablatdinov/gotemir/compare/0.0.1-a1...0.0.1-a2
[0.0.1a1]: https://github.com/blablatdinov/gotemir/releases/tag/0.0.1-a1
