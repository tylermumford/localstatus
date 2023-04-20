package brock

// Holds the data parsed from the TOML configuration file. Very important.
type BrockConfig struct {
	Checks []Definition
}

type Definition struct {
	Check string
	Url   string
}
