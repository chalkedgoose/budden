-- name: ListTodos :many
SELECT * FROM todos ORDER BY due_date;

-- name: CreateTodo :exec
INSERT INTO todos (title, description, due_date) VALUES ($1, $2, $3);