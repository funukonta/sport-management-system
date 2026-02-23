package repository

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/sport-management-system/internal/models"
	"github.com/sport-management-system/internal/utils"
)

type MatchRepository struct {
	db *sqlx.DB
}

func NewMatchRepository(db *sqlx.DB) *MatchRepository {
	return &MatchRepository{db: db}
}

func (r *MatchRepository) Create(match *models.Match) error {
	query := `INSERT INTO matches (home_team_id, away_team_id, match_date, match_time)
			  VALUES ($1, $2, $3, $4)
			  RETURNING id, home_score, away_score, status, created_at, updated_at`
	err := r.db.QueryRow(query, match.HomeTeamID, match.AwayTeamID, match.MatchDate, match.MatchTime).Scan(&match.ID, &match.HomeScore, &match.AwayScore, &match.Status, &match.CreatedAt, &match.UpdatedAt)
	if err != nil {
		if utils.IsUniqueDataError(err) {
			return utils.NewBadRequestError(ErrMatchExists)
		}
		return err
	}
	return nil
}

func (r *MatchRepository) FindAll() ([]models.Match, error) {
	var matches []models.Match
	query := `SELECT id, home_team_id, away_team_id, match_date, match_time, home_score, away_score, status, created_at, updated_at
			  FROM matches WHERE deleted_at IS NULL ORDER BY match_date, match_time`
	err := r.db.Select(&matches, query)
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func (r *MatchRepository) FindAllPaginated(limit int, offset int) ([]models.Match, int, error) {
	var matches []models.Match
	var total int

	// Get total count
	countQuery := `SELECT COUNT(*) FROM matches WHERE deleted_at IS NULL`
	err := r.db.Get(&total, countQuery)
	if err != nil {
		return nil, 0, err
	}

	// Get paginated data
	query := `SELECT id, home_team_id, away_team_id, match_date, match_time, home_score, away_score, status, created_at, updated_at
			  FROM matches WHERE deleted_at IS NULL ORDER BY match_date, match_time LIMIT $1 OFFSET $2`
	err = r.db.Select(&matches, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return matches, total, nil
}

func (r *MatchRepository) FindByID(id int) (*models.Match, error) {
	var match models.Match
	query := `SELECT id, home_team_id, away_team_id, match_date, match_time, home_score, away_score, status, created_at, updated_at
			  FROM matches WHERE id = $1 AND deleted_at IS NULL`
	err := r.db.Get(&match, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.NewBadRequestError(ErrMatchNotFound)
		}
		return nil, err
	}
	return &match, nil
}

func (r *MatchRepository) Update(match *models.Match) error {
	query := `UPDATE matches SET home_team_id = $1, away_team_id = $2, match_date = $3, match_time = $4, updated_at = NOW()
			  WHERE id = $5 AND deleted_at IS NULL AND status = 'scheduled'`
	result, err := r.db.Exec(query, match.HomeTeamID, match.AwayTeamID, match.MatchDate, match.MatchTime, match.ID)
	if err != nil {
		if utils.IsUniqueDataError(err) {
			return utils.NewBadRequestError(ErrMatchExists)
		}
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return utils.NewBadRequestError(ErrMatchNotFound)
	}
	return nil
}

func (r *MatchRepository) Delete(id int) error {
	query := `UPDATE matches SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return utils.NewBadRequestError(ErrMatchNotFound)
	}
	return nil
}

func (r *MatchRepository) UpdateResult(id int, homeScore int, awayScore int) error {
	query := `UPDATE matches SET home_score = $1, away_score = $2, status = 'finished', updated_at = NOW()
			  WHERE id = $3 AND deleted_at IS NULL AND status = 'scheduled'`
	result, err := r.db.Exec(query, homeScore, awayScore, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return utils.NewBadRequestError(ErrMatchNotFound)
	}
	return nil
}
