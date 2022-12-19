package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

type outputFormat string

const (
	defaultFormat outputFormat = "default"
	jsonFormat    outputFormat = "json"
)

func (o *outputFormat) String() string {
	return string(*o)
}

func (o *outputFormat) Set(v string) error {
	switch v {
	case "default", "json":
		*o = outputFormat(v)
		return nil
	default:
		return errors.New(`unsupported output format. accepted values are "default", "json"`)
	}
}

func (o *outputFormat) Type() string {
	return "outputFormat"
}

var silenceOutput bool
var format outputFormat

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
	rootCmd.PersistentFlags().VarP(&format, "output", "o", "Output format type (default, json)")
	rootCmd.PersistentFlags().StringP("project-id", "p", "", "Datastore project id")
	rootCmd.MarkPersistentFlagRequired("project-id")
}
