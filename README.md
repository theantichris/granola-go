# granola-to-markdown

A Go application that exports Granola meeting notes to Markdown files.

## Granola cache scheme

The cache file is in JSON with a wrapper property `cache` that's value is a JSON string. That contains a `state` object that has all the information including Google calendar events and people information.

The only part in `state` this project is currently concerned with are the `documents` which are the actual meeting notes. `documents` is an mapped object, mapped by the UUID.

```json
{
  "state": {
    "documents": {
      "UUID": {                                                    // UUID of the meeting
        "id": "string",                                            // UUID of the meeting
        "title": "string",                                         // The title of the meeting
        "created_at": "string",                                    // # ISO 8601
        "updated_at": "string",                                    // # ISO 8601
        "notes_markdown": "string",                                // Meeting notes in Markdown format, might be missing
        "notes_plain": "string",                                   // Meeting notes in plain text format, might be missing
        "notes": {                                                 // TipTap rich content
          "type": "doc",                                           // The type of notes; doc, heading, paragraph
          "content": [
            {
              "type": "heading",                                   // The type of block, heading|paragraph
              "attrs": {                                           // Optional attributes, varies by block
              "level": 1,                                          // The level for headings, 1|2|3
              "content": [{"type": "text", "text", "string"}]
              }
            },
            {
              "type": "paragraph",                                 // The type of block, heading|paragraph
              "content": [{ "type": "text", "text": "string" }]
            }
          ]
        }
      }
    }
  }
}
```
