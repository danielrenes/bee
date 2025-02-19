package bee

import (
	"flag"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

var (
	noColor bool
)

func init() {
	flag.BoolVar(&noColor, "nocolor", false, "Disable color")
}

type Bee struct {
	tb  testing.TB
	cfg config
}

func New(tb testing.TB, opts ...option) *Bee {
	lipgloss.SetColorProfile(termenv.TrueColor)
	cfg := newConfig()
	if noColor {
		opts = append(opts, NoColor())
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	return &Bee{tb: tb, cfg: cfg}
}
