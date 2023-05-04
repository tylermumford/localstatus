# LocalStatus checks

This Go package is for all of the checks that localstatus can perform.
Each check is named with a loose hierarchy.

Checks conform to the `Check` interface, which is defined in
the `interface.go` file.

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
