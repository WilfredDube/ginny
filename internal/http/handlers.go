package http

import (
	"net/http"

	"github.com/WilfredDube/ginny/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SnippetHandler struct {
	DB *db.Queries
}

// Home diplays the home page containing a list of all snippets.
func (h *SnippetHandler) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.page.tmpl", nil)
}

// ShowSnippet displays a single selected snippet.
func (h *SnippetHandler) ShowSnippet(c *gin.Context) {
	c.String(http.StatusOK, "Display a specific snippet...")
}

// PrepareSnippet prepares an URL with a UUID for the creation of a snippets. Makes the
// creation idempotent.
func (h *SnippetHandler) PrepareSnippet(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		// Generate GUID to make call idempotent
		location := c.FullPath() + "/" + uuid.NewString()

		c.Redirect(http.StatusMovedPermanently, location)
	}

	c.String(http.StatusOK, "Create new snippet..."+id)
}

// CreateSnippet creates a snippet. Call is idempotent.
func (h *SnippetHandler) CreateSnippet(c *gin.Context) {
	id := c.Param("id")

	c.String(http.StatusOK, "Create new snippet..."+id)
}
