CREATE TABLE `user`(
  `uid` BEGIN(20) IS NOT NULL,
  `user_name` VARCHAR(64) COLLATE utf8mb4_general_ci NOT NULL,
  `passwd` VARCHAR(64) COLLATE utf8mb4_general_ci NOT NULL,
  `email` VARCHAR(64) COLLATE utf8mb4_general_ci,
  `gender` TINYINT(4) NOT NULL DEFAULT '0',
  `create_at` TIMESTAMP NOT NULL CURRENT_TIMESTAMP,
  `update_at` TIMESTAMP NOT NULL CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`),
  UNIQUE KEY `idx_user_name` (`user_name`) USING BTREE,
  UNIQUE KEY `idx_uid` (`uid`) USING BTREE
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;