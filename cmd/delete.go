package cmd

import (
	"github.com/Gooner91/cloudmirror/internal/config"
	"github.com/spf13/cobra"
)

var configDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a config mapping",
	Long:  "Given a pair of src and dest glob, it deletes the respective config",
	RunE: func(cmd *cobra.Command, args []string) error {
		srcGlob, err := cmd.Flags().GetString("srcGlob")
		if err != nil {
			return err
		}

		dest, err := cmd.Flags().GetString("dest")
		if err != nil {
			return err
		}

		cfg := config.Config{
			SrcGlob: srcGlob,
			Dest:    dest,
		}

		cfg

	},
}

func init() {
	configCmd.AddCommand(configDeleteCmd)
	configDeleteCmd.Flags().StringVar("srcGlob", "", "Glob pattern for source directories/files (required)")
	configDeleteCmd.MarkFlagRequired("srcGlob")
	configDeleteCmd.Flags().StringVar("dest", "", "Destination path for the provided source (required)")
	configDeleteCmd.MarkFlagRequired("dest")
}
