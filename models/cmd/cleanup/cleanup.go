/*
Copyright © 2025 dharma <EMAIL ADDRESS>
*/
package cleanup

import (

	//"gb/cmd"

	"github.com/spf13/cobra"
)

// cleanupCmd represents the cleanup command
var CleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "cleanup repos, branches and pull requests",
	Long:  `Allows user to cleanup github repositories, branches and pull requests.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("cleanup repo called")
	// },
}

func init() {
	CleanupCmd.AddCommand(cleanupRepoCmd)
	CleanupCmd.AddCommand(cleanupbranchCmd)
	CleanupCmd.AddCommand(cleanupPRCmd)

	////cleanupCmd.AddCommand(branchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	////cleanupCmd.PersistentFlags().StringP("repo", "r", "", "To cleanup github repositories (required to cleanup branch)")
	////cleanupCmd.MarkFlagRequired("repo")

	//cleanupCmd.Flags().StringP("branch", "b", "", "To cleanup github branch for a specified repository")
	//cleanupCmd.MarkFlagRequired("repo")
	//cleanupCmd.MarkFlagsRequiredTogether("repo", "branch")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
