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

## Cache file

On Windows this is stored at `~\AppData\Roaming\Granola\cache-v3.json`.

## Granola cache scheme

The cache file is in JSON with a wrapper property `cache` that's value is a JSON string. That contains a `state` object that has all the information including Google calendar events and people information.

The only part in `state` this project is currently concerned with are the `documents` which are the actual meeting notes. `documents` is an mapped object, mapped by the UUID.

It doesn't seem that the cache file contains notes made from the transcripts, only the notes you type in yourself.

```json
{
  "state": {
    "documents": {
      "UUID": {
        // UUID of the meeting
        "id": "string", // UUID of the meeting
        "title": "string", // The title of the meeting
        "created_at": "string", // # ISO 8601
        "updated_at": "string", // # ISO 8601
        "notes_markdown": "string", // Meeting notes in Markdown format, might be missing
        "notes_plain": "string", // Meeting notes in plain text format, might be missing
        "notes": {
          // TipTap rich content
          "type": "doc", // The type of notes root node
          "content": [
            {
              "type": "heading", // The type of block, heading|paragraph
              "attrs": {
                // Optional attributes, varies by block
                "level": 1 // The level for headings
              },
              "content": [{ "type": "text", "text": "string" }]
            },
            {
              "type": "paragraph", // The type of block, heading|paragraph
              "content": [{ "type": "text", "text": "string" }]
            }
          ]
        }
      }
    }
  }
}
```
