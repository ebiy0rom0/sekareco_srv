## summary
`person` stores person's information.

### table information
| physical | logical | primary key | auto increment |
|:---------|:--------|:------------|:---------------|
| person   | person  | person_id   | person_id      |

### column information
| physical      | logical        | type    | unsigned | not null | default                        | extra               |
|:--------------|:---------------|:--------|:---------|:---------|:-------------------------------|:--------------------|
| person_id     | person id      | INTEGER |          | Y        |                                |                     |
| person_name   | person name    | TEXT    |          | Y        |                                |                     |
| friend_code   | friend code    | INTEGER |          | Y        |                                |                     |
| is_compare    | compare permit | INTEGER |          | Y        | 0                              | 0:private, 1:public |
| created_at    | create date    | TEXT    |          | Y        | (DATETIME('now', 'localtime')) |                     |
| updated_at    | update date    | TEXT    |          | Y        | (DATETIME('now', 'localtime')) |                     |

### unique key information
| No | definition  | extra |
|:---|:------------|-------|
| 1  | friend_code |       |

### index key information
| No | definition | extra |
|:---|:-----------|-------|
|    |            |       |

### extra
None specified.

---
[back to index](./index.md)