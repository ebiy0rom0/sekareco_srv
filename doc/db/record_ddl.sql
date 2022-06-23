DROP TABLE IF EXISTS record;
CREATE TABLE IF NOT EXISTS record (
    record_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    person_id INTEGER NOT NULL,
    music_id INTEGER NOT NULL,
    record_easy INTEGER,
    record_normal INTEGER,
    record_hard INTEGER,
    record_expert INTEGER,
    record_master INTEGER,
    created_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
    updated_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime'))
);
CREATE UNIQUE INDEX unq_record_1 ON record (
    person_id,
    music_id
);