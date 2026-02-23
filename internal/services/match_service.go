package services

import (
	"fmt"

	"github.com/sport-management-system/internal/dto"
	"github.com/sport-management-system/internal/models"
	"github.com/sport-management-system/internal/repository"
	"github.com/sport-management-system/internal/utils"
)

var validEventTypes = map[string]bool{
	"goal":        true,
	"assist":      true,
	"yellow_card": true,
	"red_card":    true,
}

type MatchService struct {
	matchRepo    *repository.MatchRepository
	teamRepo     *repository.TeamRepository
	playerRepo   *repository.PlayerRepository
	matchLogRepo *repository.MatchLogRepository
}

func NewMatchService(matchRepo *repository.MatchRepository, teamRepo *repository.TeamRepository, playerRepo *repository.PlayerRepository, matchLogRepo *repository.MatchLogRepository) *MatchService {
	return &MatchService{
		matchRepo:    matchRepo,
		teamRepo:     teamRepo,
		playerRepo:   playerRepo,
		matchLogRepo: matchLogRepo,
	}
}

func (s *MatchService) Create(req dto.CreateMatchRequest) (*models.Match, error) {
	if req.HomeTeamID == req.AwayTeamID {
		return nil, utils.NewBadRequestError(ErrSameTeam)
	}

	_, err := s.teamRepo.FindByID(req.HomeTeamID)
	if err != nil {
		return nil, utils.NewBadRequestError(ErrHomeTeamNotFound)
	}

	_, err = s.teamRepo.FindByID(req.AwayTeamID)
	if err != nil {
		return nil, utils.NewBadRequestError(ErrAwayTeamNotFound)
	}

	match := &models.Match{
		HomeTeamID: req.HomeTeamID,
		AwayTeamID: req.AwayTeamID,
		MatchDate:  req.MatchDate,
		MatchTime:  req.MatchTime,
	}

	err = s.matchRepo.Create(match)
	if err != nil {
		return nil, err
	}

	return match, nil
}

func (s *MatchService) FindAll() ([]models.Match, error) {
	return s.matchRepo.FindAll()
}

func (s *MatchService) FindAllPaginated(limit int, offset int) ([]models.Match, int, error) {
	return s.matchRepo.FindAllPaginated(limit, offset)
}

func (s *MatchService) FindByID(id int) (*models.Match, error) {
	return s.matchRepo.FindByID(id)
}

func (s *MatchService) Update(id int, req dto.UpdateMatchRequest) error {
	if req.HomeTeamID == req.AwayTeamID {
		return utils.NewBadRequestError(ErrSameTeam)
	}

	match, err := s.matchRepo.FindByID(id)
	if err != nil {
		return err
	}

	if match.Status != "scheduled" {
		return utils.NewBadRequestError(ErrMatchFinished)
	}

	_, err = s.teamRepo.FindByID(req.HomeTeamID)
	if err != nil {
		return utils.NewBadRequestError(ErrHomeTeamNotFound)
	}

	_, err = s.teamRepo.FindByID(req.AwayTeamID)
	if err != nil {
		return utils.NewBadRequestError(ErrAwayTeamNotFound)
	}

	match.HomeTeamID = req.HomeTeamID
	match.AwayTeamID = req.AwayTeamID
	match.MatchDate = req.MatchDate
	match.MatchTime = req.MatchTime

	return s.matchRepo.Update(match)
}

func (s *MatchService) Delete(id int) error {
	return s.matchRepo.Delete(id)
}

func (s *MatchService) AddMatchLog(matchID int, req dto.AddMatchLogRequest) (*models.MatchLog, error) {
	match, err := s.matchRepo.FindByID(matchID)
	if err != nil {
		return nil, err
	}

	if match.Status != "scheduled" {
		return nil, utils.NewBadRequestError(ErrMatchFinished)
	}

	if !validEventTypes[req.EventType] {
		return nil, utils.NewBadRequestError(fmt.Sprintf("%s: %s", ErrInvalidEventType, req.EventType))
	}

	player, err := s.playerRepo.FindByID(req.PlayerID)
	if err != nil {
		return nil, utils.NewBadRequestError(fmt.Sprintf("%s: %d", ErrMatchPlayerNotFound, req.PlayerID))
	}

	if player.TeamID == nil || (*player.TeamID != match.HomeTeamID && *player.TeamID != match.AwayTeamID) {
		return nil, utils.NewBadRequestError(fmt.Sprintf("%s: %d", ErrPlayerNotInMatch, req.PlayerID))
	}

	matchLog := &models.MatchLog{
		MatchID:   matchID,
		PlayerID:  req.PlayerID,
		Minute:    req.Minute,
		EventType: req.EventType,
	}

	err = s.matchLogRepo.Create(matchLog)
	if err != nil {
		return nil, err
	}

	return matchLog, nil
}

