CREATE TABLE birds (
  id SERIAL PRIMARY KEY,
  species VARCHAR(256),
  description VARCHAR(1024)
);

INSERT INTO birds (species, description) VALUES ('Canary', 'Small yellow bird');
select * from birds;
select species, description from birds;