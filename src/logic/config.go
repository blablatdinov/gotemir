// SPDX-FileCopyrightText: Copyright (c) 2024-2026 Almaz Ilaletdinov <a.ilaletdinov@yandex.ru>
// SPDX-License-Identifier: MIT

package logic

type Config struct {
	Version       int           `yaml:"version"`
	GotemirConfig GotemirConfig `yaml:"gotemir"`
}

type GotemirConfig struct {
	TestFreeFiles []string `yaml:"test-free-files"` //nolint:tagliatelle
	TestHelpers   []string `yaml:"test-helpers"`    //nolint:tagliatelle
}
