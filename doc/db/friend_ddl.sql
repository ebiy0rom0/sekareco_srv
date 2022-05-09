
DROP TABLE IF EXISTS friend;
CREATE TABLE IF NOT EXISTS friend (
    friend_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    person_id INTEGER NOT NULL,
    friend_person_id INTEGER NOT NULL,
    is_compare INTEGER NOT NULL DEFAULT 0,
    insert_date TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
    update_date TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime'))
);
CREATE INDEX idx_friend_1 ON friend (
    person_id
);