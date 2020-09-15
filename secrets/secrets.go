package secrets

import (
	"context"
	"fmt"
	"strconv"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// Client for secrets helper
type Client struct {
	Context   context.Context
	client    *secretmanager.Client
	projectID string
}

// New secrets helper
func New(projectID string) (Client, error) {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return Client{}, err
	}
	return Client{ctx, client, projectID}, nil
}

// CreateNewSecret creates a new secret
func (c *Client) CreateNewSecret(secretID string, payload []byte) (*secretmanagerpb.SecretVersion, error) {
	// Create the request to create the secret.
	createSecretReq := &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", c.projectID),
		SecretId: secretID,
		Secret: &secretmanagerpb.Secret{
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_Automatic_{
					Automatic: &secretmanagerpb.Replication_Automatic{},
				},
			},
		},
	}

	secret, err := c.client.CreateSecret(c.Context, createSecretReq)
	if err != nil {
		return nil, fmt.Errorf("failed to add secret: %v", err)
	}

	// Build the request.
	addSecretVersionReq := &secretmanagerpb.AddSecretVersionRequest{
		Parent: secret.Name,
		Payload: &secretmanagerpb.SecretPayload{
			Data: payload,
		},
	}

	// Call the API.
	version, err := c.client.AddSecretVersion(c.Context, addSecretVersionReq)
	if err != nil {
		return nil, fmt.Errorf("failed to add secret version: %v", err)
	}

	return version, nil
}

// GetSecretVersion gets an existing version of secret
func (c *Client) GetSecretVersion(secretName string, versionNumber int) ([]byte, error) {
	// Build the request.
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: "projects/" + c.projectID + "/secrets/" + secretName + "/versions/" + strconv.Itoa(versionNumber),
	}

	result, err := c.client.AccessSecretVersion(c.Context, accessRequest)
	if err != nil {
		return []byte(""), fmt.Errorf("failed to access secret version: %v", err)
	}

	return result.Payload.Data, nil
}
