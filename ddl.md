# テーブル定義

テーブル一覧　論理名（物理名）

* [従業員（employee）](#employee)

## employee

### employee_columns

name|type|not null|default|key|auto increment|extra|comment
---|---|---|---|---|---|---|---
id|int|y||primary key|y||連番
name|varchar(64)||||||名前
department_id|int||||||所属部署ID
created_at|datetime|y|current_timestamp||||作成日時。
updated_at|datetime|y|current_timestamp on update current_timestamp||||更新日時。

## department

### department_columns

name|type|not null|default|key|auto increment|extra|comment
---|---|---|---|---|---|---|---
id|int|y||primary key|y||連番
name|varchar(64)||||||名前
created_at|datetime|y|current_timestamp||||作成日時。
updated_at|datetime|y|current_timestamp on update current_timestamp||||更新日時。
