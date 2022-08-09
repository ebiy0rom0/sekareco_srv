INSERT INTO person(person_name, friend_code)
VALUES  ("name01", 2593519733),
        ("name02", 2593519734),
        ("name03", 2593519735)
;


-- password => $2a$12$J1t7JRNU4Dnq2MYLQHIrNOdpRzqD008mwQrnhPdfYdnM9sd3QmPNC
INSERT INTO person_login(login_id, person_id, password_hash)
VALUES  ('login_id1', 1, '$2a$12$J1t7JRNU4Dnq2MYLQHIrNOdpRzqD008mwQrnhPdfYdnM9sd3QmPNC'),
        ('login_id2', 2, '$2a$12$J1t7JRNU4Dnq2MYLQHIrNOdpRzqD008mwQrnhPdfYdnM9sd3QmPNC'),
        ('login_id3', 3, '$2a$12$J1t7JRNU4Dnq2MYLQHIrNOdpRzqD008mwQrnhPdfYdnM9sd3QmPNC')
;


-- primary key(music_id) is auto increment value
INSERT INTO master_music(artist_id, music_name, jacket_url, level_easy, level_normal, level_hard, level_expert, level_master)
VALUES  (1, "test_music001", "jacket/m_001.png", 1, 2, 3, 4, 5),
        (2, "test_music002", "jacket/m_002.png", 2, 3, 4, 5, 6),
        (1, "test_music003", "jacket/m_003.png", 3, 4, 5, 6, 7)
;


-- person_id: 3 is not registered record person
INSERT INTO record(person_id, music_id, record_easy, record_normal, record_hard, record_expert, record_master)
VALUES  (2, 1, 3, 3, 3, 2, 2),
        (2, 2, 3, 3, 2, 1, 2)
;
