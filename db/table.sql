/*
 Navicat Premium Data Transfer
 Source Server         : local-dev
 Source Server Type    : MySQL
 Source Server Version : 50733
 Source Host           : localhost:3306
 Source Schema         : dns-dev
 Target Server Type    : MySQL
 Target Server Version : 50733
 File Encoding         : 65001
 Date: 26/03/2022 14:53:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint(20)  NOT NULL AUTO_INCREMENT,
  `email` varchar(200) DEFAULT NULL,
  `password` varchar(200) DEFAULT NULL,
  `token` varchar(24) DEFAULT NULL,
  `forbidden` tinyint(1) DEFAULT NULL,
  `roles` longtext,
  `permissions` longtext,
  `updated_unixtime` bigint(20) DEFAULT NULL,
  `created_unixtime` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `token` (`token`),
  KEY `idx_users_email` (`email`),
  KEY `idx_users_token` (`token`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;

