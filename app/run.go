// This file is for the TOML handling.

package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
	"github.com/tylermumford/localstatus/checks"
)

// Runs the program. This is called from the main package.
// Any error should be displayed to the user.
func Run() error {
	// 1. Determine path to config file
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot determine home dir: %w", err)
	}

	tomlPath := filepath.Join(home, ".config", "localstatus.toml")
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
	fails := 0
	for v := range ctx.Config.Checks {
		result := runCheckDefinition(ctx.Config.Checks[v], &ctx)

		var prefix string
		var message string
		if !result.IsOkay() {
			prefix = color.RedString("! ")
			message = result.Label()
			fails += 1
		} else {
			prefix = color.GreenString("OK")
			message = result.Label()
		}

		fmt.Printf("%s  %s\n", prefix, message)
	}

	if fails > 0 {
		return fmt.Errorf("%d checks failed", fails)
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
	main := GatheredConfig{}
	_, err := toml.Decode(string(conf.TomlBytes), &main)
	if err != nil {
		return err
	}

	conf.Config = &main

	return nil
}
