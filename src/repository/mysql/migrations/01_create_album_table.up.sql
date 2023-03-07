CREATE TABLE IF NOT EXISTS album (
    id int NOT NULL AUTO_INCREMENT,
    title varchar(255) NOT NULL,
    artist varchar(255) NOT NULL,
    price float,
    album_description varchar(255),
    
    PRIMARY KEY (id),
    CONSTRAINT unique_title_artist UNIQUE (title, artist)
);

