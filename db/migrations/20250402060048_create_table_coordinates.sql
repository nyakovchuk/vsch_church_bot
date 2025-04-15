-- +goose Up
-- +goose StatementBegin
CREATE TABLE coordinates (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    platform_id INTEGER,
    external_id TEXT UNIQUE NOT NULL,
    latitude REAL,
    longitude REAL,
    is_on_text boolean DEFAULT false,
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (platform_id) REFERENCES platforms(id)
        ON DELETE NO ACTION
        ON UPDATE CASCADE,
    FOREIGN KEY (external_id) REFERENCES users(external_id) 
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE UNIQUE INDEX idx_coordinates_platform_external_id
        ON coordinates(platform_id, external_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS coordinates;
-- +goose StatementEnd
