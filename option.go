package bee

import "github.com/charmbracelet/lipgloss"

type option func(cfg *config)

func ColumnWidth(w int) option {
	return func(cfg *config) {
		cfg.expectedTextStyle = cfg.expectedTextStyle.MaxWidth(w)
		cfg.actualTextStyle = cfg.actualTextStyle.MaxWidth(w)
		cfg.expectedColumnStyle = cfg.expectedColumnStyle.Width(w).PaddingLeft(1)
		cfg.actualColumnStyle = cfg.actualColumnStyle.Width(w).PaddingRight(1)
	}
}

func NoColor() option {
	return func(cfg *config) {
		cfg.whatTextStyle = cfg.whatTextStyle.Foreground(lipgloss.NoColor{})
		cfg.expectedTextStyle = cfg.expectedTextStyle.Foreground(lipgloss.NoColor{})
		cfg.actualTextStyle = cfg.actualTextStyle.Foreground(lipgloss.NoColor{})
		cfg.expectedColumnStyle = cfg.expectedColumnStyle.Foreground(lipgloss.NoColor{})
		cfg.actualColumnStyle = cfg.actualColumnStyle.Foreground(lipgloss.NoColor{})
	}
}

func WhatColor(r, g, b uint8) option {
	return func(cfg *config) {
		cfg.whatTextStyle = cfg.whatTextStyle.Foreground(rgb(r, g, b))
	}
}

func ExpectedColor(r, g, b uint8) option {
	return func(cfg *config) {
		cfg.expectedTextStyle = cfg.expectedTextStyle.Foreground(rgb(r, g, b))
		cfg.expectedColumnStyle = cfg.expectedColumnStyle.Foreground(rgb(r, g, b))
	}
}

func ActualColor(r, g, b uint8) option {
	return func(cfg *config) {
		cfg.actualTextStyle = cfg.actualTextStyle.Foreground(rgb(r, g, b))
		cfg.actualColumnStyle = cfg.actualColumnStyle.Foreground(rgb(r, g, b))
	}
}
