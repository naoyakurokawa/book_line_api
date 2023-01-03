DROP SCHEMA IF EXISTS book_line;
CREATE SCHEMA book_line;
USE book_line;

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
  "$2a$10$8UMx9XoCCYEkdGoaFySvV.fN6HUNofYzPMJQCgh6j6mePNFVXhFsm",
  "admin",
  "2022-12-22 00:00:00.000000",
  "2022-12-22 00:00:00.000000"
),
(
  2, 
  "user", 
  "$2a$10$dx8NNMcM0VBaeb2qk05da.7Dx0Cu3rUpaNJbinbYP75D4SxC0X2uu",
  "user",
  "2022-12-22 00:00:00.000000",
  "2022-12-22 00:00:00.000000"
);

DROP TABLE IF EXISTS books;
CREATE TABLE `books`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '本の識別子',
    `isbn`   VARCHAR(20)  NOT NULL COMMENT '本のISBN',
    `created`  DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ブック';

INSERT INTO books (
  id,
  isbn,
  created,
  modified
) VALUES (
  1,
  "9784863543720",
  "2022-12-22 00:00:00.000000",
  "2022-12-22 00:00:00.000000"
),
 (
  2,
  "9784815607654",
  "2022-12-22 00:00:00.000000",
  "2022-12-22 00:00:00.000000"
);

DROP TABLE IF EXISTS book_memos;
CREATE TABLE `book_memos`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '本のメモの識別子',
    `book_id`  BIGINT UNSIGNED NOT NULL COMMENT '本の識別子',
    `page`     INT NOT NULL COMMENT '本のページ',
    `detail`   TEXT NOT NULL COMMENT 'メモの詳細',
    `created`  DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `modified` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    FOREIGN KEY (`book_id`) REFERENCES books(`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='本のメモ';

INSERT INTO book_memos (
  id,
  book_id,
  page,
  detail,
  created,
  modified
) VALUES (
  1,
  1,
  231,
  "embed: 埋め込むという意味", 
  "2022-12-22 00:00:00.000000",
  "2022-12-22 00:00:00.000000"
);