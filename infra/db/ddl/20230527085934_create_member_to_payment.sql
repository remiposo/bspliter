-- +goose Up
-- +goose StatementBegin
CREATE TABLE member_to_payment (
    member_id CHAR(36) NOT NULL,
    payment_id CHAR(36) NOT NULL,
    PRIMARY KEY (member_id, payment_id),
    FOREIGN KEY (member_id) REFERENCES member (id) ON DELETE CASCADE,
    FOREIGN KEY (payment_id) REFERENCES payment (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE member_to_payment;
-- +goose StatementEnd
