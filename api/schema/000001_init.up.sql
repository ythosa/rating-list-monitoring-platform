CREATE TABLE users
(
    id          serial       not null unique,
    username    varchar(255) not null unique,
    password    varchar(255) not null,
    first_name  varchar(255) not null,
    middle_name varchar(255) not null,
    last_name   varchar(255) not null,
    snils       varchar(255) not null
);

CREATE TABLE universities
(
    id                  serial       not null unique,
    name                varchar(255) not null,
    full_name           varchar(255) not null,
    directions_page_url varchar(255) not null
);

CREATE TABLE directions
(
    id            serial                           not null unique,
    name          varchar(255)                     not null,
    url           varchar(255)                     not null,
    university_id int references universities (id) not null
);

CREATE TABLE users_universities
(
    id            serial                           not null unique,
    user_id       int references users (id)        not null,
    university_id int references universities (id) not null
);

CREATE TABLE users_directions
(
    id           serial                         not null unique,
    user_id      int references users (id)      not null,
    direction_id int references directions (id) not null
);

INSERT INTO universities (name, full_name, directions_page_url)
VALUES ('СПБГУ',
        'Санкт-Петербургский государственный университет',
        'https://cabinet.spbu.ru/Lists/1k_EntryLists/index_comp_groups.html'),
       ('ЛЭТИ',
        'Санкт-Петербургский государственный электротехнический университет «ЛЭТИ»',
        'https://etu.ru/ru/abiturientam/priyom-na-1-y-kurs/podavshie-zayavlenie/');
