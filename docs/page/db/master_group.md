## summary
`master_group` is the master table that stores group information.

### table information
| physical      | logical       | primary key | auto increment |
|:--------------|:--------------|:------------|:---------------|
| master_group  | group master  | group_id   |                |

### column information
| physical    | logical     | type    | unsigned | not null | default                        | extra |
|:------------|:------------|:--------|:---------|:---------|:-------------------------------|:------|
| group_id    | group id    | INTEGER |          | Y        |                                |       |
| group_name  | group name  | TEXT    |          | Y        |                                |       |
| logo_url    | logo url    | TEXT    |          | Y        |                                |       |
| created_at  | create date | TEXT    |          | Y        | (DATETIME('now', 'localtime')) |       |
| updated_at  | update date | TEXT    |          | Y        | (DATETIME('now', 'localtime')) |       |

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