CREATE TABLE IF NOT EXISTS camps (
	id serial PRIMARY KEY,
	name text,
	last_updated timestamp DEFAULT now()
);
