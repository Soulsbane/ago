package main

type ProgramArgs struct {
	Ugly   bool `arg:"-u, --ugly" default:"false" help:"Remove colorized output. Yes it's ugly."`
	Hidden bool `arg:"-a, --all" default:"false" help:"Show hidden files also."`
	Sort   bool `arg:"-s, --sort" default:"false" help:"Sorts the files by file modification time."`
}
