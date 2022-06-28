package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func TestHomeController(t *testing.T) {
	router := gin.Default()
	router.GET("/", Home)
	router.LoadHTMLGlob("../../ui/html/*")

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	p, err := ioutil.ReadAll(w.Body)
	pageOK := err == nil && strings.Index(string(p), "<title>Home - Snippetbox</title>") > 0

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, true, pageOK)
}
