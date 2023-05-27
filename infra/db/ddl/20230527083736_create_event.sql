-- +goose Up
-- +goose StatementBegin
CREATE TABLE event (
    id CHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE event;
-- +goose StatementEnd
