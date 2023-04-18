package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	flagTag string
)

var rootCmd = &cobra.Command{
	Use: "cli",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(likesCmd)
	likesCmd.AddCommand(
		likesDailyCmd,
		likesWeeklyCmd,
	)

	for _, cmd := range []*cobra.Command{likesDailyCmd, likesWeeklyCmd} {
		cmd.Flags().StringVar(&flagTag, "tag", "", "")
	}
}
