/*
Copyright © 2025 dharma <EMAIL ADDRESS>
*/
package delete

import (

	//"gb/cmd"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete repos, branches and pull requests",
	Long:  `Allows user to list all github repositories, branches and pull requests.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("create repo called")
	// },
}

func init() {
	DeleteCmd.AddCommand(repoCmd)
	DeleteCmd.AddCommand(deletebranchCmd)

	////CreateCmd.AddCommand(branchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	////CreateCmd.PersistentFlags().StringP("repo", "r", "", "To create github repositories (required to create branch)")
	////CreateCmd.MarkFlagRequired("repo")

	//createCmd.Flags().StringP("branch", "b", "", "To create github branch for a specified repository")
	//createCmd.MarkFlagRequired("repo")
	//createCmd.MarkFlagsRequiredTogether("repo", "branch")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
