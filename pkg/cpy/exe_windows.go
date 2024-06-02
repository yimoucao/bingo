// Copyright (c) Bartłomiej Płotka @bwplotka
// Licensed under the Apache License 2.0.

//go:build windows

package cpy

import (
	"strings"
)

func ensureExe(f string) string {
	if strings.HasSuffix(f, ".exe") {
		return f
	}
	return f + ".exe"
}
