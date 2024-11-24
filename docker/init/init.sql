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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for table_recommend_coin
-- ----------------------------
DROP TABLE IF EXISTS `table_recommend_coin`;
CREATE TABLE `table_recommend_coin` (
                                        `id` bigint NOT NULL AUTO_INCREMENT,
                                        `uuid` varchar(255) NOT NULL,
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

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
                                    `tx_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                                    `tx_status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '1:交易成功, 0:交易失败, 2:交易进行中',
                                    `tx_info` text,
                                    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

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
                              `expire_time` datetime DEFAULT NULL,
                              `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              PRIMARY KEY (`id`),
                              UNIQUE KEY `address` (`address`),
                              KEY `role` (`role`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
