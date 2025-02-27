package fileutils

import (
	"errors"
	"log"
	"os"

	"github.com/dustin/go-humanize"
	hidden "github.com/tobychui/goHidden"
)

// GetListOfFiles returns a list of files in the current directory
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
				file.IsLink = true
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

// FileOrPathExists returns true if the file or path exists and false otherwise
func FileOrPathExists(fileName string) bool {
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

// IsFileExecutable returns true if the file is executable and false otherwise
func IsFileExecutable(info os.FileInfo) bool {
	fileMode := info.Mode()
	return fileMode.IsRegular() && fileMode.Perm()&0111 == 0111
}

// IsLink returns true if the file is a symbolic link and false otherwise
func IsLink(info os.DirEntry) bool {
	f, _ := info.Info()
	return f.Mode()&os.ModeSymlink != 0
}

// GetModifiedTime returns a human readable modified time
func GetModifiedTime(entry os.DirEntry) string {
	info, _ := entry.Info()
	return humanize.Time(info.ModTime())
}

// GetFileSize returns a human readable file size
func GetFileSize(entry os.DirEntry) string {
	info, _ := entry.Info()
	return humanize.Bytes(uint64(info.Size()))
}

// GetLinkPath returns the path of the link and a boolean indicating if the link destination path exists
func GetLinkPath(name string) (string, bool) {
	realPath, err := os.Readlink(name)

	if err != nil {
		return "", FileOrPathExists(realPath)
	}

	return realPath, FileOrPathExists(realPath)
}
