// This file is for the TOML handling.

package brock

import (
	"fmt"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/tylermumford/friendly-broccoli/checks"
)

// Runs the program. This is called from the main package.
// Any error should be displayed to the user.
func Run() error {
	fmt.Println("Running brock")

	// 1. Determine path to config file
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot determine home dir: %w", err)
	}

	tomlPath := path.Join(home, "brock_checks.toml")
	conf := RunConfig{
		TomlPath: tomlPath, // [1]
	}

	// 2. Read and parse
	tomlBytes, err := os.ReadFile(conf.TomlPath)
	if err != nil {
		return fmt.Errorf("cannot get toml from %v: %w", conf.TomlPath, err)
	}

	conf.TomlBytes = tomlBytes
	err = parseToml(&conf)
	if err != nil {
		return fmt.Errorf("cannot parse toml: %w", err)
	}

	// 3. Prepare code for checks
	conf.Registry = checks.NewCheckRegistry()
	conf.Registry.AddAllChecks()

	// 4. Run the specified checks
	fmt.Printf("    %d checks to run...\n", len(conf.Main.Checks))
	for v := range conf.Main.Checks {
		result, err := runCheckDefinition(conf.Main.Checks[v], &conf)

		var prefix string
		var message string
		if err != nil {
			prefix = "! "
			message = err.Error()
		} else if !result.IsOkay() {
			prefix = "! "
			message = result.Label()
		} else {
			prefix = "OK"
			message = result.Label()
		}

		fmt.Printf("%s  %s\n", prefix, message)
	}

	return nil
}

func runCheckDefinition(def Definition, conf *RunConfig) (checks.CheckResult, error) {
	name := def.Check
	check := conf.Registry.Get(name)
	options := map[string]any{"url": def.Url}
	return check.Run(options)
}

// [1] This should match the Readme example.

func parseToml(conf *RunConfig) error {
	main := BrockConfig{}
	_, err := toml.Decode(string(conf.TomlBytes), &main)
	if err != nil {
		return err
	}

	fmt.Printf("Parsed: %+v\n", main)
	conf.Main = &main

	return nil
}
