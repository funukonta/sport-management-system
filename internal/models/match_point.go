package models

import "time"

type MatchLog struct {
	ID        int        `db:"id"`
	MatchID   int        `db:"match_id"`
	PlayerID  int        `db:"player_id"`
	Minute    int        `db:"minute"`
	EventType string     `db:"event_type"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
