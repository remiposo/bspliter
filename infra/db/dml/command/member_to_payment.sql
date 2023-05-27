-- name: CreateMemberToPayment :execresult
INSERT INTO member_to_payment (member_id, payment_id) VALUES (?, ?);
