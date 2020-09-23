package cli

import (
	"github.com/spf13/cobra"
)

// SetupRootCommand creates the cobra root command based on various options
func SetupRootCommand(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print usage")
}
