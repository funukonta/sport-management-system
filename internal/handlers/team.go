package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sport-management-system/internal/dto"
	"github.com/sport-management-system/internal/services"
	"github.com/sport-management-system/internal/utils"
)

type TeamHandler struct {
	teamService *services.TeamService
}

func NewTeamHandler(teamService *services.TeamService) *TeamHandler {
	return &TeamHandler{teamService: teamService}
}

func (h *TeamHandler) Create(c *gin.Context) {
	var req dto.CreateTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.NewValidationError(err.Error()))
		return
	}

	team, err := h.teamService.Create(req)
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, http.StatusCreated, "team created successfully", dto.TeamResponse{
		ID:          team.ID,
		Name:        team.Name,
		Logo:        team.Logo,
		FoundedYear: team.FoundedYear,
		Address:     team.Address,
		City:        team.City,
	})
}

func (h *TeamHandler) FindAll(c *gin.Context) {
	teams, err := h.teamService.FindAll()
	if err != nil {
		_ = c.Error(err)
		return
	}

	var response []dto.TeamResponse
	for _, t := range teams {
		response = append(response, dto.TeamResponse{
			ID:          t.ID,
			Name:        t.Name,
			Logo:        t.Logo,
			FoundedYear: t.FoundedYear,
			Address:     t.Address,
			City:        t.City,
		})
	}

	utils.Success(c, http.StatusOK, "success", response)
}

func (h *TeamHandler) FindByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(utils.NewBadRequestError("invalid team ID"))
		return
	}

	team, err := h.teamService.FindByID(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}

	utils.Success(c, http.StatusOK, "success", dto.TeamResponse{
		ID:          team.ID,
		Name:        team.Name,
		Logo:        team.Logo,
		FoundedYear: team.FoundedYear,
		Address:     team.Address,
		City:        team.City,
	})
}

func (h *TeamHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(utils.NewBadRequestError("invalid team ID"))
		return
	}

	var req dto.UpdateTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(utils.NewValidationError(err.Error()))
		return
	}

	if err := h.teamService.Update(uint(id), req); err != nil {
		_ = c.Error(err)
		return
	}

	utils.Success(c, http.StatusOK, "team updated successfully", nil)
}

func (h *TeamHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(utils.NewBadRequestError("invalid team ID"))
		return
	}

	if err := h.teamService.Delete(uint(id)); err != nil {
		_ = c.Error(err)
		return
	}

	utils.Success(c, http.StatusOK, "team deleted successfully", nil)
}
