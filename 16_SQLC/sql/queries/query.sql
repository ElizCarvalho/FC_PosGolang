-- name: ListCategories :many
SELECT * FROM categories;

-- name: GetCategoryByID :one
SELECT * FROM categories WHERE id = ?;

-- name: CreateCategory :exec
INSERT INTO categories (id, name, description) VALUES (?, ?, ?);

-- name: UpdateCategory :exec
UPDATE categories SET name = ?, description = ? WHERE id = ?;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = ?;

-- name: CreateCourse :exec
INSERT INTO courses (id, name, description, category_id, price) VALUES (?, ?, ?, ?, ?);

-- -- name: ListCourses :many
-- SELECT * FROM courses;

-- -- name: GetCourseByID :one
-- SELECT * FROM courses WHERE id = ?;

-- -- name: UpdateCourse :exec
-- UPDATE courses SET name = ?, description = ?, category_id = ?, price = ? WHERE id = ?;

-- -- name: DeleteCourse :exec
-- DELETE FROM courses WHERE id = ?;