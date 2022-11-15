
DROP TABLE IF EXISTS person_friend;
CREATE TABLE IF NOT EXISTS person_friend (
    friend_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    person_id INTEGER NOT NULL,
    friend_person_id INTEGER NOT NULL,
    created_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
    updated_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime'))
);
CREATE INDEX idx_friend_1 ON person_friend (
    person_id
);