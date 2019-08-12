/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50726
Source Host           : localhost:3306
Source Database       : test

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2019-08-12 11:30:26
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for menus
-- ----------------------------
DROP TABLE IF EXISTS `menus`;
CREATE TABLE `menus` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `url` varchar(255) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `pid` int(11) NOT NULL DEFAULT '0',
  `sort` int(11) NOT NULL DEFAULT '9999',
  `is_menu` tinyint(1) NOT NULL DEFAULT '1',
  `remark` varchar(255) NOT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of menus
-- ----------------------------
INSERT INTO `menus` VALUES ('1', '首页', '/index', '0', '0', '0', '0', '', '2019-07-16 19:19:04', '2019-07-16 19:19:04', null);
INSERT INTO `menus` VALUES ('2', '用户中心', '/user', '0', '0', '0', '0', '', '2019-07-16 19:22:28', '2019-07-16 19:22:28', null);
INSERT INTO `menus` VALUES ('3', '用户管理', '/user/list', '0', '2', '4', '0', '', '2019-07-19 10:58:46', '2019-07-16 19:24:23', null);
INSERT INTO `menus` VALUES ('4', '角色管理', '/role/list', '0', '2', '0', '0', '', '2019-07-19 11:04:00', '2019-07-19 11:04:00', null);

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `menu_ids` text,
  `remark` varchar(255) NOT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of roles
-- ----------------------------
INSERT INTO `roles` VALUES ('1', '超级管理员', '0', '', '', '2019-07-16 19:35:01', '2019-07-16 19:35:01', null);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) DEFAULT NULL,
  `password` varchar(64) DEFAULT NULL,
  `status` tinyint(1) DEFAULT '1',
  `token` varchar(64) DEFAULT NULL,
  `role_id` int(11) DEFAULT '0',
  `remark` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('5', '111', '$2a$10$uiTPKfnChHJ5XOcIAQ.BJuEjy.s56d8OOsOcvNJWsxDJIcZBTX7Ti', '0', '31231', '0', '', null, '2019-07-15 01:54:34', '2019-07-15 09:54:26');
INSERT INTO `users` VALUES ('6', 'admin', '$2a$10$uiTPKfnChHJ5XOcIAQ.BJuEjy.s56d8OOsOcvNJWsxDJIcZBTX7Ti', '1', '31231222', '0', '22222', '2019-07-01 19:23:45', '2019-07-16 07:42:39', null);
INSERT INTO `users` VALUES ('7', '14123121', '123456@', '1', '31231222', '0', '22222', '2019-07-11 16:52:34', '2019-07-16 07:42:41', '2019-07-15 15:57:31');
INSERT INTO `users` VALUES ('8', '14123121', '123456@', '1', '31231222', '0', '22222', '2019-07-11 16:53:03', '2019-07-16 07:42:42', '2019-07-15 16:00:39');
INSERT INTO `users` VALUES ('9', '14123121', '123456@', '1', '31231222', '0', '22222', '2019-07-11 17:04:17', '2019-07-16 07:42:43', null);
