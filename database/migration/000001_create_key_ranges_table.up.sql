create table key_ranges (
    id serial primary key,
    start integer not null,
    "end" integer not null,
    taken boolean not null default false
);