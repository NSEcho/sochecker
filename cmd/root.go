package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sochecker",
	Short: "check username on multiple websites",
}

func Execute() error {
	return rootCmd.Execute()
}
