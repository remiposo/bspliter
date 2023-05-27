-- name: CreateEvent :execresult
INSERT INTO event (id, name) VALUES (?, ?);

-- name: DeleteEvent :execresult
DELETE FROM event WHERE id = ?;
