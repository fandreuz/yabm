create table bookmarks (
    Id bigserial primary key not null,
    Url text not null,
    Title text not null,
    CreationDate timestamp not null
);

create table tags (
    Id bigserial primary key not null,
    Label varchar(100) not null unique,
    CreationDate timestamp not null
);
