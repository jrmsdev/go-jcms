package fsutils

import (
	"os"
)

func DirExists(p string) bool {
	if f, err := os.Stat(p); os.IsNotExist(err) {
		return false
	} else {
		return f.IsDir()
	}
}

func FileExists(p string) bool {
	if f, err := os.Stat(p); os.IsNotExist(err) {
		return false
	} else {
		return !f.IsDir()
	}
}
