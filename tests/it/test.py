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
from collections.abc import Generator
from pathlib import Path

import pytest
from _pytest.legacypath import TempdirFactory


@pytest.fixture(scope="module")
def current_dir() -> Path:
    """Current directory for installing actual gotemir."""
    return Path().absolute()


@pytest.fixture(scope="module")
def _test_dir(tmpdir_factory: TempdirFactory, current_dir: str) -> Generator[Path, None, None]:
    """Directory with test structure."""
    tmp_path = tmpdir_factory.mktemp("test")
    subprocess.run(
        ["go", "build", "-o", str(tmp_path / "gotemir"), str(current_dir / "src" / "cmd" / "gotemir.go")],
        check=True,
    )
    os.chdir(tmp_path)
    Path("src").mkdir(exist_ok=True, parents=True)
    Path("src/handlers").mkdir(exist_ok=True, parents=True)
    Path("tests/handlers").mkdir(exist_ok=True, parents=True)
    Path("src/entry.py").write_bytes(b"")
    Path("src/handlers/users.py").write_bytes(b"")
    Path("tests/test_entry.py").write_bytes(b"")
    Path("tests/handlers/test_users.py").write_bytes(b"")
    yield tmp_path
    os.chdir(current_dir)


@pytest.mark.usefixtures("_test_dir")
def test() -> None:
    """Test run gotemir."""
    got = subprocess.run(["./gotemir", "src", "tests"], check=True)

    assert got.returncode == 0
