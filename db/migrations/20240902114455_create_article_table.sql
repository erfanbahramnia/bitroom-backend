-- +goose Up
-- +goose StatementBegin

CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    title VARCHAR(225) NOT NULL,
    description TEXT NOT NULL,
    summary TEXT NOT NULL,
    image VARCHAR(225) NOT NULL,
    status VARCHAR(225) DEFAULT 'InProgress',
    category_id INT NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON UPDATE CASCADE ON DELETE SET NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE articles;

-- +goose StatementEnd
