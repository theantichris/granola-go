package main

import "time"

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
