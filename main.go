package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// TODO: lead file from flag
	// TODO: choose format from flag
	// TODO: choose location from flag

	data, err := os.ReadFile("granola-cache.json")
	if err != nil {
		fmt.Printf("error reading file: %v", err)
	}

	cache, err := createCache(data)
	if err != nil {
		fmt.Printf("error creating cache: %v", err)
		os.Exit(1)
	}

	fmt.Print(cache)

	os.Exit(0)
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
