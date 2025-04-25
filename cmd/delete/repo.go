/*
Copyright © 2025 dharma <EMAIL ADDRESS>
*/
package delete

import (
	"fmt"
	"gb/gbapi"

	"github.com/spf13/cobra"
)

// delete /repos/{owner}/{repo}
// createCmd represents the create command
var repoCmd = &cobra.Command{
	Use:     "repo",
	Aliases: []string{"dr"},
	Short:   "Delete repository",
	Long:    `Allows user to delete github repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		ownerName, _ := cmd.Flags().GetString("owner")
		repoName, _ := cmd.Flags().GetString("repo")

		client := gbapi.NewRestClient(gbapi.BaseURL, nil)

		var resp any
		err := client.Delete("/repos/"+ownerName+"/"+repoName, resp)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(repoName + ": repo got deleted successfully.")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	repoCmd.Flags().StringP("owner", "o", "", "Speciy github owner name")
	repoCmd.Flags().StringP("repo", "r", "", "Specify github repository name(required to delete branch)")
	repoCmd.MarkFlagRequired("owner")
	repoCmd.MarkFlagRequired("repo")
	repoCmd.MarkFlagsRequiredTogether("owner", "repo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
