package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

func init() {
	//
}

func Execute() error {
	return rootCmd.Execute()
}
