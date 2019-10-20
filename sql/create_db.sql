CREATE EXTENSION pgcrypto;
create table events
(
    date_time        timestamp,
    description      text,
    duration         bigint,
    owner            bigint,
    title            text,
    before_time_pull bigint,
    id               uuid default gen_random_uuid() not null
        constraint events_pkey
            primary key,
    send             boolean
);

alter table events
    owner to postgres;

