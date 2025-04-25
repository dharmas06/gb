/*
Copyright © 2025 dharma <EMAIL ADDRESS>
*/
package cleanup

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cleanupCmd represents the cleanup command
var cleanupPRCmd = &cobra.Command{
	Use: "pull-request",
	//	Aliases: []string{"cpr"},
	Short: "cleanup pull request",
	Long:  `Allows user to cleanup pull requests.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cleanup pull request called")
	},
}

func init() {

	//gh pr cleanup --base main --head feature-branch --title "Your PR title" --body "Detailed

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	cleanupPRCmd.Flags().StringP("repo", "r", "", "Specify github repository name (required to cleanup branch)")

	cleanupPRCmd.Flags().StringP("source-branch", "s", "", "Speciy feature branch name from the repository")
	cleanupPRCmd.Flags().StringP("target-branch", "d", "", "Speciy target/destination feature branch name from the repository")
	cleanupPRCmd.Flags().StringP("title", "t", "", "Speciy pull Request (PR) title")
	cleanupPRCmd.Flags().StringP("message", "b", "", "Speciy detailed description for PR")
	cleanupPRCmd.MarkFlagsRequiredTogether("repo", "source-branch", "target-branch", "title", "message")
	cleanupPRCmd.MarkFlagRequired("repo")
	cleanupPRCmd.MarkFlagRequired("source-branch")
	cleanupPRCmd.MarkFlagRequired("target-branch")
	cleanupPRCmd.MarkFlagRequired("title")
	cleanupPRCmd.MarkFlagRequired("message")
	//cleanupCmd.MarkFlagsRequiredTogether("repo", "branch")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
