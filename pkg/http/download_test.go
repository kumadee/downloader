package http

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDownloadFile(t *testing.T) {
	var cases = []struct {
		url          string
		expectedFile string
		description  string
	}{
		{
			url:          "http://public.animeout.xyz/nimbus.animeout.com/series/Ongoing/Overlord Ple Ple Pleiades/[AnimeOut] Overlord Ple Ple Pleiades - 01 - Special [720p][AAC][DeadFish][RapidBot].mkv",
			expectedFile: "[AnimeOut] Overlord Ple Ple Pleiades - 01 - Special [720p][AAC][DeadFish][RapidBot].mkv",
			description:  "Download file at non-existent path",
		},
	}
	for _, tc := range cases {
		path, err := os.MkdirTemp("", "")
		require.NoError(t, err, "Create temp directory")
		err = DownloadFile(path, tc.url)
		require.NoError(t, err, tc.description)
		assert.DirExists(t, path, tc.description)
		assert.FileExists(t, filepath.Join(path, tc.expectedFile), tc.description)
	}
}
