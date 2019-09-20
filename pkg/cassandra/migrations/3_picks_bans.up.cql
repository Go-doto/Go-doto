CREATE TABLE IF NOT EXISTS picks_bans
(
    PRIMARY KEY (match_id, action_order),
    match_id     bigint,
    is_pick      Boolean,
    hero_id      int,
    team         int,
    action_order int
);