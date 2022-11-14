## summary
`person_record` stores person's score and clear status for each music.

### table information
| physical      | logical       | primary key | auto increment |
|:--------------|:--------------|:------------|:---------------|
| person_record | person record | record_id   | record_id      |

### column information
| physical      | logical      | type    | unsigned | not null | default                        | extra |
|:--------------|:-------------|:--------|:---------|:---------|:-------------------------------|:------|
| record_id     | record id    | INTEGER |          | Y        |                                |       |
| person_id     | person id    | INTEGER |          | Y        |                                |       |
| music_id      | music id     | INTEGER |          | Y        |                                |       |
| record_easy   | easy clear   | INTEGER |          |          |                                | *1    |
| score_easy    | easy score   | INTEGER |          |          |                                |       |
| record_normal | normal clear | INTEGER |          |          |                                | *1    |
| score_normal  | normal score | INTEGER |          |          |                                |       |
| record_hard   | hard clear   | INTEGER |          |          |                                | *1    |
| score_hard    | hard score   | INTEGER |          |          |                                |       |
| record_expert | expert clear | INTEGER |          |          |                                | *1    |
| score_expert  | expert score | INTEGER |          |          |                                |       |
| record_master | master clear | INTEGER |          |          |                                | *1    |
| score_master  | master score | INTEGER |          |          |                                |       |
| created_at    | create date  | TEXT    |          | Y        | (DATETIME('now', 'localtime')) |       |
| updated_at    | update date  | TEXT    |          | Y        | (DATETIME('now', 'localtime')) |       |

### unique key information
| No | definition         | extra |
|:---|:-------------------|-------|
| 1  | person_id,music_id |       |

### index key information
| No | definition | extra   |
|:---|:-----------|---------|
| 1  | music_id   | for kpi |

### extra
*1 ... Type of status  
Defined as

| value | type mean   |
|:------|:------------|
| 0     | not cleared |
| 1     | cleared     |
| 2     | full combo  |
| 3     | all perfect |

---
[back to index](./index.md)