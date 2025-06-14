# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go utility called `ps-claude` that serves as a Windows PowerShell wrapper for the Claude CLI tool. It bridges the gap between Windows environments and WSL (Windows Subsystem for Linux) where Claude runs.

## Architecture

The application consists of a single Go binary that:
1. Converts Windows paths to WSL mount paths (`/mnt/c/...` format)
2. Launches WSL with an interactive bash session
3. Locates the Claude CLI executable in the WSL environment
4. Forwards all arguments and maintains interactive I/O streams

Key functions:
- `convertWindowsPathToWSL()`: Handles path conversion from Windows format to WSL mount points
- `getCurrentDirectory()`: Gets current working directory and converts to WSL format
- Main execution flow manages WSL command construction and process handling

## Common Commands

### Build
```bash
make build
# or manually:
GOOS=windows GOARCH=amd64 go build -o ps-claude.exe
```

### Clean
```bash
make clean
```

### Development
Since this is a simple Go utility with no external dependencies:
- Standard Go tooling applies (`go run`, `go test`, etc.)
- Cross-compilation target is Windows AMD64
- No special test frameworks or linting tools configured

### Testing
To test from WSL:
```bash
make build
cp ps-claude.exe /mnt/c/temp/
powershell.exe -Command "cd C:\temp; .\ps-claude.exe --help"
```

## Known Issues
- Original version had performance issue with `find /home` command that caused hanging
- Fixed by using simpler claude path resolution (checks `which claude` first, then fallback path)