package dto

type CreateMatchRequest struct {
	HomeTeamID int    `json:"home_team_id" binding:"required"`
	AwayTeamID int    `json:"away_team_id" binding:"required"`
	MatchDate  string `json:"match_date" binding:"required"`
	MatchTime  string `json:"match_time" binding:"required"`
}

type UpdateMatchRequest struct {
	HomeTeamID int    `json:"home_team_id" binding:"required"`
	AwayTeamID int    `json:"away_team_id" binding:"required"`
	MatchDate  string `json:"match_date" binding:"required"`
	MatchTime  string `json:"match_time" binding:"required"`
}

type MatchResponse struct {
	ID         int    `json:"id"`
	HomeTeamID int    `json:"home_team_id"`
	AwayTeamID int    `json:"away_team_id"`
	MatchDate  string `json:"match_date"`
	MatchTime  string `json:"match_time"`
	HomeScore  int    `json:"home_score"`
	AwayScore  int    `json:"away_score"`
	Status     string `json:"status"`
}
