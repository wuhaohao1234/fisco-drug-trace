const urlsTemplate = {
    "user": {
        "get": {
            "info": "user/me",
            "logout": "user/logout",
        },
        "post": {
            "register": "user/register",
            "login": "user/login",
            "list": "users",
            "updateMe": "user/update/me",
            "update": "user/update",
            "dealer": "user/dealer",
        }
    },
    "drug": {
        "post": {
            "add": "drug/add",
            "list": "drug/list",
            "sale": "drug/sale",
            "dealer": "drug/dealer",
            "accept": "drug/accept",
            "buy": "drug/buy",
            "buyList": "drug/buy/list",
        }
    }

};

基于以上mock 接口
/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80033 (8.0.33)
 Source Host           : localhost:3306
 Source Schema         : fisco_drug_trace

 Target Server Type    : MySQL
 Target Server Version : 80033 (8.0.33)
 File Encoding         : 65001

 Date: 13/04/2024 16:54:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `username` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `password_digest` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `nickname` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `status` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `role` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `register_time` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '2024-04-12 16:39:07.143', '2024-04-12 17:14:40.032', NULL, '123456', '$2a$12$GOgqa/y1ECAfYcfHjYoe8eto9MhHhc5JkGyU0x5TmIOhzmdf4pMhq', '123231', 'enabled', 'admin', '');
INSERT INTO `users` VALUES (2, '2024-04-12 17:08:59.986', '2024-04-12 17:09:31.490', NULL, '987', '$2a$12$ukyRAinNHfwtDOjXNbsR9uG0favWGggk51/MXXvLJa4FzR06pVWx2', '1239', 'enabled', 'dealer', '');
INSERT INTO `users` VALUES (3, '2024-04-12 19:03:55.582', '2024-04-12 20:27:15.868', NULL, '123', '$2a$12$HtMtHweHQzmM/Tl7b0xSdO.mzfgw/eqgh9Wq17yo/NHTjON/4gBgu', '1233243', 'enabled', 'merchant', '');
INSERT INTO `users` VALUES (4, '2024-04-12 20:09:00.770', '2024-04-12 20:09:14.944', NULL, '321', '$2a$12$YnBLuDQJ6pHyAbZlAnTJa.FTS1CgAaTedss/Q.r/ogz3YS6W9B2xS', '321', 'enabled', 'user', '');
INSERT INTO `users` VALUES (5, '2024-04-12 20:46:39.742', '2024-04-12 20:46:54.073', NULL, '456', '$2a$12$w79s27bCasZwnlo38wjPr.b32UfP5AMza208YpC5o70mUjJyQklJi', '456', 'enabled', 'dealer', '');

SET FOREIGN_KEY_CHECKS = 1;
