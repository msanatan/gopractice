package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create DB and table",
	Long:  "Creates the SQLite DB and the contacts table",
	Run: func(cmd *cobra.Command, args []string) {
		database.InitDB()
		fmt.Println("Database initialised")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
