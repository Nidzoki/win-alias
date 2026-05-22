# win-alias: Windows CMD Alias Manager

Lightweight Go tool for Linux-like aliases in Windows CMD using Registry persistence and `doskey`.

## Project Overview
- **Type:** Go CLI Tool
- **Language:** Go (1.21+)
- **Architecture:** 
  - `cmd/win-alias/`: Entry point.
  - `internal/alias/`: Core logic (Parser, Registry CRUD, doskey Apply).
- **Storage:** Windows Registry (`HKCU\Software\win-alias\aliases`).
- **Integration:** CMD `AutoRun` runs `alias --load`.

## Building and Running
- **Build:** `go build -o alias.exe ./cmd/win-alias`
- **Test:** `go test ./...`
- **Install (Admin required):** Run `install.bat`.
- **Uninstall (Admin required):** Run `uninstall.bat`.
- **Setup:** `alias --setup` (injects AutoRun into Registry).
- **Disable:** `alias --disable` (removes AutoRun from Registry).
- **Load:** `alias --load` (applies `doskey` macros).

## Development Conventions
- **Command Syntax:** `alias name="command"` (e.g., `alias gs="git status"`).
- **Subcommands:**
  - `alias`: List all from Registry.
  - `unalias name`: Remove from Registry.
  - `--setup`: Configure persistent AutoRun.
  - `--disable`: Remove persistent AutoRun.
  - `--load`: Export/Apply macros for current session.
- **Error Handling:** Exit code 1 on failure.
