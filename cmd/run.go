package cmd

import (
	"errors"
	"strings"

	"github.com/lateralusd/sochecker/checker"
	_ "github.com/lateralusd/sochecker/plugins/deezer"
	_ "github.com/lateralusd/sochecker/plugins/facebook"
	_ "github.com/lateralusd/sochecker/plugins/flickr"
	_ "github.com/lateralusd/sochecker/plugins/github"
	_ "github.com/lateralusd/sochecker/plugins/instagram"
	_ "github.com/lateralusd/sochecker/plugins/pinterest"
	_ "github.com/lateralusd/sochecker/plugins/reddit"
	_ "github.com/lateralusd/sochecker/plugins/twitter"
	"github.com/spf13/cobra"
)

var errNotProvided = errors.New("username not provided")

var runCmd = &cobra.Command{
	Use:   "run [username]",
	Short: "Search the user",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errNotProvided
		}

		timeout, err := cmd.Flags().GetInt("timeout")
		if err != nil {
			return err
		}

		user := strings.Join(args, "")
		checker.RunAll(user, timeout)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().IntP("timeout", "t", 10, "number of seconds until timeout for HTTP client")
}
