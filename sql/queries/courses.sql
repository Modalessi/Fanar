-- name: CreateCourse :one
INSERT INTO courses (title, description, course_code, credit_hours, contact_hours) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetCourseByID :one
SELECT * FROM courses WHERE id = $1;

-- name: UpdateCourseDescription :one
UPDATE courses SET description = $1 WHERE id = $2 RETURNING *;

-- name: DeleteCourseByID :one
DELETE FROM courses WHERE id = $1 RETURNING *;

-- name: UpdateCourse :one
UPDATE courses 
SET 
  title = $1, 
  course_code = $2, 
  description = $3, 
  credit_hours = $4, 
  contact_hours = $5, 
  updated_at = now()
WHERE id = $6 
RETURNING id, title, description, course_code, credit_hours, contact_hours, created_at, updated_at;