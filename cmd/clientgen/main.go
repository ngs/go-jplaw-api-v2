package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func main() {
	var (
		inputFile   = flag.String("input", "lawapi-v2.yaml", "OpenAPI specification file")
		outputDir   = flag.String("output", ".", "Output directory for generated client")
		packageName = flag.String("package", "lawapi", "Package name for generated code")
	)
	flag.Parse()

	// Read OpenAPI specification file
	yamlData, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("Failed to read input file %s: %v", *inputFile, err)
	}

	var spec OpenAPISpec
	if err := yaml.Unmarshal(yamlData, &spec); err != nil {
		log.Fatalf("Failed to parse OpenAPI spec: %v", err)
	}

	// Create output directory
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory %s: %v", *outputDir, err)
	}

	generator := NewGenerator(&spec, *packageName)

	// Generate type definitions file
	typesContent := generator.GenerateTypes()
	typesFile := filepath.Join(*outputDir, "types.go")
	if err := ioutil.WriteFile(typesFile, []byte(typesContent), 0644); err != nil {
		log.Fatalf("Failed to write types file: %v", err)
	}
	fmt.Printf("Generated types: %s\n", typesFile)

	// Generate client file
	clientContent := generator.GenerateClient()
	clientFile := filepath.Join(*outputDir, "client.go")
	if err := ioutil.WriteFile(clientFile, []byte(clientContent), 0644); err != nil {
		log.Fatalf("Failed to write client file: %v", err)
	}
	fmt.Printf("Generated client: %s\n", clientFile)

	fmt.Printf("Client library generated successfully in %s/\n", *outputDir)
	fmt.Println("\nUsage example:")
	fmt.Printf("  client := %s.NewClient()\n", *packageName)
	fmt.Println("  // Use client methods to call API endpoints")
}
