CREATE TABLE IF NOT EXISTS contacts(
   contact_id serial PRIMARY KEY,
   name VARCHAR (50) NOT NULL,
   phone VARCHAR (50) NOT NULL,
   group_name VARCHAR (50) NOT NULL
);