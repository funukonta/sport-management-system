-- Delete in reverse order of creation (handle foreign key constraints)
DELETE FROM match_logs WHERE match_id > 2;
DELETE FROM matches WHERE id > 2;
DELETE FROM players WHERE team_id > 3;
DELETE FROM teams WHERE id > 3;

