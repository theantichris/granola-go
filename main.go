package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

type Wrapper struct {
	Cache string `json:"cache"`
}

type Cache struct {
	State   State   `json:"state"`
	Version float64 `json:"version"`
}

type State struct {
	Documents map[string]Document `json:"documents"`
}

type Document struct {
	ID string `json:"id"`
}

func main() {
	// Read file
	data, err := os.ReadFile("granola-cache.json")
	if err != nil {
		fmt.Printf("error reading file: %v", err)
	}

	var wrapper Wrapper
	if err := json.Unmarshal(data, &wrapper); err != nil {
		fmt.Printf("error unmarshalling outer JSON: %v\n", err)
		os.Exit(1)
	}

	var cache Cache
	if err := json.Unmarshal([]byte(wrapper.Cache), &cache); err != nil {
		fmt.Printf("error unmarshalling cache: %v", err)
		os.Exit(1)
	}

	fmt.Print("\n")
	spew.Dump(cache)

	// Loop through struct
	// Write file

	os.Exit(0)
}
