package cmd

import (
	"github.com/msanatan/gopractice/clicontacts/db"
	"github.com/spf13/cobra"
)

var database db.Database

var rootCmd = &cobra.Command{
	Use:   "contacts",
	Short: "Manage your contacts on the command line",
	Long:  "This is a sample Go app that manages user contacts via the CLI. Data is saved in SQLite",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// This app is simple, so we can just set the SQLite adapter here
	// The interface is better for testing
	database = &db.SQLite{}
}
