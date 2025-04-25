/*
Copyright © 2025 dharma <EMAIL ADDRESS>
*/
package cleanup

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cleanupCmd represents the cleanup command
var cleanupRepoCmd = &cobra.Command{
	Use: "repo",
	//	Aliases: []string{"cr"},
	Short: "cleanup repositories",
	Long:  `Allows user to cleanup github repositories.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cleanup repository called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	cleanupRepoCmd.Flags().StringP("repo", "r", "", "Specify github repository name (required to cleanup branch)")
	cleanupRepoCmd.MarkFlagRequired("repo")
	//createCmd.MarkFlagsRequiredTogether("repo", "branch")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
