-- +goose Up
-- +goose StatementBegin
CREATE TABLE languages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    code TEXT NOT NULL UNIQUE
);

INSERT OR IGNORE INTO languages (name, code) VALUES 
    ('Українська', 'uk'),
    ('Русский', 'ru'),
    ('English', 'en');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS languages;
-- +goose StatementEnd
