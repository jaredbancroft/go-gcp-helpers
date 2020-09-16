package firestore

import (
	"context"
	"fmt"

	"google.golang.org/api/iterator"

	fs "cloud.google.com/go/firestore"
)

// Client for secrets helper
type Client struct {
	Context   context.Context
	client    *fs.Client
	projectID string
}

// KeyValue is a slice of keyvalue pairs
type KeyValue map[string]interface{}

// New secrets helper
func New(projectID string) (Client, error) {
	ctx := context.Background()
	client, err := fs.NewClient(ctx, projectID)
	if err != nil {
		return Client{}, err
	}
	return Client{ctx, client, projectID}, nil
}

// NewDocument add a new document to a collection, creating the collection if it doesn't exist
func (c *Client) NewDocument(collectionID string, document KeyValue) error {
	_, _, err := c.client.Collection(collectionID).Add(c.Context, document)
	if err != nil {
		return fmt.Errorf("failed adding document: %v", err)
	}
	return nil
}

// GetAllDocumentsInCollection returns all documents in a given collection
func (c *Client) GetAllDocumentsInCollection(collectionID string) ([]KeyValue, error) {
	iter := c.client.Collection(collectionID).Documents(c.Context)
	var data []KeyValue
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Failed to iterate: %v", err)
		}
		data = append(data, doc.Data())
	}
	return data, nil
}
