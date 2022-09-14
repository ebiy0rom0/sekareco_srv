DROP TABLE IF EXISTS person;
CREATE TABLE IF NOT EXISTS person (
    person_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    person_name TEXT NOT NULL,
    friend_code INTEGER NOT NULL,
    is_compare INTEGER NOT NULL DEFAULT 0,
    created_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
    updated_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime'))
);
CREATE UNIQUE INDEX unq_person_1 ON person (
    friend_code
);

DROP TABLE IF EXISTS person_login;
CREATE TABLE IF NOT EXISTS person_login (
    person_id INTEGER NOT NULL PRIMARY KEY,
    login_id TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
    updated_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime'))
);
CREATE UNIQUE INDEX unq_login_1 ON person_login (
    login_id
);