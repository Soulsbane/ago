package main

import (
	"github.com/Soulsbane/ago/internal/fileutils"
	"github.com/alexflint/go-arg"
	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/v6/table"
	hidden "github.com/tobychui/goHidden"
	"log"
	"os"
)

// TODO: Check for links
func getListOfFiles(showHidden bool) []os.DirEntry {
	var fileList []os.DirEntry
	files, err := os.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !f.IsDir() {
			isHidden, _ := hidden.IsHidden(f.Name(), false)

			if isHidden {
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
			dirDataTable.AppendRow(table.Row{
				fileutils.GetModifiedTime(f, false),
				fileutils.GetFileSize(f, false),
				fileutils.GetFileName(f, false, showLinks),
			})

			info, _ := f.Info()
			totalFileSize += info.Size()

		} else {
			dirDataTable.AppendRow(table.Row{
				fileutils.GetModifiedTime(f, true),
				fileutils.GetFileSize(f, true),
				fileutils.GetFileName(f, true, showLinks),
			})

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
