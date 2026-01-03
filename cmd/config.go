package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use: "config",
	Short: "Manage configurations",
	Long: "Use add or remove subcommands to add or remove mappings",
}

func init(){
	rootCmd.AddCommand(configCmd)
}
