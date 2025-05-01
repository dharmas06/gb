/*
Copyright © 2025 dharma
*/
package cmd

import (
	// _ "gb/cmd/create"
	// _ "gb/cmd/delete"
	// _ "gb/cmd/list"
	"gb/cmd/cleanup"
	"gb/cmd/close"
	"gb/cmd/create"
	"gb/cmd/delete"
	"gb/cmd/list"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gb",
	Short: "CLI for managing Github ",
	Long:  `gb is a CLI library for Github that allows users to manage their git repos.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(create.CreateCmd)
	RootCmd.AddCommand(list.ListCmd)
	RootCmd.AddCommand(delete.DeleteCmd)
	RootCmd.AddCommand(cleanup.CleanupCmd)
	RootCmd.AddCommand(close.CloseCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gb.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
