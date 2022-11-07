## summary
`master_music` is the master table that stores music information.

### table information
| physical     | logical      | primary key | auto increment |
|:-------------|:-------------|:------------|:---------------|
| master_music | music master | music_id    |                |

### column information
| physical     | logical      | type    | unsigned | not null | default                        | extra                           |
|:-------------|:-------------|:--------|:---------|:---------|:-------------------------------|:--------------------------------|
| music_id     | music id     | INTEGER |          | Y        |                                |                                 |
| artist_id    | artist id    | INTEGER |          | Y        |                                |                                 |
| music_name   | music name   | TEXT    |          | Y        |                                |                                 |
| jacket_url   | jacket url   | TEXT    |          | Y        |                                | description from root directory |
| level_easy   | easy level   | INTEGER |          | Y        |                                |                                 |
| notes_easy   | easy notes   | INTEGER |          | Y        |                                |                                 |
| level_normal | normal level | INTEGER |          | Y        |                                |                                 |
| notes_normal | nomal notes  | INTEGER |          | Y        |                                |                                 |
| level_hard   | hard level   | INTEGER |          | Y        |                                |                                 |
| notes_hard   | hard notes   | INTEGER |          | Y        |                                |                                 |
| level_expert | expert level | INTEGER |          | Y        |                                |                                 |
| notes_expert | expert notes | INTEGER |          | Y        |                                |                                 |
| level_master | master level | INTEGER |          | Y        |                                |                                 |
| notes_master | master notes | INTEGER |          | Y        |                                |                                 |
| created_at   | create date  | TEXT    |          | Y        | (DATETIME('now', 'localtime')) |                                 |
| updated_at   | update date  | TEXT    |          | Y        | (DATETIME('now', 'localtime')) |                                 |

### unique key information
| No | definition | extra |
|:---|:-----------|-------|
|    |            |       |

### index key information
| No | definition | extra |
|:---|:-----------|-------|
|    |            |       |

### extra
None specified.

---
[back to index](./index.md)