package table

import (
	"github.com/4k1k0/escuelacli/pkg/user"
	"testing"
)

func createFakeUser(t *testing.T, name string, note float32) user.User {
	t.Helper()
	
	return user.User{
		Name:  name,
		Notes: []user.Note{
			{ClassName: "Maths", Note: note},
			{ClassName: "Spanish", Note: note},
			{ClassName: "Algorithms", Note: note},
		},
	}
}

func TestSetUserData(t *testing.T) {
	name := "Pinky"
	user := createFakeUser(t, name, 0)

	res := setUserData(user)

	if len(res) != 4 {
		t.Error("len should be of 4 items")
	}

	if res[0] != name {
		t.Errorf("index zero should be %q", name)
	}
}

func TestUsersData(t *testing.T) {
	slappy := "Slappy"
	skippy := "Skippy"
	users := []user.User{
		createFakeUser(t, slappy, 8),
		createFakeUser(t, skippy, 9),
	}

	res := setUsersData(users)

	if len(res) != 2 {
		t.Error("res should be a matrix with 2 slices")
	}

	if res[0][0] != slappy {
		t.Errorf("res[0][0] should be %q", slappy)
	}

	if res[1][0] != skippy {
		t.Errorf("res[1][0] should be %q", skippy)
	}
}

func TestConvertNote(t *testing.T) {
	res := convertNote(10)
	if res != "10.00" {
		t.Error("note should be 10.00")
	}
}

func TestPrint(t *testing.T) {
	var data [][]string
	PrintAll()
	printTable(data)
}

func TestPrintByID(t *testing.T) {
	t.Run("print a user data", func(t *testing.T) {
		PrintByID("01")
	})
	t.Run("user not found", func(t *testing.T) {
		PrintByID("fakeUSER")
	})
}
