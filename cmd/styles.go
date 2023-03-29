package cmd

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var blue = lipgloss.NewStyle().Foreground(lipgloss.Color("#0000FF")).Bold(true)
var red = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true)
var white = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true)

func Logo() {
	fmt.Printf("%s%s%s%s%s%s\n",
		white.Render("archiv"), blue.Render("i"), red.Render("i"), blue.Render("i"), red.Render("f"), white.Render("y"))
}
