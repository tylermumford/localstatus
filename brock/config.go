package brock

import "github.com/tylermumford/friendly-broccoli/checks"

// Holds the data parsed from the TOML configuration file. Very important.
type BrockConfig struct {
	Checks []checks.Params
}

type Definition struct {
	Check              string
	Url                string
	Variables_Required []string
}
