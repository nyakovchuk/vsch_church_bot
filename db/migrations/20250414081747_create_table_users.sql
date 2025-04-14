-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    platform_id INTEGER,
    external_id TEXT NOT NULL,
    lang_id INTEGER,
    radius INTEGER,
    username TEXT NOT NULL,
    first_name TEXT DEFAULT "",
    last_name TEXT DEFAULT "",
    language_code TEXT NOT NULL,
    is_bot boolean DEFAULT false,
    is_premium boolean DEFAULT false,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (platform_id) REFERENCES platforms(tid)
        ON DELETE NO ACTION
        ON UPDATE CASCADE,
    FOREIGN KEY (lang_id) REFERENCES languages(id)
);

CREATE UNIQUE INDEX idx_users_platform_external_id
        ON users(platform_id, external_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
--PRAGMA foreign_keys = OFF;
-- Сначала удаляем зависимые таблицы
--DROP TABLE IF EXISTS telegram_users;
--DROP TABLE IF EXISTS coordinates;
--DROP TABLE IF EXISTS languages;

DROP TABLE IF EXISTS users;

--PRAGMA foreign_keys = ON;
-- +goose StatementEnd
