/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/verminio/datastore-cmd/namespace"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List namespaces",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectId := getProjectId(cmd)
		namespaces, err := namespace.List(projectId)
		if err != nil {
			return err
		}
		render(namespaces)
		return nil
	},
}

func init() {
	namespaceCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
