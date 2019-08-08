/*
Navicat MySQL Data Transfer

Source Server         : 虚拟机
Source Server Version : 50722
Source Host           : 10.20.29.148:13306
Source Database       : jenkins

Target Server Type    : MYSQL
Target Server Version : 50722
File Encoding         : 65001

Date: 2019-08-08 14:52:19
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT NULL,
  `mobile` varchar(20) DEFAULT NULL,
  `timestamp` datetime DEFAULT NULL COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', '张三', '1346169461', '2019-08-08 14:16:04');
