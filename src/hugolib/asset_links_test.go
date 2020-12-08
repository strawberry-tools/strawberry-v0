// Copyright 2020 Gotham Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package hugolib

import (
	"testing"

	"github.com/gothamhq/gotham/deps"

	qt "github.com/frankban/quicktest"
)

func TestAssetLinksOutput(t *testing.T) {

	t.Parallel()

	testCases := []struct {
		pkgName     string
		fingerprint string
		passing     bool
	}{
		{
			"com.gothamhq.android",
			"a0df0154fa3fa814f19a1fdf8ba8f8a6330a5f8e1355de6fa65717b80387f0ec",
			true,
		},
		{
			"com.gothamhq.android",
			"",
			false,
		},
		{
			"",
			"a0df0154fa3fa814f19a1fdf8ba8f8a6330a5f8e1355de6fa65717b80387f0ec",
			false,
		},
		{
			"",
			"",
			false,
		},
	}

	for _, tc := range testCases {

		c := qt.New(t)
		cfg, fs := newTestCfg()
		cfg.Set("baseURL", "http://gotham/test/")
		cfg.Set("assetLinksPackageName", tc.pkgName)
		cfg.Set("assetLinksFingerprint", tc.fingerprint)

		depsCfg := deps.DepsCfg{Fs: fs, Cfg: cfg}

		writeSourcesToSource(t, "content", fs, weightedSources...)
		s := buildSingleSite(t, depsCfg, BuildCfg{})
		th := newTestHelper(s.Cfg, s.Fs, t)
		outputAssetLinks := "public/.well-known/assetlinks.json"

		if !tc.passing {

			th.assertFileNotExist(outputAssetLinks)
			return
		}

		th.assertFileContent(outputAssetLinks,
			tc.pkgName,
			tc.fingerprint,
		)

		content := readDestination(th, th.Fs, outputAssetLinks)
		c.Assert(content, qt.Contains, tc.pkgName)
		c.Assert(content, qt.Contains, tc.fingerprint)
	}
}
