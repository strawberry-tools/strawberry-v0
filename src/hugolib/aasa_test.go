// Copyright 2020 Gotham Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package hugolib

import (
	"testing"

	"github.com/gothamhq/gotham/deps"

	qt "github.com/frankban/quicktest"
)

func TestAASAOutput(t *testing.T) {

	t.Parallel()

	testVersions := []int{0, 1, 2, 42} // assuming that 42 is not a valid version

	testCases := []struct {
		prefix  string
		bundle  string
		passing bool
	}{
		{
			"ABCDE12345",
			"com.gothamhq.ios",
			true,
		},
		{
			"",
			"com.gothamhq.ios",
			false,
		},
		{
			"ABCDE12345",
			"",
			false,
		},
		{
			"",
			"",
			false,
		},
	}

	for _, tc := range testCases {

		for _, tv := range testVersions {
			c := qt.New(t)
			cfg, fs := newTestCfg()
			cfg.Set("baseURL", "http://gotham/test/")
			cfg.Set("aasaPrefix", tc.prefix)
			cfg.Set("aasaBundle", tc.bundle)

			if tv != 0 {
				cfg.Set("aasaVersion", tv)
			}

			depsCfg := deps.DepsCfg{Fs: fs, Cfg: cfg}

			writeSourcesToSource(t, "content", fs, weightedSources...)

			var s *Site

			if tv == 42 {
				s = buildSingleSiteExpected(t, false, true, depsCfg, BuildCfg{})
				return
			} else {
				s = buildSingleSite(t, depsCfg, BuildCfg{})
			}

			th := newTestHelper(s.Cfg, s.Fs, t)
			outputAASA := "public/.well-known/apple-app-site-association"

			if !tc.passing {

				th.assertFileNotExist(outputAASA)
				return
			}

			realVersion := cfg.Get("aasaVersion")

			if realVersion == 2 {
				th.assertFileContent(outputAASA, "\"components\":")
			} else if realVersion == 1 {
				th.assertFileContent(outputAASA, "\"apps\": []")
			} else {
				t.Errorf("Error: %d is not a valid AASA version.", realVersion)
			}

			th.assertFileContent(outputAASA,
				tc.prefix,
				tc.bundle,
			)

			content := readDestination(th, th.Fs, outputAASA)
			c.Assert(content, qt.Contains, tc.prefix)
			c.Assert(content, qt.Contains, tc.bundle)
		}
	}
}
