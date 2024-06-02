// Copyright (c) Yimou Cao @yimoucao
// Licensed under the Apache License 2.0.

//go:build linux || darwin

package cpy

func ensureExe(f string) string {
	return f
}
