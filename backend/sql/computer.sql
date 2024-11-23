CREATE SCHEMA IF NOT EXISTS computers_schema;

CREATE TABLE IF NOT EXISTS computers_schema.computer (
     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
     os TEXT,
     cpu TEXT,
     ram INTEGER,
     status BOOLEAN,
     ssh TEXT
);