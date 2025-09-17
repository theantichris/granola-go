package main

import (
	"strings"
	"testing"
)

func TestCreateCache(t *testing.T) {
	t.Run("creates the cache from valid JSON", func(t *testing.T) {
		t.Parallel()

		testJSON := `{"cache": "{\"state\":{\"documents\":{\"doc1\":{\"id\":\"abc123\"}}}}"}`

		expected := Cache{
			State: State{
				Documents: map[string]Document{
					"doc1": {
						ID: "abc123",
					},
				},
			},
		}

		cache, err := createCache([]byte(testJSON))
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if cache.State.Documents["doc1"].ID != expected.State.Documents["doc1"].ID {
			t.Errorf("expected document ID %q, got %q", expected.State.Documents["doc1"].ID, cache.State.Documents["doc1"].ID)
		}
	})

	t.Run("returns error for invalid outer JSON", func(t *testing.T) {
		t.Parallel()

		testJSON := `{"cache": {\"state\":{\"documents\":{\"doc1\":{\"id\":\"abc123\"}}}}"}`

		_, err := createCache([]byte(testJSON))
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		if !strings.Contains(err.Error(), "error unmarshalling outer JSON") {
			t.Errorf("expected 'error unmarshalling outer JSON', got %q", err.Error())
		}
	})

	t.Run("returns error for invalid cache JSON", func(t *testing.T) {
		t.Parallel()

		testJSON := `{"cache": "{\"state\":{documents\":{\"doc1\":{\"id\":\"abc123\"}}}}"}`

		_, err := createCache([]byte(testJSON))
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		if !strings.Contains(err.Error(), "error unmarshalling cache") {
			t.Errorf("expected 'error unmarshalling cache', got %q", err.Error())
		}
	})
}
