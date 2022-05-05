
DROP TABLE IF EXISTS friend;
CREATE TABLE IF NOT EXISTS friend (
    friend_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    person_id INTEGER NOT NULL,
    friend_person_id INTEGER NOT NULL,
    is_conpare INTEGER NOT NULL DEFAULT 0,
    insert_date DATETIME NOT NULL,
    update_date DATETIME NOT NULL
);
CREATE INDEX idx_friend_1 ON friend (
    person_id
);