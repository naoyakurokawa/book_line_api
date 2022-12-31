DROP SCHEMA IF EXISTS book_line;
CREATE SCHEMA book_line;
USE book_line;

DROP TABLE IF EXISTS books;
CREATE TABLE `books`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '本の識別子',
    `title`    VARCHAR(128) NOT NULL COMMENT '本のタイトル',
    `isbn`   VARCHAR(20)  NOT NULL COMMENT '本のISBN',
    `created`  DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ブック';

INSERT INTO books (
  id,
  title,
  isbn,
  created,
  modified
) VALUES (
  1, 
  "サンプルブック", 
  "9784863543720",
  "2022-12-22 00:00:00.000000",
  "2022-12-22 00:00:00.000000"
);

DROP TABLE IF EXISTS users;
CREATE TABLE `users`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザーの識別子',
    `name`     varchar(20) NOT NULL COMMENT 'ユーザー名',
    `password`    VARCHAR(128) NOT NULL COMMENT 'パスワードハッシュ',
    `role`     VARCHAR(80) NOT NULL COMMENT 'ロール(admin or normal)',
    `created`  DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';

INSERT INTO users (
  id,
  name,
  password,
  role,
  created,
  modified
) VALUES (
  1, 
  "admin", 
  "admin3294",
  "admin",
  "2022-12-22 00:00:00.000000",
  "2022-12-22 00:00:00.000000"
),
(
  2, 
  "normal", 
  "normal9834",
  "normal",
  "2022-12-22 00:00:00.000000",
  "2022-12-22 00:00:00.000000"
);