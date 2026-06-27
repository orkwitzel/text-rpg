# Claude Guidelines

## Build and Verify

Always verify changes compile before reporting done:

```bash
go build ./...
go vet ./...
```

## Code Style

- No comments unless the WHY is non-obvious
- Prefer explicit error returns over panics in library code
- CLI code (`cmd/`) may call `os.Exit` directly
- Keep `cmd` and `internal` packages strictly separated — `internal` must not import `cmd`

## Architecture

The game loop in `main.go` is intentionally simple:
1. `cmd.SaveFilesMenu()` — pick or create a game
2. Loop: `cmd.GameInput(&game)` then `savefiles.SaveGame(game)`

Extend `GameInput` in `cmd/gameCmd.go` to add new commands. Add domain logic to the appropriate `internal` package.

## Adding Tile Types or Directions

- Tile types: add a constant in `internal/game/world/tiles/tile.go` and include it in the `types` slice inside `NewRandom()`
- Directions: add a constant in `internal/game/game.go` and handle it in `MovePlayer()`

## Save Format

Save files are JSON at `~/rpg/saveFiles/<game-id>.json`. The format is stable — adding new fields is safe (they deserialize to zero values); removing or renaming fields is a breaking change.

## Testing

No tests exist yet. When writing tests:
- Place `*_test.go` files next to the code under test
- Run with `go test ./...`
- Use table-driven tests for command parsing and tile/direction logic
