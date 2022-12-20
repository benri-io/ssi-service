package cmd

import "github.com/spf13/cobra"

var credential = &cobra.Command{
	Use:   "credential",
	Short: "credential subcommand",
	Long:  `credential`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

func GetSeoncdaryCommands() []*cobra.Command {
	return []*cobra.Command{credential}
}
