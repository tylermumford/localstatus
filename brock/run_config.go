package brock

// Holds per-execution configuration.
type RunConfig struct {
	TomlPath  string
	TomlBytes []byte
	Main      *BrockConfig
}