package models

import "time"

type Player struct {
	ID           int        `db:"id"`
	TeamID       *int       `db:"team_id"`
	Name         string     `db:"name"`
	Height       float64    `db:"height"`
	Weight       float64    `db:"weight"`
	Position     string     `db:"position"`
	JerseyNumber int        `db:"jersey_number"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at"`
}
