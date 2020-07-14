package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
)

func getDifference() int64 {
	info, err := os.Stat("ago")

	if err != nil {
		log.Fatal(err)
	}

	modifiedTime := info.ModTime()
	now := time.Now()
	difference := now.Unix() - modifiedTime.Unix()

	return difference
}

func getModifedTime(info os.FileInfo, colorize bool) string {
	modifiedTime := info.ModTime()

	if colorize == true {
		return color.BlueString(humanize.Time(modifiedTime))
	}

	return humanize.Time(modifiedTime)
}

func getFileSize(info os.FileInfo, colorize bool) string {
	if colorize == true {
		return color.RedString(humanize.Bytes(uint64(info.Size())))
	}

	return humanize.Bytes(uint64(info.Size()))
}

func getFileName(info os.FileInfo, colorize bool) string {
	if colorize == true {
		return color.YellowString(info.Name())
	}

	return info.Name()
}

func listFiles(ugly bool, showHidden bool) {
	files, err := ioutil.ReadDir(".")
	writer := tabwriter.NewWriter(os.Stdout, 1, 4, 1, ' ', 0)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !f.IsDir() {
			if ugly {
				fmt.Fprintf(writer, "%s\t%s\t%s\t\n", getFileName(f, false), getFileSize(f, false), getModifedTime(f, false))
			} else {
				fmt.Fprintf(writer, "%s\t%s\t%s\t\n", getFileName(f, true), getFileSize(f, true), getModifedTime(f, true))
			}
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
