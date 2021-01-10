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

	testCases := []struct {
		prefix  string
		bundle  string
		version int
		passing bool
	}{
		{
			"ABCDE12345",
			"com.gothamhq.ios",
			2,
			true,
		},
		{
			"",
			"com.gothamhq.ios",
			2,
			false,
		},
		{
			"ABCDE12345",
			"",
			2,
			false,
		},
		{
			"",
			"",
			2,
			false,
		},
		{
			"ABCDE12345",
			"com.gothamhq.ios",
			1,
			true,
		},
		{
			"ABCDE12345",
			"",
			1,
			false,
		},
		{
			"",
			"com.gothamhq.ios",
			1,
			false,
		},
		{
			"",
			"",
			1,
			false,
		},
		{
			"",
			"",
			0,
			false,
		},
	}

	for _, tc := range testCases {

		c := qt.New(t)
		cfg, fs := newTestCfg()
		cfg.Set("baseURL", "http://gotham/test/")
		cfg.Set("aasaPrefix", tc.prefix)
		cfg.Set("aasaBundle", tc.bundle)
		cfg.Set("aasaVersion", tc.version)

		depsCfg := deps.DepsCfg{Fs: fs, Cfg: cfg}

		writeSourcesToSource(t, "content", fs, weightedSources...)
		s := buildSingleSite(t, depsCfg, BuildCfg{})
		th := newTestHelper(s.Cfg, s.Fs, t)
		outputAASA := "public/.well-known/apple-app-site-association"

		if !tc.passing {

			th.assertFileNotExist(outputAASA)
			return
		}

		if cfg.Get("aasaVersion") == 2 {
			th.assertFileContent(outputAASA, "\"components\":")
		} else if cfg.Get("aasaVersion") == 1 {
			th.assertFileContent(outputAASA, "\"apps\": []")
		} else {
			t.Error("Error: Not a valid AASA version.")
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
