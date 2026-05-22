package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	oapicodegen "github.com/atye/golichess/oapi-codegen"
)

func main() {
	// Personal Access Token authentication
	// Generate a token at: https://lichess.org/account/oauth/token
	// accessToken := "your-personal-access-token" // Replace with your token
	accessToken := "lip_IjEHjCvO7sv6X19VMvfp2"

	// Create client with Bearer token authentication
	client, err := oapicodegen.NewClientWithResponses("https://lichess.org",
		oapicodegen.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
			return nil
		}))
	if err != nil {
		log.Fatalf("creating client: %v", err)
	}

	resp, err := client.AccountMeWithResponse(context.Background())
	if err != nil {
		log.Fatalf("executing request: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		log.Fatalf("getting account: %s", string(resp.Body))
	}

	fmt.Println(*resp.JSON200)
}
