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

const (
	defaultKeybar = "<u>pdate si<d>ebar <t>op <b>ot <m>oon <s>un <p>osition"
	D0            = "\x1b[1;38;2;120;120;120m" // ANSI foreground color (= dark)
	M0            = "\x1b[1;38;2;110;110;140m" // ANSI foreground color (= middle)
	L0            = "\x1b[1;38;2;255;255;255m" // ANSI foreground color (= light)
	N0            = "\x1b[0m"                  // ANSI clear formatting.
	R1            = "\x1b[1;38;2;255;0;0m"     // ANSI foreground color (= red).
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
	// print.
	NL = fmt.Sprintln()
)

type model struct {
	keybar  string
	width   int
	height  int
	geoData geo.GeoData
}

func aIfToggleOtherwiseB(toggle bool, a, b string) string {
	if toggle {
		return a
	} else {
		return b
	}
}

// getToggle gets the value of a toggle, but defaults to false.
func (m model) getToggle(toggleStr string) bool {
	str, _ := m.geoData.GetToggle(toggleStr)
	return str
}

// getKeybar highlights each letter <x>, depending on the toggle status (except update, which depends on a web request).
func (m model) getKeybar() string {
	// TODO: highlight each letter <u>pdate depending on if there is a running web request for the update.
	N := N0 + D0
	on := N0 + L0       // on.
	off := N0 + M0      // off.
	disabled := N0 + R1 // off.
	updateC := disabled // TODO: implement.
	sideC := aIfToggleOtherwiseB(m.getToggle("side"), on, off)
	topC := aIfToggleOtherwiseB(m.getToggle("top"), on, off)
	botC := aIfToggleOtherwiseB(m.getToggle("bot"), on, off)
	moonC := disabled // TODO: implement.
	sunC := disabled  // TODO: implement.
	posC := disabled  // TODO: implement.
	return fmt.Sprintf("%s<%su%s>pdate si<%sd%s>ebar <%st%s>op <%sb%s>ot <%sm%s>oon <%ss%s>un <%sp%s>osition"+NL, N, updateC, N, sideC, N, topC, N, botC, N, moonC, N, sunC, N, posC, N)
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
		case "t": // toggle show top.
			m.geoData.Toggle("top")
			return m, nil
		case "b": // toggle show bot.
			m.geoData.Toggle("bot")
			return m, nil
		case "d": // toggle show side.
			m.geoData.Toggle("side")
			return m, nil
		}
	}
	return m, cmd
}

func (m model) View() string {
	var str string
	str += m.getKeybar()
	if infoonly {
		// TODO: implement.
	} else if maponly {
		// TODO: implement.
	} else {
		str += m.geoData.PrintDataHorizontally()
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
	kb := defaultKeybar
	m := model{kb, 0, 0, geo.New()}
	m.geoData.UpdateData()

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
