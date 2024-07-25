package cmd

import (
	"fmt"
	"os"

	"github.com/msanatan/gopractice/clicontacts/models"
	"github.com/spf13/cobra"
)

var name, email string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new contact",
	Long:  "Provie the name and email of a new contact",
	Run: func(cmd *cobra.Command, args []string) {
		contact := models.Contact{Name: name, Email: email}
		err := database.InitDB()
		if err != nil {
			fmt.Printf("Could not connect to DB: %v\n", err)
			os.Exit(1)
		}
		err = database.CreateContact(contact)
		if err != nil {
			fmt.Printf("Could not create a contact: %v\n", err)
		} else {
			fmt.Printf("Created a new contact for %v\n", name)
		}
	},
}

func init() {
	createCmd.Flags().StringVarP(&name, "name", "n", "", "Name of your contact")
	createCmd.Flags().StringVarP(&email, "email", "e", "", "Email of your contact")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("email")
	rootCmd.AddCommand(createCmd)
}
