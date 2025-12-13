package cmd

import(
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cloudmirror",
	Short: "A bridge between your computer and cloud",
	Long: "With cloudmirror you can sync your data from your system to cloud e.g. google drive (supported for now)",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
