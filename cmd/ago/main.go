package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/alexflint/go-arg"
	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
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

func outputResults(files []os.DirEntry, ugly bool, sortByModTime bool) {
	if sortByModTime {
		files = sortResults(files)
	}

	writer := tabwriter.NewWriter(os.Stdout, 1, 4, 1, ' ', 0)

	for _, f := range files {
		if ugly {
			fmt.Fprintf(writer, "%s\t%s\t%s\t\n", getFileName(f, false), getFileSize(f, false), getModifedTime(f, false))
		} else {
			fmt.Fprintf(writer, "%s\t%s\t%s\t\n", getFileName(f, true), getFileSize(f, true), getModifedTime(f, true))
		}
	}

	writer.Flush()
}

func main() {
	var args struct {
		Ugly   bool `arg:"-u, --ugly" default:"false" help:"Remove colorized output. Yes it's ugly."`
		Hidden bool `arg:"-a, --all" default:"false" help:"Show hidden files also."`
		Sort   bool `arg:"-s, --sort" default:"false" help:"Sorts the files by file modification time."`
	}

	arg.MustParse(&args)
	files := getListOfFiles(args.Hidden)

	if args.Sort {
		files = sortResults(files)
	}

	outputResults(files, args.Ugly, args.Sort)
}
