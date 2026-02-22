CREATE TABLE IF NOT EXISTS match_logs (
    id SERIAL PRIMARY KEY,
    match_id INT NOT NULL REFERENCES matches(id),
    player_id INT NOT NULL REFERENCES players(id),
    minute INT NOT NULL,
    event_type VARCHAR(50) NOT NULL CHECK (event_type IN ('goal', 'assist', 'yellow_card', 'red_card')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL
);
