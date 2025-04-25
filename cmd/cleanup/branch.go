/*
Copyright © 2025 dharma <EMAIL ADDRESS>
*/
package cleanup

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cleanupCmd represents the cleanup command
var cleanupbranchCmd = &cobra.Command{
	Use: "branch",
	//Aliases: []string{"cb"},
	Short: "Cleanup branches",
	Long:  `Allows user to cleanup github branches.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cleanup branch called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	cleanupbranchCmd.Flags().StringP("repo", "r", "", "Specify github repository name (required to cleanup branch)")

	cleanupbranchCmd.Flags().StringP("branch", "b", "", "Speciy branch name for a specified repository")
	cleanupbranchCmd.MarkFlagsRequiredTogether("repo", "branch")
	cleanupbranchCmd.MarkFlagRequired("repo")
	cleanupbranchCmd.MarkFlagRequired("branch")
	//cleanupCmd.MarkFlagsRequiredTogether("repo", "branch")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
