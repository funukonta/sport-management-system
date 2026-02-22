package models

import "time"

type Team struct {
	ID          uint       `db:"id"`
	Name        string     `db:"name"`
	Logo        string     `db:"logo"`
	FoundedYear int        `db:"founded_year"`
	Address     string     `db:"address"`
	City        string     `db:"city"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}
