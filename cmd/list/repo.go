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
var repoCmd = &cobra.Command{
	Use:     "repo",
	Aliases: []string{"lr"},
	Short:   "List all the repositories",
	Long:    `Allows user to list all github repositories.`,
	//Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		//	orgName := args[1]

		orgName, _ := cmd.Flags().GetString("org")
		//	fmt.Println(orgName)
		client := gbapi.NewRestClient(gbapi.BaseURL, nil)

		// GET  /orgs/{org}/Repos
		var getResponse []models.RepoResponse
		err := client.Get("/orgs/"+orgName+"/repos", &getResponse)
		if err != nil {
			fmt.Println(err)
			return
		}

		// if getResponseresp.StatusCode != http.StatusOK {
		// 	fmt.Printf("GitHub API returned status: %s\n", resp.Status)
		// 	return
		// }

		if len(getResponse) == 0 {
			fmt.Println("No repositories found.")
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Description", "User"})

		for _, repo := range getResponse {
			table.Append([]string{
				repo.Name,
				repo.Description,
				repo.OwnerInfo.Login,
			})
		}
		table.Render()

		//	fmt.Println("GET response:", getResponse)

		//fmt.Println("list repository called")
	},
}

//var org string

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// //get  /orgs/{org}/Repos

	repoCmd.Flags().StringP("org", "o", "", "Specify github organization name")
	repoCmd.MarkFlagRequired("org")
	//createCmd.MarkFlagsRequiredTogether("repo", "branch")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
