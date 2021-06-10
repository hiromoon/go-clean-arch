CREATE DATABASE IF NOT EXISTS `test`;
GRANT ALL ON root.* TO 'root'@'%';

CREATE TABLE IF NOT EXISTS `test`.`users` (
  id varchar(64),
  name varchar(64),
  password varchar(64)
);
