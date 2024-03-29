INSERT INTO person(person_name, friend_code)
VALUES  ("name01", 2593519733),
        ("name02", 2593519734),
        ("name03", 2593519735)
;


-- password => $2a$12$J1t7JRNU4Dnq2MYLQHIrNOdpRzqD008mwQrnhPdfYdnM9sd3QmPNC
INSERT INTO person_login(person_id, login_id, password_hash)
VALUES  (1, 'login_id1', '$2a$12$J1t7JRNU4Dnq2MYLQHIrNOdpRzqD008mwQrnhPdfYdnM9sd3QmPNC'),
        (2, 'login_id2', '$2a$12$J1t7JRNU4Dnq2MYLQHIrNOdpRzqD008mwQrnhPdfYdnM9sd3QmPNC'),
        (3, 'login_id3', '$2a$12$J1t7JRNU4Dnq2MYLQHIrNOdpRzqD008mwQrnhPdfYdnM9sd3QmPNC')
;


-- primary key(music_id) is auto increment value
INSERT INTO master_music(
    music_id,
    group_id,
    music_name,
    jacket_url,
    level_easy,
    notes_easy,
    level_normal,
    notes_normal,
    level_hard,
    notes_hard,
    level_expert,
    notes_expert,
    level_master,
    notes_master
)
VALUES  (1, 1, "test_music001", "jacket/m_001.png", 1, 100, 2, 200, 3, 300, 4, 400, 5, 500),
        (2, 2, "test_music002", "jacket/m_002.png", 2, 200, 3, 300, 4, 400, 5, 500, 6, 600),
        (3, 1, "test_music003", "jacket/m_003.png", 3, 300, 4, 400, 5, 500, 6, 600, 7, 700)
;


-- person_id: 3 is not registered record person
INSERT INTO person_record(
    person_id,
    music_id,
    record_easy,
    score_easy,
    record_normal,
    score_normal,
    record_hard,
    score_hard,
    record_expert,
    score_expert,
    record_master,
    score_master
)
VALUES  (2, 1, 3, 300, 3, 600, 3, 900,  2, 1195, 2, 1480),
        (2, 2, 3, 600, 3, 900, 2, 1180, 1, 1460, 2, 1700)
;
