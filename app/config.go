package app

import "github.com/tylermumford/localstatus/checks"

// Holds the data parsed from the TOML configuration file. Very important.
type GatheredConfig struct {
	Checks []checks.Params
}
