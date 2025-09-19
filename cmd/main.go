package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/flytam/filenamify"
	"github.com/theantichris/granola-to-markdown/granola"
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

	granola, err := granola.New(data)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	// Write to files
	for _, doc := range granola.State.Documents {
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
