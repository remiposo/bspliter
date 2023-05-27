-- name: CreateMember :execresult
INSERT INTO member (id, name, event_id) VALUES (?, ?, ?);

-- name: ListMembersByEventID :many
SELECT * FROM member WHERE event_id = ?;
