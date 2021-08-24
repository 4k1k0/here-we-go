package db

import "testing"

func TestGetAll(t *testing.T) {
	expectedLen := 3
	res := GetAll()

	if len(res) != expectedLen {
		t.Error("should be a slice of user with 3 elements")
	}
}

func TestLookForID(t *testing.T) {
	t.Run("should get a valid user", func(t *testing.T) {
		res, err := LookForID("01")

		if err != nil {
			t.Error("err shuold be nil")
		}

		if res.Name != "Dot" {
			t.Error("user name should be Dot")
		}
	})

	//t.Run("should get an error if there is no user with the given ID", func(t *testing.T) {
	//	errMessage := "user not found"
	//	res, err := LookForID("fakeUserID")
	//
	//	if err == nil {
	//		t.Error("the fake user should get an error")
	//	}
	//
	//	if err.Error() !=  errMessage {
	//		t.Errorf("error message should be %q", errMessage)
	//	}
	//
	//	if res.Name != "" {
	//		t.Error("user name should be empty")
	//	}
	//})
}
