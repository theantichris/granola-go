# Repository Guidelines

## Project Structure & Modules
- `cmd/main.go` – CLI entry (flag parsing, I/O).
- `granola/` – core logic (cache parsing, transforms).
- `granola/*_test.go` – unit tests colocated with code.
- Outputs are written to the directory provided by `--output`.

## Build, Test, and Development
- Build: `go build -o granola-to-markdown.exe ./cmd` (use a POSIX binary name on macOS/Linux).
- Run: `go run ./cmd --cache path/to/cache-v3.json --output outdir`.
- Test (verbose): `go test -v ./...`.
- Coverage: `go test -cover ./...`.
- Static checks: `go vet ./...` (run before PRs).

## Coding Style & Naming
- Format with `go fmt ./...` (or `gofmt -s -w .`) and keep diffs clean.
- Use idiomatic Go: packages lowercase (`granola`), exported identifiers `CamelCase` with doc comments starting with the name, unexported `camelCase`.
- Prefer early returns; handle errors explicitly and wrap with context.
- File names lowercase with underscores as needed; test files end with `_test.go`.

## Testing Guidelines
- Use the standard `testing` package; prefer subtests with `t.Run(...)` for readability.
- When a subtest is independent, call `t.Parallel()` inside that subtest to run them concurrently.
- Name tests `TestXxx` matching the exported behavior (e.g., `TestParseCache`).
- Keep unit tests fast and deterministic; avoid touching the real cache file — use fixtures/strings.
- Example:
  ```go
  func TestParseCache(t *testing.T) {
    t.Run("empty cache", func(t *testing.T) {
      t.Parallel()
      // ... assertions
    })
    t.Run("valid doc", func(t *testing.T) {
      t.Parallel()
      // ... assertions
    })
  }
  ```
- Run `go test -v ./...` locally before pushing; include coverage when changing parsing logic.

## Commit & Pull Request Guidelines
- Commits: short, imperative subject (≤72 chars), optional scope. Examples:
  - `cmd: add --output flag`
  - `granola: extract title creation`
- Reference issues/PRs using `#123` when applicable.
- PRs: include a clear description, motivation, before/after behavior, and test coverage notes; add screenshots only if user-facing output changes.

## Security & Configuration Tips
- Do not commit `cache-v3.json` or any `supabase.json`/tokens. Keep secrets out of code and logs.
- Default cache paths (reference only):
  - Windows: `~\AppData\Roaming\Granola\cache-v3.json`
  - macOS: `~/Library/Application Support/Granola/cache-v3.json`

## Agent Notes
- Keep changes minimal and focused; do not break existing flags or output structure.
- Touch only relevant files; mirror existing patterns in `granola/` and update tests when behavior changes.

## CI Checklist
- `gofmt -s -l .` reports no files.
- `go vet ./...` passes.
- `go test -v ./...` passes (consider `-race` locally).
- Add/adjust tests for behavior changes; note coverage in PR description.

## Pre-push Hook (Optional but Recommended)
- Enable repo hooks: `git config core.hooksPath .githooks`.
- The hook blocks pushes if formatting, vet, or tests fail.
- Run manually if needed: `.githooks/pre-push`.
