-- name: CreateMember :execresult
INSERT INTO member (id, name, event_id) VALUES (?, ?, ?);
