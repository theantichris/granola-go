# granola-go

A Go CLI application that exports Granola meeting notes to Markdown files. Reads from Granola's local cache (JSON) and writes one Markdown file per meeting document.

## Usage

```sh
# Build
go build -o granola-to-markdown.exe ./cmd

# Run with default cache file and output directory
./granola-to-markdown.exe

# Run with a specific cache file
./granola-to-markdown.exe --cache=path/to/cache-v3.json

# Run with a specific output directory
./granola-to-markdown.exe --output=outputdir

# Combine flags
./granola-to-markdown.exe --cache=path/to/cache-v3.json --output=outputdir
```

## Granola Cache

### Cache file locations

- **Windows:** `~\AppData\Roaming\Granola\cache-v3.json`
- **macOS:** `~/Library/Application Support/Granola/cache-v3.json`

### Cache schema

The cache file is JSON with a wrapper property `cache` (a JSON string). The inner JSON contains a `state` object:

```text
state
├── documents: Map (keyed by UUID)
│   ├── [UUID]
│   │   ├── id: String (UUID)
│   │   ├── title: String
│   │   ├── created_at: Timestamp (ISO 8601)
│   │   ├── updated_at: Timestamp (ISO 8601)
│   │   ├── notes_plain: String
│   │   ├── notes_markdown: String
│   │   ├── notes: Object (TipTap rich content)
│   │   │   ├── type: String
│   │   │   └── content: Array
├── transcripts: Map (keyed by document UUID)
│   ├── [UUID]: Array of transcript entries
│   │   ├── id: String
│   │   ├── document_id: String
│   │   ├── text: String
│   │   ├── source: String
│   │   ├── start_timestamp: String
│   │   ├── end_timestamp: String
│   │   ├── is_final: Bool
version: float64
```

## Project Structure

- `cmd/main.go`: CLI entry point (parses flags, handles I/O)
- `granola/cache.go`: Core cache parsing and data structures
- `granola/cache_test.go`: Unit tests for cache logic
- `.github/workflows/go.yml`: CI workflow (builds and tests on push/PR)

## API & Authentication

The Granola API uses a bearer token from `supabase.json` (same directory as the cache file). Base URL: <https://api.granola.ai/v2/>

Set these headers:

```go
http.SetHeader("Content-Type", "application/json")
http.SetHeader("Authorization", "Bearer "+token)
http.SetHeader("User-Agent", "Granola/5.354.0")
http.SetHeader("X-Client-Version", "5.354.0")
```
