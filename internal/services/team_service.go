package services

import (
	"github.com/sport-management-system/internal/dto"
	"github.com/sport-management-system/internal/models"
	"github.com/sport-management-system/internal/repository"
)

type TeamService struct {
	teamRepo *repository.TeamRepository
}

func NewTeamService(teamRepo *repository.TeamRepository) *TeamService {
	return &TeamService{teamRepo: teamRepo}
}

func (s *TeamService) Create(req dto.CreateTeamRequest) (*models.Team, error) {
	team := &models.Team{
		Name:        req.Name,
		Logo:        req.Logo,
		FoundedYear: req.FoundedYear,
		Address:     req.Address,
		City:        req.City,
	}

	if err := s.teamRepo.Create(team); err != nil {
		return nil, err
	}

	return team, nil
}

func (s *TeamService) FindAll() ([]models.Team, error) {
	return s.teamRepo.FindAll()
}

func (s *TeamService) FindByID(id uint) (*models.Team, error) {
	return s.teamRepo.FindByID(id)
}

func (s *TeamService) Update(id uint, req dto.UpdateTeamRequest) error {
	team := &models.Team{
		ID:          id,
		Name:        req.Name,
		Logo:        req.Logo,
		FoundedYear: req.FoundedYear,
		Address:     req.Address,
		City:        req.City,
	}

	return s.teamRepo.Update(team)
}

func (s *TeamService) Delete(id uint) error {
	return s.teamRepo.Delete(id)
}
