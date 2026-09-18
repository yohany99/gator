-- +goose Up
CREATE TABLE feed_follows (
    id UUID primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    user_id UUID not null,
    feed_id UUID not null,
    CONSTRAINT fk_userid
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    CONSTRAINT fk_feedid
        FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE
    CONSTRAINT user_feed_uq UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;