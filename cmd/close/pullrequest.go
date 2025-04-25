/*
Copyright © 2025 dharma <EMAIL ADDRESS>
*/
package close

import (
	"fmt"
	"gb/gbapi"
	"gb/models"

	"github.com/spf13/cobra"
)

// /repos/{owner}/{repo}/pulls/{pull_number}
// closePRCmd represents the delete branch command
var closePRCmd = &cobra.Command{
	Use:     "pull-request",
	Aliases: []string{"pr"},
	Short:   "Close pull request from a branch.",
	Long:    `Allows user to close pull request from a branch.`,
	Run: func(cmd *cobra.Command, args []string) {
		ownerName, _ := cmd.Flags().GetString("owner")
		repoName, _ := cmd.Flags().GetString("repo")
		pullNumber, _ := cmd.Flags().GetString("pull-number")
		title, _ := cmd.Flags().GetString("title")
		body, _ := cmd.Flags().GetString("body")
		state, _ := cmd.Flags().GetString("state")
		base, _ := cmd.Flags().GetString("base")

		client := gbapi.NewRestClient(gbapi.BaseURL, nil)

		var resp any

		closePRReq := models.PRRequest{
			Title: title, Body: body, State: state, Base: base,
		}
		err := client.Patch("/repos/"+ownerName+"/"+repoName+"/pulls/"+pullNumber, closePRReq, resp)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(repoName + ":" + pullNumber + " got closed successfully.")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	closePRCmd.Flags().StringP("owner", "o", "", "Speciy github owner name")
	closePRCmd.Flags().StringP("repo", "r", "", "Specify github repository name (required to delete branch)")
	closePRCmd.Flags().StringP("pull-number", "", "", "Specify number that identifies the pull request.")
	closePRCmd.Flags().StringP("title", "t", "", "Specify title of the pull request.")
	closePRCmd.Flags().StringP("body", "b", "", "Specify contents of the pull request.")
	closePRCmd.Flags().StringP("state", "s", "", "State of this Pull Request. i.e., closed.")
	closePRCmd.Flags().StringP("base", "", "", "Specify name of the branch you want your changes pulled into.")

	closePRCmd.MarkFlagRequired("owner")
	closePRCmd.MarkFlagRequired("repo")
	closePRCmd.MarkFlagRequired("pull-number")
	closePRCmd.MarkFlagRequired("title")
	closePRCmd.MarkFlagRequired("state")
	closePRCmd.MarkFlagsRequiredTogether("owner", "repo", "pull-number", "title", "state")

	//createCmd.MarkFlagsRequiredTogether("repo", "branch")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
