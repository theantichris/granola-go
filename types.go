package main

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
	ID string `json:"id"` // UUID of the document
}
