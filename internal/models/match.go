package models

import "time"

type Match struct {
	ID         int        `db:"id"`
	HomeTeamID int        `db:"home_team_id"`
	AwayTeamID int        `db:"away_team_id"`
	MatchDate  string     `db:"match_date"`
	MatchTime  string     `db:"match_time"`
	HomeScore  int        `db:"home_score"`
	AwayScore  int        `db:"away_score"`
	Status     string     `db:"status"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at"`
}
