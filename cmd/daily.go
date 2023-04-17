package cmd

import "github.com/spf13/cobra"

var dailyCmd = &cobra.Command{
	Use: "daily",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
