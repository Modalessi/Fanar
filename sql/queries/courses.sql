-- name: CreateCourse :one
INSERT INTO courses (title, description, course_code, credit_hours, contact_hours) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetCourseByID :one
SELECT * FROM courses WHERE id = $1;

-- name: UpdateCourseDescription :one
UPDATE courses SET description = $1 WHERE id = $2 RETURNING *;