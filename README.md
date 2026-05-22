# win-alias

![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)

Fast and lightweight alias manager for Windows CMD, bringing Linux-like `alias` functionality with Registry-backed persistence.

## Features
- **Registry Storage:** Aliases are stored in `HKCU\Software\win-alias\aliases`.
- **Session Support:** Uses `doskey` for immediate alias application.
- **Opt-in Persistence:** Automated CMD `AutoRun` integration via `--setup`.
- **Clean Architecture:** Built with the Standard Go Project Layout.

## Installation

### Via WinGet (Recommended)
Once the manifest is accepted by the community repository:
```cmd
winget install Nidzoki.win-alias
```

### Via Scoop
```cmd
scoop bucket add nidzoki https://github.com/Nidzoki/scoop-bucket
scoop install alias
```

### Via Go
```cmd
go install github.com/Nidzoki/win-alias/cmd/win-alias@latest
```
*Note: This will install as `win-alias`. You may want to rename it to `alias` in your GOPATH/bin.*

### Manual Installation
1. Open CMD as **Administrator**.
2. Run `install.bat`. This builds the tool as `alias.exe` and moves it to `C:\Windows`.

## Usage
Add or update an alias:
```cmd
alias gs="git status"
```

List active aliases:
```cmd
alias
```

Remove an alias:
```cmd
unalias gs
```

## Persistence
By default, aliases are only available in the current session. To make them persist across all new CMD windows:
```cmd
alias --setup
```

To disable persistence:
```cmd
alias --disable
```

## Uninstallation

### Via WinGet
```cmd
winget uninstall Nidzoki.win-alias
```

### Manual
1. Open CMD as **Administrator**.
2. Run `uninstall.bat`.

## Development
Build locally:
```cmd
go build -o alias.exe ./cmd/win-alias
```

Run tests:
```cmd
go test ./...
```
