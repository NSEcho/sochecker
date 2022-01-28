package cmd

import (
	"fmt"

	"github.com/lateralusd/sochecker/checker"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info for specific module",
	RunE: func(cmd *cobra.Command, args []string) error {
		m, err := cmd.Flags().GetString("module")
		if err != nil {
			return err
		}

		info := checker.Info(m)
		fmt.Println(info)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
	infoCmd.Flags().StringP("module", "m", "all", "get info for module")
}
