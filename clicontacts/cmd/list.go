package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List contacts",
	Long:  "Provie the name and email of all contacts in the DB",
	Run: func(cmd *cobra.Command, args []string) {
		err := database.InitDB()
		if err != nil {
			fmt.Printf("Could not connect to DB: %v\n", err)
			os.Exit(1)
		}
		contacts, err := database.ListContacts()
		if err != nil {
			fmt.Printf("Could not create a contact: %v\n", err)
		} else {
			fmt.Println("ID\t\tName\t\t\tEmail")
			fmt.Println("_____________________________________________")
			for _, contact := range contacts {
				fmt.Printf("%v.\t\t%v\t\t%v\n", contact.ID, contact.Name, contact.Email)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
