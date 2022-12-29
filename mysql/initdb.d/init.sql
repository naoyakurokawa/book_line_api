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