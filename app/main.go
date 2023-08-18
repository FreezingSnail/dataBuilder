package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	datatypes "github.com/freezingsnail/creatureBuilder/models/dataTypes"
	tableView "github.com/freezingsnail/creatureBuilder/models/table"
)

func main() {
	path := os.Args[1]
	if len(path) == 0 {
		panic("no path")
	}
	t := tableView.TableView{}
	d := datatypes.Creature{}
	err := datatypes.ReadFromFile(d, path)
	if err != nil {
		panic(err)
	}
	t.LoadTableData(d)

	p := tea.NewProgram(t)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
