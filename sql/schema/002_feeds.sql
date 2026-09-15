-- +goose Up
CREATE TABLE feeds (
    id UUID primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    name text not null,
    url text unique not null,
    user_id UUID not null,
    CONSTRAINT fk_id
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;