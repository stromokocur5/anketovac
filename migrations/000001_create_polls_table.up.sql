CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE poll_type AS ENUM ('multiple', 'ranking', 'image');

CREATE TABLE IF NOT EXISTS polls_anketovac 
(
    id VARCHAR(16) PRIMARY KEY DEFAULT substr(gen_random_uuid()::text, 1, 16), 
    title VARCHAR(100) NOT NULL CHECK (LENGTH(title) > 0),
    description VARCHAR(255), 
    ptype poll_type NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS poll_options
(
    poll_id VARCHAR(16) REFERENCES polls_anketovac(id) ON DELETE CASCADE, 
    option_order INT NOT NULL CHECK (option_order >= 0),
    name VARCHAR(100) NOT NULL CHECK (LENGTH(name) > 0), 
    votes INT NOT NULL DEFAULT 0 CHECK (votes >= 0),
    PRIMARY KEY (poll_id, name, option_order)
);

CREATE INDEX idx_poll_options_poll_id ON poll_options(poll_id);
