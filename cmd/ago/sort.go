package main

import (
	"github.com/Soulsbane/ago/internal/fileutils"
	"sort"
)

func sortByModTime(files []fileutils.FileInfo) []fileutils.FileInfo {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Modified > files[j].Modified
	})

	return files
}

func sortBySize(files []fileutils.FileInfo) []fileutils.FileInfo {
	sort.Slice(files, func(i, j int) bool {
		return files[i].RawSize > files[j].RawSize
	})

	return files
}

func sortByFileName(files []fileutils.FileInfo) []fileutils.FileInfo {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name > files[j].Name
	})

	return files
}
