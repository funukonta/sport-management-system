package dto

type CreateTeamRequest struct {
	Name        string `json:"name" binding:"required,max=100"`
	Logo        string `json:"logo"`
	FoundedYear int    `json:"founded_year"`
	Address     string `json:"address"`
	City        string `json:"city"`
}

type UpdateTeamRequest struct {
	Name        string `json:"name" binding:"required,max=100"`
	Logo        string `json:"logo"`
	FoundedYear int    `json:"founded_year"`
	Address     string `json:"address"`
	City        string `json:"city"`
}

type TeamResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Logo        string `json:"logo"`
	FoundedYear int    `json:"founded_year"`
	Address     string `json:"address"`
	City        string `json:"city"`
}
