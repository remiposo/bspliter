-- name: CreateEvent :execresult
INSERT INTO event (id, name) VALUES (?, ?);

-- name: GetEventByID :one
SELECT * FROM event WHERE id = ?;

-- name: DeleteEvent :execresult
DELETE FROM event WHERE id = ?;
