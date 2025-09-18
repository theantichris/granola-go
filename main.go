package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/flytam/filenamify"
)

// Wrapper is the outer structure that contains the cache as a JSON string.
type Wrapper struct {
	Cache string `json:"cache"`
}

// Cache represents the main cache structure.
type Cache struct {
	State State `json:"state"`
}

// State holds the documents in the cache.
type State struct {
	Documents map[string]Document `json:"documents"`
}

// Document represents a single document in the cache.
type Document struct {
	ID            string    `json:"id"`             // UUID of the document
	Title         string    `json:"title"`          // Title of the document
	CreatedAt     time.Time `json:"created_at"`     // Creation timestamp
	UpdatedAt     time.Time `json:"updated_at"`     // Last updated timestamp
	NotesMarkdown string    `json:"notes_markdown"` // Notes in Markdown format
	NotesPlain    string    `json:"notes_plain"`    // Notes in plain text format
	Notes         Notes     `json:"notes"`          // Notes in TipTap format
}

// Notes represents the notes in TipTap format.
type Notes struct {
	Type    string    `json:"type"`    // Type of the note, e.g., "doc"
	Content []Content `json:"content"` // Content of the note
}

// Content represents a piece of content in the TipTap notes.
type Content struct {
	Type    string         `json:"type"`              // Type of the content, e.g., "heading", "paragraph", "text"
	Attrs   map[string]any `json:"attrs,omitempty"`   // Attributes for the content, e.g., level for headings
	Content []Content      `json:"content,omitempty"` // Nested content
	Text    string         `json:"text,omitempty"`    // Text content for text nodes
}

func main() {
	cacheFile := flag.String("cache", "granola-cache.json", "Path to the Granola cache JSON file")
	outputFolder := flag.String("output", "output", "Directory to save the output markdown files")
	flag.Parse()

	data, err := os.ReadFile(*cacheFile)
	if err != nil {
		fmt.Printf("error reading file: %v", err)
		os.Exit(1)
	}

	cache, err := createCache(data)
	if err != nil {
		fmt.Printf("error creating cache: %v", err)
		os.Exit(1)
	}

	for _, doc := range cache.State.Documents {
		contents := doc.Title + "\n" + doc.NotesMarkdown

		safeTitle, err := getSafeTitle(doc)
		if err != nil {
			fmt.Printf("error creating safe filename: %v", err)
			os.Exit(1)
		}

		err = os.MkdirAll(*outputFolder, 0755)
		if err != nil {
			fmt.Printf("error creating output directory: %v", err)
			os.Exit(1)
		}

		filename := fmt.Sprintf("%s-%s.md", safeTitle, doc.ID)
		if err := os.WriteFile("output/"+filename, []byte(contents), 0644); err != nil {
			fmt.Printf("error writing file %s: %v", filename, err)
			os.Exit(1)
		}
	}

	os.Exit(0)
}

// getSafeTitle generates a filesystem-safe title for a document.
func getSafeTitle(doc Document) (string, error) {
	safeTitle, err := filenamify.Filenamify(doc.Title, filenamify.Options{
		Replacement: "-",
	})
	return safeTitle, err
}

// createCache takes a byte slice of JSON data and unmarshals it into a Cache struct.
func createCache(data []byte) (Cache, error) {
	var wrapper Wrapper
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return Cache{}, fmt.Errorf("error unmarshalling outer JSON: %v", err)
	}

	var cache Cache
	if err := json.Unmarshal([]byte(wrapper.Cache), &cache); err != nil {
		return Cache{}, fmt.Errorf("error unmarshalling cache: %v", err)
	}

	return cache, nil
}
