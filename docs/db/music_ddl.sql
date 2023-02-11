DROP TABLE IF EXISTS master_music;
CREATE TABLE IF NOT EXISTS master_music (
    music_id INTEGER NOT NULL PRIMARY KEY,
    group_id INTEGER NOT NULL,
    music_name TEXT NOT NULL,
    jacket_url TEXT NOT NULL,
    level_easy INTEGER NOT NULL,
    notes_easy INTEGER NOT NULL,
    level_normal INTEGER NOT NULL,
    notes_normal INTEGER NOT NULL,
    level_hard INTEGER NOT NULL,
    notes_hard INTEGER NOT NULL,
    level_expert INTEGER NOT NULL,
    notes_expert INTEGER NOT NULL,
    level_master INTEGER NOT NULL,
    notes_master INTEGER NOT NULL,
    created_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
    updated_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime'))
);

DROP TABLE IF EXISTS master_group;
CREATE TABLE IF NOT EXISTS master_group (
    group_id INTEGER NOT NULL PRIMARY KEY,
    group_name TEXT NOT NULL,
    logo_url TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
    updated_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime'))
);
