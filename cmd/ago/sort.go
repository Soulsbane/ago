package main

import (
	"os"
	"sort"
)

func sortByModTime(files []os.DirEntry) []os.DirEntry {
	sort.Slice(files, func(i, j int) bool {
		infoI, _ := files[i].Info()
		infoJ, _ := files[j].Info()
		return infoI.ModTime().Unix() > infoJ.ModTime().Unix()
	})

	return files
}

func sortBySize(files []os.DirEntry) []os.DirEntry {
	sort.Slice(files, func(i, j int) bool {
		infoI, _ := files[i].Info()
		infoJ, _ := files[j].Info()
		return infoI.Size() > infoJ.Size()
	})

	return files
}

func sortByFileName(files []os.DirEntry) []os.DirEntry {
	sort.Slice(files, func(i, j int) bool {
		infoI, _ := files[i].Info()
		infoJ, _ := files[j].Info()
		return infoI.Name() > infoJ.Name()
	})

	return files
}
