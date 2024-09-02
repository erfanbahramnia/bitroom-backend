-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    phone VARCHAR(225)  UNIQUE NOT NULL,
    first_name VARCHAR(225) DEFAULT '',
    last_name VARCHAR(225) DEFAULT '',
    role VARCHAR(225) DEFAULT 'user',
    password VARCHAR(225) NOT NULL
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE users;

-- +goose StatementEnd
