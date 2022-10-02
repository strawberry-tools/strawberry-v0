// Copyright 2020 Gotham Authors. All rights reserved.
// Copyright 2021 Ricardo N Feliciano. All rights reserved.
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
// Versioning version number. Used for the StrawberryVersion as well as any other
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

// PrintStrawberryVersion - Build complete version string for user via CLI
func PrintStrawberryVersion(vType VersionType) string {

	if vType == VersionShort {
		return StrawberryVersion.String()
	}

	version := "Strawberry v" + StrawberryVersion.String()

	bi := getBuildInfo()

	if strings.Contains(version, "-dev") {
		version += "-" + strings.ToUpper(bi.Revision)
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

	date := bi.RevisionTime
	if date == "" {
		date = "unknown"
	}

	version += "BuildDate: " + date + "\n"
	version += "Platform: " + runtime.GOOS + "/" + runtime.GOARCH

	return version
}
