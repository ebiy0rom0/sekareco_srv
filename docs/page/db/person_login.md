## summary
`person_login` stores person's login information.

### table information
| physical     | logical      | primary key | auto increment |
|:-------------|:-------------|:------------|:---------------|
| person_login | person login | person_id   |                |

### column information
| physical      | logical       | type    | unsigned | not null | default                        | extra                 |
|:--------------|:--------------|:--------|:---------|:---------|:-------------------------------|:----------------------|
| person_id     | person id     | INTEGER |          | Y        |                                |                       |
| login_id      | login id      | TEXT    |          | Y        |                                | user specified string |
| password_hash | password hash | TEXT    |          | Y        |                                |                       |
| created_at    | create date   | TEXT    |          | Y        | (DATETIME('now', 'localtime')) |                       |
| updated_at    | update date   | TEXT    |          | Y        | (DATETIME('now', 'localtime')) |                       |

### unique key information
| No | definition | extra |
|:---|:-----------|-------|
| 1  | login_id   |       |

### index key information
| No | definition | extra |
|:---|:-----------|-------|
|    |            |       |

### extra
None specified.