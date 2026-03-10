-- liquibase formatted sql

-- changeset manan:001
CREATE TABLE teams (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    age INT,
    player_count INT,
    points INT DEFAULT 0,

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_teams_deleted_at ON teams(deleted_at);