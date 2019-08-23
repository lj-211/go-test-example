package nettest

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

func Test_VerifyDataOK(t *testing.T) {
	dataStr, _ := json.Marshal(&struct {
		Result string
	}{
		Result: "OK",
	})
	errServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
		r *http.Request) {
		w.WriteHeader(200)
		io.WriteString(w, string(dataStr))
	}))
	defer errServer.Close()

	err := VerifyData(errServer.URL)
	assert.Nil(t, err)
}

func Test_VerifyDataErrCode(t *testing.T) {
	errServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
		r *http.Request) {
		w.WriteHeader(404)
	}))
	defer errServer.Close()

	err := VerifyData(errServer.URL)
	assert.NotNil(t, err)
}

func Test_VerifyDataErrResult(t *testing.T) {
	dataStr, _ := json.Marshal(&struct {
		Result string
	}{
		Result: "fail",
	})
	errServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
		r *http.Request) {
		w.WriteHeader(200)
		io.WriteString(w, string(dataStr))
	}))
	defer errServer.Close()

	err := VerifyData(errServer.URL)
	assert.NotNil(t, err)
}
