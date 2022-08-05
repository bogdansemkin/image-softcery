CREATE TABLE images
(
    id serial not null unique,
    path varchar(255) not null,
    seventy_five_path varchar(255) not null,
    half_path varchar(255) not null,
    twenty_five_path varchar(255) not null
)