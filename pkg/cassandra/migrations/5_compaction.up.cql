ALTER TABLE player_matches
    WITH compaction = { 'class' : 'LeveledCompactionStrategy' };
