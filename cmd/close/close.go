/*
Copyright © 2025 dharma <EMAIL ADDRESS>
*/
package close

import (

	//"gb/cmd"

	"github.com/spf13/cobra"
)

// closeCmd represents the create command
var CloseCmd = &cobra.Command{
	Use:   "close",
	Short: "Close pull requests",
	Long:  `Allows user to close a pull requests.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("create repo called")
	// },
}

func init() {
	CloseCmd.AddCommand(closePRCmd)

}
