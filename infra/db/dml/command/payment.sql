-- name: CreatePayment :execresult
INSERT INTO payment (id, name, amount, event_id, member_id) VALUES (?, ?, ?, ?, ?);

-- name: ListPaymentsByEventID :many
SELECT * FROM payment WHERE event_id = ?;
