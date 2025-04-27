package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gogineni1998/oolio-assignment-backend/handlers"
	"github.com/stretchr/testify/assert"
)

func mockHandlerOK() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}

func TestNewServer_Routes_WithMockedHandlers(t *testing.T) {
	handlers.Token = mockHandlerOK
	handlers.Register = mockHandlerOK
	handlers.Products = func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusOK)
		}
	}
	handlers.Product = handlers.Products
	handlers.Order = handlers.Products

	server := NewServer()

	tests := []struct {
		name         string
		method       string
		url          string
		withAuth     bool
		expectedCode int
	}{
		{
			name:         "Token Endpoint",
			method:       http.MethodPost,
			url:          "/token",
			withAuth:     false,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Register Endpoint",
			method:       http.MethodPost,
			url:          "/register",
			withAuth:     false,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Products Endpoint - No Token",
			method:       http.MethodGet,
			url:          "/products",
			withAuth:     false,
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "Product By ID Endpoint - No Token",
			method:       http.MethodGet,
			url:          "/product/123",
			withAuth:     false,
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "Order Endpoint - No Token",
			method:       http.MethodPost,
			url:          "/order",
			withAuth:     false,
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.url, nil)
			if tc.withAuth {
				req.Header.Set("Authorization", "Bearer mocktoken")
			}
			rec := httptest.NewRecorder()

			server.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
