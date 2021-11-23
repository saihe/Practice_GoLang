drop table if exists `employee`;
create table `employee` (
  `id` int not null  auto_increment primary key comment '連番',
  `name` varchar(64)     comment '名前',
  `department_id` int     comment '所属部署ID',
  `created_at` datetime not null default current_timestamp   comment '作成日時。',
  `updated_at` datetime not null default current_timestamp on update current_timestamp   comment '更新日時。'
) ENGINE = InnoDB DEFAULT CHARSET utf8;
drop table if exists `department`;
create table `department` (
  `id` int not null  auto_increment primary key comment '連番',
  `name` varchar(64)     comment '名前',
  `created_at` datetime not null default current_timestamp   comment '作成日時。',
  `updated_at` datetime not null default current_timestamp on update current_timestamp   comment '更新日時。'
) ENGINE = InnoDB DEFAULT CHARSET utf8;
