/*
Copyright © 2025 dharma <EMAIL ADDRESS>
*/
package delete

import (
	"fmt"
	"gb/gbapi"

	"github.com/spf13/cobra"
)

// /repos/{owner}/{repo}/git/refs/{ref}
// deletebranchCmd represents the delete branch command
var deletebranchCmd = &cobra.Command{
	Use:     "branch",
	Aliases: []string{"db"},
	Short:   "Delete branches for a repository",
	Long:    `Allows user to delete github branches for a specified repository.`,
	Run: func(cmd *cobra.Command, args []string) {

		ownerName, _ := cmd.Flags().GetString("owner")
		repoName, _ := cmd.Flags().GetString("repo")
		ref, _ := cmd.Flags().GetString("ref")

		client := gbapi.NewRestClient(gbapi.BaseURL, nil)

		var resp any
		err := client.Delete("/repos/"+ownerName+"/"+repoName+"/git/refs/"+ref, resp)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(repoName + ":" + ref + " got deleted successfully.")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	deletebranchCmd.Flags().StringP("owner", "o", "", "Speciy github owner name")
	deletebranchCmd.Flags().StringP("repo", "r", "", "Specify github repository name (required to delete branch)")
	deletebranchCmd.Flags().StringP("ref", "", "", "Speciy reference/branch name (ie: master).")
	deletebranchCmd.MarkFlagRequired("owner")
	deletebranchCmd.MarkFlagRequired("repo")
	deletebranchCmd.MarkFlagRequired("ref")
	deletebranchCmd.MarkFlagsRequiredTogether("owner", "repo", "ref")

	//createCmd.MarkFlagsRequiredTogether("repo", "branch")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
