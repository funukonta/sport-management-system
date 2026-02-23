package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sport-management-system/internal/dto"
	"github.com/sport-management-system/internal/services"
	"github.com/sport-management-system/internal/utils"
)

type PlayerHandler struct {
	playerService *services.PlayerService
}

func NewPlayerHandler(playerService *services.PlayerService) *PlayerHandler {
	return &PlayerHandler{playerService: playerService}
}

func (h *PlayerHandler) Create(c *gin.Context) {
	var req dto.CreatePlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.NewValidationError(err.Error()))
		return
	}

	player, err := h.playerService.Create(req)
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, http.StatusCreated, "player created successfully", dto.PlayerResponse{
		ID:           player.ID,
		TeamID:       player.TeamID,
		Name:         player.Name,
		Height:       player.Height,
		Weight:       player.Weight,
		Position:     player.Position,
		JerseyNumber: player.JerseyNumber,
	})
}

func (h *PlayerHandler) FindAll(c *gin.Context) {
	var pagination dto.PaginationQuery
	if err := c.ShouldBindQuery(&pagination); err != nil {
		c.Error(utils.NewValidationError(err.Error()))
		return
	}

	limit := pagination.GetLimit()
	offset := pagination.GetOffset()
	page := pagination.GetPage()

	players, total, err := h.playerService.FindAllPaginated(limit, offset)
	if err != nil {
		c.Error(err)
		return
	}

	var response []dto.PlayerResponse
	for _, p := range players {
		response = append(response, dto.PlayerResponse{
			ID:           p.ID,
			TeamID:       p.TeamID,
			Name:         p.Name,
			Height:       p.Height,
			Weight:       p.Weight,
			Position:     p.Position,
			JerseyNumber: p.JerseyNumber,
		})
	}

	meta := &utils.Meta{
		Total:    total,
		Page:     page,
		PageSize: limit,
	}

	utils.SuccessWithMeta(c, http.StatusOK, "success", response, meta)
}

func (h *PlayerHandler) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(utils.NewBadRequestError("invalid player ID"))
		return
	}

	player, err := h.playerService.FindByID(int(id))
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, http.StatusOK, "success", dto.PlayerResponse{
		ID:           player.ID,
		TeamID:       player.TeamID,
		Name:         player.Name,
		Height:       player.Height,
		Weight:       player.Weight,
		Position:     player.Position,
		JerseyNumber: player.JerseyNumber,
	})
}

func (h *PlayerHandler) FindByTeamID(c *gin.Context) {
	teamID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(utils.NewBadRequestError("invalid team ID"))
		return
	}

	players, err := h.playerService.FindByTeamID(int(teamID))
	if err != nil {
		c.Error(err)
		return
	}

	var response []dto.PlayerResponse
	for _, p := range players {
		response = append(response, dto.PlayerResponse{
			ID:           p.ID,
			TeamID:       p.TeamID,
			Name:         p.Name,
			Height:       p.Height,
			Weight:       p.Weight,
			Position:     p.Position,
			JerseyNumber: p.JerseyNumber,
		})
	}

	utils.Success(c, http.StatusOK, "success", response)
}

func (h *PlayerHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(utils.NewBadRequestError("invalid player ID"))
		return
	}

	var req dto.UpdatePlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.NewValidationError(err.Error()))
		return
	}

	if err := h.playerService.Update(id, req); err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, http.StatusOK, "player updated successfully", nil)
}

func (h *PlayerHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(utils.NewBadRequestError("invalid player ID"))
		return
	}

	if err := h.playerService.Delete(id); err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, http.StatusOK, "player deleted successfully", nil)
}
