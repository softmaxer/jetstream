package server

import (
	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

type Server struct {
	Songs        []string
	filepicker   filepicker.Model
	selectedSong string
	quitting     bool
	err          error
}

func (server Server) Init() tea.Cmd {
	return nil
}

func (server Server) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (server Server) View() string {
	return "Hello"
}
