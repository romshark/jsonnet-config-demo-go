// Package config provides a jsonnet configuration file evaluator.
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/go-jsonnet"
)

var ErrFileNotFound = errors.New("config file not found")

// Parse evaluates and unmarshals the contents of the given file
// to config and returns an error if the file doesn't exist,
// couldn't be evaluated, or can't be unmarshaled to config.
func Parse[T any](filePath string) (config T, parsing time.Duration, err error) {
	start := time.Now()
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return config, 0, fmt.Errorf("%w: %q", ErrFileNotFound, filePath)
	}

	vm := jsonnet.MakeVM()
	jsonStr, err := vm.EvaluateFile(filePath)
	if err != nil {
		return config, 0, fmt.Errorf("evaluating Jsonnet: %w", err)
	}
	if err := json.Unmarshal([]byte(jsonStr), &config); err != nil {
		return config, 0, fmt.Errorf("unmarshaling JSON to %T: %w", config, err)
	}

	v := validator.New(validator.WithRequiredStructEnabled())
	err = v.Struct(config)
	return config, time.Since(start), err
}

// MustParse panics if the file doesn't exist, couldn't be evaluated,
// or can't be unmarshaled to config.
func MustParse[T any](filePath string) (config T, parsing time.Duration) {
	config, parsing, err := Parse[T](filePath)
	if err != nil {
		panic(err)
	}
	return config, parsing
}
