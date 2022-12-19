package cmd

import (
	"encoding/json"
	"fmt"
	"os"

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
	switch format {
	case jsonFormat:
		encoder := json.NewEncoder(os.Stdout)
		if err := encoder.Encode(res); err != nil {
			panic(err)
		}
		return
	default:
		fmt.Printf("%v\n", res)
		return
	}
}
