create table bookmarks (
    id bigserial primary key not null,
    url text not null,
    title text not null,
    creationDate timestamp not null
);

create table tags (
    id bigserial primary key not null,
    label varchar(100) not null,
    creationDate timestamp not null
);
