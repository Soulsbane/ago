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

func getListOfFiles(showHidden bool) []fileutils.FileInfo {
	var fileList []fileutils.FileInfo
	files, err := os.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !f.IsDir() {
			var file fileutils.FileInfo
			isHidden, _ := hidden.IsHidden(f.Name(), false)
			info, _ := f.Info()

			file.Name = f.Name()
			file.HumanizeSize = fileutils.GetFileSize(f, true)
			file.RawSize = info.Size()
			file.HumanizeModified = fileutils.GetModifiedTime(f, true)
			file.Modified = info.ModTime().Unix()

			if fileutils.IsLink(f) {
				file.LinkPath, _ = fileutils.GetLinkPath(f.Name())
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

func outputResults(files []fileutils.FileInfo, ugly bool, noTable bool, showLinks bool) {
	var totalFileSize int64
	dirDataTable := table.NewWriter()

	dirDataTable.SetOutputMirror(os.Stdout)

	if !noTable {
		dirDataTable.AppendHeader(table.Row{"Modified", "Size", "Name"})
	}

	for _, f := range files {
		if ugly {
			dirDataTable.AppendRow(table.Row{
				f.HumanizeModified,
				f.HumanizeSize,
				f.Name,
			})

			totalFileSize += f.RawSize

		} else {
			dirDataTable.AppendRow(table.Row{
				f.HumanizeModified,
				f.HumanizeSize,
				f.Name,
			})

			totalFileSize += f.RawSize
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
