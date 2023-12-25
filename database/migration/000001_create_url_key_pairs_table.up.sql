create table if not exists url_key_pairs (
    id uuid not null primary key,
    url text not null,
    key varchar(11) not null,
    created_at timestamp with time zone not null default current_timestamp,
    updated_at timestamp with time zone not null default current_timestamp,
    expired_at date not null
);