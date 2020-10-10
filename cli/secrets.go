package cli

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/jaredbancroft/go-gcp-helpers/pkg/secrets"
	"github.com/spf13/cobra"
)

func init() {
	secretsCmd.AddCommand(secretsGetCmd)
	secretsCmd.AddCommand(secretsNewCmd)
	secretsCmd.AddCommand(secretsUpdateCmd)
}

var secretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "Run secrets commands",
	Long:  `Run basic secrets commands via the cli`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var secretsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a secret out of secret manager",
	Long: `Get a secret out of secret manager optionally specifiying
	the version number`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		name := args[0]
		version := "latest"
		if len(args) > 1 {
			version = args[1]
		}
		secret := getSecret(ctx, name, version)
		fmt.Println(string(secret))
	},
}

var secretsNewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new secret in secret manager",
	Long:  `Create a new secret in secret manager`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		name := args[0]
		data := args[1]
		secret := createSecret(ctx, name, []byte(data))
		fmt.Println(secret)
	},
}

var secretsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a secret with a new version",
	Long:  `Update a secret with a new version`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		name := args[0]
		data := args[1]
		secret := updateSecret(ctx, name, []byte(data))
		fmt.Println(secret)
	},
}

func getSecret(ctx context.Context, name string, version string) []byte {
	s := newSecretsClient(ctx)
	if version != "latest" {
		intVersion, err := strconv.Atoi(version)
		if err != nil {
			log.Fatalf("Unable to parse version number: %v", err)
		}
		secret, err := s.GetSecretVersion(ctx, name, intVersion)
		if err != nil {
			log.Fatalf("Error accessing secret: %v", err)
		}

		return secret
	}

	secret, err := s.GetSecretLatest(ctx, name)
	if err != nil {
		log.Fatalf("Error accessing secret: %v", err)
	}

	return secret
}

func createSecret(ctx context.Context, name string, data []byte) string {
	s := newSecretsClient(ctx)
	secret, err := s.CreateNewSecret(ctx, name, data)
	if err != nil {
		log.Fatalf("Failed to create secret: %v", err)
	}

	return secret.GetName()
}

func updateSecret(ctx context.Context, name string, data []byte) string {
	s := newSecretsClient(ctx)
	secret, err := s.UpdateSecret(ctx, name, data)
	if err != nil {
		log.Fatalf("Failed to create secret: %v", err)
	}

	return secret.GetName()
}

func newSecretsClient(ctx context.Context) secrets.Client {
	f, err := secrets.New(ctx, projectID)
	if err != nil {
		log.Fatalf("Unable to create new Secrets Client: %v", err)
	}

	return f
}
