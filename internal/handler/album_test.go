package handler

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestGetAlbums(t *testing.T) {
    // Set Gin to Test Mode
    gin.SetMode(gin.TestMode)

    // Setup router
    router := gin.Default()
    RegisterRoutes(router)

    // Create request
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/albums", nil)
    
    // Serve request
    router.ServeHTTP(w, req)

    // Assertions
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Blue Train")
}
