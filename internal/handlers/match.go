package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sport-management-system/internal/dto"
	"github.com/sport-management-system/internal/services"
	"github.com/sport-management-system/internal/utils"
)

type MatchHandler struct {
	matchService *services.MatchService
}

func NewMatchHandler(matchService *services.MatchService) *MatchHandler {
	return &MatchHandler{matchService: matchService}
}

func (h *MatchHandler) Create(c *gin.Context) {
	var req dto.CreateMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.NewValidationError(err.Error()))
		return
	}

	match, err := h.matchService.Create(req)
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, http.StatusCreated, "match created successfully", dto.MatchResponse{
		ID:         match.ID,
		HomeTeamID: match.HomeTeamID,
		AwayTeamID: match.AwayTeamID,
		MatchDate:  match.MatchDate,
		MatchTime:  match.MatchTime,
		HomeScore:  match.HomeScore,
		AwayScore:  match.AwayScore,
		Status:     match.Status,
	})
}

func (h *MatchHandler) FindAll(c *gin.Context) {
	var pagination dto.PaginationQuery
	if err := c.ShouldBindQuery(&pagination); err != nil {
		c.Error(utils.NewValidationError(err.Error()))
		return
	}

	limit := pagination.GetLimit()
	offset := pagination.GetOffset()
	page := pagination.GetPage()

	matches, total, err := h.matchService.FindAllPaginated(limit, offset)
	if err != nil {
		c.Error(err)
		return
	}

	var response []dto.MatchResponse
	for _, m := range matches {
		response = append(response, dto.MatchResponse{
			ID:         m.ID,
			HomeTeamID: m.HomeTeamID,
			AwayTeamID: m.AwayTeamID,
			MatchDate:  m.MatchDate,
			MatchTime:  m.MatchTime,
			HomeScore:  m.HomeScore,
			AwayScore:  m.AwayScore,
			Status:     m.Status,
		})
	}

	meta := &utils.Meta{
		Total:    total,
		Page:     page,
		PageSize: limit,
	}

	utils.SuccessWithMeta(c, http.StatusOK, "success", response, meta)
}

func (h *MatchHandler) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(utils.NewBadRequestError("invalid match ID"))
		return
	}

	match, err := h.matchService.FindByID(id)
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, http.StatusOK, "success", dto.MatchResponse{
		ID:         match.ID,
		HomeTeamID: match.HomeTeamID,
		AwayTeamID: match.AwayTeamID,
		MatchDate:  match.MatchDate,
		MatchTime:  match.MatchTime,
		HomeScore:  match.HomeScore,
		AwayScore:  match.AwayScore,
		Status:     match.Status,
	})
}

func (h *MatchHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(utils.NewBadRequestError("invalid match ID"))
		return
	}

	var req dto.UpdateMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.NewValidationError(err.Error()))
		return
	}

	if err := h.matchService.Update(id, req); err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, http.StatusOK, "match updated successfully", nil)
}

func (h *MatchHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(utils.NewBadRequestError("invalid match ID"))
		return
	}

	if err := h.matchService.Delete(id); err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, http.StatusOK, "match deleted successfully", nil)
}

func (h *MatchHandler) AddMatchLog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(utils.NewBadRequestError("invalid match ID"))
		return
	}

	var req dto.AddMatchLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.NewValidationError(err.Error()))
		return
	}

	matchLog, err := h.matchService.AddMatchLog(id, req)
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, http.StatusCreated, "match log added successfully", matchLog)
}

func (h *MatchHandler) FinishMatch(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(utils.NewBadRequestError("invalid match ID"))
		return
	}

	if err := h.matchService.FinishMatch(id); err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, http.StatusOK, "match finished successfully", nil)
}

func (h *MatchHandler) GetMatchReport(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(utils.NewBadRequestError("invalid match ID"))
		return
	}

	report, err := h.matchService.GetMatchReport(id)
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, http.StatusOK, "success", report)
}
