// Package table display the info on the terminal and parse the data from the db.
package table

import (
	"fmt"
	"github.com/4k1k0/escuelacli/pkg/db"
	"github.com/4k1k0/escuelacli/pkg/user"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

// PrintByID looks for a student using an ID and display the info on the screen.
//   +-------+-------+---------+------------+
//   | NAME  | MATHS | SPANISH | ALGORITHMS |
//   +-------+-------+---------+------------+
//   | Yakko |  8.00 |    8.00 |       8.00 |
//   +-------+-------+---------+------------+
func PrintByID(ID string) {
	var data [][]string
	userDetails, err := db.LookForID(ID)

	if err != nil {
		log.Printf("User %q not found.\n", ID)
		return
	}

	userData := setUserData(userDetails)
	data = append(data, userData)
	printTable(data)
}

// PrintAll gets all the data from the db and displays this data on the screen.
//   +-------+-------+---------+------------+
//   | NAME  | MATHS | SPANISH | ALGORITHMS |
//   +-------+-------+---------+------------+
//   | Yakko |  8.00 |    8.00 |       8.00 |
//   | Dot   | 10.00 |   10.00 |      10.00 |
//   | Wakko |  6.00 |    6.00 |       6.00 |
//   +-------+-------+---------+------------+
func PrintAll() {
	users := db.GetAll()
	data := setUsersData(users)
	printTable(data)
}

func printTable(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Maths", "Spanish", "Algorithms"})
	table.AppendBulk(data)

	table.Render()
}

func convertNote(note float32) string {
	return fmt.Sprintf("%.2f", note)
}

func setUsersData(users []user.User) [][]string {
	var data [][]string

	for _, u := range users {
		userData := setUserData(u)
		data = append(data, userData)
	}

	return data
}

func setUserData(user user.User) []string {
	name := user.Name
	var data []string
	data = append(data, name)

	for _, note := range user.Notes {
		n := convertNote(note.Note)
		data = append(data, n)
	}
	return data
}
