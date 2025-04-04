-- +goose Up
-- +goose StatementBegin
CREATE TABLE telegram_users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    tg_id INTEGER NOT NULL,
    username TEXT NOT NULL,
    first_name TEXT DEFAULT "",
    last_name TEXT DEFAULT "",
    language_code TEXT NOT NULL,
    is_bot boolean DEFAULT false,
    is_premium boolean DEFAULT false
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS telegram_users;
-- +goose StatementEnd
