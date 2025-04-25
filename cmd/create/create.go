/*
Copyright © 2025 dharma <EMAIL ADDRESS>
*/
package create

import (

	//"gb/cmd"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create repos, branches and pull requests",
	Long:  `Allows user to create github repositories, branches and pull requests.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("create repo called")
	// },
}

func init() {
	CreateCmd.AddCommand(createRepoCmd)
	CreateCmd.AddCommand(createbranchCmd)
	CreateCmd.AddCommand(createPRCmd)

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
