package main

import (
	"github.com/msanatan/gopractice/clicontacts/cmd"
	"github.com/msanatan/gopractice/clicontacts/db"
)

func main() {
	cmd.SetDB(&db.SQLite{})
	cmd.Execute()
}
