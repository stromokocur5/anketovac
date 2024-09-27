CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE poll_type AS ENUM ('multiple', 'ranking', 'image');

CREATE TABLE IF NOT EXISTS polls_anketovac 
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), 
    title VARCHAR(100) NOT NULL CHECK (LENGTH(title) > 0),
    description VARCHAR(255), 
    type poll_type NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
);

CREATE TABLE IF NOT EXISTS poll_options
(
    id SERIAL PRIMARY KEY,
    poll_id UUID REFERENCES polls_anketovac(id) ON DELETE CASCADE, 
    name VARCHAR(100) NOT NULL CHECK (LENGTH(name) > 0), 
    votes INT DEFAULT 0 CHECK (votes >= 0) 
);

CREATE INDEX idx_poll_options_poll_id ON poll_options(poll_id);
