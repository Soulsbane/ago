package main

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

var (
	//  INFO: This is for the default bordered table style.
	defaultTableOptions = table.Options{
		DoNotColorBordersAndSeparators: false,
		DrawBorder:                     true,
		SeparateColumns:                true,
		SeparateFooter:                 true,
		SeparateHeader:                 true,
		SeparateRows:                   false,
	}

	defaultTitleOptions = table.TitleOptions{
		Align: text.AlignLeft,
	}

	agoDefaultStyle = table.Style{
		Name:    "AgoDefaultStyle",
		Box:     table.StyleBoxRounded,
		Color:   table.ColorOptionsDefault,
		Format:  table.FormatOptionsDefault,
		HTML:    table.DefaultHTMLOptions,
		Options: defaultTableOptions,
		Title:   defaultTitleOptions,
	}

	//  INFO: This is for the --no-table option. It will still appear aligned but borderless.
	noTableOptions = table.Options{
		DoNotColorBordersAndSeparators: false,
		DrawBorder:                     false,
		SeparateColumns:                false,
		SeparateFooter:                 false,
		SeparateHeader:                 false,
		SeparateRows:                   false,
	}

	agoNoStyle = table.Style{
		Name:    "AgoNoStyle",
		Box:     table.StyleBoxRounded,
		Color:   table.ColorOptionsDefault,
		Format:  table.FormatOptionsDefault,
		HTML:    table.DefaultHTMLOptions,
		Options: noTableOptions,
		Title:   defaultTitleOptions,
	}
)
