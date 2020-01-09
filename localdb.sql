CREATE TABLE [adopters] (
    [id] INTEGER NOT NULL PRIMARY KEY,
    [name] NVARCHAR(55)
    );
CREATE TABLE [adoptions] (
    [id] INTEGER NOT NULL PRIMARY KEY,
    [adopter_id] INTEGER,
    [animal_id] INTEGER
    );
CREATE TABLE [animals] (
    [id] INTEGER NOT NULL PRIMARY KEY,
    [specie] NVARCHAR(55),
    [adoptable] VARCHAR(5)
    );
CREATE TABLE [users] (
    [id] INTEGER NOT NULL PRIMARY KEY,
    [adopter_id] INTEGER
    );

INSERT INTO adopters (id, name) VALUES (1, "ti8m");

INSERT INTO adoptions (id, adopter_id, animal_id) VALUES (1, 1, 1);
INSERT INTO adoptions (id, adopter_id, animal_id) VALUES (2, 1, 2);

INSERT INTO animals (id, specie, adoptable) VALUES (1, "cat", "false");
INSERT INTO animals (id, specie, adoptable) VALUES (2, "dog", "false");
INSERT INTO animals (id, specie, adoptable) VALUES (3, "crocodile", "true");
INSERT INTO animals (id, specie, adoptable) VALUES (4, "kangaroo", "true");
INSERT INTO animals (id, specie, adoptable) VALUES (5, "giraffe", "true");
INSERT INTO animals (id, specie, adoptable) VALUES (6, "human", "true");
INSERT INTO animals (id, specie, adoptable) VALUES (7, "fish", "true");

INSERT INTO users (id, adopter_id) VALUES (1, 1);