package fileutils

import (
	"errors"
	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	hidden "github.com/tobychui/goHidden"
	"log"
	"os"
)

func GetListOfFiles(showHidden bool) []FileInfo {
	var fileList []FileInfo
	files, err := os.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !f.IsDir() {
			var file FileInfo
			isHidden, _ := hidden.IsHidden(f.Name(), false)
			info, _ := f.Info()

			file.Name = f.Name()
			file.Executable = IsFileExecutable(info)
			file.HumanizeSize = GetFileSize(f)
			file.RawSize = info.Size()
			file.HumanizeModified = GetModifiedTime(f)
			file.Modified = info.ModTime().Unix()

			if IsLink(f) {
				file.LinkPath, _ = GetLinkPath(f.Name())
			}

			if isHidden {
				if showHidden {
					fileList = append(fileList, file)
				}
			} else {
				fileList = append(fileList, file)
			}
		}
	}

	return fileList
}

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

func GetModifiedTime(entry os.DirEntry) string {
	info, _ := entry.Info()
	return humanize.Time(info.ModTime())
}

func GetFileSize(entry os.DirEntry) string {
	info, _ := entry.Info()
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