func (s *MatchService) FinishMatch(matchID int) error {
	match, err := s.matchRepo.FindByID(matchID)
	if err != nil {
		return err
	}

	if match.Status != "scheduled" {
		return utils.NewBadRequestError(ErrResultAlreadyReported)
	}

	homeScore, err := s.matchLogRepo.CountGoalsByTeamInMatch(matchID, match.HomeTeamID)
	if err != nil {
		return err
	}

	awayScore, err := s.matchLogRepo.CountGoalsByTeamInMatch(matchID, match.AwayTeamID)
	if err != nil {
		return err
	}

	return s.matchRepo.UpdateResult(matchID, homeScore, awayScore)
}

func (s *MatchService) GetMatchReport(matchID int) (*dto.MatchReportResponse, error) {
	match, err := s.matchRepo.FindByID(matchID)
	if err != nil {
		return nil, err
	}

	if match.Status != "finished" {
		return nil, utils.NewBadRequestError(ErrMatchNotFinished)
	}

	homeTeam, err := s.teamRepo.FindByID(match.HomeTeamID)
	if err != nil {
		return nil, err
	}

	awayTeam, err := s.teamRepo.FindByID(match.AwayTeamID)
	if err != nil {
		return nil, err
	}

	matchLogs, err := s.matchLogRepo.FindByMatchID(matchID)
	if err != nil {
		return nil, err
	}

	var result string
	if match.HomeScore > match.AwayScore {
		result = "Home Win"
	} else if match.AwayScore > match.HomeScore {
		result = "Away Win"
	} else {
		result = "Draw"
	}

	goalCount := make(map[int]int)
	playerNames := make(map[int]string)
	var logResponses []dto.MatchLogResponse

	for _, ml := range matchLogs {
		player, err := s.playerRepo.FindByID(ml.PlayerID)
		if err != nil {
			continue
		}

		logResponses = append(logResponses, dto.MatchLogResponse{
			PlayerID:   ml.PlayerID,
			PlayerName: player.Name,
			Minute:     ml.Minute,
			EventType:  ml.EventType,
		})

		if ml.EventType == "goal" {
			goalCount[ml.PlayerID]++
			playerNames[ml.PlayerID] = player.Name
		}
	}

	var topScorer *dto.TopScorer
	maxGoals := 0
	for playerID, goals := range goalCount {
		if goals > maxGoals {
			maxGoals = goals
			topScorer = &dto.TopScorer{
				PlayerID: playerID,
				Name:     playerNames[playerID],
				Goals:    goals,
			}
		}
	}

	homeTeamWins, err := s.matchLogRepo.CountWins(match.HomeTeamID, matchID)
	if err != nil {
		return nil, err
	}

	awayTeamWins, err := s.matchLogRepo.CountWins(match.AwayTeamID, matchID)
	if err != nil {
		return nil, err
	}

	report := &dto.MatchReportResponse{
		MatchID:   match.ID,
		MatchDate: match.MatchDate,
		MatchTime: match.MatchTime,
		HomeTeam: dto.TeamInfo{
			ID:   homeTeam.ID,
			Name: homeTeam.Name,
		},
		AwayTeam: dto.TeamInfo{
			ID:   awayTeam.ID,
			Name: awayTeam.Name,
		},
		HomeScore:         match.HomeScore,
		AwayScore:         match.AwayScore,
		Status:            match.Status,
		Result:            result,
		TopScorer:         topScorer,
		MatchLogs:         logResponses,
		HomeTeamTotalWins: homeTeamWins,
		AwayTeamTotalWins: awayTeamWins,
	}

	return report, nil
}
