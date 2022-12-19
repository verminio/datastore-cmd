package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func getProjectId(cmd *cobra.Command) (projectId string) {
	projectId, _ = cmd.Root().PersistentFlags().GetString("project-id")
	if !silenceOutput {
		fmt.Printf("Project: %v\n", projectId)
	}
	return
}

func render(res interface{}) {
	fmt.Printf("%v\n", res)
}
