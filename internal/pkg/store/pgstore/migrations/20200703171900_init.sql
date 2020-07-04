-- +goose Up

create table if not exists "user"
(
    id            bigserial primary key,
    login         varchar(100) unique not null,
    email         varchar(100),
    password_hash varchar(256) not null
);

-- +goose Down

drop table if exists "user";
