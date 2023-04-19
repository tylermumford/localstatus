// This file is for the TOML handling.

package brock

import (
	"fmt"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

// Runs the program. This is called from the main package.
func Run() error {
	fmt.Println("Running brock")

	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot determine home dir: %w", err)
	}

	tomlPath := path.Join(home, "brock_checks.toml")
	conf := RunConfig{
		TomlPath: tomlPath, // [1]
	}

	return innerRun(conf)
}

// [1] This should match the Readme example.

func innerRun(conf RunConfig) error {
	fmt.Printf("Running with toml %v\n", conf.TomlPath)

	tomlBytes, err := os.ReadFile(conf.TomlPath)
	if err != nil {
		return fmt.Errorf("cannot get toml from %v: %w", conf.TomlPath, err)
	}

	conf.TomlBytes = tomlBytes
	err = parseToml(&conf)
	if err != nil {
		return fmt.Errorf("cannot parse toml: %w", err)
	}

	return nil
}

func parseToml(conf *RunConfig) error {
	main := BrockConfig{}
	_, err := toml.Decode(string(conf.TomlBytes), &main)
	if err != nil {
		return err
	}

	fmt.Printf("Parsed: %+v", main)
	conf.Main = &main

	return nil
}
