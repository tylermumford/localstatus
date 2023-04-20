package brock

import "github.com/tylermumford/friendly-broccoli/checks"

// Holds per-execution configuration.
type RunConfig struct {
	TomlPath  string
	TomlBytes []byte
	Main      *BrockConfig
	Registry  *checks.CheckRegistry
}
