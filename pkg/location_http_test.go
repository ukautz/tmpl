package tmpl

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildHTTPLocation(t *testing.T) {
	s := BuildHTTPLocation("http://valid.tld/")
	assert.NotNil(t, s, "HTTPLocation should build from HTTP URL")
	s = BuildHTTPLocation("https://valid.tld/")
	assert.NotNil(t, s, "HTTPLocation should build from HTTPS URL")
	s = BuildHTTPLocation("ftp://valid.tld/")
	assert.Nil(t, s, "HTTPLocation should NOT build from FTP URL")
	s = BuildHTTPLocation("file://valid.tld/")
	assert.Nil(t, s, "HTTPLocation should NOT build from file URL")
}

func TestHTTPLocation_Load(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`THE DATA RESPONSE`))
	}))
	s := HTTPLocation(ts.URL)
	data, err := s.Load()
	assert.Nil(t, err, "Location could be loaded")
	assert.Equal(t, `THE DATA RESPONSE`, string(data), "Data is delegated")
}

func TestHTTPLocation_Load_FailWithNonOKResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`THE DATA RESPONSE`))
	}))
	s := HTTPLocation(ts.URL)
	_, err := s.Load()
	assert.NotNil(t, err, "Location should NOT be loaded")
	assert.Contains(t, err.Error(), "500 Internal Server Error", "Error should be delegated")
}

func TestHTTPLocation_Load_FailFromInvalid(t *testing.T) {
	s := HTTPLocation("foo://bar.baz")
	_, err := s.Load()
	assert.NotNil(t, err, "Location should NOT be loaded")
	assert.Contains(t, err.Error(), "unsupported protocol scheme \"foo\"", "Error should be delegated")
}
