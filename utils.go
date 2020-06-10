package bbpak

import (
	"os"
	"path"
	"strings"
)

// CleanDir removes trailing spaces safely
func CleanDir(val string) string {
	out := make([]string, 0)
	for _, pch := range strings.Split(val, "/") {
		if pch != "" {
			out = append(out, pch)
		}
	}
	return path.Join(out...)
}

// ResolveIfSymlink checks the input if it is a symlink and then resolves it accordingly
func ResolveIfSymlink(pth string) string {
	pth = CleanDir(pth)
	nfo, err := os.Lstat(pth)
	if err != nil {
		panic(err)
	}
	if nfo.Mode()&os.ModeSymlink != 0 {
		ref, err := os.Readlink(pth)
		if err != nil {
			panic(err)
		}
		return ref
	}
	return pth
}
