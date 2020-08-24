package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"sort"
	"text/tabwriter"

	"github.com/alexflint/go-arg"
	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
)

func getModifedTime(info os.FileInfo, colorize bool) string {
	modifiedTime := info.ModTime()

	if colorize {
		return color.HiBlueString(humanize.Time(modifiedTime))
	}

	return humanize.Time(modifiedTime)
}

func getFileSize(info os.FileInfo, colorize bool) string {
	if colorize {
		return color.HiYellowString(humanize.Bytes(uint64(info.Size())))
	}

	return humanize.Bytes(uint64(info.Size()))
}

func getFileName(info os.FileInfo, colorize bool) string {
	if colorize {
		if isFileExecutable(info) {
			return color.HiRedString(info.Name())
		}
	}

	return color.WhiteString(info.Name())
}

// INFO: Always returns false on windows as it's not supported.
func isFileHidden(info os.FileInfo) bool {
	if runtime.GOOS != "windows" {
		if info.Name()[0:1] == "." {
			return true
		}

		return false
	}

	return false
}

func isFileExecutable(info os.FileInfo) bool {
	mode := info.Mode()
	exec := mode & 0111

	if exec != 0 {
		return true
	}

	return false
}

func listFiles(ugly bool, showHidden bool, sortByModTime bool) {
	var filteredFiles []os.FileInfo
	files, err := ioutil.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !f.IsDir() {
			if isFileHidden(f) {
				if showHidden {
					filteredFiles = append(filteredFiles, f)
				}
			} else {
				filteredFiles = append(filteredFiles, f)
			}
		}
	}

	outputResults(filteredFiles, ugly, sortByModTime)
}

// TODO Possibly add more sorting options
func sortResults(files []os.FileInfo) []os.FileInfo {
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().Unix() < files[j].ModTime().Unix()
	})

	return files
}

// FIXME: Still some problems lining up correctly.
func outputResults(files []os.FileInfo, ugly bool, sortByModTime bool) {
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
	listFiles(args.Ugly, args.Hidden, args.Sort)
}
