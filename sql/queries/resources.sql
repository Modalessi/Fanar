-- name: CreateResource :one
INSERT INTO resources (course_id, title, description, file_ext, s3_url, tags, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: DeleteResourceByID :one
DELETE FROM resources WHERE id = $1 RETURNING *;

-- name: GetResourceByID :one
SELECT * FROM resources WHERE id = $1;

-- name: GetCourseResources :many
SELECT * FROM resources WHERE course_id = $1;