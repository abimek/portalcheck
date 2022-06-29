package main

import (
	"fmt"
	"strconv"

	"github.com/jroimartin/gocui"
)

const (
	boardSessionInfoView    = "board_session_info"
	boardMarkingPeriodsView = "board_marking_periods"
	boardClassListView      = "board_class_list"
)

func boardLayout(gui *gocui.Gui) error {
	if !userSession.LoggedIn {
		return nil
	}
	maxX, maxY := gui.Size()

	if v, err := gui.SetView(boardSessionInfoView, 1, 1, maxX/7, maxY/8); err != nil {
		if err != gocui.ErrUnknownView {
			return nil
		}
		v.Title = "Session Info"

		// we print the Session Info
		id := studentInfo.PermID
		name := studentInfo.FormattedName
		grade := studentInfo.Grade
		homeroom := studentInfo.HomeRoom
		school := studentInfo.CurrentSchool
		fmt.Fprintln(v, "\u001b[31;1mName: \u001b[34m"+name)
		fmt.Fprintln(v, "\u001b[31;1mID: \u001b[34m"+strconv.Itoa(id))
		fmt.Fprintln(v, "\u001b[31;1mGRADE: \u001b[34m"+strconv.Itoa(int(grade)))
		fmt.Fprintln(v, "\u001b[31;1mHome Room: \u001b[34m"+homeroom)
		fmt.Fprintln(v, "\u001b[31;1mSCHOOL: \u001b[34m"+school+"\u001b[32m")
	}
	if v, err := gui.SetView(boardMarkingPeriodsView, 1, maxY/8+1, maxX/7, maxY/4); err != nil {
		if err != gocui.ErrUnknownView {
			return nil
		}
		v.Title = "Marking Periods"
		markingPeriods := []string{"MP1", "MP2", "MP3", "MP4"}

		for _, mp := range markingPeriods {
			fmt.Fprintln(v, "\u001b[32m"+mp)
		}
	}
	if v, err := gui.SetView(boardClassListView, 1, maxY/4+1, maxX/7, maxY/2); err != nil {
		if err != gocui.ErrUnknownView {
			return nil
		}
		v.Title = "Classes"

		gradebook := markingPeriodGrades[userSession.currentMarkingPeriod]
		for _, course := range gradebook.Courses {
			fmt.Fprintln(v, "\u001b[32m"+course.Title)
		}
	}
	if firstBoardRun {
		firstBoardRun = false
		gui.SetCurrentView(boardMarkingPeriodsView)
	}
	return nil
}
