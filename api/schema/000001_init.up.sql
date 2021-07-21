CREATE TABLE users
(
    id          serial       not null unique,
    nickname    varchar(255) not null unique,
    password    varchar(255) not null,
    first_name  varchar(255) not null,
    middle_name varchar(255) not null,
    last_name   varchar(255) not null,
    snils       varchar(255) not null
);
