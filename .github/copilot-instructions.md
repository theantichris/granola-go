# Copilot Instructions for granola-to-markdown

## Project Overview

- **Purpose:** Converts Granola meeting notes (JSON) into Markdown files.
- **Language:** Go (see `main.go`).
- **Entry Point:** `main.go` contains the main logic and is the only source file.

## Architecture & Data Flow

- Reads a Granola cache file (JSON format) with a wrapper property `cache` (a JSON string).
- Unmarshals the outer JSON, then the inner JSON string, into Go structs (see `Cache`, `State`, `Document`).
- The `notes` field in each document is a TipTap JSON structure (see README for schema).
- Iterates through the struct to generate Markdown output.
- Writes the Markdown to a file.

## Developer Workflows

- **Build:**
  - Run `go build -o granola-to-markdown.exe` to build the executable for Windows.
- **Run:**
  - Execute `./granola-to-markdown.exe` (or `go run main.go` for development).
- **Testing:**
  - Run `go test` to execute unit tests in `main_test.go`. Tests use in-memory JSON, not files.

## Project Conventions

- All logic is currently in `main.go`.
- Input/output file paths and formats are hardcoded.

## Integration Points

- Expects input in the form of a Granola cache JSON file.
- Output is a Markdown file, format.

## Key Files

- `main.go`: All application logic.
- `README.md`: Project summary, data schema, and usage.

## Example Usage

```sh
# Build
$ go build -o granola-to-markdown.exe
# Run
$ ./granola-to-markdown.exe
# Test
$ go test
```

---

**Update this file if you add new files, workflows, or conventions.**
