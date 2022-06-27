package ui

import (
	"log"

	"github.com/jroimartin/gocui"
)

// Size of the layout
var (
	firstRun bool = true
)

func StartGUI(g *gocui.Gui) {
	g.SetManagerFunc(layout)
	g.Highlight = true
	g.SelBgColor = gocui.ColorDefault
	g.SelFgColor = gocui.ColorYellow
	g.BgColor = gocui.ColorBlack
	g.FgColor = gocui.ColorCyan
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	// this checks to see if we've logged in and if we haven't it will create a menu showing that we have
	err := loginLayout(g)
	if firstRun {
		firstRun = false
		g.SetCurrentView(LoginPortalAddressView)
	}
	if err != nil {
		return err
	}

	err = statusLayout(g)
	if err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
