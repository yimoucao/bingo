//go:build linux || darwin

package cpy

func ensureExe(f string) string {
	return f
}
