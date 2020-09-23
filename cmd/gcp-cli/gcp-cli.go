package main

import (
	"fmt"
	"os"

	"github.com/jaredbancroft/go-gcp-helpers/cli"
	"github.com/spf13/cobra"
)

func newCLICommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "gcp-cli [command]",
		Short: "gcp-cli is a CLI for the go-gcp-helper library",
		Long:  "gcp-cli is a CLI for the go-gcp-helper library",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
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
	if err := cmd.Execute(); err != nil {
		onError(err)
	}

	/*s, err := secrets.New("jared-go-react")
	if err != nil {
		log.Fatalf("whoops %v", err)
	}
	version, err := s.CreateNewSecret("my-new-secret3", []byte("secret datas!"))
	fmt.Println(err)
	fmt.Println(version)

	secret, err := s.GetSecretVersion("my-new-secret2", 1)
	fmt.Println(err)
	fmt.Println(string(secret))

	f, err := firestore.New("jared-go-react")
	if err != nil {
		log.Fatalf("whoops %v", err)
	}
	err = f.NewDocument("test", firestore.KeyValue{"fn": "jared", "ln": "bancroft"})
	if err != nil {
		log.Fatalf("whoops %v", err)
	}

	results, err := f.GetAllDocumentsInCollection("test")
	if err != nil {
		log.Fatalf("whoops %v", err)
	}

	fmt.Println(string(results))*/
}
