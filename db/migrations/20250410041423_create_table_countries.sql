-- +goose Up
-- +goose StatementBegin
CREATE TABLE countries (
    id INTEGER UNIQUE,
    country_ru TEXT UNIQUE NOT NULL,
    country_en TEXT UNIQUE NOT NULL
);
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (1, 'Беларусь', 'Belarus');
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (4, 'США', 'USA');
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (9, 'Россия', 'Russia');
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (12, 'Украина', 'Ukraine');
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (59, 'Эстония', 'Estonia');
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (62, 'Литва', 'Lithuania');
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (70, 'Германия', 'Germany');
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (89, 'Канада', 'Canada');
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (94, 'Румыния', 'Romania');
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (110, 'Молдова', 'Moldova');
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (116, 'Латвия', 'Latvia');
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (151, 'Финляндия', 'Finland');
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (152, 'Великобритания', 'UK');
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (153, 'Испания', 'Spain');
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (154, 'Казахстан', 'Kazakhstan');
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (155, 'Польша', 'Poland');
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (156, 'Чехия', 'Czechia');
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (157, 'Киргизия', 'Kyrgyzstan');
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (158, 'Бельгия', 'Belgium');
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (159, 'Австралия', 'Australia');
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (160, 'Узбекистан', 'Uzbekistan');
INSERT INTO `countries` (`id`, `country_ru`, `country_en`) VALUES (161, 'Азербайджан', 'Azerbaijan');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS countries;
-- +goose StatementEnd
