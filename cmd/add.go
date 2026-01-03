package cmd

import (
	"github.com/Gooner91/cloudmirror/internal/config"
	"github.com/spf13/cobra"
)

var (
	srcGlob string
	dest string
)

var configAddCmd = &cobra.Command{
	Use: "add",
	Short: "Adds a mapping to the config",
	Long: "Allows to add a source and destination glob",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.Config{
			SrcGlob: srcGlob,
			Dest:    dest,
		}

		return config.Save(cfg)
	},
}

func init(){
	configCmd.AddCommand(configAddCmd)
	configAddCmd.Flags().StringVar(&srcGlob, "srcGlob", "", "Glob pattern for source directories/files (required)")
	configAddCmd.MarkFlagRequired("srcGlob")
	configAddCmd.Flags().StringVar(&dest, "dest", "", "Destination path for the provided source")
	configAddCmd.MarkFlagRequired("dest")

}
