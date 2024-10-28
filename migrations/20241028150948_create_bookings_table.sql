-- +goose Up
-- +goose StatementBegin
create table bookings
(
    id            character varying primary key,
    first_name    character varying        not null,
    last_name     character varying        not null,
    gender        character varying        not null,
    date_of_birth date                     not null,
    launchpad_id  character varying        not null,
    destination   character varying        not null,
    launch_date   timestamp with time zone not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table bookings;
-- +goose StatementEnd
