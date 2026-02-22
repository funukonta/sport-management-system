CREATE TABLE IF NOT EXISTS players (
    id SERIAL PRIMARY KEY,
    team_id INT REFERENCES teams(id),
    name VARCHAR(100) NOT NULL,
    height DECIMAL(5,2),
    weight DECIMAL(5,2),
    position VARCHAR(20) NOT NULL CHECK (position IN ('forward', 'midfielder', 'defender', 'goalkeeper')),
    jersey_number INT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL
);

CREATE UNIQUE INDEX unique_jersey_per_team
    ON players (team_id, jersey_number)
    WHERE deleted_at IS NULL;
