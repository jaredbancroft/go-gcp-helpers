package secrets

import (
	"context"
	"fmt"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// Secrets helper
type Secrets struct {
	client    *secretmanager.Client
	ctx       context.Context
	ProjectID string
}

// New secrets helper.
func New(projectID string) *Secrets {
	ctx := initSecretManagerContext()
	client := initSecretManagerClient(ctx)
	s := Secrets{client, ctx, projectID}
	return &s
}

func initSecretManagerClient(ctx context.Context) *secretmanager.Client {
	// Create the client.
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}

	return client
}

func initSecretManagerContext() context.Context {
	ctx := context.Background()

	return ctx
}

// CreateSecret creates a new secret
func (s *Secrets) CreateNewSecret(secretID string, payload []byte) *secretmanagerpb.SecretVersion {
	// Create the request to create the secret.
	createSecretReq := &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", s.ProjectID),
		SecretId: secretID,
		Secret: &secretmanagerpb.Secret{
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_Automatic_{
					Automatic: &secretmanagerpb.Replication_Automatic{},
				},
			},
		},
	}

	secret, err := s.client.CreateSecret(s.ctx, createSecretReq)
	if err != nil {
		log.Fatalf("failed to add secret: %v", err)
	}

	// Build the request.
	addSecretVersionReq := &secretmanagerpb.AddSecretVersionRequest{
		Parent: secret.Name,
		Payload: &secretmanagerpb.SecretPayload{
			Data: payload,
		},
	}

	// Call the API.
	version, err := s.client.AddSecretVersion(s.ctx, addSecretVersionReq)
	if err != nil {
		log.Fatalf("failed to add secret version: %v", err)
	}

	return version
}

/*


func (s *Secrets) Create {}

/*

	// Build the request.
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
			Name: version.Name,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
			log.Fatalf("failed to access secret version: %v", err)
	}

	// Print the secret payload.
	//
	// WARNING: Do not print the secret in a production environment - this
	// snippet is showing how to access the secret material.
	log.Printf("Plaintext: %s", result.Payload.Data)
}
*/
