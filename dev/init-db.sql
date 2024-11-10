create table bookmarks (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    url text NOT NULL,
    title text NOT NULL,
    creationDate DATE NOT NULL
);

create table tags (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    label varchar(100) NOT NULL,
    creationDate DATE NOT NULL
);
