package studentvue

import (
	"errors"

	"github.com/abimek/synergy"
)

var client *synergy.Client

func CreateClient(portal string, id int, password string) error {
	client := synergy.New(portal, id, password)

	si, err := client.SchoolInfo()
	if err != nil {
		return err
	}
	if si.School != "" {
		return nil
	}
	return errors.New("issue creating client")
}
