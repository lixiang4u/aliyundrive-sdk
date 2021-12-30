package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "user login",
	Run: func(cmd *cobra.Command, args []string) {

		login()
	},
}

var (
	username string
	password string
)

func init() {
	authCmd.Flags().StringVarP(&username, "username", "u", "", "login phone number")
	authCmd.Flags().StringVarP(&password, "password", "p", "", "login password")
	rootCmd.AddCommand(authCmd)
}

func login() {
	log.Println("dsdsf")
}
