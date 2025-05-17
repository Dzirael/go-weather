-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS subscription (
    id                UUID PRIMARY KEY,
    confirmation_code UUID,
    status            VARCHAR(32)  NOT NULL,
    frequency         VARCHAR(16)  NOT NULL,
    email             VARCHAR(255) NOT NULL UNIQUE,
    city              VARCHAR(100) NOT NULL,
    sended_at         TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS subscription;
-- +goose StatementEnd
