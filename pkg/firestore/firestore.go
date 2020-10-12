package firestore

import (
	"context"
	"encoding/json"
	"fmt"

	"google.golang.org/api/iterator"

	fs "cloud.google.com/go/firestore"
)

// Client for firestore helper
type Client struct {
	client    *fs.Client
	projectID string
}

// KeyValue is a slice of keyvalue pairs
type KeyValue map[string]interface{}

// New firestore helper
func New(ctx context.Context, projectID string) (*Client, error) {
	client, err := fs.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return &Client{client, projectID}, nil
}

// NewDocument add a new document to a collection, creating the collection if it doesn't exist
func (c *Client) NewDocument(ctx context.Context, collectionID string, document KeyValue) error {
	_, _, err := c.client.Collection(collectionID).Add(ctx, document)
	if err != nil {
		return fmt.Errorf("failed adding document: %v", err)
	}
	return nil
}

// GetAllDocumentsInCollection returns all documents in a given collection
func (c *Client) GetAllDocumentsInCollection(ctx context.Context, collectionID string) ([]byte, error) {
	iter := c.client.Collection(collectionID).Documents(ctx)
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

	jsonData, _ := json.MarshalIndent(data, "", "    ")
	return jsonData, nil
}

//MakeKeyValue create a Firestore KeyValue map
func (c *Client) MakeKeyValue(s string) (*KeyValue, error) {
	kv := make(KeyValue)
	err := json.Unmarshal([]byte(s), &kv)
	if err != nil {
		return nil, fmt.Errorf("Error creating Firestore Key Value pairs: %v", err)
	}
	return &kv, nil
}
