package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sport-management-system/internal/models"
)

type MatchLogRepository struct {
	db *sqlx.DB
}

func NewMatchLogRepository(db *sqlx.DB) *MatchLogRepository {
	return &MatchLogRepository{db: db}
}

func (r *MatchLogRepository) Create(matchLog *models.MatchLog) error {
	query := `INSERT INTO match_logs (match_id, player_id, minute, event_type)
			  VALUES ($1, $2, $3, $4)
			  RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(query, matchLog.MatchID, matchLog.PlayerID, matchLog.Minute, matchLog.EventType).Scan(&matchLog.ID, &matchLog.CreatedAt, &matchLog.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *MatchLogRepository) FindByMatchID(matchID int) ([]models.MatchLog, error) {
	var matchLogs []models.MatchLog
	query := `SELECT id, match_id, player_id, minute, event_type, created_at, updated_at
			  FROM match_logs WHERE match_id = $1 AND deleted_at IS NULL ORDER BY minute`
	err := r.db.Select(&matchLogs, query, matchID)
	if err != nil {
		return nil, err
	}
	return matchLogs, nil
}

func (r *MatchLogRepository) CountWins(teamID int, matchID int) (int, error) {
	var count int
	query := `
		SELECT COUNT(*) FROM matches
		WHERE deleted_at IS NULL
		  AND status = 'finished'
		  AND id <= $2
		  AND (
		    (home_team_id = $1 AND home_score > away_score)
		    OR
		    (away_team_id = $1 AND away_score > home_score)
		  )`
	err := r.db.Get(&count, query, teamID, matchID)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *MatchLogRepository) CountGoalsByTeamInMatch(matchID int, teamID int) (int, error) {
	var count int
	query := `
		SELECT COUNT(*) FROM match_logs ml
		JOIN players p ON ml.player_id = p.id
		WHERE ml.match_id = $1
		  AND ml.deleted_at IS NULL
		  AND p.team_id = $2
		  AND ml.event_type = 'goal'`
	err := r.db.Get(&count, query, matchID, teamID)
	if err != nil {
		return 0, err
	}
	return count, nil
}
