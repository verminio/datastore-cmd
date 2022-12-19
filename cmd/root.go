/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var silenceOutput bool

var rootCmd = &cobra.Command{
	Use:   "datastore-cmd",
	Short: "Command line interface for accessing GCP Datastore",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&silenceOutput, "silent", "s", false, "Hide command output")
	rootCmd.PersistentFlags().StringP("project-id", "p", "", "Datastore project id")
	rootCmd.MarkPersistentFlagRequired("project-id")
}
