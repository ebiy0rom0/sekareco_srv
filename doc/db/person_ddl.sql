DROP TABLE IF EXISTS person;
CREATE TABLE IF NOT EXISTS person (
    person_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    person_name TEXT NOT NULL,
    friend_code TEXT NOT NULL,
    insert_date DATETIME NOT NULL,
    update_date DATETIME NOT NULL,
);

DROP TABLE IF EXISTS person_login;
CREATE TABLE IF NOT EXISTS person_login (
    login_id TEXT NOT NULL PRIMARY KEY,
    person_id INTEGER NOT NULL,
    password_hash TEXT NOT NULL,
    insert_date DATETIME NOT NULL,
    update_date DATETIME NOT NULL,
);