package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gogineni1998/oolio-assignment-backend/configuration"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	configuration.Address = ":0"

	configuration.DBClient = nil

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	go func() {
		err := Run()
		assert.NoError(t, err)
	}()
}
