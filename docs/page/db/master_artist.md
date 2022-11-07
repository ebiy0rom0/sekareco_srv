## summary
`master_artist` is the master table that stores artist information.

### table information
| physical      | logical       | primary key | auto increment |
|:--------------|:--------------|:------------|:---------------|
| master_artist | artist master | artist_id   |                |

### column information
| physical    | logical     | type    | unsigned | not null | default                        | extra |
|:------------|:------------|:--------|:---------|:---------|:-------------------------------|:------|
| artist_id   | artist id   | INTEGER |          | Y        |                                |       |
| artist_name | artist name | TEXT    |          | Y        |                                |       |
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