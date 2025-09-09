package main

type ProgramArgs struct {
	Ugly      bool   `arg:"-u, --ugly" default:"false" help:"Remove colorized output. Yes it's ugly."`
	NoTable   bool   `arg:"-t, --no-table" default:"false" help:"Don't display output as a table."`
	Hidden    bool   `arg:"-a, --all" default:"false" help:"Show hidden files also."`
	NoLinks   bool   `arg:"-l, --no-links" default:"false" help:"Don't mark files as symbolic links."`
	SortBy    string `arg:"-s, --sort-by" default:"modified" help:"Sorts the files by name, size, or modified."`
	SortOrder string `arg:"-o, --sort-order" default:"desc" help:"Sorts the files in ascending(asc) or descending(desc) order."`
}

func (args ProgramArgs) Description() string {
	return "List files of a directory in a human readable format with colorized output optionally included"
}
