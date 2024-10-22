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

import os
import subprocess
from collections.abc import Callable, Generator
from pathlib import Path

import pytest
from _pytest.legacypath import TempdirFactory


@pytest.fixture(scope="module")
def current_dir() -> Path:
    """Current directory for installing actual gotemir."""
    return Path().absolute()


@pytest.fixture
def test_dir(tmpdir_factory: TempdirFactory, current_dir: str) -> Generator[Path, None, None]:
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

    └── src
        └── handlers
            └── users.py
    """
    def _create_path(path: str) -> None:
        dir_ = "/".join(path.split("/")[:-1])
        Path(test_dir / dir_).mkdir(exist_ok=True, parents=True)
        Path(test_dir / path).write_bytes(b"")
    return _create_path


@pytest.mark.usefixtures("test_dir")
@pytest.mark.parametrize("file_structure", [
    (
        "src/handlers/users.py",
        "src/entry.py",
        "tests/handlers/test_users.py",
        "tests/test_entry.py",
    ),
])
def test_correct(create_path: Callable[[str], None], file_structure: tuple[str, ...]) -> None:
    """Test run gotemir."""
    [create_path(file) for file in file_structure]
    got = subprocess.run(
        ["./gotemir", "src", "tests"],
        stdout=subprocess.PIPE,
    )

    assert got.returncode == 0
    assert got.stdout.decode("utf-8").strip() == "Complete!"


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
    [create_path(file) for file in file_structure]
    got = subprocess.run(
        ["./gotemir", "src", "tests"],
        stdout=subprocess.PIPE, check=False,
    )

    assert got.returncode == 1
    assert got.stdout.decode("utf-8").strip() == "\n".join([
        "Files without tests:",
        " - {0}".format(str(Path("handlers/users.py"))),
    ])
