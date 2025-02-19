package bee

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func rgb(r, g, b uint8) lipgloss.Color {
	return lipgloss.Color(fmt.Sprintf("#%02x%02x%02x", r, g, b))
}
