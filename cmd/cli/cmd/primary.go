package cmd

import "github.com/spf13/cobra"

// rootCmd represents the base command when called without any subcommands
var add = &cobra.Command{
	Use:   "add",
	Short: "add action",
	Long:  `add action`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

var update = &cobra.Command{
	Use:   "update",
	Short: "update action",
	Long:  `update action`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

var delete = &cobra.Command{
	Use:   "delete",
	Short: "delete action",
	Long:  `delete action`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

var get = &cobra.Command{
	Use:   "get",
	Short: "get action",
	Long:  `get action`,
}

func GetPrimaryCommands() []*cobra.Command {
	return []*cobra.Command{add, update, get, delete}
}
