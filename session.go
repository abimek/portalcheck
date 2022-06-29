package main

type session struct {
	LoggedIn             bool
	currentMarkingPeriod int
}

var userSession session
