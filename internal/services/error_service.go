package services

const (
	ErrTokenGenerationFailed = "failed to generate token"

	ErrInvalidPosition = "invalid position"

	ErrSameTeam              = "home team and away team cannot be the same"
	ErrHomeTeamNotFound      = "home team not found"
	ErrAwayTeamNotFound      = "away team not found"
	ErrMatchFinished         = "cannot update a finished match"
	ErrResultAlreadyReported = "match result already reported"
	ErrInvalidEventType      = "invalid event type"
	ErrPlayerNotInMatch      = "player does not belong to either team in this match"
	ErrMatchPlayerNotFound   = "player not found"
	ErrMatchNotFinished      = "match is not finished yet"
)
