create table galaxies
(
    id bigserial primary key ,
    name varchar(255) not null
);
create table stars
(
    id bigserial primary key ,
    name varchar(255) not null,
    galaxy_id bigint references galaxies on update cascade on delete cascade
);
create table planets
(
    id bigserial primary key ,
    name varchar(255) not null,
    star_id bigint references stars on update cascade on delete cascade
);