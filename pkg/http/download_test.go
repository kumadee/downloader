package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownloadFile(t *testing.T) {
	var cases = []struct {
		url          string
		path         string
		expectedFile string
		description  string
	}{
		{
			url:          "http://public.animeout.xyz/nimbus.animeout.com/series/Ongoing/Overlord Ple Ple Pleiades/[AnimeOut] Overlord Ple Ple Pleiades - 01 - Special [720p][AAC][DeadFish][RapidBot].mkv",
			path:         "/tmp/overlord-specials",
			expectedFile: "/tmp/overlord-specials/[AnimeOut] Overlord Ple Ple Pleiades - 01 - Special [720p][AAC][DeadFish][RapidBot].mkv",
			description:  "Download file at non-existent path",
		},
	}
	for _, tc := range cases {
		DownloadFile(tc.path, tc.url)
		assert.DirExists(t, tc.path, tc.description)
		assert.FileExists(t, tc.expectedFile, tc.description)
	}
}
