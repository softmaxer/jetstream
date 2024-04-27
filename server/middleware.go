package server

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/softmaxer/jetstream/views"
)

func TeaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, _ := s.Pty()
	renderer := bubbletea.MakeRenderer(s)
	txtStyle := renderer.NewStyle().Foreground(lipgloss.Color("10"))
	quitStyle := renderer.NewStyle().Foreground(lipgloss.Color("8"))
	m := views.Model{
		Term:      pty.Term,
		Width:     pty.Window.Width,
		Height:    pty.Window.Height,
		TxtStyle:  txtStyle,
		QuitStyle: quitStyle,
	}
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}
