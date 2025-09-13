package sortfuncs

import (
	"cmp"
	"slices"

	"github.com/Soulsbane/ago/internal/fileutils"
)

func SortByModTime(files []fileutils.FileInfo, sortOrder string) []fileutils.FileInfo {
	slices.SortFunc(files, func(a, b fileutils.FileInfo) int {
		if sortOrder == "desc" {
			return cmp.Compare(a.Modified, b.Modified)
		}

		return -cmp.Compare(a.Modified, b.Modified)
	})

	return files
}

func SortBySize(files []fileutils.FileInfo, sortOrder string) []fileutils.FileInfo {
	slices.SortFunc(files, func(a, b fileutils.FileInfo) int {
		if sortOrder == "desc" {
			return -cmp.Compare(a.RawSize, b.RawSize)
		}

		return cmp.Compare(a.RawSize, b.RawSize)
	})

	return files
}

func SortByFileName(files []fileutils.FileInfo, sortOrder string) []fileutils.FileInfo {
	slices.SortFunc(files, func(a, b fileutils.FileInfo) int {
		if sortOrder == "desc" {
			return -cmp.Compare(a.Name, b.Name)
		}

		return cmp.Compare(a.Name, b.Name)
	})

	return files
}

// import (
// 	"fmt"
// )

// func SortBySizeAgain() {
// 	fmt.Print("Hello, World!")
// }
