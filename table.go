package main

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

var headers []string = []string{"Id", "Name", "Description", "Language", "Link"}

func createTable(r []Repo) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.SetRowLine(true)
	table.SetRowSeparator("-")

	for _, v := range r {
		item := []string{
			strconv.Itoa(v.Id), v.Name, v.Description, v.Language, v.Html_url,
		}
		table.Append(item)
	}

	table.Render()
}
