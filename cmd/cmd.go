package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func getProjectId(cmd *cobra.Command) (projectId string) {
	projectId, _ = cmd.Root().PersistentFlags().GetString("project-id")
	fmt.Printf("Project: %v\n", projectId)
	return
}
