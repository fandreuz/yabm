create table bookmarks (
    id BIGSERIAL PRIMARY KEY,
    url text NOT NULL,
    title text,
    creationDate bigint NOT NULL
);

create table tags (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    label varchar(100) NOT NULL,
    creationDate bigint NOT NULL
);
