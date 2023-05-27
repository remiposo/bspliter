-- +goose Up
-- +goose StatementBegin
CREATE TABLE member (
    id CHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    event_id CHAR(36) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (event_id) REFERENCES event (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE member;
-- +goose StatementEnd
