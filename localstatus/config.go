package localstatus

import "github.com/tylermumford/friendly-broccoli/checks"

// Holds the data parsed from the TOML configuration file. Very important.
type GatheredConfig struct {
	Checks []checks.Params
}
