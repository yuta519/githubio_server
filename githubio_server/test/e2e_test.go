package test

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestFetchPages(t *testing.T) {
	// Mock Notion API response
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"GET",
		"http://localhost:3000/",
		httpmock.NewBytesResponder(
			200,
			[]byte("Ping, \"/\"")),
	)

	// To test a key in object page
}
