package firestore

import (
	"context"

	fs "cloud.google.com/go/firestore"
)

// Client for secrets helper
type Client struct {
	Context   context.Context
	client    *fs.Client
	projectID string
}

// New secrets helper
func New(projectID string) (Client, error) {
	ctx := context.Background()
	client, err := fs.NewClient(ctx, projectID)
	if err != nil {
		return Client{}, err
	}
	return Client{ctx, client, projectID}, nil
}

// NewCollection creates a new firestore collection
func (c *Client) NewCollection(collectionID string) *fs.CollectionRef {
	newCollection := c.client.Collection(collectionID)
	return newCollection
}

// NewDocument creates a new document in a the supplied collection
func (c *Client) NewDocument(collection *fs.CollectionRef, document string) *fs.DocumentRef {
	newDocument := collection.Doc(document)
	return newDocument
}

// GetSingleDocument gets a single doc from a collection
func (c *Client) GetSingleDocument(document *fs.DocumentRef) (map[string]interface{}, error) {
	documentSnapshot, err := document.Get(c.Context)
	if err != nil {
		return nil, err
	}

	dataMap := documentSnapshot.Data()
	return dataMap, nil
}
