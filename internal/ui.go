package internal

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	output     string
	confirming bool
	approved   bool
	done       bool
}

var (
	green = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Bold(true)
	cyan  = lipgloss.NewStyle().Foreground(lipgloss.Color("81"))
	bold  = lipgloss.NewStyle().Bold(true)
)

var (
	blockStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Margin(1, 0).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#6C91BF")).
			Width(80)

	yesStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("10")).
			Bold(true)

	noStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("1")).
		Bold(true)
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.done = true
			return m, tea.Quit
		case "y":
			if m.confirming {
				m.approved = true
				m.done = true
				return m, tea.Quit
			}
		case "n":
			if m.confirming {
				m.approved = false
				m.done = true
				return m, tea.Quit
			}
		case "enter":
			m.confirming = true
		}
	}
	return m, nil
}

func (m model) View() string {
	if !m.confirming {
		return blockStyle.Render(fmt.Sprintf(
			"🧠 Capsule Shell:\n\n%s\n\n%s\n",
			m.output,
			"Press [Enter] to confirm execution...",
		))
	}
	confirm := "⚠️  Execute this command? (y/n)"
	if m.approved {
		confirm = "🚀 Execute shell command..."
	}

	if !m.approved && m.done {
		confirm = "🚫 Command execution canceled."
	}

	confirmStyled := lipgloss.NewStyle().
		Padding(1).
		Foreground(lipgloss.Color("2")).
		Bold(true).
		Render(confirm)

	return blockStyle.Render(fmt.Sprintf(
		"🧠 Capsule Shell:\n\n%s\n\n%s",
		m.output,
		confirmStyled,
	))
}

func RunUI(output string) bool {
	m := &model{output: output, confirming: false}
	p := tea.NewProgram(m)

	finalModel, err := p.Run()
	if err != nil {
		fmt.Println("TUI failed:", err)
		return false
	}

	// Casting dari interface{} ke *model
	if result, ok := finalModel.(*model); ok {
		return result.approved
	}

	return false
}

func RenderOnly(output string) {
	model := NewReadonlyModel(output)
	p := tea.NewProgram(model)
	_, err := p.Run()
	if err != nil {
		fmt.Println("TUI failed:", err)
		return
	}
}
