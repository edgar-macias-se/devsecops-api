package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHealthHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/health", nil)
    w := httptest.NewRecorder()
    healthHandler(w, req)
    if w.Code != http.StatusOK {
        t.Errorf("esperaba 200, obtuve %d", w.Code)
    }
}
