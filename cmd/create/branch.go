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

// createCmd represents the create command
var createbranchCmd = &cobra.Command{
	Use:     "branch",
	Aliases: []string{"cb"},
	Short:   "Create branch",
	Long:    `Allows user to create github branch and pull request.`,
	Run: func(cmd *cobra.Command, args []string) {

		ownerName, _ := cmd.Flags().GetString("owner")
		repoName, _ := cmd.Flags().GetString("repo")
		ref, _ := cmd.Flags().GetString("ref")
		sha, _ := cmd.Flags().GetString("sha")

		client := gbapi.NewRestClient(gbapi.BaseURL, nil)

		var createBranchResp models.CreateBranchResponse
		// {"ref":"refs/heads/featureA","sha":"sdh2h3n2id82hsal891f9e74264a383fa43fefbd"}
		createBranchReq := models.CreateBranchRequest{
			Ref: ref,
			SHA: sha,
		}

		err := client.Post("/repos/"+ownerName+"/"+repoName+"/git/refs", createBranchReq, &createBranchResp)

		if err != nil {
			fmt.Println(err)
			return
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Branch Name", "URL", "SHA"})

		//for _, repo := range createRepoResp {
		table.Append([]string{
			createBranchResp.Ref,
			createBranchResp.URL,
			createBranchResp.Object.SHA,
		})
		//}
		table.Render()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	//post /Repos/{owner}/{Repo}/git/refs
	// {"ref":"refs/heads/featureA","sha":"aa218f56b14c9653891f9e74264a383fa43fefbd"}
	createbranchCmd.Flags().StringP("repo", "r", "", "Specify github repository name (required to create branch)")
	createbranchCmd.Flags().StringP("owner", "o", "", "Speciy github owner name")
	createbranchCmd.Flags().StringP("ref", "", "", "Speciy name of the fully qualified reference(ie: refs/heads/master).")
	createbranchCmd.Flags().StringP("sha", "", "", "Speciy SHA1 value for this reference")

	createbranchCmd.MarkFlagsRequiredTogether("repo", "owner", "ref", "sha")
	createbranchCmd.MarkFlagRequired("repo")
	createbranchCmd.MarkFlagRequired("owner")
	createbranchCmd.MarkFlagRequired("ref")
	createbranchCmd.MarkFlagRequired("sha")
	//createCmd.MarkFlagsRequiredTogether("repo", "branch")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
