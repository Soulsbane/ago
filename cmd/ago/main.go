package main

import (
	"log"
	"os"
	"runtime"
	"sort"
	"strings"

	"github.com/alexflint/go-arg"
	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

func getModifedTime(entry os.DirEntry, colorize bool) string {
	info, _ := entry.Info()
	modifiedTime := info.ModTime()

	if colorize {
		return color.HiBlueString(humanize.Time(modifiedTime))
	}

	return humanize.Time(modifiedTime)
}

func getFileSize(entry os.DirEntry, colorize bool) string {
	info, _ := entry.Info()

	if colorize {
		return color.HiYellowString(humanize.Bytes(uint64(info.Size())))
	}

	return humanize.Bytes(uint64(info.Size()))
}

func getFileName(entry os.DirEntry, colorize bool) string {
	info, _ := entry.Info()

	if colorize {
		if isFileExecutable(info) {
			return color.HiRedString(info.Name())
		}
	}

	return info.Name()
}

// INFO: Always returns false on windows as it's not supported.
func isFileHidden(name string) bool {
	if runtime.GOOS != "windows" {
		if strings.HasPrefix(name, ".") {
			return true
		}
	}

	return false
}

func isFileExecutable(info os.FileInfo) bool {
	return info.Mode()&0111 != 0
}

// TODO: Check for links
func getListOfFiles(showHidden bool) []os.DirEntry {
	var fileList []os.DirEntry
	files, err := os.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !f.IsDir() {
			if isFileHidden(f.Name()) {
				if showHidden {
					fileList = append(fileList, f)
				}
			} else {
				fileList = append(fileList, f)
			}
		}
	}

	return fileList
}

// TODO Possibly add more sorting options
func sortResults(files []os.DirEntry) []os.DirEntry {
	sort.Slice(files, func(i, j int) bool {
		infoI, _ := files[i].Info()
		infoJ, _ := files[i].Info()
		return infoI.ModTime().Unix() > infoJ.ModTime().Unix()
	})

	return files
}

func outputResults(files []os.DirEntry, ugly bool, sortByModTime bool, noTable bool) {
	dirDataTable := table.NewWriter()
	dirDataTable.SetOutputMirror(os.Stdout)

	if !noTable {
		dirDataTable.AppendHeader(table.Row{"Name", "Size", "Modified"})
	}

	if sortByModTime {
		files = sortResults(files)
	}

	for _, f := range files {
		if ugly {
			dirDataTable.AppendRow(table.Row{getFileName(f, false), getFileSize(f, false), getModifedTime(f, false)})
		} else {
			dirDataTable.AppendRow(table.Row{getFileName(f, true), getFileSize(f, true), getModifedTime(f, true)})
		}
	}

	dirDataTable.SetStyle(table.StyleRounded)

	if noTable {
		dirDataTable.Style().Options.DrawBorder = false
		dirDataTable.Style().Options.SeparateColumns = false
		dirDataTable.Style().Options.SeparateRows = false
		dirDataTable.Style().Options.SeparateHeader = false
	}

	dirDataTable.Render()
}

func main() {
	var args ProgramArgs

	arg.MustParse(&args)
	files := getListOfFiles(args.Hidden)

	if args.Sort {
		files = sortResults(files)
	}

	outputResults(files, args.Ugly, args.Sort, args.NoTable)
}
