package bee

import "github.com/charmbracelet/lipgloss"

var (
	defaultColumnWidth   = ColumnWidth(60)
	defaultWhatColor     = WhatColor(2, 118, 250)
	defaultExpectedColor = ExpectedColor(18, 181, 32)
	defaultActualColor   = ActualColor(250, 40, 25)
	defaultOpts          = []option{
		defaultColumnWidth,
		defaultWhatColor,
		defaultExpectedColor,
		defaultActualColor,
	}
)

type config struct {
	whatTextStyle       lipgloss.Style
	expectedTextStyle   lipgloss.Style
	actualTextStyle     lipgloss.Style
	expectedColumnStyle lipgloss.Style
	actualColumnStyle   lipgloss.Style
}

func newConfig() config {
	cfg := config{}
	for _, opt := range defaultOpts {
		opt(&cfg)
	}
	return cfg
}
