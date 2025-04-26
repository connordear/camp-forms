CREATE TABLE IF NOT EXISTS version (
	current_version text PRIMARY KEY,
	last_updated timestamp DEFAULT now()
);

INSERT INTO version (current_version) VALUES ('0.0.0');

CREATE TABLE IF NOT EXISTS camps (
	id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	name string NOT NULL
);

INSERT INTO camps (name) VALUES ('New Camp');

