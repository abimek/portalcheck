package main

import (
	"errors"

	"github.com/abimek/synergy"
)

var (
	markingPeriodGrades [4]*synergy.GradeBook
	studentInfo         *synergy.StudentInfo
)

var client *synergy.Client

// creates a client and builds all the app data
func createClient(portal string, id int, password string) error {
	client1 := synergy.New(portal, id, password)

	si, err := client1.SchoolInfo()
	if err != nil {
		return err
	}
	if si.School != "" {
		client = &client1
		buildAppData()
		return nil
	}
	return errors.New("issue creating client")
}

// This will take some time
func buildAppData() {
	// create all the gradebooks
	for i := 0; i < 4; i++ {
		pb := synergy.NewParamaterBuilder()
		pb.Add(&synergy.ReportPeriodParamater{Period: i})

		gb, err := client.GradeBook(&pb)
		if err != nil {
			panic(err)
		}
		markingPeriodGrades[i] = gb
	}

	// get student info
	si, err := client.StudentInfo()
	if err != nil {
		panic(err)
	}
	studentInfo = si
}
