-- schema.sql

DROP DATABASE IF EXISTS posservice;

CREATE DATABASE posservice;

USE posservice;

GRANT SELECT, INSERT, UPDATE, DELETE ON posservice.* TO 'posservice'@'%' IDENTIFIED BY 'posservice';

DROP TABLE IF EXISTS `dst_user`;
CREATE TABLE `dst_user` (
    `id`            varchar(50)     NOT NULL,
    `name`          varchar(50)     NOT NULL,
    `create_at`     mediumint(10)   NOT NULL,
    `create_at_str` varchar(20)     NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `dst_user` (`id`,`name`,`create_at`,`create_at_str`) VALUES ('404504209241669642','dumh',0,'1970-1-1 00:00:00');
INSERT INTO `dst_user` (`id`,`name`,`create_at`,`create_at_str`) VALUES ('402631387577974797','stevenwong2017',0,'1970-1-1 00:00:00');
INSERT INTO `dst_user` (`id`,`name`,`create_at`,`create_at_str`) VALUES ('401916285929127947','Parker Lee',0,'1970-1-1 00:00:00');
INSERT INTO `dst_user` (`id`,`name`,`create_at`,`create_at_str`) VALUES ('396837819550662668','mako jr',0,'1970-1-1 00:00:00');
INSERT INTO `dst_user` (`id`,`name`,`create_at`,`create_at_str`) VALUES ('411932460344016896','cat lmao',0,'1970-1-1 00:00:00');
INSERT INTO `dst_user` (`id`,`name`,`create_at`,`create_at_str`) VALUES ('407552893806182411','lucky168',0,'1970-1-1 00:00:00');
INSERT INTO `dst_user` (`id`,`name`,`create_at`,`create_at_str`) VALUES ('403478549379678211','baobao',0,'1970-1-1 00:00:00');
INSERT INTO `dst_user` (`id`,`name`,`create_at`,`create_at_str`) VALUES ('403341228176965633','JWKY',0,'1970-1-1 00:00:00');
INSERT INTO `dst_user` (`id`,`name`,`create_at`,`create_at_str`) VALUES ('385061500034875392','RAY',0,'1970-1-1 00:00:00');

DROP TABLE IF EXISTS `dst_user_addr`;
CREATE TABLE `dst_user_addr` (
    `addr`            varchar(50)     NOT NULL,
    `userid`          varchar(50)     NOT NULL,
    `username`        varchar(50)     NOT NULL,
    PRIMARY KEY (`addr`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 主空投地址
INSERT INTO `dst_user_addr` (`addr`,`userid`,`username`) VALUES('fMZictACJc9dKhrMxCKRMkWNpq8Ni2ZBx8','404504209241669642','dumh');
-- 15.00
INSERT INTO `dst_user_addr` (`addr`,`userid`,`username`) VALUES('F9Kb7kzVa8ormWWJMwoyvpBZmG5W1LqUrs','404504209241669642','dumh');
-- 2.00
INSERT INTO `dst_user_addr` (`addr`,`userid`,`username`) VALUES('FCEj57L3DkVP2N8vccnZPQU6ciL4m3QrPZ','404504209241669642','dumh');
-- 10.00
INSERT INTO `dst_user_addr` (`addr`,`userid`,`username`) VALUES('FLC3yVLsv2pUSof4w9uZzf3ZosgYxZLiW1','404504209241669642','dumh');
-- 2.00
INSERT INTO `dst_user_addr` (`addr`,`userid`,`username`) VALUES('FFgiXPCVH47KwjqZjoV9cjpo2UfoVTnKyF','404504209241669642','dumh');
-- -1366.5001
-- INSERT INTO `dst_user_addr` (`addr`,`userid`,`username`) VALUES('FCd1oY8aMH2daf8YvBnSfiFafunj85RZfT','404504209241669642','dumh');
-- 11
INSERT INTO `dst_user_addr` (`addr`,`userid`,`username`) VALUES('FJaCXYHBd9gnWX4ZKvBmUWPqJatqD1X9aX','404504209241669642','dumh');

-- 10500
INSERT INTO `dst_user_addr` (`addr`,`userid`,`username`) VALUES('FB1uuSzdh5mscz8KwzgFSN4TAe6rhz2xj2','402631387577974797','stevenwong2017');
-- 39
INSERT INTO `dst_user_addr` (`addr`,`userid`,`username`) VALUES('F7pLR3gWjr4NQMMPLw3ZxBDC22q2ATmQwU','402631387577974797','stevenwong2017');

INSERT INTO `dst_user_addr` (`addr`,`userid`,`username`) VALUES('F6AjHLwgL6Yn2pPpL3762AZ5ZqkFuGvjZ7','401916285929127947','Parker Lee');

INSERT INTO `dst_user_addr` (`addr`,`userid`,`username`) VALUES('F6tyXagtuzhitN1Sjbg1PGedzTkDuVHDDp','396837819550662668','mako jr');

INSERT INTO `dst_user_addr` (`addr`,`userid`,`username`) VALUES('FKokKAbCw6cTm2iwBe3hhDBrCAv1EvaND6','411932460344016896','cat lmao');

INSERT INTO `dst_user_addr` (`addr`,`userid`,`username`) VALUES('F9TiwuFvp4oF6QksWBhA6qNpYFRkY2AS3n','407552893806182411','lucky168');

INSERT INTO `dst_user_addr` (`addr`,`userid`,`username`) VALUES('FQC361sfYtNrLiQjBDonzgFVTbCPMGdVKK','403478549379678211','baobao');

INSERT INTO `dst_user_addr` (`addr`,`userid`,`username`) VALUES('F8BcxZresqYT3pA1NugGdJ3VRqwd6ot3rR','403341228176965633','JWKY');

INSERT INTO `dst_user_addr` (`addr`,`userid`,`username`) VALUES('FTddbJCfWY6h5VfgZ1yzjbVCXhMuq5kx6X','385061500034875392','RAY');

CREATE TABLE dst_transaction (
  `txid`        VARCHAR(80)     NOT NULL,
  `category`    VARCHAR(20)     NOT NULL,
  `amount`      DOUBLE(16,8)    NOT NULL,
  `txtime`      mediumint(10)   NOT NULL,
  `txtime_str`  VARCHAR(20)     NOT NULL,
  PRIMARY KEY (`txid`),
  KEY `idx_dst_transactions_txtime` (`txtime`)
) ENGINE = innodb DEFAULT CHARSET = utf8;
