-- +goose Up
create table if not exists Trips (
    id serial primary key,
    trip_status int not null,
    created_at timestamp NOT NULL DEFAULT NOW()
);

-- +goose Down
drop table Trips;