package services

import (
	"github.com/sport-management-system/internal/dto"
	"github.com/sport-management-system/internal/models"
	"github.com/sport-management-system/internal/repository"
	"github.com/sport-management-system/internal/utils"
)

type PlayerService struct {
	playerRepo *repository.PlayerRepository
	teamRepo   *repository.TeamRepository
}

func NewPlayerService(playerRepo *repository.PlayerRepository, teamRepo *repository.TeamRepository) *PlayerService {
	return &PlayerService{playerRepo: playerRepo, teamRepo: teamRepo}
}

func (s *PlayerService) Create(req dto.CreatePlayerRequest) (*models.Player, error) {
	if !s.isValidPosition(req.Position) {
		return nil, utils.NewBadRequestError(ErrInvalidPosition)
	}

	if req.TeamID != nil {
		_, err := s.teamRepo.FindByID(*req.TeamID)
		if err != nil {
			return nil, err
		}
	}

	player := &models.Player{
		TeamID:       req.TeamID,
		Name:         req.Name,
		Height:       req.Height,
		Weight:       req.Weight,
		Position:     req.Position,
		JerseyNumber: req.JerseyNumber,
	}

	err := s.playerRepo.Create(player)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func (s *PlayerService) FindAll() ([]models.Player, error) {
	return s.playerRepo.FindAll()
}

func (s *PlayerService) FindAllPaginated(limit int, offset int) ([]models.Player, int, error) {
	return s.playerRepo.FindAllPaginated(limit, offset)
}

func (s *PlayerService) FindByID(id int) (*models.Player, error) {
	return s.playerRepo.FindByID(id)
}

func (s *PlayerService) FindByTeamID(teamID int) ([]models.Player, error) {
	// Check if team exists before fetching players
	_, err := s.teamRepo.FindByID(teamID)
	if err != nil {
		return nil, err
	}

	return s.playerRepo.FindByTeamID(teamID)
}

func (s *PlayerService) Update(id int, req dto.UpdatePlayerRequest) error {
	if !s.isValidPosition(req.Position) {
		return utils.NewBadRequestError(ErrInvalidPosition)
	}

	if req.TeamID != nil {
		if _, err := s.teamRepo.FindByID(*req.TeamID); err != nil {
			return err
		}
	}

	player := &models.Player{
		ID:           id,
		TeamID:       req.TeamID,
		Name:         req.Name,
		Height:       req.Height,
		Weight:       req.Weight,
		Position:     req.Position,
		JerseyNumber: req.JerseyNumber,
	}

	return s.playerRepo.Update(player)
}

func (s *PlayerService) Delete(id int) error {
	return s.playerRepo.Delete(id)
}

func (s *PlayerService) isValidPosition(position string) bool {
	valid := map[string]bool{
		"forward":    true,
		"midfielder": true,
		"defender":   true,
		"goalkeeper": true,
	}
	return valid[position]
}
