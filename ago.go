package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"text/tabwriter"

	"github.com/alexflint/go-arg"
	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
)

func getModifedTime(info os.FileInfo, colorize bool) string {
	modifiedTime := info.ModTime()

	if colorize == true {
		return color.HiBlueString(humanize.Time(modifiedTime))
	}

	return humanize.Time(modifiedTime)
}

func getFileSize(info os.FileInfo, colorize bool) string {
	if colorize == true {
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

	return info.Name()
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

func listFiles(ugly bool, showHidden bool) {
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

	outputResults(filteredFiles, ugly)
}

func outputResults(files []os.FileInfo, ugly bool) {
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
		Ugly   bool `arg:"-u" default:"false" help:"Remove colorized output. Yes it's ugly."`
		Hidden bool `arg:"-i" default:"false" help:"Show hidden files."`
	}

	arg.MustParse(&args)
	listFiles(args.Ugly, args.Hidden)
}
