## summary
`friend` stores person's friend information.

### table information
| physical | logical | primary key | auto increment |
|:---------|:--------|:------------|:---------------|
| friend   | friend  | friend_id   | friend_id      |

### column information
| physical         | logical          | type    | unsigned | not null | default                        | extra |
|:-----------------|:-----------------|:--------|:---------|:---------|:-------------------------------|:------|
| friend_id        | friend id        | INTEGER |          | Y        |                                |       |
| person_id        | person id        | INTEGER |          | Y        |                                |       |
| friend_person_id | friend person id | INTEGER |          | Y        |                                |       |
| created_at       | create date      | TEXT    |          | Y        | (DATETIME('now', 'localtime')) |       |
| updated_at       | update date      | TEXT    |          | Y        | (DATETIME('now', 'localtime')) |       |

### unique key information
| No | definition | extra |
|:---|:-----------|-------|
|    |            |       |

### index key information
| No | definition | extra |
|:---|:-----------|-------|
| 1  | person_id  |       |

### extra
None specified.