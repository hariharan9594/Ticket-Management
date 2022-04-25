-- +migrate Up

create table if not exists tickets
(
    Desk      text,
    ticket_type text,
    ticket_id integer PRIMARY KEY,
    subject text,
    u_id integer,
    CONSTRAINT fk_tickets_users FOREIGN KEY(u_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS tickets;