# ps-claude

A Windows PowerShell wrapper for the Claude CLI tool that bridges Windows environments with WSL (Windows Subsystem for Linux). You must have already have installed claude in your default WSL2 distro.

## What it does

`ps-claude` allows you to run the Claude CLI from Windows PowerShell/Command Prompt by automatically:
- Converting Windows paths to WSL mount paths (`C:\path` → `/mnt/c/path`)
- Launching WSL with proper environment
- Forwarding all arguments to Claude CLI
- Maintaining interactive I/O streams

## Installation

### One-line PowerShell Install
```powershell
iwr -useb https://github.com/standardbeagle/ps-claude/releases/latest/download/ps-claude.exe -outfile ps-claude.exe
```

### NPM Install
```bash
npm install -g @standardbeagle/ps-claude
```

## Usage

Once installed, use `ps-claude` exactly like you would use `claude`:

```powershell
# Start Claude
ps-claude

# Pass arguments
ps-claude --help
ps-claude -p "Hello Claude"

# Works from any Windows directory
cd C:\MyProject
ps-claude
```

## Requirements

- Windows 10/11 with WSL enabled
- Claude CLI installed in WSL environment
- Go 1.19+ (for building from source)

## Building from Source

```bash
make build
```

Or manually:
```bash
GOOS=windows GOARCH=amd64 go build -o ps-claude.exe
```

## How it Works

1. Detects current Windows working directory
2. Converts path to WSL mount format (`/mnt/c/...`)
3. Launches WSL with interactive bash session
4. Locates Claude CLI in WSL environment
5. Forwards all arguments and maintains I/O streams