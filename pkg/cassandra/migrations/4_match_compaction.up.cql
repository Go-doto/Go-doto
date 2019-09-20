ALTER TABLE matches
    WITH compaction = { 'class' : 'LeveledCompactionStrategy' };