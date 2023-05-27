-- +goose Up
-- +goose StatementBegin
CREATE TABLE payment (
    id CHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    amount BIGINT UNSIGNED NOT NULL,
    event_id CHAR(36) NOT NULL,
    member_id CHAR(36) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (event_id) REFERENCES event (id) ON DELETE CASCADE,
    FOREIGN KEY (member_id) REFERENCES member (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE payment;
-- +goose StatementEnd
