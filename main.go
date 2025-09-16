package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

type Granola struct {
	Cache string `json:"cache"`
}

type GranolaCache struct {
	State GranolaState `json:"state"`
}

type GranolaState struct {
	Events []GranolaEvent `json:"events"`
}

type GranolaEvent struct {
	ID string `json:"id"`
}

func main() {
	// Read file
	data, err := os.ReadFile("granola-cache.json")
	if err != nil {
		fmt.Printf("error reading file: %v", err)
	}

	var granola Granola
	if err := json.Unmarshal(data, &granola); err != nil {
		fmt.Printf("error unmarshalling outer JSON: %v\n", err)
		os.Exit(1)
	}

	var cache GranolaCache
	if err := json.Unmarshal([]byte(granola.Cache), &cache); err != nil {
		fmt.Printf("error unmarshalling cache: %v", err)
		os.Exit(1)
	}

	fmt.Print("\n")
	spew.Dump(cache)

	// Loop through struct
	// Write file

	os.Exit(0)
}
