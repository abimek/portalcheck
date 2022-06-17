package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	portalcheck "portalcheck/src"
)

type User struct {
	credentials portalcheck.Credentials
}

var user User

func main() {
	credentials := request_credentials()
	user := User{credentials}
}

func request_credentials() portalcheck.Credentials {
	var identifier portalcheck.Identifier

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Student Identifer")
		id, _ := reader.ReadString('\n')
		id2, ok := strconv.Atoi(id)

		if ok != nil {
			fmt.Println("Invalid Identifier")
			continue
		}
		identifier = portalcheck.Identifier(id2)
		break
	}

	password, _ := reader.ReadString('\n')

	return portalcheck.Credentials{Identifier: identifier, Password: password}
}
