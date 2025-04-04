-- +goose Up
-- +goose StatementBegin
CREATE TABLE languages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    abbr TEXT NOT NULL UNIQUE
);

INSERT OR IGNORE INTO languages (name, abbr) VALUES 
    ('Українська', 'ua'),
    ('Русский', 'ru'),
    ('English', 'en');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS languages;
-- +goose StatementEnd
