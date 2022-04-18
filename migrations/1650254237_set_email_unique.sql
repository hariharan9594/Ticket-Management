-- +migrate Up

ALTER TABLE ONLY users
    ADD CONSTRAINT users_email_key UNIQUE (email);

-- +migrate Down

ALTER TABLE IF EXISTS ONLY users 
    DROP CONSTRAINT IF EXISTS users_email_key;
