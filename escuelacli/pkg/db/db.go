// Package db give access to our db and its data.
package db

import (
	"errors"
	"github.com/4k1k0/escuelacli/pkg/user"
)

var schoolDB = map[string]user.User{
	"01": {
		Name: "Dot",
		Notes: []user.Note{
			{ClassName: "Maths", Note: 10},
			{ClassName: "Spanish", Note: 10},
			{ClassName: "Algorithms", Note: 10},
		},
	},
	"02": {
		Name: "Wakko",
		Notes: []user.Note{
			{ClassName: "Maths", Note: 6},
			{ClassName: "Spanish", Note: 6},
			{ClassName: "Algorithms", Note: 6},
		},
	},
	"03": {
		Name: "Yakko",
		Notes: []user.Note{
			{ClassName: "Maths", Note: 8},
			{ClassName: "Spanish", Note: 8},
			{ClassName: "Algorithms", Note: 8},
		},
	},
}

// GetAll returns all the data from our db
func GetAll() []user.User {
	users := make([]user.User, 0)
	for _, u := range schoolDB {
		users = append(users, u)
	}
	return users
}

// LookForID use the ID to return the user data if our database contains a user with that ID
func LookForID(ID string) (user.User, error) {
	if user, ok := schoolDB[ID]; ok {
		return user, nil
	}
	return user.User{}, errors.New("user not found")
}

