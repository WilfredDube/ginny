-- name: CreateSnippet :one
INSERT INTO snippets (
    guid, title, content, created, expires
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: Get :one
SELECT * FROM snippets
WHERE expires > NOW() AND guid = $1;

-- name: All :many
SELECT * FROM snippets
WHERE expires > NOW() ORDER BY created DESC LIMIT 10;