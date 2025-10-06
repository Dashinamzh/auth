-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS auth (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table auth;
-- +goose StatementEnd
