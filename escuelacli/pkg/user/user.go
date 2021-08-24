// Package user only contains the basic structs for our students.
package user

// User reprensents a student from this school.
type User struct     {
	Name   string
	Notes     []Note
}

// Note represents the grade obtained by a student in a specific subject.
type Note struct    {
	ClassName                 string
	Note  float32
}
