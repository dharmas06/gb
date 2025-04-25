/*
Copyright © 2025 dharma <EMAIL ADDRESS>
*/
package list

import (
	"fmt"
	"gb/gbapi"
	"gb/models"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// listPRCmd represents the list PRs command
var listPRCmd = &cobra.Command{
	Use:     "pull-request",
	Aliases: []string{"lpr"},
	Short:   "List all pull request for branch from a repository.",
	Long:    `Allows user to list all pull request for branch from a repository.`,
	Run: func(cmd *cobra.Command, args []string) {

		ownerName, _ := cmd.Flags().GetString("owner")
		repoName, _ := cmd.Flags().GetString("repo")

		client := gbapi.NewRestClient(gbapi.BaseURL, nil)

		// GET  /orgs/{org}/Repos
		var getPRsResponse []models.PRResponse
		// repos/gbuser/gbrepo/branches
		err := client.Get("/repos/"+ownerName+"/"+repoName+"/pulls", &getPRsResponse)
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(getPRsResponse) == 0 {
			fmt.Println("No PRs found.")
			return
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Pull-request ID", "Name", "Repo Name", "Base Branch", "Head Branch", "PR Title", "State"})
		for _, data := range getPRsResponse {
			pullID := strconv.Itoa(data.ID)

			table.Append([]string{
				pullID,
				data.User.Login,
				data.Base.Repo.Name,
				data.Base.Ref,
				data.Head.Ref,
				data.Title,
				data.State,
			})
		}
		table.Render()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	/// repos/gbuser/gbrepo/pulls
	listPRCmd.Flags().StringP("owner", "o", "", "Speciy owner name")

	listPRCmd.Flags().StringP("repo", "r", "", "Specify github repository name (required to list branch))")
	listPRCmd.MarkFlagsRequiredTogether("repo", "owner")
	listPRCmd.MarkFlagRequired("repo")
	listPRCmd.MarkFlagRequired("owner")
	//createCmd.MarkFlagsRequiredTogether("repo", "branch")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
