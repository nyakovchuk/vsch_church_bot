-- +goose Up
-- +goose StatementBegin
CREATE TABLE confessions (
    id INTEGER UNIQUE,
    name_ru TEXT UNIQUE NOT NULL,
    name_en TEXT UNIQUE NOT NULL
);

INSERT INTO `confessions` (`id`, `name_ru`, `name_en`) VALUES (1, 'Пятидесятники', 'Pentecostals');
INSERT INTO `confessions` (`id`, `name_ru`, `name_en`) VALUES (2, 'Баптисты', 'Baptists');
INSERT INTO `confessions` (`id`, `name_ru`, `name_en`) VALUES (3, 'Харизматы', 'Charismatics');
INSERT INTO `confessions` (`id`, `name_ru`, `name_en`) VALUES (4, 'Мессианский иудаизм', 'Messianic judaism');
INSERT INTO `confessions` (`id`, `name_ru`, `name_en`) VALUES (5, 'Адвентисты седьмого дня', 'Seventh day Adventists');
INSERT INTO `confessions` (`id`, `name_ru`, `name_en`) VALUES (6, 'Греко-Католики', 'Greco Catholics');
INSERT INTO `confessions` (`id`, `name_ru`, `name_en`) VALUES (7, 'Католики', 'Catholics');
INSERT INTO `confessions` (`id`, `name_ru`, `name_en`) VALUES (8, 'Православие', 'Orthodoxy');
INSERT INTO `confessions` (`id`, `name_ru`, `name_en`) VALUES (9, 'Другое', 'Other');
INSERT INTO `confessions` (`id`, `name_ru`, `name_en`) VALUES (10, 'Христиане Субботнего Дня', 'Christians of the Sabbath Day');
INSERT INTO `confessions` (`id`, `name_ru`, `name_en`) VALUES (11, 'Евангельские Христианы', 'Evangelical Christians');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS confessions;
-- +goose StatementEnd
