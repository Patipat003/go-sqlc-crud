-- name: GetAllUsers :many
SELECT id, first_name, last_name, email, gender, ip_address, date 
  FROM mockdata ORDER BY id;

-- name: GetUserByID :one
SELECT id, first_name, last_name, email, gender, ip_address, date 
  FROM mockdata WHERE id = $1;

-- name: CreateUser :one
INSERT INTO mockdata (first_name, last_name, email, gender, ip_address ,date)
  VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING *;

-- name: DeleteUser :one
DELETE FROM mockdata WHERE id = $1
  RETURNING *;

-- name: UpdateUser :one
UPDATE mockdata SET first_name=$1, last_name=$2 WHERE id=$3
  RETURNING *;