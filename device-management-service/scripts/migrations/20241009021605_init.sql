-- +goose Up
-- +goose StatementBegin
create database device_management_service;

create table if not exists devices (
    id serial primary key,
    name varchar(255) not null,
    description text,
    status_id int not null,
    home_id int not null,
    user_id int not null,
    attributes jsonb,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);

create table if not exists device_status (
    id serial primary key,
    status varchar(255) not null,
    created_at timestamp not null default current_timestamp
);

create table if not exists device_type (
    id serial primary key,
    type varchar(255) not null,
    created_at timestamp not null default current_timestamp
);

create table if not exists users (
     id serial primary key,
     name varchar(255) not null,
     email varchar(255) not null,
     password varchar(255) not null,
     created_at timestamp not null default current_timestamp,
     updated_at timestamp not null default current_timestamp
);

create table if not exists homes (
    id serial primary key,
    name varchar(255) not null,
    description text,
    user_id    int not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);

insert into device_status (status) values ('on');
insert into device_status (status) values ('off');


insert into device_type (type) values ('light');
insert into device_type (type) values ('temperature');
insert into device_type (type) values ('door');
insert into device_type (type) values ('window');
insert into device_type (type) values ('camera');
insert into device_type (type) values ('gate');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop database device_management_service;

drop table if exists devices;

drop table if exists users;

drop table if exists homes;
-- +goose StatementEnd
