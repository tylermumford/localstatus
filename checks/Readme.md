# Brock checks

This Go package is for all of the checks that brock can perform.
Each check is named with a loose hierarchy, and most (maybe all)
start with `brock`.

Checks conform to the `Check` interface, which is defined in
the `check.go` file.

Eventually, there will be many checks.

In scope:

- Network requests
- Files and directories
- Environment variables
- Running custom scripts
- SQL queries
- Language dependencies (npm, Gems, etc.)
- Git branches, remotes, etc.
- Programs on PATH, including versions

Not in scope:

- Cloud providers (AWS, etc.)
- SSH/remote shells
