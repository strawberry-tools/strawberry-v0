// Copyright 2020 Gotham Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package hugo

import (
	"runtime"
	"testing"
)

func TestPrintGothamVersion(t *testing.T) {

	testCases := []struct {
		vType      VersionType
		version    SemVerVersion
		date       string
		commitHash string
		expected   string
	}{
		{
			VersionShort,
			SemVerVersion{
				Major:  0,
				Minor:  1,
				Patch:  2,
				Suffix: "dev",
			},
			"",
			"",
			"0.1.2-dev",
		},
		{
			VersionRegular,
			SemVerVersion{
				Major:  0,
				Minor:  5,
				Patch:  0,
				Suffix: "dev",
			},
			"1991-01-21T00:00:00-0500",
			"IMaHASH",
			"Gotham v0.5.0-dev-IMAHASH (compatible with Hugo v" + CurrentVersion.String() + "/extended)",
		},
		{
			VersionRegular,
			SemVerVersion{
				Major:  0,
				Minor:  5,
				Patch:  0,
				Suffix: "",
			},
			"1991-01-21T00:00:00-0500",
			"IMaHASH",
			"Gotham v0.5.0 (compatible with Hugo v" + CurrentVersion.String() + "/extended)",
		},
		{
			VersionDetailed,
			SemVerVersion{
				Major:  1,
				Minor:  0,
				Patch:  2,
				Suffix: "",
			},
			"2020-06-04T00:00:00-0500",
			"",
			"Gotham v1.0.2 (compatible with Hugo v" + CurrentVersion.String() + "/extended)\nBuildDate: 2020-06-04T00:00:00-0500\nPlatform: " + runtime.GOOS + "/" + runtime.GOARCH,
		},
	}

	for i, tc := range testCases {

		if tc.date == "" {
			buildDate = "unknown"
		} else {
			buildDate = tc.date
		}

		commitHash = tc.commitHash

		GothamVersion = tc.version
		actual := PrintGothamVersion(tc.vType)

		if actual != tc.expected {
			t.Errorf("PrintGothamVersion test[%d]:\nexpected %s,\n\nactual %s", i, tc.expected, actual)
		}
	}
}
