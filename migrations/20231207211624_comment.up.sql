CREATE TABLE IF NOT EXISTS comments (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    message VARCHAR(255) NOT NULL,
    post_id INT REFERENCES posts(id),
    user_id INT REFERENCES users(id),
    parent_id INT,
    created_at timestamp,
    updated_at timestamp
);