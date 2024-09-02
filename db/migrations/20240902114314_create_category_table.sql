-- +goose Up
-- +goose StatementBegin

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(225) NOT NULL,
    parent_id INT,
    FOREIGN KEY (parent_id) REFERENCES categories(id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE categories;

-- +goose StatementEnd
