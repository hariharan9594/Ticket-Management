-- +migrate Up

CREATE TABLE if not exists gorp_migrations (
    id text NOT NULL,
    applied_at timestamp with time zone
);

ALTER TABLE IF EXISTS ONLY gorp_migrations DROP CONSTRAINT IF EXISTS gorp_migrations_pkey;

ALTER TABLE ONLY gorp_migrations
    ADD CONSTRAINT gorp_migrations_pkey PRIMARY KEY (id);

-- +migrate Down
DROP TABLE IF EXISTS gorp_migrations;