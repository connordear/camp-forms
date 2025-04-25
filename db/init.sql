CREATE TABLE IF NOT EXISTS version (
	current_version text PRIMARY KEY,
	last_updated timestamp DEFAULT now()
);

INSERT INTO version (current_version) VALUES ('0.0.0');

