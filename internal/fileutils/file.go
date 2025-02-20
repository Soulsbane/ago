package fileutils

import (
	"errors"
	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
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

func GetModifiedTime(entry os.DirEntry, colorize bool) string {
	info, _ := entry.Info()
	modifiedTime := info.ModTime()

	if colorize {
		return color.HiBlueString(humanize.Time(modifiedTime))
	}

	return humanize.Time(modifiedTime)
}

func GetFileSize(entry os.DirEntry, colorize bool) string {
	info, _ := entry.Info()

	if colorize {
		return color.HiYellowString(humanize.Bytes(uint64(info.Size())))
	}

	return humanize.Bytes(uint64(info.Size()))
}

func GetFileName(entry os.DirEntry, colorize bool, noLinks bool) string {
	info, _ := entry.Info()
	name := info.Name()
	linkText := ""
	path, _ := GetLinkPath(name)

	if !noLinks && IsLink(entry) {
		linkText = "-> " + path
	}

	if colorize {
		// If the path doesn't exist due to a broken link then display it in red
		if !FileOrPathExists(path) {
			return color.HiRedString(name + " " + linkText)
		}

		if IsFileExecutable(info) {
			return color.HiBlueString(name + " " + linkText)
		}
	}

	return name + " " + linkText
}

// GetLinkPath returns the path of the link and a boolean indicating if the link destination path exists
func GetLinkPath(name string) (string, bool) {
	realPath, err := os.Readlink(name)

	if err != nil {
		return "", FileOrPathExists(realPath)
	}

	return realPath, FileOrPathExists(realPath)
}
