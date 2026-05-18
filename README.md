# Lichess Go Client

Auto-generated Go client for the [Lichess API](https://lichess.org/api) using [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator).

## Overview

This repository automatically generates a Go client for the Lichess API. The client is regenerated whenever the Lichess API specification is updated via GitHub Actions, ensuring you always have access to the latest API endpoints and features.

## Features

- **Auto-generated**: Client code is generated from the official Lichess OpenAPI specification
- **Auto-updating**: GitHub Action checks daily for updates to the Lichess API and regenerates the client
- **Type-safe**: Full Go type definitions for all API models and responses
- **Complete**: Covers all Lichess API endpoints
- **No manual spec management**: The OpenAPI spec is cloned directly from the lichess-org/api repository

## Installation

```bash
go get golichess
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    "golichess"
    "golichess/api"
)

func main() {
    // Create a new client
    cfg := lichess.NewConfiguration()
    client := lichess.NewAPIClient(cfg)
    
    // Use the client to make API calls
    // Example: Get user profile
    ctx := context.Background()
    user, _, err := client.AccountApi.GetAccount(ctx)
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Username: %s\n", user.Username)
}
```

## Automated Updates

This repository uses a GitHub Action to automatically check for updates to the Lichess API:

- **Schedule**: Runs daily at 2 AM UTC
- **Manual Trigger**: Can be triggered manually from the Actions tab
- **Force Update**: Can force an update even if the version hasn't changed

The workflow:
1. Clones the lichess-org/api repository
2. Copies the OpenAPI spec (`doc/specs/lichess-api.yaml`) and dependencies (schemas, tags directories)
3. Compares the version with the previous one
4. If changed (or forced), regenerates the Go client using openapi-generator
5. Commits and pushes the changes

All OpenAPI spec files are managed by the GitHub Action and not committed to this repository.

## Configuration

The client generation is configured via `openapi-config.json`:

- `packageName`: Name of the generated Go package
- `packageVersion`: Version of the generated package
- `goPackage`: Go module path
- `modelPackage`: Package name for models
- `apiPackage`: Package name for API clients

## Lichess API Documentation

- [Official API Documentation](https://lichess.org/api)
- [Lichess API GitHub Repository](https://github.com/lichess-org/api)
- [API Demo App](https://lichess-org.github.io/api-demo/)

## License

This generated client follows the same license as the Lichess API: AGPL-3.0-or-later

## Contributing

Since this is an auto-generated client, direct code modifications are not recommended. Instead:

1. If you find issues with the generated code, they may stem from the OpenAPI spec
2. Report issues to the [Lichess API repository](https://github.com/lichess-org/api/issues)
3. For generator-specific issues, report to [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator/issues)
