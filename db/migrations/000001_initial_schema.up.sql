CREATE TABLE IF NOT EXISTS example 
(
    id varchar PRIMARY KEY NOT NULL,
    name varchar NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL
);