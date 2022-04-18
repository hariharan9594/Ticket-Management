-- +migrate Up

create table if not exists users
(
    id              serial primary key,
    created_at      timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    user_name       text,
    email           text, 
    hash_password   bytea NOT NULL,
    is_admin        boolean
);

-- +migrate Down
DROP TABLE IF EXISTS users;