package query

const AddUser = `
INSERT INTO users (name, role)
VALUES ($1, $2)
RETURNING id, name, role, last_updated_at
`

const DeleteUser = `
DELETE
FROM users
WHERE id = $1
RETURNING id, name, role, last_updated_at
`
const GetAllUsers = `
SELECT id, name, role, last_updated_at
FROM users
`
const GetUser = `
SELECT id, name, role, last_updated_at
FROM users
WHERE id = $1
LIMIT 1
`
const UpdateUserRole = `
UPDATE users
SET role            = $2,
    last_updated_at = now()
WHERE id = $1
RETURNING id, name, role, last_updated_at
`
