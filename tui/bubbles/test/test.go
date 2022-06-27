package test

// A simple program demonstrating the spinner component from the Bubbles
// component library.

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding  = 2
	maxWidth = 80
)

var pad = strings.Repeat(" ", padding)
var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

var (
	subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	good   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	bad    = lipgloss.AdaptiveColor{Light: "#bf2a34", Dark: "#ef1a28"}

	listHeader = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(subtle).
			MarginRight(2).
			Render

	listItem = lipgloss.NewStyle().PaddingLeft(2).Render

	progressMark = func(s string) string {
		return lipgloss.NewStyle().SetString(s).
			PaddingRight(1).
			String()
	}
	checkMark = lipgloss.NewStyle().SetString("✓").
			Foreground(good).
			PaddingRight(1).
			String()

	xMark = lipgloss.NewStyle().SetString("✘").
		Foreground(bad).
		PaddingRight(1).
		String()

	listRunning = func(char, s string) string {
		return char + s
	}

	listFail = func(s string) string {
		return line(s, xMark, "#bf2a34")
	}

	listSuccess = func(s string) string {
		return line(s, checkMark, "#43BF6D")
	}

	line = func(s, icon string, color lipgloss.Color) string {
		return icon + lipgloss.NewStyle().
			//Strikethrough(true).
			Foreground(color).
			Render(s)
	}
)

type errMsg error

type EventState int

const (
	Hidden EventState = iota
	Pending
	Running
	Success
	Fail
)

type Event struct {
	Msg     string
	Visible time.Duration
	Start   time.Duration
	Dur     time.Duration
	state   EventState
	Final   EventState
}

type EventConfig struct {
	Title  string
	Events []Event
}

type model struct {
	start    time.Time
	config   EventConfig
	done     uint64
	spinner  spinner.Model
	quitting bool
	err      error
}

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{spinner: s, start: time.Now()}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) handleEvents() (events []Event, done bool) {
	events = m.config.Events
	runtime := time.Since(m.start)
	doneCount := 0
	for i, evt := range events {
		newState := Hidden
		if evt.Start+evt.Dur < runtime {
			newState = evt.Final
			doneCount++
		} else if evt.Start < runtime {
			newState = Running
		} else if evt.Visible < runtime {
			newState = Pending
		}
		events[i].state = newState
	}

	done = len(events) == doneCount

	return
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		m.quitting = true
		return m, tea.Quit

	case errMsg:
		m.err = msg
		return m, nil

	case spinner.TickMsg:
		m.config.Events, m.quitting = m.handleEvents()
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		if m.quitting {
			return m, tea.Quit
		}
		return m, cmd

	default:
		return m, nil
	}
}

func (m model) View() string {
	postfix := ""
	if m.quitting {
		postfix = ""
	} else {
		postfix = pad + helpStyle("Press any key to quit")
	}

	buf := strings.Builder{}

	spin := m.spinner.View()
	for _, evt := range m.config.Events {
		switch evt.state {
		case Pending:
			buf.WriteString(listItem(evt.Msg))
			buf.WriteString("\n")
		case Running:
			buf.WriteString(listRunning(spin, evt.Msg))
			buf.WriteString("\n")
		case Success:
			buf.WriteString(listSuccess(evt.Msg))
			buf.WriteString("\n")
		case Fail:
			buf.WriteString(listFail(evt.Msg))
			buf.WriteString("\n")
		case Hidden:
		}
	}

	return listHeader(m.config.Title) + "\n" +
		buf.String() + "\n" +
		postfix
}

func Test(config EventConfig) {
	m := initialModel()
	m.config = config
	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
