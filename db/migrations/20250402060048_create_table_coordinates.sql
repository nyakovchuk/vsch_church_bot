-- +goose Up
-- +goose StatementBegin
CREATE TABLE coordinates (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    latitude REAL,
    longitude REAL,
    is_on_text boolean DEFAULT false,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS coordinates;
-- +goose StatementEnd
