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
    directions_page_url varchar(255) not null
);

CREATE TABLE directions
(
    id   serial       not null unique,
    name varchar(255) not null,
    url  varchar(255) not null
);

CREATE TABLE users_universities
(
    id            serial                           not null unique,
    user_id       int references users (id)        not null,
    university_id int references universities (id) not null
);

CREATE TABLE universities_directions
(
    id            serial                           not null unique,
    university_id int references universities (id) not null,
    direction_id  int references directions (id)   not null,
);
