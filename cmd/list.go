package cmd

import (
	"github.com/spf13/cobra"
	"github.com/verminio/datastore-cmd/namespace"
)

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
}
