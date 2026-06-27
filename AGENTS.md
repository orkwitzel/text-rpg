# Agent Guidelines

This document describes conventions for AI agents contributing to this repository.

## Language and Toolchain

- Go 1.21+, module path `rpg`
- Single external dependency: `github.com/google/uuid`
- Run `go build ./...` and `go vet ./...` before considering any change complete

## Project Layout

```
cmd/          # CLI layer — input helpers, menus, command dispatch
internal/     # Domain logic — not importable outside this module
  game/       # Game state, player, world, tiles
  saveFiles/  # JSON persistence
main.go       # Wires cmd and saveFiles together
```

Keep the split clean: `cmd` handles user I/O and calls `internal` packages; `internal` packages never import `cmd`.

## Conventions

- No unused imports or variables — the Go compiler enforces this
- Error handling: print to `os.Stderr` and exit with `os.Exit(1)` in CLI code; return errors from library code
- JSON tags on all exported struct fields that cross the persistence boundary
- Tile types, directions, and other domain enumerations are typed `string` constants — do not use bare strings in switch statements

## Testing

There are no tests yet. When adding tests, place them alongside the code they test (`*_test.go`), use `go test ./...` to run all tests, and prefer table-driven tests.

## Save File Location

Save files live at `~/rpg/saveFiles/`. Do not commit save files or test data that contains personal paths.

## Out of Scope

- Do not add a UI framework or external terminal library without discussing first
- Do not change the save file format without a migration path
