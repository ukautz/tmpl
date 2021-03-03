package tmpl

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildHTTPSource(t *testing.T) {
	s := BuildHTTPSource("http://valid.tld/")
	assert.NotNil(t, s, "HTTPSource should build from HTTP URL")
	s = BuildHTTPSource("https://valid.tld/")
	assert.NotNil(t, s, "HTTPSource should build from HTTPS URL")
	s = BuildHTTPSource("ftp://valid.tld/")
	assert.Nil(t, s, "HTTPSource should NOT build from FTP URL")
	s = BuildHTTPSource("file://valid.tld/")
	assert.Nil(t, s, "HTTPSource should NOT build from file URL")
}

func TestSource_Load(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`THE DATA RESPONSE`))
	}))
	s := HTTPSource(ts.URL)
	data, err := s.Load()
	assert.Nil(t, err, "Source could be loaded")
	assert.Equal(t, `THE DATA RESPONSE`, string(data), "Data is delegated")
}

func TestSource_Load_FailWithNonOKResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`THE DATA RESPONSE`))
	}))
	s := HTTPSource(ts.URL)
	_, err := s.Load()
	assert.NotNil(t, err, "Source should NOT be loaded")
	assert.Contains(t, err.Error(), "500 Internal Server Error", "Error should be delegated")
}

func TestSource_Load_FailFromInvalid(t *testing.T) {
	s := HTTPSource("foo://bar.baz")
	_, err := s.Load()
	assert.NotNil(t, err, "Source should NOT be loaded")
	assert.Contains(t, err.Error(), "unsupported protocol scheme \"foo\"", "Error should be delegated")
}
