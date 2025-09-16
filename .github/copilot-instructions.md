# Copilot Instructions for granola-to-markdown

## Project Overview

- **Purpose:** Converts Granola meeting notes (JSON) into Markdown files.
- **Language:** Go (see `main.go`).
- **Entry Point:** `main.go` contains the main logic and is the only source file.

## Architecture & Data Flow

- Reads a Granola meeting notes file (JSON format).
- Unmarshals JSON into Go structs (see `GranolaMeeting` in `main.go`).
- Iterates through the struct to generate Markdown output.
- Writes the Markdown to a file.

## Developer Workflows

- **Build:**
  - Run `go build -o granola-to-markdown.exe` to build the executable for Windows.
- **Run:**
  - Execute `./granola-to-markdown.exe` (or `go run main.go` for development).
- **Dependencies:**
  - No external Go dependencies are currently declared in `go.mod`.
- **Testing:**
  - No test files or test framework present as of now.

## Project Conventions

- All logic is currently in `main.go`.
- The struct `GranolaMeeting` is the placeholder for the data model; expand as needed to match the JSON schema.
- Input/output file paths and formats are not hardcoded—update `main.go` to handle arguments or config as needed.

## Integration Points

- Expects input in the form of a Granola meeting notes JSON file (see `granola-cache.json` for an example or placeholder).
- Output is a Markdown file, format and location to be defined in code.

## Key Files

- `main.go`: All application logic.
- `granola-cache.json`: Example or cache of input data.
- `README.md`: Project summary and usage.

## Example Usage

```sh
# Build
$ go build -o granola-to-markdown.exe
# Run
$ ./granola-to-markdown.exe
```

---

**Update this file if you add new files, workflows, or conventions.**
