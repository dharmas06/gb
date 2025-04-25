/*
Copyright © 2025 dharma <EMAIL ADDRESS>
*/
package create

import (
	"fmt"
	"gb/gbapi"
	"gb/models"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// /orgs/{org}/repos
// createCmd represents the create command
var createRepoCmd = &cobra.Command{
	Use:     "repo",
	Aliases: []string{"cr"},
	Short:   "Create repositories",
	Long:    `Allows user to create github repositories.`,
	Run: func(cmd *cobra.Command, args []string) {
		orgName, _ := cmd.Flags().GetString("org")
		repoName, _ := cmd.Flags().GetString("repo")
		description, _ := cmd.Flags().GetString("desc")

		client := gbapi.NewRestClient(gbapi.BaseURL, nil)

		var createRepoResp models.RepoResponse
		repoReq := models.CreateRepoRequest{
			Name:        repoName,
			Description: description,
		}

		err := client.Post("/orgs/"+orgName+"/repos", repoReq, &createRepoResp)
		if err != nil {
			fmt.Println(err)
			return
		}

		// if getResponseresp.StatusCode != http.StatusOK {
		// 	fmt.Printf("GitHub API returned status: %s\n", resp.Status)
		// 	return
		// }

		// if len(createRepoResp) == 0 {
		// 	fmt.Println("No repositories found.")
		// 	return
		// }

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Description", "User"})

		//for _, repo := range createRepoResp {
		table.Append([]string{
			createRepoResp.Name,
			createRepoResp.Description,
			createRepoResp.OwnerInfo.Login,
		})
		//}
		table.Render()

	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	createRepoCmd.Flags().StringP("org", "o", "", "Specify github organization name (required to create branch)")
	createRepoCmd.Flags().StringP("repo", "r", "", "Specify github repository name (required to create branch)")
	createRepoCmd.Flags().StringP("desc", "d", "", "Specify github repository description (required to create branch)")
	createRepoCmd.MarkFlagRequired("org")
	createRepoCmd.MarkFlagRequired("repo")
	createRepoCmd.MarkFlagRequired("desc")
	createRepoCmd.MarkFlagsRequiredTogether("repo", "org", "desc")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
