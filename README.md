# granola-to-markdown

A Go application that exports Granola meeting notes to Markdown files.

## Usage

```sh
# Build
$ go build -o granola-to-markdown.exe

# Run with default cache file and output directory
$ ./granola-to-markdown.exe

# Run with a specific cache file
$ ./granola-to-markdown.exe --cache=my-cache.json

# Run with a specific output directory
$ ./granola-to-markdown.exe --output=outdir

# Combine flags
$ ./granola-to-markdown.exe --cache=my-cache.json --output=outdir
```

## Granola Cache

### Cache file

On Windows this is stored at `~\AppData\Roaming\Granola\cache-v3.json`.

On MacOS this is stored at `Library/Application Support/Granola/cache-v3.json`.

### Granola cache scheme

The cache file is in JSON with a wrapper property `cache` that's value is a JSON string. That contains a `state` object that has all the information.

Documents contain the written notes a user takes but not notes generated from the transcript.

Transcripts is an array of the different parts of the transcript.

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
│   ├── [UUID]
│   │   └── entries: Array
│   │       ├── id: String
│   │       ├── document_id: String
│   │       ├── text: String
│   │       ├── source: String
│   │       ├── start_timestamp: String
│   │       ├── end_timestamp: String
│   │       ├── is_final: Bool
version: float64
```

## API

There is a "secret" API that can be accessed. The bearer token is stored at `supabase.json` in the same directory as the cache file. The base URL is <https://api.granola.ai/v2/>.

Set these headers.

```go
http.SetHeader("Content-Type", "application/json")
http.SetHeader("Authorization", "Bearer "+token)
http.SetHeader("User-Agent", "Granola/5.354.0")
http.SetHeader("X-Client-Version", "5.354.0")
```

### Endpoints

- get-documents
