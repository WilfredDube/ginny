package http

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/WilfredDube/ginny/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SnippetHandler struct {
	DB *db.Queries
}

// Home diplays the home page containing a list of all snippets.
func (h *SnippetHandler) Home(c *gin.Context) {
	s, err := h.DB.All(context.Background())
	if err != nil {
		ServerError(c, err)
		return
	}

	c.HTML(http.StatusOK, "home.page.tmpl", gin.H{
		"Page":     "Home",
		"Snippets": s,
	})
}

// ShowSnippet displays a single selected snippet.
func (h *SnippetHandler) ShowSnippet(c *gin.Context) {
	id := c.Param("id")

	s, err := h.DB.Get(context.Background(), uuid.MustParse(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			notFound(c)
		} else {
			ServerError(c, err)
		}
		return
	}

	c.HTML(http.StatusOK, "show.page.tmpl", gin.H{
		"Page":    "Snippet #" + s.Guid.String(),
		"Snippet": s,
	})
}

// PrepareSnippet prepares an URL with a UUID for the creation of a snippets. Makes the
// creation idempotent.
func (h *SnippetHandler) PrepareSnippet(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		// Generate GUID to make call idempotent
		newid := uuid.NewString()

		location := url.URL{Path: c.FullPath() + "/" + newid}

		c.Header("Cache-Control", "no-store")
		c.Redirect(http.StatusMovedPermanently, location.RequestURI())
		return
	}

	c.Header("Cache-Control", "no-store")
	c.HTML(http.StatusOK, "create.page.tmpl", gin.H{
		"Page": "New ",
	})
}

type snippetRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Expires int64  `json:"expires,string"`
}

// CreateSnippet creates a snippet. Call is idempotent.
func (h *SnippetHandler) CreateSnippet(c *gin.Context) {
	id := c.Param("id")

	guid := uuid.MustParse(id)

	sn, err := h.DB.Get(context.Background(), guid)
	if err == sql.ErrNoRows {
		var r snippetRequest
		if err := c.ShouldBindJSON(&r); err != nil {
			ServerError(c, err)
			return
		}

		created := time.Now()

		arg := db.CreateSnippetParams{
			Guid:    guid,
			Title:   r.Title,
			Content: r.Content,
			Created: created,
			Expires: created.AddDate(0, 0, int(r.Expires)),
		}

		sn, err = h.DB.CreateSnippet(context.Background(), arg)
		if err != nil {
			ServerError(c, err)
			return
		}
	}

	location := url.URL{Path: fmt.Sprintf("/snippets/%s", sn.Guid)}
	c.Redirect(http.StatusMovedPermanently, location.RequestURI())
}
