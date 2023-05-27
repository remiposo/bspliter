-- name: CreateMemberToPayment :execresult
INSERT INTO member_to_payment (event_id, member_id, payment_id) VALUES (?, ?, ?);

-- name: ListMemberToPaymentsByEventID :many
SELECT * FROM member_to_payment WHERE event_id = ?;
