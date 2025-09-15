package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <markdown-file>")
	}

	filename := os.Args[1]

	// Read the markdown file
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Create goldmark parser with frontmatter extension
	md := goldmark.New(
		goldmark.WithExtensions(
			&frontmatter.Extender{},
		),
	)

	// Create a parser context to capture frontmatter
	ctx := parser.NewContext()

	// Parse the markdown (this extracts frontmatter)
	var htmlOutput bytes.Buffer
	if err := md.Convert(content, &htmlOutput, parser.WithContext(ctx)); err != nil {
		log.Fatalf("Error converting markdown: %v", err)
	}

	// Get the frontmatter data
	fm := frontmatter.Get(ctx)
	if fm == nil {
		fmt.Println("No frontmatter found in the document")
		return
	}

	// Decode frontmatter into a map
	var metadata map[string]interface{}
	if err := fm.Decode(&metadata); err != nil {
		log.Fatalf("Error decoding frontmatter: %v", err)
	}

	// Pretty print the frontmatter as JSON
	jsonData, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling to JSON: %v", err)
	}

	fmt.Println("Extracted Frontmatter:")
	fmt.Println(string(jsonData))

	// Also demonstrate decoding into a struct
	var blogPost struct {
		Title       string   `yaml:"title"`
		Description string   `yaml:"description"`
		Tags        []string `yaml:"tags"`
		Date        string   `yaml:"date"`
		Draft       bool     `yaml:"draft"`
	}

	if err := fm.Decode(&blogPost); err == nil {
		fmt.Printf("\nDecoded into struct:\n")
		fmt.Printf("Title: %s\n", blogPost.Title)
		fmt.Printf("Description: %s\n", blogPost.Description)
		fmt.Printf("Tags: %v\n", blogPost.Tags)
		fmt.Printf("Date: %s\n", blogPost.Date)
		fmt.Printf("Draft: %t\n", blogPost.Draft)
	}
}
