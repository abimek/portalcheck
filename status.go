package main

import (
	"fmt"
	"time"

	"github.com/jroimartin/gocui"
)

type statusInfo struct {
	message  string
	duration int
	gui      *gocui.Gui
}

const (
	statusView = "error_view"
)

var currentStatus *statusInfo = nil

func statusLayout(g *gocui.Gui) error {
	status := currentStatus
	if currentStatus != nil {
		maxX, _ := g.Size()
		if v, err := g.SetView(statusView, maxX/3+2, 5, maxX-maxX/3-2, 10); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			v.FgColor = gocui.ColorYellow
			fmt.Fprintln(v, "\u001b[31m"+status.message)
			v.Title = "Status"
			v.Editable = true
		}

	}
	return nil
}

// duration is in seconds
func createStatus(status statusInfo) error {
	g := status.gui
	v, err := g.View(statusView)
	if err == nil {
		fmt.Fprint(v, "\u001b[31m"+status.message)
		return nil
	}

	maxX, _ := g.Size()
	if v, err := g.SetView(statusView, maxX/2-15, 5, maxX/2+15, 7); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.FgColor = gocui.ColorYellow
		fmt.Fprintln(v, "\u001b[31m"+status.message)
		v.Title = "Status"
		v.Editable = true
	}

	// scheduale the error to be removed
	go func(g *gocui.Gui, duration int) {
		time.Sleep(time.Duration(duration) * time.Second)
		g.Update(func(gui *gocui.Gui) error {
			// wait before removing
			v, err := gui.View(statusView)
			if err != nil {
				return err
			}
			gui.DeleteView(v.Name())
			currentStatus = nil
			return nil
		})
	}(g, status.duration)

	return nil
}
