/*
Copyright © 2025 dharma <EMAIL ADDRESS>
*/
package create

import (
	"fmt"
	"gb/gbapi"
	"gb/models"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// post /repos/{owner}/{repo}/pulls
// createCmd represents the create command
var createPRCmd = &cobra.Command{
	Use:     "pull-request",
	Aliases: []string{"cpr"},
	Short:   "Create pull request",
	Long:    `Allows user to create pull requests.`,
	Run: func(cmd *cobra.Command, args []string) {

		ownerName, _ := cmd.Flags().GetString("owner")
		repoName, _ := cmd.Flags().GetString("repo")
		title, _ := cmd.Flags().GetString("title")
		body, _ := cmd.Flags().GetString("body")
		head, _ := cmd.Flags().GetString("head")
		base, _ := cmd.Flags().GetString("base")

		client := gbapi.NewRestClient(gbapi.BaseURL, nil)

		var createPRResp models.PRResponse

		createPRReq := models.PRRequest{
			Title: title,
			Body:  body,
			Head:  head,
			Base:  base,
		}

		err := client.Post("/repos/"+ownerName+"/"+repoName+"/pulls", createPRReq, &createPRResp)

		if err != nil {
			fmt.Println(err)
			return
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Pull-request ID", "Name", "Repo Name", "Base Branch", "Head Branch", "PR Title", "State"})

		//for _, repo := range createRepoResp {
		pullID := strconv.Itoa(createPRResp.ID)
		table.Append([]string{
			pullID,
			createPRResp.User.Login,
			createPRResp.Base.Repo.Name,
			createPRResp.Base.Ref,
			createPRResp.Head.Ref,
			createPRResp.Title,
			createPRResp.State,
		})
		//}
		table.Render()

	},
}

func init() {

	//gh pr create --base main --head feature-branch --title "Your PR title" --body "Detailed

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	createPRCmd.Flags().StringP("owner", "o", "", "Speciy github owner name")
	createPRCmd.Flags().StringP("repo", "r", "", "Specify github repository name (required to delete branch)")

	createPRCmd.Flags().StringP("title", "t", "", "Specify title of the new pull request.")
	createPRCmd.Flags().StringP("body", "b", "", "Specify contents of the pull request.")
	createPRCmd.Flags().StringP("head", "", "", "The name of the branch where your changes are implemented.(ie.,username:branch)")
	createPRCmd.Flags().StringP("base", "", "", "Specify name of the branch you want your changes pulled into.")
	createPRCmd.MarkFlagRequired("owner")
	createPRCmd.MarkFlagRequired("repo")
	createPRCmd.MarkFlagRequired("title")
	createPRCmd.MarkFlagRequired("base")
	createPRCmd.MarkFlagRequired("head")
	createPRCmd.MarkFlagRequired("body")
	createPRCmd.MarkFlagsRequiredTogether("owner", "repo", "title", "body", "head", "base")

	//createCmd.MarkFlagsRequiredTogether("repo", "branch")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
