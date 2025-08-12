# Japan Law API v2 Client Generator

Go Client library for Japanese [Laws API Version 2].


## Project Structure

- `cmd/clientgen/` - Tool for generating Go client from OpenAPI specification
- `types.go` - Generated API type definitions and schema structures
- `client.go` - Generated HTTP client and API methods
- `example/` - Usage examples

## Client Generation Tool Usage

```bash
wget https://laws.e-gov.go.jp/api/2/swagger-ui/lawapi-v2.yaml
# Build the tool
go run ./cmd/clientgen
```

### Options

- `-input`: Path to OpenAPI specification file (default: `lawapi-v2.yaml`)
- `-output`: Output directory (default: `.`)
- `-package`: Go package name (default: `lawapi`)

## Generated Files

- `types.go`: API type definitions and schema structures
- `client.go`: HTTP client and API methods

## Client Library Usage

```go
package main

import (
    "fmt"
    "log"

    "go.ngs.io/jplaw-api-v2"
)

func main() {
    // Create client
    client := lawapi.NewClient()

    // Get laws list
    params := &lawapi.GetLawsParams{
        LawTitle: lawapi.StringPtr("Constitution"),
        Limit:    lawapi.Int32Ptr(10),
    }

    result, err := client.GetLaws(params)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Count: %d\n", result.Count)
}
```

## Key Features

- Automatic code generation from OpenAPI 3.0 specification
- Type-safe Go client
- Automatic query parameter construction
- Error handling
- Pointer helper functions

## Generated API Methods

- `GetLaws()` - Retrieve laws list
- `GetLawData()` - Retrieve law text content
- `GetRevisions()` - Retrieve law revision history
- `GetKeyword()` - Keyword search
- `GetLawFile()` - Retrieve law content file
- `GetAttachment()` - Retrieve attachments

## Notes

- API methods requiring path parameters currently have TODO comments
- When calling actual APIs, make sure to set appropriate path parameters

[Laws API Version 2]: https://laws.e-gov.go.jp/api/2/swagger-ui#/
