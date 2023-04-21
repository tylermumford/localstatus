package app

import "github.com/tylermumford/localstatus/checks"

// Holds per-execution configuration.
type RunContext struct {
	// TOML file to load.
	TomlPath string
	// Raw bytes loaded.
	TomlBytes []byte
	// Parsed configuration from the TOML file.
	Config *GatheredConfig
	// Checks available to run, if configured.
	Registry *checks.CheckRegistry
}
