package main

import (
	"github.com/Soulsbane/ago/internal/fileutils"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/alexflint/go-arg"
	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

func getModifiedTime(entry os.DirEntry, colorize bool) string {
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

func getFileName(entry os.DirEntry, colorize bool, noLinks bool) string {
	info, _ := entry.Info()
	linkText := ""

	if !noLinks && fileutils.IsLink(entry) {
		linkText = color.HiMagentaString("(link)")
	}

	if colorize {
		if isFileExecutable(info) {
			return color.HiRedString(info.Name() + " " + linkText)
		}
	}

	return info.Name() + " " + linkText
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

func outputResults(files []os.DirEntry, ugly bool, noTable bool, showLinks bool) {
	var totalFileSize int64
	dirDataTable := table.NewWriter()

	dirDataTable.SetOutputMirror(os.Stdout)

	if !noTable {
		dirDataTable.AppendHeader(table.Row{"Modified", "Size", "Name"})
	}

	for _, f := range files {
		if ugly {
			dirDataTable.AppendRow(table.Row{getModifiedTime(f, false), getFileSize(f, false), getFileName(f, false, showLinks)})
			info, _ := f.Info()
			totalFileSize += info.Size()

		} else {
			dirDataTable.AppendRow(table.Row{getModifiedTime(f, true), getFileSize(f, true), getFileName(f, true, showLinks)})
			info, _ := f.Info()
			totalFileSize += info.Size()
		}
	}

	if noTable {
		dirDataTable.SetStyle(agoNoStyle)
	} else {
		dirDataTable.SetStyle(agoDefaultStyle)
	}

	dirDataTable.AppendSeparator()
	dirDataTable.AppendFooter(table.Row{"TOTAL", humanize.Bytes(uint64(totalFileSize))})

	dirDataTable.Render()
}

func main() {
	var args ProgramArgs

	parser := arg.MustParse(&args)
	files := getListOfFiles(args.Hidden)

	if args.SortBy == "name" {
		files = sortByFileName(files)
	} else if args.SortBy == "size" {
		files = sortBySize(files)
	} else if args.SortBy == "modified" {
		files = sortByModTime(files)
	} else {
		parser.Fail("Invalid sort option! Valid options are: 'name', 'size', or 'modified'.")
	}

	outputResults(files, args.Ugly, args.NoTable, args.NoLinks)
}
