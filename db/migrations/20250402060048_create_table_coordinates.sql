-- +goose Up
-- +goose StatementBegin
CREATE TABLE coordinates (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    latitude REAL,
    longitude REAL,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE coordinates;
-- +goose StatementEnd
