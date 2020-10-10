package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jaredbancroft/go-gcp-helpers/cli"
	"github.com/spf13/cobra"
)

func newCLICommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "gcp-cli",
		Short: "gcp-cli is a CLI for the go-gcp-helper library",
		Long:  "gcp-cli is a CLI for the go-gcp-helper library",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
		Version: "v0.0.0",
	}

	cli.SetupRootCommand(cmd)

	return cmd, nil
}

func main() {

	onError := func(err error) {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	cmd, err := newCLICommand()
	if err != nil {
		onError(err)
		os.Exit(1)
	}

	cmd.SetOut(os.Stdout)
	ctx := context.Background()
	if err := cmd.ExecuteContext(ctx); err != nil {
		onError(err)
	}
}
