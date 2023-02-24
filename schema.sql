CREATE TABLE "demo-sql".album
(
    id serial,
    title text,
    artist text,
    price integer DEFAULT 0,
    PRIMARY KEY (id),
    CONSTRAINT unique_artist_album UNIQUE (title, artist)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS "demo-sql".album
    OWNER to "demo-user";

CREATE DATABASE demo;

USE demo;

CREATE TABLE album (
    id int NOT NULL AUTO_INCREMENT,
    title varchar(255) NOT NULL,
    artist varchar(255) NOT NULL,
    price float,
    album_description varchar(255),
    
    PRIMARY KEY (id),
    CONSTRAINT unique_title_artist UNIQUE (title, artist) 
);

INSERT INTO album (title, artist, price, album_description) VALUES
    ('Blue Train', 'John Coltrane', 56.99, ''),
    ('Jeru', 'Gerry Mulligan', 17.99, 'jazz'),
    ('Sarah Vaughan and Clifford Brown', 'Sarah Vaughan', 39.99, 'single'),
    ('Renaissance', 'Beyonce', 85.99, 'very expensive'),
    ('#52', 'Quevedo', 20.99, 'session');

SELECT * FROM album;