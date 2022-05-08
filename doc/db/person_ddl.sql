DROP TABLE IF EXISTS person;
CREATE TABLE IF NOT EXISTS person (
    person_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    person_name TEXT NOT NULL,
    friend_code INTEGER NOT NULL,
    insert_date TEXT NOT NULL DEFAULT (DATETIME(`now`, `localtime`)),
    update_date TEXT NOT NULL DEFAULT (DATETIME(`now`, `localtime`))
);
CREATE INDEX idx_person_1 ON person (
    friend_code
);

DROP TABLE IF EXISTS person_login;
CREATE TABLE IF NOT EXISTS person_login (
    login_id TEXT NOT NULL PRIMARY KEY,
    person_id INTEGER NOT NULL,
    password_hash TEXT NOT NULL,
    insert_date TEXT NOT NULL DEFAULT (DATETIME(`now`, `localtime`)),
    update_date TEXT NOT NULL DEFAULT (DATETIME(`now`, `localtime`))
);