CREATE DATABASE IF NOT EXISTS `ton-server`;
/*
 Navicat MySQL Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80040
 Source Host           : localhost:3306
 Source Schema         : ton-server

 Target Server Type    : MySQL
 Target Server Version : 80040
 File Encoding         : 65001

 Date: 20/11/2024 21:38:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
USE `ton-server`;


-- ----------------------------
-- Table structure for table_coin_info
-- ----------------------------
DROP TABLE IF EXISTS `table_coin_info`;
CREATE TABLE `table_coin_info` (
                                   `id` bigint NOT NULL AUTO_INCREMENT,
                                   `uuid` varchar(255) NOT NULL,
                                   `detail` text NOT NULL,
                                   PRIMARY KEY (`id`),
                                   UNIQUE KEY `uuid` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for table_coin_price
-- ----------------------------
DROP TABLE IF EXISTS `table_coin_price`;
CREATE TABLE `table_coin_price` (
                                    `id` bigint NOT NULL AUTO_INCREMENT,
                                    `contract_address` varchar(255) NOT NULL,
                                    `price` varchar(255) NOT NULL,
                                    `record_time` varchar(255) NOT NULL,
                                    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                    PRIMARY KEY (`id`),
                                    KEY `address` (`contract_address`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1903 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for table_recommend_coin
-- ----------------------------
DROP TABLE IF EXISTS `table_recommend_coin`;
CREATE TABLE `table_recommend_coin` (
                                        `id` bigint NOT NULL AUTO_INCREMENT,
                                        `uuid` varchar(255) NOT NULL,
                                        `nick_name` varchar(255) DEFAULT NULL,
                                        `symbol` varchar(255) NOT NULL,
                                        `decimals` tinyint unsigned NOT NULL,
                                        `total_supply` varchar(255) NOT NULL,
                                        `contract_address` varchar(255) NOT NULL,
                                        `index` int NOT NULL,
                                        `expire_time` datetime DEFAULT NULL,
                                        `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                        `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                        PRIMARY KEY (`id`),
                                        UNIQUE KEY `uuid` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for table_task
-- ----------------------------
DROP TABLE IF EXISTS `table_task`;
CREATE TABLE `table_task` (
                              `id` bigint NOT NULL AUTO_INCREMENT,
                              `uuid` varchar(255) NOT NULL,
                              `contract_address` varchar(255) NOT NULL,
                              `rate` tinyint NOT NULL DEFAULT '1' COMMENT '1:时 2:分',
                              `active` tinyint NOT NULL DEFAULT '1' COMMENT '1：活跃，0：停止',
                              `expire_time` datetime DEFAULT NULL,
                              `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              PRIMARY KEY (`id`),
                              UNIQUE KEY `uuid` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for table_tx_history
-- ----------------------------
DROP TABLE IF EXISTS `table_tx_history`;
CREATE TABLE `table_tx_history` (
                                    `id` bigint NOT NULL AUTO_INCREMENT,
                                    `from_address` varchar(255) NOT NULL,
                                    `to_address` varchar(255) NOT NULL,
                                    `contract_address` varchar(255) NOT NULL,
                                    `amount` varchar(255) NOT NULL,
                                    `tx_id` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
                                    `tx_status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '1:交易成功, 0:交易失败, 2:交易进行中',
                                    `tx_info` text,
                                    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for table_user
-- ----------------------------
DROP TABLE IF EXISTS `table_user`;
CREATE TABLE `table_user` (
                              `id` bigint NOT NULL AUTO_INCREMENT,
                              `nickname` varchar(255) NOT NULL,
                              `address` varchar(255) NOT NULL,
                              `role` int NOT NULL DEFAULT '0' COMMENT '1:vip, 0:normal, 2:ing',
                              `stake_tx` varchar(255) DEFAULT NULL,
                              `stake_amount` varchar(255) DEFAULT NULL,
                              `utime` bigint DEFAULT '0',
                              `expire_time` datetime DEFAULT NULL,
                              `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              PRIMARY KEY (`id`),
                              UNIQUE KEY `address` (`address`),
                              KEY `role` (`role`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;

