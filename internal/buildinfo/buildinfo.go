// Package buildinfo exposes compile-time metadata shared across the server.
package buildinfo

import (
	"bytes"
	_ "embed"
	"strings"
)

//go:embed VERSION.md
var versionFile []byte

// The following variables are overridden via ldflags during release builds.
// Defaults cover local development builds.
var (
	// Version is the semantic version or git describe output of the binary.
	// It falls back to the embedded VERSION.md file if not set via ldflags.
	Version = "dev"

	// Commit is the git commit SHA baked into the binary.
	Commit = "none"

	// BuildDate records when the binary was built in UTC.
	BuildDate = "unknown"
)

func init() {
	if Version == "dev" || Version == "" {
		if len(versionFile) > 0 {
			Version = strings.TrimSpace(string(bytes.TrimSpace(versionFile)))
		}
	}
}
