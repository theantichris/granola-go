package granola

import (
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	t.Run("creates the cache with document", func(t *testing.T) {
		t.Parallel()

		testJSON := `{"cache": "{\"state\":{\"documents\":{\"abc123\":{\"id\":\"abc123\",\"title\":\"Test Document\",\"created_at\":\"2025-09-12T18:59:15.595Z\",\"updated_at\":\"2025-09-12T19:15:33.102Z\",\"notes_markdown\":\"# Heading\\nSome notes here.\",\"notes_plain\":\"Heading: Some notes here.\",\"notes\":{\"type\":\"doc\",\"content\":[{\"type\":\"heading\",\"attrs\":{\"level\":1},\"content\":[{\"type\":\"text\",\"text\":\"Meeting Title\"}]},{\"type\":\"paragraph\",\"content\":[{\"type\":\"text\",\"text\":\"Some notes here.\"}]}]}}}}}"}`

		createdAt, _ := time.Parse(time.RFC3339, "2025-09-12T18:59:15.595Z")
		updatedAt, _ := time.Parse(time.RFC3339, "2025-09-12T19:15:33.102Z")
		expected := Cache{
			State: State{
				Documents: map[string]Document{
					"abc123": {
						ID:            "abc123",
						Title:         "Test Document",
						CreatedAt:     createdAt,
						UpdatedAt:     updatedAt,
						NotesMarkdown: "# Heading\nSome notes here.",
						NotesPlain:    "Heading: Some notes here.",
						Notes: Notes{
							Type: "doc",
							Content: []Content{
								{
									Type:  "heading",
									Attrs: map[string]any{"level": float64(1)},
									Content: []Content{
										{
											Type: "text",
											Text: "Meeting Title",
										},
									}},
								{
									Type: "paragraph",
									Content: []Content{
										{
											Type: "text",
											Text: "Some notes here."},
									}},
							},
						},
					},
				},
			},
		}

		cache, err := New([]byte(testJSON))
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if cache.State.Documents["abc123"].ID != expected.State.Documents["abc123"].ID {
			t.Errorf("expected document ID %q, got %q", expected.State.Documents["doc1"].ID, cache.State.Documents["doc1"].ID)
		}

		if cache.State.Documents["doc1"].Title != expected.State.Documents["doc1"].Title {
			t.Errorf("expected document Title %q, got %q", expected.State.Documents["doc1"].Title, cache.State.Documents["doc1"].Title)
		}

		if cache.State.Documents["doc1"].CreatedAt != expected.State.Documents["doc1"].CreatedAt {
			t.Errorf("expected created time %q, got %q", expected.State.Documents["doc1"].CreatedAt, cache.State.Documents["doc1"].CreatedAt)
		}

		if cache.State.Documents["doc1"].UpdatedAt != expected.State.Documents["doc1"].UpdatedAt {
			t.Errorf("expected updated time %q, got %q", expected.State.Documents["doc1"].UpdatedAt, cache.State.Documents["doc1"].UpdatedAt)
		}

		if cache.State.Documents["doc1"].NotesMarkdown != expected.State.Documents["doc1"].NotesMarkdown {
			t.Errorf("expected notes markdown %q, got %q", expected.State.Documents["doc1"].NotesMarkdown, cache.State.Documents["doc1"].NotesMarkdown)
		}

		if cache.State.Documents["doc1"].NotesPlain != expected.State.Documents["doc1"].NotesPlain {
			t.Errorf("expected notes plain %q, got %q", expected.State.Documents["doc1"].NotesPlain, cache.State.Documents["doc1"].NotesPlain)
		}

		if !cmp.Equal(cache.State.Documents["abc123"].Notes, expected.State.Documents["abc123"].Notes) {
			t.Errorf("expected notes %+v, got %+v", expected.State.Documents["abc123"].Notes, cache.State.Documents["abc123"].Notes)
		}
	})

	t.Run("returns error for invalid wrapper JSON", func(t *testing.T) {
		t.Parallel()

		invalidWrapperJSON := `{"cache": {\"state\":{\"documents\":{\"doc1\":{\"id\":\"abc123\"}}}}"}`

		_, err := New([]byte(invalidWrapperJSON))
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		if !strings.Contains(err.Error(), "error unmarshalling outer JSON") {
			t.Errorf("expected 'error unmarshalling outer JSON', got %q", err.Error())
		}
	})

	t.Run("returns error for invalid cache JSON", func(t *testing.T) {
		t.Parallel()

		invalidCacheJSON := `{"cache": "{\"state\":{documents\":{\"doc1\":{\"id\":\"abc123\"}}}}"}`

		_, err := New([]byte(invalidCacheJSON))
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		if !strings.Contains(err.Error(), "error unmarshalling cache") {
			t.Errorf("expected 'error unmarshalling cache', got %q", err.Error())
		}
	})
}
