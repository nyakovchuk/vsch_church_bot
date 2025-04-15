-- +goose Up
-- +goose StatementBegin
CREATE TABLE platforms (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE NOT NULL
);

INSERT OR IGNORE INTO platforms (name) VALUES 
    ('telegram'), ('viber');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS platforms;
-- +goose StatementEnd
