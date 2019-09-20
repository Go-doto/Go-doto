ALTER TABLE picks_bans
    WITH compaction = { 'class' : 'LeveledCompactionStrategy' };