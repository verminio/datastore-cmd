package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var namespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "Namespace related actions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use one of the available sub-commands: list")
	},
}

func init() {
	rootCmd.AddCommand(namespaceCmd)
}
