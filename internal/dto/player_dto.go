package dto

type CreatePlayerRequest struct {
	TeamID       *int    `json:"team_id"`
	Name         string  `json:"name" binding:"required,max=100"`
	Height       float64 `json:"height"`
	Weight       float64 `json:"weight"`
	Position     string  `json:"position" binding:"required"`
	JerseyNumber int     `json:"jersey_number"`
}

type UpdatePlayerRequest struct {
	TeamID       *int    `json:"team_id"`
	Name         string  `json:"name" binding:"required,max=100"`
	Height       float64 `json:"height"`
	Weight       float64 `json:"weight"`
	Position     string  `json:"position" binding:"required"`
	JerseyNumber int     `json:"jersey_number"`
}

type PlayerResponse struct {
	ID           int     `json:"id"`
	TeamID       *int    `json:"team_id"`
	Name         string  `json:"name"`
	Height       float64 `json:"height"`
	Weight       float64 `json:"weight"`
	Position     string  `json:"position"`
	JerseyNumber int     `json:"jersey_number"`
}
