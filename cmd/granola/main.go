package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/flytam/filenamify"
	"github.com/theantichris/granola-to-markdown/internal/granola"
)

func main() {
	cacheFile := flag.String("cache", "granola-cache.json", "Path to the Granola cache JSON file")
	outputFolder := flag.String("output", "output", "Directory to save the output markdown files")
	flag.Parse()

	data, err := os.ReadFile(*cacheFile)
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
		os.Exit(1)
	}

	cache, err := createCache(data)
	if err != nil {
		fmt.Printf("error creating cache: %v\n", err)
		os.Exit(1)
	}

	// Write to files
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
func getSafeTitle(doc granola.Document) (string, error) {
	safeTitle, err := filenamify.Filenamify(doc.Title, filenamify.Options{
		Replacement: "-",
	})
	return safeTitle, err
}

// createCache takes a byte slice of JSON data and unmarshals it into a Cache struct.
func createCache(data []byte) (granola.Cache, error) {
	var wrapper granola.Wrapper
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return granola.Cache{}, fmt.Errorf("error unmarshalling outer JSON: %v", err)
	}

	var cache granola.Cache
	if err := json.Unmarshal([]byte(wrapper.Cache), &cache); err != nil {
		return granola.Cache{}, fmt.Errorf("error unmarshalling cache: %v", err)
	}

	return cache, nil
}
