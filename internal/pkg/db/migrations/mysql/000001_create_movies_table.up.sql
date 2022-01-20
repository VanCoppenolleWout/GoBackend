CREATE TABLE IF NOT EXISTS Movies(
    ID INT NOT NULL UNIQUE AUTO_INCREMENT,
    Title VARCHAR(127) NOT NULL,
    Genre VARCHAR(127) NOT NULL,
    imgURL VARCHAR(200) NOT NULL,
    Description VARCHAR(500),
    ReleaseDate INT(255) NOT NULL,
    Length VARCHAR(127) NOT NULL,
    Likes INT(255) NOT NULL,
    Comments INT(255) NOT NULL,
    PRIMARY KEY (ID)
)