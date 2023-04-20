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
	ctx := RunContext{
		TomlPath: tomlPath, // [1]
	}

	// 2. Read and parse
	tomlBytes, err := os.ReadFile(ctx.TomlPath)
	if err != nil {
		return fmt.Errorf("cannot get toml from %v: %w", ctx.TomlPath, err)
	}

	ctx.TomlBytes = tomlBytes
	err = parseToml(&ctx)
	if err != nil {
		return fmt.Errorf("cannot parse toml: %w", err)
	}

	// 3. Prepare code for checks
	ctx.Registry = checks.NewCheckRegistry()
	ctx.Registry.AddAllChecks()

	// 4. Run the specified checks
	fmt.Printf("    %d checks to run...\n", len(ctx.Config.Checks))
	for v := range ctx.Config.Checks {
		result := runCheckDefinition(ctx.Config.Checks[v], &ctx)

		var prefix string
		var message string
		if !result.IsOkay() {
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

func runCheckDefinition(p checks.Params, conf *RunContext) checks.CheckResult {
	name := p.GetString("check")
	check := conf.Registry.Get(name)
	return check.Run(p)
}

// [1] This should match the Readme example.

func parseToml(conf *RunContext) error {
	main := BrockConfig{}
	_, err := toml.Decode(string(conf.TomlBytes), &main)
	if err != nil {
		return err
	}

	fmt.Printf("Parsed: %+v\n", main)
	conf.Config = &main

	return nil
}
