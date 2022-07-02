package http

import (
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/WilfredDube/ginny/internal/db"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func openDB(dsn string) (*db.Queries, error) {
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("pgx.Connect%w", err)
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	db := db.New(conn)

	return db, nil
}
func TestHomeController(t *testing.T) {
	db, err := openDB(os.Getenv("$SNIPPET_DSN"))
	if err != nil {
		return
	}

	handler := &SnippetHandler{
		DB: db,
	}

	router := gin.Default()

	router.GET("/", handler.Home)

	router.SetFuncMap(template.FuncMap{
		"TrimGuid":       TrimGuid,
		"HumanDate":      HumanDate,
		"HumanDateShort": HumanDateShort,
	})

	router.LoadHTMLGlob("../../ui/html/*")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)

	p, err := ioutil.ReadAll(w.Body)
	pageOK := err == nil && strings.Index(string(p), "<title>Home - Snippetbox</title>") > 0

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, true, pageOK)
}
