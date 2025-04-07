-- +goose Up
-- +goose StatementBegin
CREATE TABLE coordinates (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    tg_user_id INTEGER UNIQUE NOT NULL,
    latitude REAL,
    longitude REAL,
    is_on_text boolean DEFAULT false,
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (tg_user_id) REFERENCES users(tg_user_id) ON DELETE CASCADE
            ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS coordinates;
-- +goose StatementEnd
