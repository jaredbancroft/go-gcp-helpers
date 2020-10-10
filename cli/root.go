package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")

// SetupRootCommand creates the cobra root command based on various options
func SetupRootCommand(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print usage")
	rootCmd.AddCommand(firestoreCmd, secretsCmd)
}
