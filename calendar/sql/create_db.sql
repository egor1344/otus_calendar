create table events (
    id serial primary key,
    owner bigint,
    title text,
    description text,
    date_time TIMESTAMP,
    duration bigint,
    before_duration bigint
);
create index owner_idx ON events (owner);