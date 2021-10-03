CREATE TABLE `user`(
  `uid` BIGINT(20) NOT NULL,
  `user_name` VARCHAR(64) NOT NULL,
  `passwd` VARCHAR(64) NOT NULL,
  `email` VARCHAR(64),
  `gender` TINYINT(4) NOT NULL DEFAULT '0',
  `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`),
  UNIQUE KEY `idx_user_name` (`user_name`) USING BTREE,
  UNIQUE KEY `idx_uid` (`uid`) USING BTREE
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

CREATE TABLE `category`(
    `cid` int not null AUTO_INCREMENT,
    `c_name` varchar(64) not null,
    primary key (`cid`),
    unique key idx_cid (`cid`) using btree
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

CREATE TABLE `posts`(
  `pid` BIGINT(20) not null,
  `title` varchar(255) not null,
  `content` varchar(6000) not null,
  `cid` int not null,
  `auth_id` BIGINT(20) not null,
  `create_at` TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP,
  `update_at` TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  primary key (pid),
  unique key idx_pid (`pid`) using btree,
  unique key idx_cid (`pid`) using btree
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;