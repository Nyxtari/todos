CREATE SCHEMA IF NOT EXISTS app;

SET search_path TO app;

CREATE TABLE IF NOT EXISTS todos (
  id SERIAL PRIMARY KEY,
  isDone      BOOLEAN,
	title       VARCHAR(50),
	description VARCHAR(256),
	startDate   timestamp without time zone DEFAULT NOW() NOT NULL,
	endDate     timestamp without time zone DEFAULT NOW() NOT NULL,
	difficulty  INTEGER
);

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50),
  email VARCHAR(100) UNIQUE
);

INSERT INTO todos (isDone, title, description, startDate, endDate, difficulty) VALUES
  (true, 'hello', 'hello first', now(), now(), 5);

INSERT INTO users (name, email) VALUES
  ('Alpha', 'alpha@example.com'),
  ('Beta', 'beta@example.com'),
  ('Gamma', 'gamma@example.com')
ON CONFLICT (email) DO NOTHING;

SELECT * FROM todos;