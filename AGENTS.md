# AGENTS Guidelines for This Repository

This repository contains a Go CLI application that exports Granola meeting notes to Markdown files. The application reads from Granola's local cache files (JSON format) and converts meeting documents to individual Markdown files. When working on this project, please follow the guidelines below to maintain code quality and consistency.

## 1. Follow TDD and Go Best Practices

- **Always write tests first** using the Test-Driven Development approach. Write unit tests for all new functionality before implementing features.
- **Use descriptive variable and function names** - prefer `cache` over `c`, `transcriptData` over `td`, etc. Code should be self-documenting.
- **Prefer standard Go packages** unless there's a clear advantage to using external dependencies.
- Run `gofmt` to ensure consistent formatting across the codebase.

## 2. Error Handling Standards

- **Use sentinel errors** for expected error conditions that callers need to handle.
- **Wrap errors with `%w`** to maintain error chains and provide context.
- **Use `errors.Is()`** for error assertions and comparisons, not string matching.

Example:

```go
var ErrCacheNotFound = errors.New("cache file not found")

if err := readCache(); err != nil {
    return fmt.Errorf("failed to process cache: %w", err)
}

if errors.Is(err, ErrCacheNotFound) {
    // handle missing cache
}
```

## 3. Logging Requirements

- **Use `log/slog` exclusively** for all logging needs.
- **Never use `fmt` for logging in library code** - this includes `fmt.Printf`, `fmt.Println`, etc.
- **Output separation is critical**:
  - All logs must go to `stderr`
  - Normal output goes to `stdout`

## 4. Project Structure Notes

The current structure has core functionality in `granola`. Keep the CLI entry point in `cmd`.

## 5. Testing and CI

- Run `go test ./...` before every commit to ensure all tests pass.
- The project uses GitHub Actions (`.github/workflows/go.yml`) for automated building and testing.
- All CI checks must pass before merging any PR.

## 6. Branching and Commits

- **Branch naming**: Use descriptive prefixes like `feature/transcripts` or `bug/fix-bad-code`.
- **Commit messages**: Use imperative tense - "add transcript ID", "fix unmarshalling error", not "dded" or "fixed".
- **PR titles**: Follow the same imperative style as commit messages.

## 7. Application Context

The application works with Granola cache files located at:

- **Windows**: `~\AppData\Roaming\Granola\cache-v3.json`
- **macOS**: `~/Library/Application Support/Granola/cache-v3.json`

Authentication for the Granola API uses a bearer token from `supabase.json` in the same directory as the cache file.

## 8. Useful Commands

| Command                                               | Purpose                 |
| ----------------------------------------------------- | ----------------------- |
| `go build -o granola-to-markdown.exe ./cmd/granola`   | Build the binary        |
| `go test ./...`                                       | Run all tests           |
| `go test -cover ./...`                                | Run tests with coverage |
| `./granola-to-markdown.exe --cache=path --output=dir` | Run with custom options |

---

Following these guidelines ensures consistent, maintainable code that integrates well with the existing codebase and CI pipeline. When in doubt, prioritize descriptive naming and comprehensive testing over brevity.
