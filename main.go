package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
)

// convertWindowsPathToWSL converts a Windows path to WSL mount path
func convertWindowsPathToWSL(windowsPath string) string {
	// Handle UNC paths (\\server\share)
	if strings.HasPrefix(windowsPath, "\\\\") {
		return windowsPath // Keep UNC paths as-is for now
	}
	
	// Handle drive letters (C:\path -> /mnt/c/path)
	driveRegex := regexp.MustCompile(`^([A-Za-z]):(.*)`)
	if matches := driveRegex.FindStringSubmatch(windowsPath); matches != nil {
		drive := strings.ToLower(matches[1])
		path := strings.ReplaceAll(matches[2], "\\", "/")
		return fmt.Sprintf("/mnt/%s%s", drive, path)
	}
	
	// If it's already a Unix-style path, return as-is
	return windowsPath
}

// getCurrentDirectory gets the current working directory and converts it to WSL format
func getCurrentDirectory() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	
	// Convert to WSL path format
	wslPath := convertWindowsPathToWSL(cwd)
	return wslPath, nil
}

func main() {
	// Get current directory in WSL format
	wslCwd, err := getCurrentDirectory()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting current directory: %v\n", err)
		os.Exit(1)
	}
	
	// Prepare WSL command with interactive login shell
	args := []string{"wsl", "-e", "bash", "-l", "-c"}
	
	// Build the command to run in WSL with proper environment loading
	// Try common claude locations without expensive find operation
	claudeCmd := fmt.Sprintf("cd '%s' && CLAUDE_PATH=$(which claude 2>/dev/null || echo '/home/beagle/.npm-global/bin/claude') && exec \"$CLAUDE_PATH\"", wslCwd)
	
	// Add any additional arguments passed to ps-claude
	if len(os.Args) > 1 {
		// Escape and append arguments
		escapedArgs := make([]string, len(os.Args)-1)
		for i, arg := range os.Args[1:] {
			escapedArgs[i] = fmt.Sprintf("'%s'", strings.ReplaceAll(arg, "'", "'\"'\"'"))
		}
		claudeCmd += " " + strings.Join(escapedArgs, " ")
	}
	
	args = append(args, claudeCmd)
	
	// Create the command
	cmd := exec.Command(args[0], args[1:]...)
	
	// Connect stdin, stdout, stderr to maintain interactive session
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	// Run the command and wait for it to complete
	err = cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			// Get the exit code from WSL/claude
			if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
				os.Exit(status.ExitStatus())
			}
		}
		fmt.Fprintf(os.Stderr, "Error running command: %v\n", err)
		os.Exit(1)
	}
}