docker run --name my_postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 postgres
-- Create a `snippets` table.
CREATE TABLE snippets (
id SERIAL PRIMARY KEY,
title VARCHAR(100) NOT NULL,
content TEXT NOT NULL,
created TIMESTAMP NOT NULL,
expires TIMESTAMP NOT NULL
);

-- Add an index on the created column.
CREATE INDEX idx_snippets_created ON snippets(created);

-- Add some dummy records (which we'll use in the next couple of chapters).
INSERT INTO snippets (title, content, created, expires) VALUES (
'An old silent pond',
'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
CURRENT_TIMESTAMP,
CURRENT_TIMESTAMP + INTERVAL '365 days'
);

INSERT INTO snippets (title, content, created, expires) VALUES (
'Over the wintry forest',
'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
CURRENT_TIMESTAMP,
CURRENT_TIMESTAMP + INTERVAL '365 days'
);

INSERT INTO snippets (title, content, created, expires) VALUES (
'First autumn morning',
'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
CURRENT_TIMESTAMP,
CURRENT_TIMESTAMP + INTERVAL '7 days'
);



-- Create the user 'web' with a password.
CREATE USER web WITH PASSWORD 'pass';

-- Grant the necessary permissions on the 'snippetbox' database.
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO web;

-- Ensure future tables also get the correct permissions.
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO web;
