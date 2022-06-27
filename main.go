package main

import (
	"log"

	"portalcheck/session"
	"portalcheck/ui"

	"github.com/jroimartin/gocui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	// create session
	session.UserSession = session.Session{LoggedIn: false}

	ui.StartGUI(g)
}
