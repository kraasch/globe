package main

import (
	"flag"
	"fmt"
	"os"

	// for making the TUI.
	tea "github.com/charmbracelet/bubbletea"
	lip "github.com/charmbracelet/lipgloss"

	// local packages.
	geo "github.com/kraasch/geo/pkg/geoview"
)

var (
	// return value.
	output = ""
	// flags.
	maponly  = false
	infoonly = false
	suppress = false
	// styles.
	styleBox = lip.NewStyle().
			BorderStyle(lip.NormalBorder()).
			BorderForeground(lip.Color("56"))
)

type model struct {
	width   int
	height  int
	geoData geo.GeoData
}

func (m model) Init() tea.Cmd {
	return func() tea.Msg { return nil }
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "q": // quit.
			output = "You quit."
			return m, tea.Quit
		case "u": // update.
			m.geoData.UpdateData()
			return m, nil
		}
	}
	return m, cmd
}

func (m model) View() string {
	var str string
	if infoonly {
		// TODO: implement.
	} else if maponly {
		// TODO: implement.
	} else {
		str = m.geoData.PrintDataHorizontally()
	}
	str = styleBox.Render(str) // To add an outer box.
	return lip.Place(m.width, m.height, lip.Center, lip.Center, str)
}

func main() {
	// parse flags.
	flag.BoolVar(&maponly, "maponly", false, "Only show the map section.")
	flag.BoolVar(&infoonly, "infoonly", false, "Only show the info section.")
	flag.BoolVar(&suppress, "suppress", false, "Suppress final output.")
	flag.Parse()

	// init model.
	m := model{0, 0, geo.New()}

	// start bubbletea.
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	// print the last highlighted value in calendar to stdout.
	if !suppress {
		fmt.Println(output)
	}
} // fin.
