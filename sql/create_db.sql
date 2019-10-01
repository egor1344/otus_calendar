-- Table: public.events

-- DROP TABLE public.events;

CREATE TABLE public.events
(
    date_time timestamp without time zone,
    description text COLLATE pg_catalog."default",
    duration bigint,
    id integer NOT NULL DEFAULT nextval('events_id_seq'::regclass),
    owner bigint,
    title text COLLATE pg_catalog."default",
    CONSTRAINT events_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.events
    OWNER to postgres;