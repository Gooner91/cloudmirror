package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var update bool

var configCmd = &cobra.Command{
	Use: "config",
	Short: "Set configuration for media to be synced",
	Long: "Set configuration for the media you intend to sync, set all sources and their destination on cloud",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("command called")
		fmt.Println("flag value test update=%v", update)

	},
}

func init(){
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().BoolVarP(&update, "update", "u", false, "Update an existing mapping")
}
