package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var id int

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a contact",
	Long:  "Remove a contact from the DB given its ID",
	Run: func(cmd *cobra.Command, args []string) {
		err := database.InitDB()
		if err != nil {
			fmt.Printf("Could not connect to DB: %v\n", err)
			os.Exit(1)
		}
		err = database.DeleteContact(id)
		if err != nil {
			fmt.Printf("Could not delete contact: %v\n", err)
		} else {
			fmt.Printf("Deleted contact with ID: %v", id)
		}
	},
}

func init() {
	deleteCmd.Flags().IntVarP(&id, "id", "i", 0, "ID of contact to delete")
	rootCmd.AddCommand(deleteCmd)
}
