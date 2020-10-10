package cli

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/jaredbancroft/go-gcp-helpers/pkg/firestore"
	"github.com/spf13/cobra"
)

func init() {
	firestoreCmd.AddCommand(firestoreDocumentCmd)
	firestoreDocumentCmd.AddCommand(firestoreDocumentNewCmd)
	firestoreDocumentCmd.AddCommand(firestoreDocumentGetCmd)
}

var firestoreCmd = &cobra.Command{
	Use:   "firestore",
	Short: "Run firestore commands",
	Long:  `Run basic firestore commands via the cli`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var firestoreDocumentCmd = &cobra.Command{
	Use:   "document",
	Short: "Firestore document commands",
	Long:  `Run basic firestore document commands via the cli`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var firestoreDocumentNewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new firestore document",
	Long:  `Create new firestore document commands via the cli`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		argStr := strings.Join(args[1:], "")
		newDocument(ctx, args[0], argStr)

		return nil
	},
}

var firestoreDocumentGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Return a firestore document",
	Long:  `Return an existing firestore document via the cli`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a collection")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		fmt.Println(getDocument(ctx, args[0]))

		return nil
	},
}

// NewDocument creates a new firestore document via cli
func newDocument(ctx context.Context, collection string, document string) bool {
	f := newFirestoreClient(ctx)
	d, err := f.MakeKeyValue(document)
	err = f.NewDocument(ctx, collection, d)
	if err != nil {
		log.Fatalf("whoops %v", err)
		return false
	}

	return true
}

// GetDocument is a thing
func getDocument(ctx context.Context, collection string) string {
	f := newFirestoreClient(ctx)
	results, err := f.GetAllDocumentsInCollection(ctx, collection)
	if err != nil {
		log.Fatalf("whoops %v", err)
	}

	return string(results)
}

func newFirestoreClient(ctx context.Context) firestore.Client {
	f, err := firestore.New(ctx, projectID)
	if err != nil {
		log.Fatalf("Unable to create new Firestore Client: %v", err)
	}

	return f
}
