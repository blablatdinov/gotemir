# The MIT License (MIT).
#
# Copyright (c) 2024 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
# EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
# MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
# IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
# DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
# OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE
# OR OTHER DEALINGS IN THE SOFTWARE.

"""Integration tests for gotemir."""

import logging
import os
import subprocess
import sys
from collections.abc import Callable, Generator
from pathlib import Path

import pytest
from _pytest.legacypath import TempdirFactory

log = logging.getLogger("tests")


@pytest.fixture(scope="module")
def current_dir() -> Path:
    """Current directory for installing actual gotemir."""
    return Path().absolute()


@pytest.fixture
def test_dir(tmpdir_factory: TempdirFactory, current_dir: Path) -> Generator[Path, None, None]:
    """Directory with test structure."""
    tmp_path = tmpdir_factory.mktemp("test")
    subprocess.run(
        ["go", "build", "-o", str(tmp_path / "gotemir"), str(current_dir / "src" / "cmd" / "gotemir.go")],
        check=True,
    )
    os.chdir(tmp_path)
    yield tmp_path
    os.chdir(current_dir)


@pytest.fixture
def create_path(test_dir: Path) -> Callable[[str], None]:
    """Creating structure via path to file.

    Exmaple:

    from "src/handlers/users.py" create_path generate next structure:

    src
    └── handlers
        └── users.py
    """
    def _create_path(path: str) -> None:
        dir_ = "/".join(path.split("/")[:-1])
        Path(test_dir / dir_).mkdir(exist_ok=True, parents=True)
        Path(test_dir / path).write_bytes(b"")
        log.debug("Created files: {0}".format(list(Path().glob("**/*"))))
    return _create_path


@pytest.mark.usefixtures("test_dir")
@pytest.mark.parametrize(("file_structure", "src_dir", "tests_dir"), [
    (
        (
            "src/handlers/users.py",
            "src/entry.py",
            "tests/handlers/test_users.py",
            "tests/test_entry.py",
        ),
        "src",
        "tests",
    ),
    (
        (
            "src/handlers/users.py",
            "src/entry.py",
            "src/README.md",
            "tests/handlers/test_users.py",
            "tests/test_entry.py",
        ),
        "src",
        "tests",
    ),
    (
        (
            "src/handlers/users.py",
            "src/entry.py",
            "src/README.md",
            "src/tests/handlers/test_users.py",
            "src/tests/test_entry.py",
        ),
        "src",
        "src/tests",
    ),
    (
        (
            "src/handlers/users.py",
            "src/entry.py",
            "src/tests/unit/handlers/test_users.py",
            "src/tests/it/test_entry.py",
        ),
        "src",
        "src/tests/unit,src/tests/it",
    ),
    # TODO: research this case:
    # dir "tcp" equal to file name without extensions "tcp.py"
    (
        (
            "src/tcp/tcp.py",
            "tests/unit/tcp/test_tcp.py",
        ),
        "src",
        "tests/unit",
    ),
    # TODO: research this case:
    # dir name contains "test_" or "_test"
    (
        (
            "src/test_server/tcp.py",
            "tests/unit/test_server/test_tcp.py",
        ),
        "src",
        "tests/unit",
    ),
])
# TODO @blablatdinov: fix test for windows
# https://github.com/blablatdinov/gotemir/issues/14
@pytest.mark.skipif(sys.platform == "win32", reason="Test fail on windows")
def test_correct(
    create_path: Callable[[str], None],
    file_structure: tuple[str, ...],
    src_dir: str,
    tests_dir: str,
) -> None:
    """Test run gotemir."""
    [create_path(file) for file in file_structure]  # type: ignore [func-returns-value]
    got = subprocess.run(
        ["./gotemir", "--ext=.py", src_dir, tests_dir],
        stdout=subprocess.PIPE,
        check=False,
    )

    assert got.returncode == 0, got.stdout.decode("utf-8").strip()
    assert got.stdout.decode("utf-8").strip() == "Complete!"


@pytest.mark.usefixtures("test_dir")
def test_help() -> None:
    """Test --help flag."""
    got = subprocess.run(
        ["./gotemir", "--help"],
        stdout=subprocess.PIPE,
        check=False,
    )

    assert got.returncode == 0


@pytest.mark.usefixtures("test_dir")
@pytest.mark.parametrize("file_structure", [
    (
        "src/handlers/users.py",
        "src/entry.py",
        "tests/test_entry.py",
    ),
])
def test_invalid(create_path: Callable[[str], None], file_structure: tuple[str, ...]) -> None:
    """Test invalid cases."""
    [create_path(file) for file in file_structure]  # type: ignore [func-returns-value]
    got = subprocess.run(
        ["./gotemir", "--ext", ".py", "src", "tests"],
        stdout=subprocess.PIPE, check=False,
    )

    assert got.returncode == 1
    assert got.stdout.decode("utf-8").strip() == "{0}:0:0 Not found test for file".format(
        str(Path("src/handlers/users.py")),
    )


@pytest.mark.usefixtures("test_dir")
@pytest.mark.parametrize("file_structure", [
    (
        "src/entry.py",
        "tests/test_entry.py",
        "tests/test_users.py",
    ),
    # (
    #     (
    #         "src/test_server/tcp.py",
    #         "tests/unit/test_server/test_tcp.py",
    #     ),
    #     "src",
    #     "tests",
    # ),
])
def test_unbinded_test_file(create_path: Callable[[str], None], file_structure: tuple[str, ...]) -> None:
    """Check test files without src code."""
    [create_path(file) for file in file_structure]  # type: ignore [func-returns-value]
    got = subprocess.run(
        ["./gotemir", "--ext", ".py", "src", "tests"],
        stdout=subprocess.PIPE, check=False,
    )

    assert got.returncode == 1, got.stdout.decode("utf-8").strip()
    assert got.stdout.decode("utf-8").strip() == "{0}:0:0 Not found source file for test".format(
        str(Path("tests/test_users.py")),
    )
