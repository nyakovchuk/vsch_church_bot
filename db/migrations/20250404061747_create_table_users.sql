-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    tg_user_id INTEGER UNIQUE NOT NULL,
    lang_id INTEGER,
    radius INTEGER,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (tg_user_id) REFERENCES telegram_users(tg_id) ON DELETE CASCADE
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
