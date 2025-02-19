package fileutils

import (
	"errors"
	"os"
)

func FileOrPathExists(fileName string) bool {
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func IsFileExecutable(info os.FileInfo) bool {
	return info.Mode()&0111 != 0
}

func IsLink(info os.DirEntry) bool {
	f, _ := info.Info()
	return f.Mode()&os.ModeSymlink != 0
}

// GetLinkPath returns the path of the link and a boolean indicating if the link destination path exists
func GetLinkPath(name string) (string, bool) {
	realPath, err := os.Readlink(name)

	if err != nil {
		return "", FileOrPathExists(realPath)
	}

	return realPath, FileOrPathExists(realPath)
}
