-- +goose Up
-- +goose StatementBegin

CREATE TABLE article_comments (
    id SERIAL PRIMARY KEY,
    comment VARCHAR(255) NOT NULL,
    article_id INT NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (article_id) REFERENCES articles(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE article_comments;

-- +goose StatementEnd
