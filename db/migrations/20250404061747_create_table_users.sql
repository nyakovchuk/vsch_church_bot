-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    telegram_users_id INTEGER NOT NULL,
    coordinates_id INTEGER,
    lang_id INTEGER,
    radius INTEGER,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (telegram_users_id) REFERENCES telegram_users(id) ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (coordinates_id) REFERENCES coordinates(id) ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (lang_id) REFERENCES languages(id)
);
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
