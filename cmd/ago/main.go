package main

import (
	"github.com/Soulsbane/ago/internal/fileutils"
	"github.com/alexflint/go-arg"
	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"path/filepath"
)

func GetColorizedName(info fileutils.FileInfo, noLinks bool) string {
	name := info.Name
	linkText := ""
	path, _ := fileutils.GetLinkPath(name)

	if !noLinks && info.LinkPath != "" {
		linkText = "-> " + path
	}

	if !fileutils.FileOrPathExists(path) {
		return color.HiRedString(name + " " + linkText)
	}

	if info.Executable {
		program := filepath.Base(path)
		dir := filepath.Dir(path)
		green := color.New(color.Bold, color.FgGreen).SprintfFunc()

		return color.HiBlueString(name) + "-> " + filepath.Join(color.HiBlueString(dir), green(program))
	}

	return name + " " + linkText
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
				color.HiBlueString(f.HumanizeModified),
				color.YellowString(f.HumanizeSize),
				GetColorizedName(f, showLinks),
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
	files := fileutils.GetListOfFiles(args.Hidden)

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
