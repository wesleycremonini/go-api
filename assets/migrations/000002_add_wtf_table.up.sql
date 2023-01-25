create table if not exists wtfs (
    id serial primary key not null,
    wtf varchar(250) not null,
    created_at timestamp not null
);