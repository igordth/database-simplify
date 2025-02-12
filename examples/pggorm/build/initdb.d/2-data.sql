insert into galaxies (id, name)
values (1, 'Milky Way'),
       (2, 'Andromeda')
;

insert into stars (id, name, galaxy_id)
values (1, 'Sun', 1),
       (2, 'Betelgeuse', 1),
       (3, 'Alpheratz', 2),
       (4, 'Mirach', 2)
;

insert into planets (name, star_id)
values ('Mercury', 1),
       ('Venus', 1),
       ('Earth', 1),
       ('Mars', 1)
;