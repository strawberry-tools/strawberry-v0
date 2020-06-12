// Copyright 2020 Gotham Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package hugo

import (
	"fmt"
	"runtime"
	"strings"
)

type VersionType int

const (
	VersionShort VersionType = iota
	VersionRegular
	VersionDetailed
)

// SemVerVersion represents a simplified representation of a Semantic
// Versioning version number. Used for the GothamVersion as well as any other
// SemVer versions.
type SemVerVersion struct {
	Major  uint
	Minor  uint
	Patch  uint
	Suffix string
}

// Return the SemVer version as a string representation.
func (v SemVerVersion) String() string {

	if v.Suffix == "" {
		return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	} else {
		return fmt.Sprintf("%d.%d.%d-%s", v.Major, v.Minor, v.Patch, v.Suffix)
	}
}

// Build complete version string for user via CLI
func PrintGothamVersion(vType VersionType) string {

	if vType == VersionShort {
		return GothamVersion.String()
	}

	version := "Gotham v" + GothamVersion.String()

	if commitHash != "" {
		version += "-" + strings.ToUpper(commitHash)
	}

	version += " (compatible with Hugo v" + CurrentVersion.String()

	if IsExtended {
		version += "/extended"
	}

	version += ")"

	if vType == VersionRegular {
		return version
	}

	version += "\n"

	date := buildDate
	if date == "" {
		date = "unknown"
	}

	version += "BuildDate: " + date + "\n"
	version += "Platform: " + runtime.GOOS + "/" + runtime.GOARCH

	return version
}
