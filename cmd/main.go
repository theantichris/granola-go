package main

import (
	"flag"
	"log/slog"
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
		slog.Error("error reading file", "err", err, "file", *cacheFile)
		os.Exit(1)
	}

	cache, err := granola.NewCache(data)
	if err != nil {
		slog.Error("error creating cache", "err", err)
		os.Exit(1)
	}

	// Write to files
	for _, doc := range cache.State.Documents {
		contents := doc.Title + "\n" + doc.NotesMarkdown

		safeTitle, err := getSafeTitle(doc)
		if err != nil {
			slog.Error("error creating safe filename", "err", err, "title", doc.Title)
			os.Exit(1)
		}

		err = os.MkdirAll(*outputFolder, 0755)
		if err != nil {
			slog.Error("error creating output directory", "err", err, "dir", *outputFolder)
			os.Exit(1)
		}

		filename := safeTitle + "-" + doc.ID + ".md"
		outPath := *outputFolder + string(os.PathSeparator) + filename
		if err := os.WriteFile(outPath, []byte(contents), 0644); err != nil {
			slog.Error("error writing file", "err", err, "file", outPath)
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
