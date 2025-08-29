-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = ? and deleted_at IS NULL;

-- name: InsertUser :exec
INSERT INTO users (username, password, role)
VALUES (?, ?, ?);
