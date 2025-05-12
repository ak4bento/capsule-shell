package internal

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type ReadonlyModel struct {
	output string
	ready  bool
}

func NewReadonlyModel(output string) ReadonlyModel {
	return ReadonlyModel{output: output}
}

func (m ReadonlyModel) Init() tea.Cmd {
	return tea.Quit
}

func (m ReadonlyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
		return m, nil
}

func (m ReadonlyModel) View() string {
	return blockStyle.Render(fmt.Sprintf("🧠 Capsule AI:\n\n%s", m.output))
}
