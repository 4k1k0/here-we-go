// Package cmd implements our business logic.
package cmd

import (
	"flag"
	"github.com/4k1k0/escuelacli/pkg/table"
)

func Execute() {
	id := flag.String("id", "fake user ID", "Student's ID to look on the db.")
	showAll := flag.Bool("all", false, "Show the notes from all students.")

	flag.Parse()

	if *showAll {
		table.PrintAll()
		return
	}

	table.PrintByID(*id)
}
