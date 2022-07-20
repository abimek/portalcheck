package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jroimartin/gocui"
)

type loginInfo struct {
	portal   string
	id       int
	password string
}

const (
	loginOutlineView       = "login_layout"
	loginPortalAddressView = "login_portal_address"
	loginIDView            = "login_id"
	loginPasswordView      = "login_password"
)

var loginViews = []string{loginOutlineView, loginPortalAddressView, loginIDView, loginPasswordView}

func loginLayout(g *gocui.Gui) error {
	// return if we've already logged in so the login View doesn't show up
	if userSession.LoggedIn {
		return nil
	}
	// This is the outline for the login view
	maxX, maxY := g.Size()
	if v, err := g.SetView(loginOutlineView, maxX/3, maxY/3, maxX-maxX/3, maxY-maxY/3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.FgColor = gocui.ColorMagenta
		v.Title = "Login Menu"
		fmt.Fprintln(v, "\u001b[31mPlease Login...")
	}

	if v, err := g.SetView(loginPortalAddressView, maxX/3+2, maxY/3+3, maxX-maxX/3-2, maxY/3+6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if err := g.SetKeybinding(loginPortalAddressView, gocui.KeyArrowDown, gocui.ModNone, keyDownLogin); err != nil {
			return err
		}
		if err := g.SetKeybinding(loginPortalAddressView, gocui.KeyEnter, gocui.ModNone, keyEnterLogin); err != nil {
			return err
		}
		v.Title = "Portal Address"
		v.Editable = true
	}

	if v, err := g.SetView(loginIDView, maxX/3+2, maxY/3+9, maxX-maxX/3-2, maxY/3+12); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if err := g.SetKeybinding(loginIDView, gocui.KeyArrowDown, gocui.ModNone, keyDownLogin); err != nil {
			return err
		}
		if err := g.SetKeybinding(loginIDView, gocui.KeyArrowUp, gocui.ModNone, keyUpLogin); err != nil {
			return err
		}
		if err := g.SetKeybinding(loginIDView, gocui.KeyEnter, gocui.ModNone, keyEnterLogin); err != nil {
			return err
		}
		v.Editable = true
		v.Title = "Student ID"
	}

	if v, err := g.SetView(loginPasswordView, maxX/3+2, maxY/3+15, maxX-maxX/3-2, maxY/3+18); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if err := g.SetKeybinding(loginPasswordView, gocui.KeyArrowUp, gocui.ModNone, keyUpLogin); err != nil {
			return err
		}
		if err := g.SetKeybinding(loginPasswordView, gocui.KeyEnter, gocui.ModNone, keyEnterLogin); err != nil {
			return err
		}
		v.Title = "Password"
		v.Editable = true
		v.Mask = '*'
	}
	return nil
}

// bidngs for whem someone goes down in the login menu
var loginDownBindings = map[string]string{
	loginPortalAddressView: loginIDView,
	loginIDView:            loginPasswordView,
}

func keyDownLogin(gui *gocui.Gui, view *gocui.View) error {
	for k, v := range loginDownBindings {
		if view.Name() == k {
			gui.SetCurrentView(v)
			return nil
		}
	}
	return nil
}

// Bindings for when someone goes up in the Login Menu
var loginUpBindings = map[string]string{
	loginPasswordView: loginIDView,
	loginIDView:       loginPortalAddressView,
}

// GO DOWN THE LOGIN MENU WITH ARROW KEYS
func keyUpLogin(gui *gocui.Gui, view *gocui.View) error {
	for k, v := range loginUpBindings {
		if view.Name() == k {
			gui.SetCurrentView(v)
			return nil
		}
	}
	return nil
}

// bindings for when someone clicks enter in a menu
var loginEnterBindings = map[string]string{
	loginPortalAddressView: loginIDView,
	loginIDView:            loginPasswordView,
}

// When They want to login
func keyEnterLogin(gui *gocui.Gui, view *gocui.View) error {
	for v, k := range loginEnterBindings {
		if view.Name() == v {
			if view.Buffer() != "" {
				gui.SetCurrentView(k)
				return nil
			}
			createStatus(statusInfo{"section is empty", 3, gui})
		}
	}
	// AUTOMATIC DETECTIONF OR PASSWORD TO CHECK IF THEY've LOGGED IN
	if view.Name() == loginPasswordView {
		views := []string{loginPortalAddressView, loginIDView, loginPasswordView}
		var logininfo [3]string
		for i, mapview := range views {
			v, err := gui.View(mapview)
			if err != nil {
				return err
			}
			if v.Buffer() == "" {
				createStatus(statusInfo{"you have an empty section", 3, gui})
				return nil
			}
			// we have to trim off the suffix
			logininfo[i] = strings.TrimSuffix(v.Buffer(), "\n")
		}
		// we check to see if the id is an int and if not we delete the ID field
		id, err := strconv.Atoi(logininfo[1])
		if err != nil {
			v, err := gui.View(loginIDView)
			if err != nil {
				return err
			}
			fmt.Fprint(v, "")
			createStatus(statusInfo{"You're ID is not a number", 4, gui})
			return nil
		}
		login(loginInfo{portal: logininfo[0], id: id, password: logininfo[2]}, gui)
	}
	return nil
}

// This method is called once the login menu is submitted, one the data is valid
func login(loginInfo loginInfo, gui *gocui.Gui) {
	err := createClient(loginInfo.portal, loginInfo.id, loginInfo.password)
	if err != nil {
		createStatus(statusInfo{"Failed to log in", 3, gui})
		return
	}
	deleteLoginLayout(gui)
	// we set it so that they've logged in so that the board menu shows up
	userSession.LoggedIn = true
	// TODO: Send To Main Menu
}

// DELETE ALL THE LOGIN UI's
func deleteLoginLayout(g *gocui.Gui) error {
	for _, v := range loginViews {
		g.DeleteKeybindings(v)
		g.DeleteView(v)
	}
	return nil
}
