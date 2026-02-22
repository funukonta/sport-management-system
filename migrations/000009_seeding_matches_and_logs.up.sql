-- Seed 2 matches (Persija Jakarta vs Persib Bandung, Arema FC vs Persija Jakarta)
INSERT INTO matches (home_team_id, away_team_id, match_date, match_time, home_score, away_score, status)
VALUES
(1, 2, '2026-01-15', '19:00:00', 3, 2, 'finished'),
(3, 1, '2026-01-22', '15:30:00', 1, 4, 'finished');

-- Match 1: Persija Jakarta (team 1) vs Persib Bandung (team 2) — 3-2
-- Assuming match_id = 1
-- Persija players: IDs 1-7, Persib players: IDs 8-14
INSERT INTO match_logs (match_id, player_id, minute, event_type)
VALUES
-- Goals
(1, 5, 12, 'goal'),       -- Bagus Ramadhan (Persija) scores
(1, 6, 13, 'assist'),     -- Ilham Maulana (Persija) assist
(1, 12, 25, 'goal'),      -- Ardiansyah Putra (Persib) equalizes
(1, 11, 25, 'assist'),    -- Mochammad Iqbal (Persib) assist
(1, 7, 38, 'goal'),       -- Yoga Prasetyo (Persija) scores
(1, 4, 38, 'assist'),     -- Dimas Saputra (Persija) assist
(1, 14, 55, 'goal'),      -- Rafli Mahendra (Persib) scores
(1, 13, 55, 'assist'),    -- Bayu Anggara (Persib) assist
(1, 5, 78, 'goal'),       -- Bagus Ramadhan (Persija) scores winner
(1, 7, 78, 'assist'),     -- Yoga Prasetyo (Persija) assist
-- Cards
(1, 3, 30, 'yellow_card'),  -- Fajar Nugroho (Persija) yellow card
(1, 9, 42, 'yellow_card'),  -- Rian Kurniawan (Persib) yellow card
(1, 10, 67, 'yellow_card'), -- Tegar Hidayat (Persib) yellow card
(1, 10, 82, 'red_card');    -- Tegar Hidayat (Persib) second yellow -> red card

-- Match 2: Arema FC (team 3) vs Persija Jakarta (team 1) — 1-4
-- Assuming match_id = 2
-- Arema players: IDs 15-21, Persija players: IDs 1-7
INSERT INTO match_logs (match_id, player_id, minute, event_type)
VALUES
-- Goals
(2, 5, 8, 'goal'),        -- Bagus Ramadhan (Persija) scores
(2, 6, 8, 'assist'),      -- Ilham Maulana (Persija) assist
(2, 7, 22, 'goal'),       -- Yoga Prasetyo (Persija) scores
(2, 4, 22, 'assist'),     -- Dimas Saputra (Persija) assist
(2, 19, 35, 'goal'),      -- Hendra Wijaya (Arema) pulls one back
(2, 18, 35, 'assist'),    -- Galang Putra (Arema) assist
(2, 5, 60, 'goal'),       -- Bagus Ramadhan (Persija) hat-trick
(2, 7, 60, 'assist'),     -- Yoga Prasetyo (Persija) assist
(2, 6, 88, 'goal'),       -- Ilham Maulana (Persija) seals it
(2, 5, 88, 'assist'),     -- Bagus Ramadhan (Persija) assist
-- Cards
(2, 16, 18, 'yellow_card'),  -- Reza Gunawan (Arema) yellow card
(2, 2, 40, 'yellow_card'),   -- Andi Setiawan (Persija) yellow card
(2, 17, 52, 'yellow_card'),  -- Wahyu Prakoso (Arema) yellow card
(2, 20, 70, 'yellow_card'),  -- Iqbal Ramadhan (Arema) yellow card
(2, 20, 85, 'red_card');     -- Iqbal Ramadhan (Arema) second yellow -> red card
