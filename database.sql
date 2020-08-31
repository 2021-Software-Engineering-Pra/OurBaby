/*
MySQL Backup
Database: shop
Backup Time: 2020-01-05 12:22:03
*/

SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `shop`.`address`;
DROP TABLE IF EXISTS `shop`.`cart`;
DROP TABLE IF EXISTS `shop`.`orders`;
DROP TABLE IF EXISTS `shop`.`product`;
DROP TABLE IF EXISTS `shop`.`user`;
CREATE TABLE `address` (
  `aid` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(10) unsigned NOT NULL,
  `recvname` varchar(18) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `phone` bigint(20) unsigned DEFAULT NULL,
  `detailed` varchar(50) DEFAULT NULL,
  `creatime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updatime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`aid`),
  KEY `uid` (`uid`),
  CONSTRAINT `address_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `user` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `cart` (
  `cid` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(10) unsigned NOT NULL,
  `pid` int(10) unsigned NOT NULL,
  `pname` varchar(100) DEFAULT NULL,
  `peizhi` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `type` varchar(30) DEFAULT NULL,
  `linkid` int(10) unsigned DEFAULT NULL,
  `count` smallint(5) unsigned DEFAULT NULL,
  `aprice` int(10) unsigned DEFAULT NULL,
  `creatime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`cid`),
  KEY `pid` (`pid`),
  KEY `uid` (`uid`),
  CONSTRAINT `cart_ibfk_1` FOREIGN KEY (`pid`) REFERENCES `product` (`pid`),
  CONSTRAINT `cart_ibfk_2` FOREIGN KEY (`uid`) REFERENCES `user` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `orders` (
  `orid` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `uid` int(10) unsigned NOT NULL,
  `pid` varchar(30) NOT NULL,
  `uname` varchar(18) DEFAULT NULL,
  `pname` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `peizhi` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `linkid` varchar(10) DEFAULT NULL,
  `count` varchar(20) DEFAULT NULL,
  `aprice` varchar(30) DEFAULT NULL,
  `status` varchar(4) DEFAULT '未支付',
  `creatime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updatime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `recvname` varchar(18) DEFAULT NULL,
  `phone` varchar(12) DEFAULT NULL,
  `detailed` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`orid`),
  KEY `uid` (`uid`),
  KEY `uname` (`uname`),
  KEY `pid` (`pid`),
  CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `user` (`uid`),
  CONSTRAINT `uname` FOREIGN KEY (`uname`) REFERENCES `user` (`uname`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `product` (
  `pid` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `pname` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `pdesp` varchar(100) DEFAULT NULL,
  `peizhi` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `kucun` smallint(5) unsigned DEFAULT NULL,
  `price` int(10) unsigned DEFAULT '0',
  `type` varchar(30) DEFAULT NULL,
  `kind` varchar(10) DEFAULT NULL,
  `sellcount` int(10) unsigned DEFAULT '0',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '1',
  `linkid` int(10) DEFAULT '0',
  `creatime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updatime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`pid`,`pname`) USING BTREE,
  KEY `pid` (`pid`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `user` (
  `uid` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uname` varchar(18) NOT NULL,
  `upasswd` varchar(18) NOT NULL,
  `paywd` varchar(7) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `email` varchar(35) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `yue` int(10) unsigned DEFAULT '100000',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '1',
  `creatime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updatime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`,`uname`) USING BTREE,
  KEY `uname` (`uname`),
  KEY `uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
BEGIN;
LOCK TABLES `shop`.`address` WRITE;
DELETE FROM `shop`.`address`;
INSERT INTO `shop`.`address` (`aid`,`uid`,`recvname`,`phone`,`detailed`,`creatime`,`updatime`) VALUES (3, 1, '李晓明', 17843215678, '山东青岛崂山', '2019-12-15 10:57:47', '2019-12-22 10:50:04'),(8, 1, '张晓明', 12345678911, '山东省青岛市崂山区松岭路238号中国海洋大学东区', '2019-12-21 14:00:47', '2019-12-21 14:00:47'),(11, 1, '老张', 1234567892, '山东省青岛市崂山区中国海洋大学', '2019-12-22 16:13:47', '2020-01-05 12:19:50'),(15, 12, '谢叶子', 17860716615, '中国海洋大学', '2019-12-31 08:21:41', '2019-12-31 08:21:41');
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `shop`.`cart` WRITE;
DELETE FROM `shop`.`cart`;
INSERT INTO `shop`.`cart` (`cid`,`uid`,`pid`,`pname`,`peizhi`,`type`,`linkid`,`count`,`aprice`,`creatime`) VALUES (28, 1, 3, '拯救者 Y7000P 2019 英特尔酷睿i5 15.6英寸专业电竞本', 'i5-9300H/Windows 10 家庭中文版/15.6英寸/8G/1T SSD/ GeForce RTX™ 1660Ti 6G独显/黑色', '拯救者Y7000P', 2, 2, 7899, '2019-12-21 20:58:14'),(34, 1, 4, 'ThinkPad X1 Carbon 2019 英特尔酷睿 笔记本电脑 20QD002BCD', 'i7-8565U/Windows 10家庭中文版/8GB/512GB SSD/UHD 620显示芯片/14英寸WQHD IPS LED背光显示屏', 'ThinkPad X1 Carbon', 3, 1, 12499, '2020-01-04 22:53:12'),(35, 17, 2, '拯救者 Y7000P 2019 英特尔酷睿i7 15.6英寸专业电竞本', 'i7-9750H/Windows 10 家庭中文版/15.6英寸/16G/1T SSD/ GeForce RTX™ 2060 6G独显/红色', '拯救者Y7000P', 2, 1, 9898, '2020-01-05 11:31:48'),(36, 17, 3, '拯救者 Y7000P 2019 英特尔酷睿i5 15.6英寸专业电竞本', 'i5-9300H/Windows 10 家庭中文版/15.6英寸/8G/1T SSD/ GeForce RTX™ 1660Ti 6G独显/黑色', '拯救者Y7000P', 2, 1, 7899, '2020-01-05 11:32:42'),(37, 17, 4, 'ThinkPad X1 Carbon 2019 英特尔酷睿 笔记本电脑 20QD002BCD', 'i7-8565U/Windows 10家庭中文版/8GB/512GB SSD/UHD 620显示芯片/14英寸WQHD IPS LED背光显示屏', 'ThinkPad X1 Carbon', 3, 1, 12499, '2020-01-05 11:35:33'),(38, 17, 5, 'ThinkPad X1 Carbon 2019 英特尔酷睿 笔记本电脑 20QD0020CD', 'i5-8265U/Windows 10家庭中文版/8GB/512GB SSD/UHD 620显示芯片/14英寸WQHD IPS LED背光显示屏', 'ThinkPad X1 Carbon', 3, 1, 9699, '2020-01-05 11:35:35');
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `shop`.`orders` WRITE;
DELETE FROM `shop`.`orders`;
INSERT INTO `shop`.`orders` (`orid`,`uid`,`pid`,`uname`,`pname`,`peizhi`,`linkid`,`count`,`aprice`,`status`,`creatime`,`updatime`,`recvname`,`phone`,`detailed`) VALUES ('202001051214221', 1, '4 3', 'test', 'ThinkPad X1 Carbon 2019 英特尔酷睿 笔记本电脑 20QD002BCD', 'i7-8565U/Windows 10家庭中文版/8GB/512GB SSD/UHD 620显示芯片/14英寸WQHD IPS LED背光显示屏', '3', '1 2', '12499', '已取消', '2020-01-05 12:14:22', '2020-01-05 12:20:00', '老王1', '1234567892', '山东省青岛市崂山区中国海洋大学'),('202001051216141', 1, '4', 'test', 'ThinkPad X1 Carbon 2019 英特尔酷睿 笔记本电脑 20QD002BCD', 'i7-8565U/Windows 10家庭中文版/8GB/512GB SSD/UHD 620显示芯片/14英寸WQHD IPS LED背光显示屏', '3', '1', '12499', '已支付', '2020-01-05 12:16:14', '2020-01-05 12:18:29', '李晓明', '17843215678', '山东青岛崂山');
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `shop`.`product` WRITE;
DELETE FROM `shop`.`product`;
INSERT INTO `shop`.`product` (`pid`,`pname`,`pdesp`,`peizhi`,`kucun`,`price`,`type`,`kind`,`sellcount`,`status`,`linkid`,`creatime`,`updatime`) VALUES (1, '小新Pro13.3英寸英特尔酷睿全面屏超轻薄笔记本电脑', '全新10代i5，“真”全面屏', 'I5-10210U 16G 512G MX250 2.5K QHD 银', 3402, 6699, '小新Pro', 'light', 0, '1', 1, '2019-11-27 21:29:57', '2020-01-05 11:32:17'),(2, '拯救者 Y7000P 2019 英特尔酷睿i7 15.6英寸专业电竞本', '为战而生，全新拯救者助你一臂之力', 'i7-9750H/Windows 10 家庭中文版/15.6英寸/16G/1T SSD/ GeForce RTX™ 2060 6G独显/红色', 3997, 9898, '拯救者Y7000P', 'game', 0, '1', 2, '2019-12-11 20:42:12', '2020-01-05 11:35:11'),(3, '拯救者 Y7000P 2019 英特尔酷睿i5 15.6英寸专业电竞本', '为战而生，全新拯救者助你一臂之力', 'i5-9300H/Windows 10 家庭中文版/15.6英寸/8G/1T SSD/ GeForce RTX™ 1660Ti 6G独显/黑色', 3987, 7899, '拯救者Y7000P', 'game', 0, '1', 2, '2019-12-11 20:46:58', '2020-01-05 12:20:00'),(4, 'ThinkPad X1 Carbon 2019 英特尔酷睿 笔记本电脑 20QD002BCD', '超轻薄精英商务本', 'i7-8565U/Windows 10家庭中文版/8GB/512GB SSD/UHD 620显示芯片/14英寸WQHD IPS LED背光显示屏', 3983, 12499, 'ThinkPad X1 Carbon', 'ThinkPad', 0, '1', 3, '2019-12-11 20:50:31', '2020-01-05 12:20:00'),(5, 'ThinkPad X1 Carbon 2019 英特尔酷睿 笔记本电脑 20QD0020CD', '超轻薄精英商务本', 'i5-8265U/Windows 10家庭中文版/8GB/512GB SSD/UHD 620显示芯片/14英寸WQHD IPS LED背光显示屏', 3998, 9699, 'ThinkPad X1 Carbon', 'ThinkPad', 0, '1', 3, '2019-12-11 20:51:35', '2020-01-05 11:36:01');
UNLOCK TABLES;
COMMIT;
BEGIN;
LOCK TABLES `shop`.`user` WRITE;
DELETE FROM `shop`.`user`;
INSERT INTO `shop`.`user` (`uid`,`uname`,`upasswd`,`paywd`,`email`,`phone`,`yue`,`status`,`creatime`,`updatime`) VALUES (1, 'test', 'test', '123456', 'ouc@stu.ouc.edu.cn', '17685842001', 34705, '1', '2019-11-27 21:11:09', '2020-01-05 12:20:00'),(12, 'xyzxx', '12345678', '123456', '3512198443@qq.com', '17860716615', 0, '1', '2019-12-31 08:19:22', '2019-12-31 08:22:17'),(14, 'test111', 'test111', '', 'e@w.c', '12345678911', 100000, '1', '2020-01-04 21:52:51', '2020-01-04 21:52:51'),(15, 'test1112', 'test111', '', '2W@e.c', '11111111111', 100000, '1', '2020-01-04 21:54:56', '2020-01-04 21:54:56'),(16, 'test1113', 'test111', '123456', 'e@qs.c', '12345678911', 87501, '1', '2020-01-04 21:55:29', '2020-01-04 21:57:29'),(17, 'nnn111', '1234567', '123456', 'ei@q.c1', '12345678910', 87501, '1', '2020-01-05 11:31:31', '2020-01-05 12:09:28');
UNLOCK TABLES;
COMMIT;
