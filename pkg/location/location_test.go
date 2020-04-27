package location

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestZipCodeLoad(t *testing.T) {

	zipCodeMap, err := LoadZipCodeMap("testdata/zip.csv")

	require.NoError(t, err)
	assert.NotEmpty(t, zipCodeMap)
}
