/*
Copyright © 2025 dharma <EMAIL ADDRESS>
*/
package list

import (
	"fmt"
	"gb/gbapi"
	"gb/models"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var listbranchCmd = &cobra.Command{
	Use:     "branch",
	Aliases: []string{"lb"},
	Short:   "List all branches for a repository",
	Long:    `Allows user to list all github branches for a repository.`,
	Run: func(cmd *cobra.Command, args []string) {

		//	orgName := args[1]
		ownerName, _ := cmd.Flags().GetString("owner")
		repoName, _ := cmd.Flags().GetString("repo")

		client := gbapi.NewRestClient(gbapi.BaseURL, nil)

		// GET  /orgs/{org}/Repos
		var getBranchResponse []models.ListBranchresponse
		// repos/gbuser/gbrepo/branches
		err := client.Get("/repos/"+ownerName+"/"+repoName+"/branches", &getBranchResponse)
		if err != nil {
			fmt.Println(err)
			return
		}

		// if getResponseresp.StatusCode != http.StatusOK {
		// 	fmt.Printf("GitHub API returned status: %s\n", resp.Status)
		// 	return
		// }

		if len(getBranchResponse) == 0 {
			fmt.Println("No repositories found.")
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name"})

		for _, repo := range getBranchResponse {
			table.Append([]string{
				repo.Name,
			})
		}
		table.Render()

		//	fmt.Println("GET response:", getResponse)

		//fmt.Println("list repository called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	listbranchCmd.Flags().StringP("owner", "o", "", "Speciy owner name")

	listbranchCmd.Flags().StringP("repo", "r", "", "Specify github repository name (required to list branch))")
	listbranchCmd.MarkFlagsRequiredTogether("repo", "owner")
	listbranchCmd.MarkFlagRequired("repo")
	listbranchCmd.MarkFlagRequired("owner")
	//createCmd.MarkFlagsRequiredTogether("repo", "branch")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
