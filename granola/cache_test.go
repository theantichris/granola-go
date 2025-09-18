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
		expected := Document{
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
		}

		cache, err := New([]byte(testJSON))
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		got := cache.State.Documents["abc123"]

		if got.ID != expected.ID {
			t.Errorf("expected document ID %q, got %q", expected.ID, got.ID)
		}

		if got.Title != expected.Title {
			t.Errorf("expected document Title %q, got %q", expected.Title, got.Title)
		}

		if got.CreatedAt != expected.CreatedAt {
			t.Errorf("expected created time %q, got %q", expected.CreatedAt, got.CreatedAt)
		}

		if got.UpdatedAt != expected.UpdatedAt {
			t.Errorf("expected updated time %q, got %q", expected.UpdatedAt, got.UpdatedAt)
		}

		if got.NotesMarkdown != expected.NotesMarkdown {
			t.Errorf("expected notes markdown %q, got %q", expected.NotesMarkdown, got.NotesMarkdown)
		}

		if got.NotesPlain != expected.NotesPlain {
			t.Errorf("expected notes plain %q, got %q", expected.NotesPlain, got.NotesPlain)
		}

		if !cmp.Equal(got.Notes, expected.Notes) {
			t.Errorf("expected notes %+v, got %+v", expected.Notes, got.Notes)
		}
	})

	t.Run("creates the cache with transcript", func(t *testing.T) {
		t.Parallel()

		testJSON := `{"cache": "{\"state\":{\"transcripts\":{\"abc123\":[{\"id\":\"abc123\",\"document_id\":\"doc1\",\"text\":\"This is a test transcript.\",\"source\":\"system\",\"start_timestamp\":\"2025-09-12T18:59:15.595Z\",\"end_timestamp\":\"2025-09-12T19:15:33.102Z\",\"is_final\":true}]}}}"}`

		startTimestamp, _ := time.Parse(time.RFC3339, "2025-09-12T18:59:15.595Z")
		endTimestamp, _ := time.Parse(time.RFC3339, "2025-09-12T19:15:33.102Z")
		expected := Transcript{
			ID:             "abc123",
			DocumentID:     "doc1",
			Text:           "This is a test transcript.",
			Source:         "system",
			StartTimestamp: startTimestamp,
			EndTimestamp:   endTimestamp,
			IsFinal:        true,
		}

		cache, err := New([]byte(testJSON))
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		got := cache.State.Transcripts["abc123"][0]

		if got.ID != expected.ID {
			t.Errorf("expected transcript ID %q, got %q", expected.ID, got.ID)
		}

		if got.DocumentID != expected.DocumentID {
			t.Errorf("expected document ID %q, got %q", expected.DocumentID, got.DocumentID)
		}

		if got.Text != expected.Text {
			t.Errorf("expected text %q, got %q", expected.Text, got.Text)
		}

		if got.Source != expected.Source {
			t.Errorf("expected source %q, got %q", expected.Source, got.Source)
		}

		if !got.StartTimestamp.Equal(expected.StartTimestamp) {
			t.Errorf("expected start timestamp %q, got %q", expected.StartTimestamp, got.StartTimestamp)
		}

		if !got.EndTimestamp.Equal(expected.EndTimestamp) {
			t.Errorf("expected end timestamp %q, got %q", expected.EndTimestamp, got.EndTimestamp)
		}

		if got.IsFinal != expected.IsFinal {
			t.Errorf("expected is_final %v, got %v", expected.IsFinal, got.IsFinal)
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
