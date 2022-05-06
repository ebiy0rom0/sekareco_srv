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
    insert_date DATETIME NOT NULL,
    update_date DATETIME NOT NULL
);
CREATE UNIQUE INDEX unq_record_1 ON record (
    person_id,
    music_id
);