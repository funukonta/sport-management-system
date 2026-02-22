package dto

type AddMatchLogRequest struct {
	PlayerID  int    `json:"player_id" binding:"required"`
	Minute    int    `json:"minute" binding:"required"`
	EventType string `json:"event_type" binding:"required"`
}

type MatchLogResponse struct {
	PlayerID   int    `json:"player_id"`
	PlayerName string `json:"player_name"`
	Minute     int    `json:"minute"`
	EventType  string `json:"event_type"`
}

type MatchReportResponse struct {
	MatchID           int                `json:"match_id"`
	MatchDate         string             `json:"match_date"`
	MatchTime         string             `json:"match_time"`
	HomeTeam          TeamInfo           `json:"home_team"`
	AwayTeam          TeamInfo           `json:"away_team"`
	HomeScore         int                `json:"home_score"`
	AwayScore         int                `json:"away_score"`
	Status            string             `json:"status"`
	Result            string             `json:"result"`
	TopScorer         *TopScorer         `json:"top_scorer,omitempty"`
	MatchLogs         []MatchLogResponse `json:"match_logs"`
	HomeTeamTotalWins int                `json:"home_team_total_wins"`
	AwayTeamTotalWins int                `json:"away_team_total_wins"`
}

type TopScorer struct {
	PlayerID int    `json:"player_id"`
	Name     string `json:"name"`
	Goals    int    `json:"goals"`
}

type TeamInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
