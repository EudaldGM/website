-- name: GetUser :one
SELECT * FROM users
WHERE name = $1 LIMIT 1;

-- name: GetFalseUsers :many
SELECT * FROM users
WHERE status = false;

-- name: GetUsers :many
SELECT * FROM users;

-- name: CreateUser :exec
INSERT INTO users (id, name, status)
VALUES ($1, $2, $3);

-- name: DeleteUser :exec
DELETE FROM users WHERE name=$1;

-- name: LastId :one
SELECT id FROM users ORDER BY ID DESC LIMIT 1;