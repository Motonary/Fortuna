
-- +migrate Up
create table users (
id serial PRIMARY KEY,
name char(255) NOT NULL,
email char(255) NOT NULL,
password char(255) NOT NULL,
created_at timestamp NOT NULL,
updated_at timestamp NOT NULL);

-- +migrate Down
drop table users;
