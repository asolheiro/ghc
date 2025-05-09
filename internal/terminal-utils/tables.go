package terminalutils

import (
	"os"

	table "github.com/jedib0t/go-pretty/v6/table"
)

func PrettyTables(rows []table.Row) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{
		"Name",
		"ClusterID",
		"CPU (%)",
		"RAM (%)",
		"PODS ",
		"Alerts",
		"Incidents",
		"Security",
		"Problems",
	})

	t.AppendRows(rows)

	t.SetStyle(table.StyleBold)
	t.SetStyle(table.StyleColoredBlackOnYellowWhite)
	t.Render()
}
