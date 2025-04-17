/*
 Navicat Premium Data Transfer

 Source Server         : docker-mysql-master
 Source Server Type    : MySQL
 Source Server Version : 80032
 Source Host           : 192.168.122.247:3306
 Source Schema         : game_platform

 Target Server Type    : MySQL
 Target Server Version : 80032
 File Encoding         : 65001

 Date: 17/04/2025 14:39:27
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for apps
-- ----------------------------
DROP TABLE IF EXISTS `apps`;
CREATE TABLE `apps`  (
  `id` mediumint UNSIGNED NOT NULL AUTO_INCREMENT,
  `app_name` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `game_class` tinyint(1) NULL DEFAULT 1 COMMENT '游戏类别',
  `company_id` int NOT NULL COMMENT '研发公司id',
  `status` tinyint NOT NULL DEFAULT 0 COMMENT '0-上线 1-下架',
  `created_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `updated_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `g`(`app_name` ASC) USING BTREE,
  INDEX `game_class`(`game_class` ASC) USING BTREE,
  INDEX `deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '游戏资料表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for channels
-- ----------------------------
DROP TABLE IF EXISTS `channels`;
CREATE TABLE `channels`  (
  `id` mediumint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `code` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '标识',
  `doc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '文档连接',
  `auth_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '授权地址',
  `params` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '授权固定参数',
  `redirect_uri` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '授权回调地址',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `opt_user` int NOT NULL DEFAULT 0 COMMENT '操作人',
  `created_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `updated_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted_at` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '媒体渠道' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for game_notice
-- ----------------------------
DROP TABLE IF EXISTS `game_notice`;
CREATE TABLE `game_notice`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `notice_type` int NOT NULL DEFAULT 1 COMMENT '状态: 0 正常  1流畅 ',
  `game_id` mediumint UNSIGNED NOT NULL DEFAULT 0,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `start_time` int UNSIGNED NOT NULL DEFAULT 0,
  `end_time` int UNSIGNED NOT NULL DEFAULT 0,
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `last_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `account` char(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '管理员账号',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态: 0 正常  1流畅 ',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for game_server
-- ----------------------------
DROP TABLE IF EXISTS `game_server`;
CREATE TABLE `game_server`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `game_id` int NOT NULL DEFAULT 0 COMMENT '游戏id',
  `zone_id` varchar(45) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '0' COMMENT '区服id',
  `zone_name` varchar(45) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '区服名称',
  `create_time` int UNSIGNED NOT NULL DEFAULT 0,
  `create_user` int UNSIGNED NOT NULL DEFAULT 0,
  `update_time` int UNSIGNED NOT NULL DEFAULT 11,
  `update_user` int UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `game_id_zone_id`(`zone_id` ASC, `game_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 674 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '游戏区服' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for games
-- ----------------------------
DROP TABLE IF EXISTS `games`;
CREATE TABLE `games`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `app_id` int NOT NULL DEFAULT 0 COMMENT '应用ID',
  `pkg_name` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '包名全局唯一',
  `game_name` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '游戏名',
  `cp_callback_url` varchar(300) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT 'CP 发货正式接口',
  `cp_test_callback_url` varchar(300) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT 'CP 发货测试接口',
  `app_key` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '发货key',
  `server_key` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '服务端key',
  `os` tinyint(1) NOT NULL DEFAULT 1 COMMENT '应用类型：1-Android 2-IOS 3-H5 4-小程序',
  `link_h5` varchar(300) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT 'H5 链接 ',
  `download_url` varchar(500) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '游戏下载地址',
  `status` tinyint(1) NOT NULL DEFAULT -1 COMMENT '状态: 对接中(0)、已上线 (1) 、已下线(2)',
  `conversion` float NOT NULL DEFAULT 0 COMMENT '人民币和游戏币转换倍率，人民币是 1',
  `icon` varchar(400) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT 'icon',
  `remark` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '备注',
  `publish_at` int NOT NULL COMMENT '发布时间',
  `is_auth_real_name` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否需要实名认证  0-是 1-否',
  `is_limit_underage` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否限制未成年 0-是 1-否',
  `signature` tinyint(1) NOT NULL DEFAULT 0 COMMENT '签名方式 0-md5',
  `created_at` int NOT NULL DEFAULT 0 COMMENT '添加时间',
  `updated_at` int NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted_at` int NOT NULL DEFAULT 0,
  `opt_user` int NOT NULL DEFAULT 0 COMMENT '操作人',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `app_pkg_name`(`pkg_name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '游戏应用包，主要用于分发' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for infra_config
-- ----------------------------
DROP TABLE IF EXISTS `infra_config`;
CREATE TABLE `infra_config`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `group` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '参数分组',
  `type` tinyint NOT NULL COMMENT '参数类型',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '参数名称',
  `key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '参数键名',
  `value` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '参数键值',
  `sensitive` bit(1) NOT NULL COMMENT '是否敏感',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '参数配置表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ios_blur
-- ----------------------------
DROP TABLE IF EXISTS `ios_blur`;
CREATE TABLE `ios_blur`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `game_id` int NOT NULL DEFAULT 0,
  `url_list` json NULL,
  `params_map` json NULL,
  `ts_url` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `function_map` json NULL,
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '状态，2=禁用。1=启用',
  `gs_url` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `game_id_blur` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
  `encrypt_type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '1 AES 2 DES 3 XOR',
  `promotion_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '推广url',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `only_key`(`game_id` ASC) USING BTREE,
  INDEX `url`(`ts_url` ASC) USING BTREE,
  INDEX `game_id_blur`(`game_id_blur` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 202 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'ios混淆表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ios_shm_url
-- ----------------------------
DROP TABLE IF EXISTS `ios_shm_url`;
CREATE TABLE `ios_shm_url`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `url` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '提审域名',
  `game_id` int NULL DEFAULT 0 COMMENT '游戏id',
  `promotion_url` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '推广url',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `url`(`url` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '提审域名与ios渠道包' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for jobs
-- ----------------------------
DROP TABLE IF EXISTS `jobs`;
CREATE TABLE `jobs`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `queue` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `payload` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `attempts` tinyint UNSIGNED NOT NULL,
  `reserved_at` int UNSIGNED NULL DEFAULT NULL,
  `available_at` int UNSIGNED NOT NULL,
  `created_at` int UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `jobs_queue_index`(`queue` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for jobs_failed
-- ----------------------------
DROP TABLE IF EXISTS `jobs_failed`;
CREATE TABLE `jobs_failed`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `connection` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `queue` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `payload` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `exception` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `failed_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for market_app_id
-- ----------------------------
DROP TABLE IF EXISTS `market_app_id`;
CREATE TABLE `market_app_id`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '名字',
  `channel_id` int NOT NULL DEFAULT 0 COMMENT '媒体渠道',
  `account` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '开发者账号',
  `app_id` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT 'app_id',
  `secret` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '秘钥',
  `butler_id` varchar(250) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '授权账户管家ID',
  `state` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '标识',
  `status` tinyint NOT NULL DEFAULT 0 COMMENT '是否可用 0 否 1是',
  `params` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '授权参数配置',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '备注',
  `opt_user` int NOT NULL DEFAULT 0 COMMENT '操作人',
  `created_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `updated_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `channel_app_id`(`channel_id` ASC, `app_id` ASC) USING BTREE,
  INDEX `state`(`state` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '媒体渠道APPID授权管理表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for market_principals
-- ----------------------------
DROP TABLE IF EXISTS `market_principals`;
CREATE TABLE `market_principals`  (
  `id` mediumint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `code` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '标识',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `opt_user` int NOT NULL DEFAULT 0 COMMENT '操作人',
  `created_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `updated_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted_at` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '开户主体' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for market_proxy_account
-- ----------------------------
DROP TABLE IF EXISTS `market_proxy_account`;
CREATE TABLE `market_proxy_account`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `project_id` int NOT NULL DEFAULT 0 COMMENT '推广id',
  `name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '广告主名称',
  `short_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '账户简称',
  `uid` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '账号id',
  `account` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '账号名',
  `password` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '账号密码',
  `owner` int NOT NULL DEFAULT 0 COMMENT '使用人',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '账号状态，0-正常 1-停止使用',
  `oauth_type` tinyint NOT NULL DEFAULT 1 COMMENT '授权类型,1=管家,2=单账户',
  `oauth_status` tinyint NOT NULL DEFAULT 1 COMMENT '授权状态,1不可使用,2可使用',
  `oauth_subject` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '授权主体标识',
  `lot` char(12) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '0' COMMENT '上传批次',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `opt_user` int NOT NULL DEFAULT 0 COMMENT '操作人',
  `created_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `updated_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `rowunique`(`uid` ASC, `deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11752 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '广告账号' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for market_proxy_company
-- ----------------------------
DROP TABLE IF EXISTS `market_proxy_company`;
CREATE TABLE `market_proxy_company`  (
  `id` mediumint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `short_name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '简称',
  `code` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '标识',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `opt_user` int NOT NULL DEFAULT 0 COMMENT '操作人',
  `created_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `updated_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted_at` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '代理商公司' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for market_proxy_project
-- ----------------------------
DROP TABLE IF EXISTS `market_proxy_project`;
CREATE TABLE `market_proxy_project`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `short_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '简称',
  `channel_id` int NOT NULL DEFAULT 0 COMMENT '媒体渠道',
  `principal_id` int NOT NULL DEFAULT 0 COMMENT '开户主体id',
  `proxy_company_id` int NOT NULL DEFAULT 0 COMMENT '代理商主体id',
  `rebate` varchar(10) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '媒体返点',
  `contract_start_time` int NOT NULL DEFAULT 0 COMMENT '合同签订时间',
  `contract_end_time` int NOT NULL DEFAULT 0 COMMENT '合同终止时间',
  `contract_status` tinyint NOT NULL DEFAULT 0 COMMENT '合同进度,1商务初审,2法务二审,3审核完成,4已归档',
  `contract_url` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '合同下载地址',
  `remark` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '备注',
  `opt_user` int NOT NULL DEFAULT 0 COMMENT '操作人',
  `created_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `updated_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted_at` int NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '合作代理项目' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for market_proxy_rebate
-- ----------------------------
DROP TABLE IF EXISTS `market_proxy_rebate`;
CREATE TABLE `market_proxy_rebate`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `proxy_project_id` int NOT NULL DEFAULT 0 COMMENT '代理计划id',
  `start_time` int NOT NULL DEFAULT 0 COMMENT '开始时间',
  `end_time` int NOT NULL DEFAULT 0 COMMENT '结束时间',
  `rate` float(11, 2) NOT NULL DEFAULT 0.00 COMMENT '返点',
  `opt_user` int NOT NULL DEFAULT 0 COMMENT '操作人',
  `created_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `updated_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '代理项目不同时间段的返点配置' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for member
-- ----------------------------
DROP TABLE IF EXISTS `member`;
CREATE TABLE `member`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `promote_code` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '推广码',
  `account_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '帐号类型：0  游客; 1  Email ; 2  FaceBook; 3  Google;4 Apple',
  `account` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '帐户名',
  `password` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '登陆密码',
  `salt` char(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '盐值',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '-1:未激活，0：正常，1：禁用',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `pu`(`account` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 699931 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for member_game_role
-- ----------------------------
DROP TABLE IF EXISTS `member_game_role`;
CREATE TABLE `member_game_role`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `promote_code` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '推广码',
  `os` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '操作系统：android、ios',
  `device_id` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '设备号',
  `ip` char(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT 'IP ',
  `area_code` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区码',
  `area` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区',
  `user_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `account` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL DEFAULT '' COMMENT '帐户名',
  `app_channel` int UNSIGNED NOT NULL DEFAULT 1 COMMENT '应用渠道：1(官方)、2(安卓)、3(苹果)',
  `game_project_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏项目ID',
  `game_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `game_name` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '游戏名',
  `zone_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '区服',
  `zone_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '区服名',
  `game_role_id` bigint NOT NULL DEFAULT 0,
  `game_role_name` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '角色名',
  `game_role_level` int NOT NULL DEFAULT 0 COMMENT '角色级别',
  `online` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '在线时长单位为秒',
  `auth_status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '认证状态：0(未认证)、1(已认证)',
  `auth_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '认证时间',
  `last_ip` char(158) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '最后登陆IP',
  `last_pay_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后充值时间',
  `last_login_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后登陆时间',
  `total_pay_money` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付总金额',
  `create_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '事件时间，即创建角色时间',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `ugzg`(`user_id` ASC, `game_id` ASC, `zone_id` ASC, `game_role_id` ASC) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `game_role_id`(`game_role_id` ASC) USING BTREE,
  INDEX `game_id`(`game_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 952350 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for member_profile
-- ----------------------------
DROP TABLE IF EXISTS `member_profile`;
CREATE TABLE `member_profile`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `nickname` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `mobile` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '绑定的手机号',
  `trade_password` char(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '安全码',
  `balance` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '余额',
  `full_name` varchar(12) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `reg_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '注册时间',
  `reg_ip` char(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '注册IP',
  `last_login_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后登陆时间',
  `last_login_ip` char(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '最后登陆IP',
  `last_login_way` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后登陆方式： 0(visitor)、1(email)、 2(facebook)、3(google)、4(apple)',
  `area_code` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '注册地区码',
  `area` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '注册地区',
  `game_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `wechat` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `email` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '用户邮箱',
  `sex` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '2:未知 1:男 0: 女',
  `avatar` varchar(400) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `vip` smallint UNSIGNED NOT NULL DEFAULT 1 COMMENT '会员VIP等级',
  `device_id` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '设备号',
  `sys_model` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '机型',
  `login_times` mediumint UNSIGNED NOT NULL DEFAULT 0 COMMENT '登陆次数',
  `remark` varchar(800) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` int NOT NULL DEFAULT 0 COMMENT '添加时间',
  `updated_at` int NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `last_ip`(`last_login_ip` ASC) USING BTREE,
  INDEX `last_time`(`last_login_time` ASC) USING BTREE,
  INDEX `reg_time`(`reg_time` ASC) USING BTREE,
  INDEX `reg_ip`(`reg_ip` ASC) USING BTREE,
  INDEX `mobile`(`mobile` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 22353 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '用户扩展属性表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for member_third
-- ----------------------------
DROP TABLE IF EXISTS `member_third`;
CREATE TABLE `member_third`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '帐号id：member_id or user_id',
  `socialite` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '应用渠道即带有SDK的渠道如 Google、FaceBook',
  `openid` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '三方开放的帐号ID 如 vivo、wx 等',
  `unionid` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '联合ID',
  `nickname` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `email` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '用户邮箱',
  `email_verified` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT 'email 是存验证过',
  `avatar` varchar(800) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '头像',
  `last_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `ao`(`socialite` ASC, `openid` ASC) USING BTREE,
  INDEX `add_time`(`add_time` ASC) USING BTREE,
  INDEX `last_time`(`last_time` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 46 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '三方用户' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for migrations
-- ----------------------------
DROP TABLE IF EXISTS `migrations`;
CREATE TABLE `migrations`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `batch` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for notifications
-- ----------------------------
DROP TABLE IF EXISTS `notifications`;
CREATE TABLE `notifications`  (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `notifiable_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `notifiable_id` bigint UNSIGNED NOT NULL,
  `data` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `read_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `notifications_notifiable_type_notifiable_id_index`(`notifiable_type` ASC, `notifiable_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for packages
-- ----------------------------
DROP TABLE IF EXISTS `packages`;
CREATE TABLE `packages`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '渠道包名称',
  `package_name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '渠道应用名称',
  `app_id` int NOT NULL DEFAULT 0 COMMENT '应用ID',
  `game_id` int NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `channel_id` int NOT NULL DEFAULT 0 COMMENT '渠道ID',
  `campaign_id` int NOT NULL DEFAULT 0 COMMENT '自然量广告ID',
  `original_package_id` int NOT NULL DEFAULT 0 COMMENT '母包ID',
  `sdk_id` int NOT NULL DEFAULT 0 COMMENT 'sdk_ID',
  `skin_id` int NOT NULL DEFAULT 0 COMMENT '换皮配置ID',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '渠道包状态 0-正常  1-下架',
  `pack_status` tinyint NOT NULL DEFAULT 0 COMMENT '打包状态',
  `last_package_time` int NOT NULL DEFAULT 0 COMMENT '最后打包时间',
  `cp_callback_url` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'cp 发货地址',
  `cp_callback_test_url` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'cp 测试发货地址',
  `is_change_pay` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否切支付 0-否 1-是',
  `is_sdk_float_on` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'SDK浮点0隐藏，1显示',
  `is_user_float_on` tinyint(1) NOT NULL DEFAULT 0 COMMENT '用户浮点，0隐藏，1显示',
  `is_reg_login_on` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否预下载 0关闭，1正常',
  `is_visitor_on` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否开启游客模式，0不开，1开',
  `is_auto_login_on` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否开启自动登录，0关闭，1开启',
  `is_log_on` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否开启错误日志上报，0不开，1开',
  `is_shm` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否切换审核服  0不是，1是',
  `switch_login` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否切渠道登录  0不切，1切',
  `online_time` int NOT NULL DEFAULT 0 COMMENT '上线时间',
  `scheme` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'urlscheme,支付后跳转url',
  `privacy` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '隐私协议',
  `rate` decimal(9, 4) NOT NULL DEFAULT 10.0000 COMMENT '游戏币比例1rmb=10yxb',
  `id_card_verify` tinyint(1) NOT NULL DEFAULT 0 COMMENT '实名认证弹窗( 0关闭 1开启可关闭 2开启不可关闭， 3绝对不开)',
  `is_limit_minor` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否限制未成年人游戏和消费 0否 1是 2 绝对不开',
  `deny_reg` tinyint(1) NOT NULL DEFAULT 0 COMMENT '禁止注册新用户，0否1是',
  `remarks` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `opt_user` int NOT NULL DEFAULT 0 COMMENT '操作人',
  `created_at` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted_at` int NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '媒体渠道包' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_area
-- ----------------------------
DROP TABLE IF EXISTS `pay_area`;
CREATE TABLE `pay_area`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '国家名',
  `name_en` varchar(80) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `area_code` varchar(12) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '国家code',
  `pay_currency_id` int NOT NULL DEFAULT 0 COMMENT '国家本土货币',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态，1-下拉开启，0-下拉关闭',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `code`(`area_code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1000 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '支付地域名' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for pay_currency
-- ----------------------------
DROP TABLE IF EXISTS `pay_currency`;
CREATE TABLE `pay_currency`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `symbol` varchar(5) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '货币符号',
  `exchange_rate` decimal(10, 4) UNSIGNED NOT NULL DEFAULT 0.0000 COMMENT '相对美元的汇率',
  `code` varchar(10) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '货币代码',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `code`(`code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '以美元兑换汇率' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for pay_order_refund
-- ----------------------------
DROP TABLE IF EXISTS `pay_order_refund`;
CREATE TABLE `pay_order_refund`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `game_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `pay_channel` varchar(20) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '支付通道对应表pay_channel的 channel_code',
  `trade_id` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '',
  `purchase_time` datetime NOT NULL COMMENT '购买时间',
  `purchase_token` varchar(1280) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '购买Token',
  `refund_time` datetime NOT NULL COMMENT '退款时间',
  `refund_reason` varchar(2048) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '退款原因',
  `refund_source` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL DEFAULT '' COMMENT '退款来源',
  `last_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `pt`(`pay_channel` ASC, `trade_id` ASC) USING BTREE,
  INDEX `purchase_time`(`purchase_time` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '退款表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for pay_orders
-- ----------------------------
DROP TABLE IF EXISTS `pay_orders`;
CREATE TABLE `pay_orders`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `merchant_id` smallint UNSIGNED NOT NULL DEFAULT 1 COMMENT '支付商户ID',
  `pay_channel` varchar(20) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '支付通道对应表pay_channel的 channel_code',
  `user_id` bigint NOT NULL DEFAULT 0 COMMENT '用户ID',
  `account` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL DEFAULT '' COMMENT '帐户名',
  `promote_code` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '推广码',
  `openid` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '三方开放的帐号ID 如 vivo、wx 等',
  `subject` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '物品描述',
  `goods_id` varchar(120) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '对就商品goods_id',
  `goods_currency` varchar(45) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT 'USD' COMMENT '支付货币代码:USD(美元)、HKG(港元) 、MAC(澳门元) 、TWD(新台币) ...',
  `order_type` tinyint(1) NULL DEFAULT 1 COMMENT '订单类型：直购(1)、应用购(2)、充余额(10)、提现(20)',
  `order_num` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '订单号',
  `order_price` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '定价,单位分',
  `order_time` int UNSIGNED NOT NULL DEFAULT 0,
  `pay_product_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付产品ID',
  `pay_money` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '实付金额,单位分',
  `pay_currency` varchar(45) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT 'USD' COMMENT '支付货币代码:USD(美元)、HKG(港元) 、MAC(澳门元) 、TWD(新台币) ...',
  `pay_status` tinyint(1) NULL DEFAULT -1,
  `pay_time` int UNSIGNED NOT NULL DEFAULT 0,
  `trade_id` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '',
  `trade_account` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '',
  `trade_status` tinyint(1) NOT NULL DEFAULT 0,
  `trade_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '三方交易时间',
  `trade_data` varchar(2048) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '交易数据',
  `trade_ticket_md5` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '交易凭证Md5',
  `cp_order_num` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT 'CP 订单号',
  `cp_order_status` tinyint(1) NULL DEFAULT -1 COMMENT '发货状态： -1 未发 0 失败 1 成功',
  `cp_order_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '发货时间',
  `app_channel` int UNSIGNED NOT NULL DEFAULT 1 COMMENT '应用渠道',
  `game_project_id` mediumint NOT NULL DEFAULT 0 COMMENT '游戏项目ID',
  `game_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `game_name` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '游戏名',
  `zone_id` varchar(50) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '0' COMMENT '区服',
  `zone_name` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '区服名',
  `game_role_id` varchar(50) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '0',
  `game_role_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名',
  `game_role_level` varchar(50) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '0' COMMENT '角色级别',
  `pkg_name` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '包名',
  `os` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '操作系统：android、ios',
  `device_id` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '设备号',
  `ip` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '',
  `area` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区',
  `sandbox` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '环境: 测试(1)、生产(0)',
  `remarks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `order_num`(`order_num` ASC) USING BTREE,
  INDEX `member_id`(`user_id` ASC) USING BTREE,
  INDEX `order_time`(`order_time` ASC) USING BTREE,
  INDEX `pay_channel`(`pay_channel` ASC) USING BTREE,
  INDEX `trade_ticket_md5`(`trade_ticket_md5` ASC) USING BTREE,
  INDEX `trade_id`(`trade_id` ASC) USING BTREE,
  INDEX `game_role_id`(`game_role_id` ASC) USING BTREE,
  INDEX `game_id`(`game_id` ASC) USING BTREE,
  INDEX `pay_time`(`pay_time` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '订单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for pay_product
-- ----------------------------
DROP TABLE IF EXISTS `pay_product`;
CREATE TABLE `pay_product`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `channel_id` bigint NOT NULL DEFAULT 102 COMMENT '渠道ID',
  `code` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '支付编号：扫码、App、Wap、小程序等',
  `display_name` varchar(45) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `icon` varchar(400) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '',
  `fee_rate` double NOT NULL DEFAULT 0 COMMENT '渠道费率，单位：百分比',
  `sorting` int NOT NULL DEFAULT 1 COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态: 关闭(0)、正常(1)',
  `config` varchar(4096) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '配置',
  `remark` varchar(800) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `last_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `tc`(`code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '支付产品' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for report_daily
-- ----------------------------
DROP TABLE IF EXISTS `report_daily`;
CREATE TABLE `report_daily`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `days` int UNSIGNED NOT NULL DEFAULT 0,
  `promote_code` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '广告码',
  `app_channel` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '应用渠道',
  `agent_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '代理ID,即主播',
  `media` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '广告媒体如：oe、tc、ks',
  `advertiser_id` bigint NULL DEFAULT 0 COMMENT '广告主ID',
  `game_id` int NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `pv` int UNSIGNED NOT NULL DEFAULT 0,
  `uv` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '独立IP',
  `click` int UNSIGNED NOT NULL DEFAULT 0,
  `active` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '激活即启动数',
  `dnu` int UNSIGNED NOT NULL DEFAULT 0,
  `dau` int UNSIGNED NOT NULL DEFAULT 0,
  `players` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏注册数：一个dnu对应N个players ',
  `enter` int UNSIGNED NOT NULL DEFAULT 0,
  `pay_money` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '付费金额',
  `fee` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '手续费',
  `pay_times` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '付费人次',
  `pay_numbers` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '付费人数',
  `ad_money` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '广告费，即消耗',
  `lt` int UNSIGNED NOT NULL DEFAULT 0 COMMENT 'LT生命周期(LT:Life Time)：一个用户从第1次到最后1次参与游戏之间的时间段，一般按月计算平均值',
  `keep1` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '次留',
  `keep2` int UNSIGNED NOT NULL DEFAULT 0,
  `keep3` int UNSIGNED NOT NULL DEFAULT 0,
  `keep7` int UNSIGNED NOT NULL DEFAULT 0,
  `keep15` int UNSIGNED NOT NULL DEFAULT 0,
  `keep30` int UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `dtaaa`(`days` ASC, `app_channel` ASC, `game_id` ASC, `advertiser_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '广告日报' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for report_daily_device_retain
-- ----------------------------
DROP TABLE IF EXISTS `report_daily_device_retain`;
CREATE TABLE `report_daily_device_retain`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `date` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '日期',
  `game_project_id` int NOT NULL DEFAULT 0 COMMENT '项目',
  `game_id` int NOT NULL DEFAULT 0 COMMENT '游戏id',
  `app_channel` tinyint(1) NOT NULL DEFAULT 1 COMMENT '应用渠道',
  `area_code` varchar(10) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '地区',
  `retain_1` int NOT NULL DEFAULT 0 COMMENT '第1日留存',
  `retain_2` int NOT NULL DEFAULT 0 COMMENT '第2日留存',
  `retain_3` int NOT NULL DEFAULT 0 COMMENT '第3日留存',
  `retain_4` int NOT NULL DEFAULT 0 COMMENT '第4日留存',
  `retain_5` int NOT NULL DEFAULT 0 COMMENT '第5日留存',
  `retain_6` int NOT NULL DEFAULT 0 COMMENT '第6日留存',
  `retain_7` int NOT NULL DEFAULT 0 COMMENT '第7日留存',
  `retain_8` int NOT NULL DEFAULT 0 COMMENT '第8日留存',
  `retain_9` int NOT NULL DEFAULT 0 COMMENT '第9日留存',
  `retain_10` int NOT NULL DEFAULT 0 COMMENT '第10日留存',
  `retain_11` int NOT NULL DEFAULT 0 COMMENT '第11日留存',
  `retain_12` int NOT NULL DEFAULT 0 COMMENT '第12日留存',
  `retain_13` int NOT NULL DEFAULT 0 COMMENT '第13日留存',
  `retain_14` int NOT NULL DEFAULT 0 COMMENT '第14日留存',
  `retain_15` int NOT NULL DEFAULT 0 COMMENT '第15日留存',
  `retain_16` int NOT NULL DEFAULT 0 COMMENT '第16日留存',
  `retain_17` int NOT NULL DEFAULT 0 COMMENT '第17日留存',
  `retain_18` int NOT NULL DEFAULT 0 COMMENT '第18日留存',
  `retain_19` int NOT NULL DEFAULT 0 COMMENT '第19日留存',
  `retain_20` int NOT NULL DEFAULT 0 COMMENT '第20日留存',
  `retain_21` int NOT NULL DEFAULT 0 COMMENT '第21日留存',
  `retain_22` int NOT NULL DEFAULT 0 COMMENT '第22日留存',
  `retain_23` int NOT NULL DEFAULT 0 COMMENT '第23日留存',
  `retain_24` int NOT NULL DEFAULT 0 COMMENT '第24日留存',
  `retain_25` int NOT NULL DEFAULT 0 COMMENT '第25日留存',
  `retain_26` int NOT NULL DEFAULT 0 COMMENT '第26日留存',
  `retain_27` int NOT NULL DEFAULT 0 COMMENT '第27日留存',
  `retain_28` int NOT NULL DEFAULT 0 COMMENT '第28日留存',
  `retain_29` int NOT NULL DEFAULT 0 COMMENT '第29日留存',
  `retain_30` int NOT NULL DEFAULT 0 COMMENT '第30日留存',
  `retain_60` int NOT NULL DEFAULT 0 COMMENT '第60日留存',
  `retain_90` int NOT NULL DEFAULT 0 COMMENT '第90日留存',
  `update_time` bigint NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `days_game_id_channel_area`(`date` ASC, `game_project_id` ASC, `game_id` ASC, `app_channel` ASC, `area_code` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '设备留存统计表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for report_daily_hour
-- ----------------------------
DROP TABLE IF EXISTS `report_daily_hour`;
CREATE TABLE `report_daily_hour`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `date` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '日期 20230101 ',
  `hour` char(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '小时',
  `game_project_id` int NOT NULL DEFAULT 0 COMMENT '项目',
  `game_id` int NOT NULL COMMENT '游戏',
  `app_channel` tinyint(1) NOT NULL COMMENT '应用渠道：1(官方)、2(安卓)、3(苹果)',
  `area_code` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地区',
  `dau` int NOT NULL DEFAULT 0 COMMENT '去重登录数  \r\n从0点到该点累计去重登陆总数',
  `register` int NOT NULL DEFAULT 0 COMMENT '新账号注册数',
  `device` int NOT NULL DEFAULT 0 COMMENT '新设备数',
  `pay_money` decimal(11, 2) NOT NULL DEFAULT 0.00 COMMENT '总付费数',
  `pay_people` int NOT NULL DEFAULT 0 COMMENT '总付费人数',
  `pay_times` int NOT NULL DEFAULT 0 COMMENT '付费次数',
  `new_pay_money` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '新增账号充值数',
  `new_pay_people` int NOT NULL DEFAULT 0 COMMENT '新增账号人数',
  `new_pay_times` int NOT NULL DEFAULT 0 COMMENT '新增账号次数',
  `update_time` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `date_hour_game_channel`(`date` ASC, `hour` ASC, `game_project_id` ASC, `game_id` ASC, `app_channel` ASC, `area_code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 30415 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '每小时数据统计表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for report_daily_ip_retain
-- ----------------------------
DROP TABLE IF EXISTS `report_daily_ip_retain`;
CREATE TABLE `report_daily_ip_retain`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `date` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '日期',
  `game_project_id` int NOT NULL DEFAULT 0 COMMENT '项目',
  `game_id` int NOT NULL DEFAULT 0 COMMENT '游戏id',
  `app_channel` tinyint(1) NOT NULL DEFAULT 1 COMMENT '应用渠道',
  `area_code` varchar(10) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '地区',
  `retain_1` int NOT NULL DEFAULT 0 COMMENT '第1日留存',
  `retain_2` int NOT NULL DEFAULT 0 COMMENT '第2日留存',
  `retain_3` int NOT NULL DEFAULT 0 COMMENT '第3日留存',
  `retain_4` int NOT NULL DEFAULT 0 COMMENT '第4日留存',
  `retain_5` int NOT NULL DEFAULT 0 COMMENT '第5日留存',
  `retain_6` int NOT NULL DEFAULT 0 COMMENT '第6日留存',
  `retain_7` int NOT NULL DEFAULT 0 COMMENT '第7日留存',
  `retain_8` int NOT NULL DEFAULT 0 COMMENT '第8日留存',
  `retain_9` int NOT NULL DEFAULT 0 COMMENT '第9日留存',
  `retain_10` int NOT NULL DEFAULT 0 COMMENT '第10日留存',
  `retain_11` int NOT NULL DEFAULT 0 COMMENT '第11日留存',
  `retain_12` int NOT NULL DEFAULT 0 COMMENT '第12日留存',
  `retain_13` int NOT NULL DEFAULT 0 COMMENT '第13日留存',
  `retain_14` int NOT NULL DEFAULT 0 COMMENT '第14日留存',
  `retain_15` int NOT NULL DEFAULT 0 COMMENT '第15日留存',
  `retain_16` int NOT NULL DEFAULT 0 COMMENT '第16日留存',
  `retain_17` int NOT NULL DEFAULT 0 COMMENT '第17日留存',
  `retain_18` int NOT NULL DEFAULT 0 COMMENT '第18日留存',
  `retain_19` int NOT NULL DEFAULT 0 COMMENT '第19日留存',
  `retain_20` int NOT NULL DEFAULT 0 COMMENT '第20日留存',
  `retain_21` int NOT NULL DEFAULT 0 COMMENT '第21日留存',
  `retain_22` int NOT NULL DEFAULT 0 COMMENT '第22日留存',
  `retain_23` int NOT NULL DEFAULT 0 COMMENT '第23日留存',
  `retain_24` int NOT NULL DEFAULT 0 COMMENT '第24日留存',
  `retain_25` int NOT NULL DEFAULT 0 COMMENT '第25日留存',
  `retain_26` int NOT NULL DEFAULT 0 COMMENT '第26日留存',
  `retain_27` int NOT NULL DEFAULT 0 COMMENT '第27日留存',
  `retain_28` int NOT NULL DEFAULT 0 COMMENT '第28日留存',
  `retain_29` int NOT NULL DEFAULT 0 COMMENT '第29日留存',
  `retain_30` int NOT NULL DEFAULT 0 COMMENT '第30日留存',
  `retain_60` int NOT NULL DEFAULT 0 COMMENT '第60日留存',
  `retain_90` int NOT NULL DEFAULT 0 COMMENT '第90日留存',
  `update_time` bigint NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `days_game_id_channel_area`(`date` ASC, `game_project_id` ASC, `game_id` ASC, `app_channel` ASC, `area_code` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = 'ip留存统计表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for report_daily_ltv
-- ----------------------------
DROP TABLE IF EXISTS `report_daily_ltv`;
CREATE TABLE `report_daily_ltv`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `date` int NOT NULL DEFAULT 0 COMMENT '日期',
  `game_project_id` int NOT NULL DEFAULT 0 COMMENT '项目',
  `game_id` int NOT NULL DEFAULT 0 COMMENT '游戏id',
  `app_channel` tinyint(1) NOT NULL DEFAULT 1 COMMENT '应用渠道',
  `area_code` varchar(10) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '地区',
  `register` int NOT NULL COMMENT '注册数',
  `all_money` int NOT NULL DEFAULT 0 COMMENT '最新充值总数，包括90~120之间的数据',
  `money_1` int NOT NULL DEFAULT 0 COMMENT '第1天充值总数',
  `money_2` int NOT NULL DEFAULT 0 COMMENT '第2天充值总数',
  `money_3` int NOT NULL DEFAULT 0 COMMENT '第3天充值总数',
  `money_4` int NOT NULL DEFAULT 0 COMMENT '第4天充值总数',
  `money_5` int NOT NULL DEFAULT 0 COMMENT '第5天充值总数',
  `money_6` int NOT NULL DEFAULT 0 COMMENT '第6天充值总数',
  `money_7` int NOT NULL DEFAULT 0 COMMENT '第7天充值总数',
  `money_8` int NOT NULL DEFAULT 0 COMMENT '第8天充值总数',
  `money_9` int NOT NULL DEFAULT 0 COMMENT '第9天充值总数',
  `money_10` int NOT NULL DEFAULT 0 COMMENT '第10天充值总数',
  `money_11` int NOT NULL DEFAULT 0 COMMENT '第11天充值总数',
  `money_12` int NOT NULL DEFAULT 0 COMMENT '第12天充值总数',
  `money_13` int NOT NULL DEFAULT 0 COMMENT '第13天充值总数',
  `money_14` int NOT NULL DEFAULT 0 COMMENT '第14天充值总数',
  `money_15` int NOT NULL DEFAULT 0 COMMENT '第15天充值总数',
  `money_16` int NOT NULL DEFAULT 0 COMMENT '第16天充值总数',
  `money_17` int NOT NULL DEFAULT 0 COMMENT '第17天充值总数',
  `money_18` int NOT NULL DEFAULT 0 COMMENT '第18天充值总数',
  `money_19` int NOT NULL DEFAULT 0 COMMENT '第19天充值总数',
  `money_20` int NOT NULL DEFAULT 0 COMMENT '第20天充值总数',
  `money_21` int NOT NULL DEFAULT 0 COMMENT '第21天充值总数',
  `money_22` int NOT NULL DEFAULT 0 COMMENT '第22天充值总数',
  `money_23` int NOT NULL DEFAULT 0 COMMENT '第23天充值总数',
  `money_24` int NOT NULL DEFAULT 0 COMMENT '第24天充值总数',
  `money_25` int NOT NULL DEFAULT 0 COMMENT '第25天充值总数',
  `money_26` int NOT NULL DEFAULT 0 COMMENT '第26天充值总数',
  `money_27` int NOT NULL DEFAULT 0 COMMENT '第27天充值总数',
  `money_28` int NOT NULL DEFAULT 0 COMMENT '第28天充值总数',
  `money_29` int NOT NULL DEFAULT 0 COMMENT '第29天充值总数',
  `money_30` int NOT NULL DEFAULT 0 COMMENT '第30天充值总数',
  `money_60` int NOT NULL DEFAULT 0 COMMENT '第60天充值总数',
  `money_90` int NOT NULL DEFAULT 0 COMMENT '第90天充值总数',
  `money_120` int NOT NULL DEFAULT 0 COMMENT '第120天充值总数',
  `money_180` int NOT NULL DEFAULT 0 COMMENT '第180天充值总数',
  `money_360` int NOT NULL DEFAULT 0 COMMENT '第360天充值总数',
  `all_pay_num` int NOT NULL DEFAULT 0 COMMENT '累计付费人数',
  `pay_num_1` int NOT NULL DEFAULT 0 COMMENT '第1天累计付费人数',
  `pay_num_2` int NOT NULL DEFAULT 0 COMMENT '第2天累计付费人数',
  `pay_num_3` int NOT NULL DEFAULT 0 COMMENT '第3天累计付费人数',
  `pay_num_4` int NOT NULL DEFAULT 0 COMMENT '第4天累计付费人数',
  `pay_num_5` int NOT NULL DEFAULT 0 COMMENT '第5天累计付费人数',
  `pay_num_6` int NOT NULL DEFAULT 0 COMMENT '第6天累计付费人数',
  `pay_num_7` int NOT NULL DEFAULT 0 COMMENT '第7天累计付费人数',
  `pay_num_8` int NOT NULL DEFAULT 0 COMMENT '第8天累计付费人数',
  `pay_num_9` int NOT NULL DEFAULT 0 COMMENT '第9天累计付费人数',
  `pay_num_10` int NOT NULL DEFAULT 0 COMMENT '第10天累计付费人数',
  `pay_num_11` int NOT NULL DEFAULT 0 COMMENT '第11天累计付费人数',
  `pay_num_12` int NOT NULL DEFAULT 0 COMMENT '第12天累计付费人数',
  `pay_num_13` int NOT NULL DEFAULT 0 COMMENT '第13天累计付费人数',
  `pay_num_14` int NOT NULL DEFAULT 0 COMMENT '第14天累计付费人数',
  `pay_num_15` int NOT NULL DEFAULT 0 COMMENT '第15天累计付费人数',
  `pay_num_16` int NOT NULL DEFAULT 0 COMMENT '第16天累计付费人数',
  `pay_num_17` int NOT NULL DEFAULT 0 COMMENT '第17天累计付费人数',
  `pay_num_18` int NOT NULL DEFAULT 0 COMMENT '第18天累计付费人数',
  `pay_num_19` int NOT NULL DEFAULT 0 COMMENT '第19天累计付费人数',
  `pay_num_20` int NOT NULL DEFAULT 0 COMMENT '第20天累计付费人数',
  `pay_num_21` int NOT NULL DEFAULT 0 COMMENT '第21天累计付费人数',
  `pay_num_22` int NOT NULL DEFAULT 0 COMMENT '第22天累计付费人数',
  `pay_num_23` int NOT NULL DEFAULT 0 COMMENT '第23天累计付费人数',
  `pay_num_24` int NOT NULL DEFAULT 0 COMMENT '第24天累计付费人数',
  `pay_num_25` int NOT NULL DEFAULT 0 COMMENT '第25天累计付费人数',
  `pay_num_26` int NOT NULL DEFAULT 0 COMMENT '第26天累计付费人数',
  `pay_num_27` int NOT NULL DEFAULT 0 COMMENT '第27天累计付费人数',
  `pay_num_28` int NOT NULL DEFAULT 0 COMMENT '第28天累计付费人数',
  `pay_num_29` int NOT NULL DEFAULT 0 COMMENT '第29天累计付费人数',
  `pay_num_30` int NOT NULL DEFAULT 0 COMMENT '第30天累计付费人数',
  `pay_num_60` int NOT NULL DEFAULT 0 COMMENT '第60天累计付费人数',
  `pay_num_90` int NOT NULL DEFAULT 0 COMMENT '第90天累计付费人数',
  `pay_num_120` int NOT NULL DEFAULT 0 COMMENT '第120天累计付费人数',
  `pay_num_180` int NOT NULL DEFAULT 0 COMMENT '第180天累计付费人数',
  `pay_num_360` int NOT NULL DEFAULT 0 COMMENT '第360天累计付费人数',
  `update_time` bigint NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `days_game_id_channel_area`(`date` ASC, `game_project_id` ASC, `game_id` ASC, `app_channel` ASC, `area_code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 291 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = 'ltv' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for report_daily_pay
-- ----------------------------
DROP TABLE IF EXISTS `report_daily_pay`;
CREATE TABLE `report_daily_pay`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `date` int NOT NULL DEFAULT 0 COMMENT '日期',
  `game_project_id` int NOT NULL DEFAULT 0 COMMENT '项目',
  `game_id` int NOT NULL DEFAULT 0 COMMENT '游戏id',
  `app_channel` tinyint(1) NOT NULL DEFAULT 1 COMMENT '应用渠道',
  `area_code` varchar(10) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '地区',
  `dau` int NOT NULL DEFAULT 0 COMMENT 'dau',
  `pay_1` int NOT NULL DEFAULT 0 COMMENT '第1天新增充值人数',
  `pay_2` int NOT NULL DEFAULT 0 COMMENT '第2天新增充值人数',
  `pay_3` int NOT NULL DEFAULT 0 COMMENT '第3天新增充值人数',
  `pay_4` int NOT NULL DEFAULT 0 COMMENT '第4天新增充值人数',
  `pay_5` int NOT NULL DEFAULT 0 COMMENT '第5天新增充值人数',
  `pay_6` int NOT NULL DEFAULT 0 COMMENT '第6天新增充值人数',
  `pay_7` int NOT NULL DEFAULT 0 COMMENT '第7天新增充值人数',
  `pay_8` int NOT NULL DEFAULT 0 COMMENT '第8天新增充值人数',
  `pay_9` int NOT NULL DEFAULT 0 COMMENT '第9天新增充值人数',
  `pay_10` int NOT NULL DEFAULT 0 COMMENT '第10天新增充值人数',
  `pay_11` int NOT NULL DEFAULT 0 COMMENT '第11天新增充值人数',
  `pay_12` int NOT NULL DEFAULT 0 COMMENT '第12天新增充值人数',
  `pay_13` int NOT NULL DEFAULT 0 COMMENT '第13天新增充值人数',
  `pay_14` int NOT NULL DEFAULT 0 COMMENT '第14天新增充值人数',
  `pay_15` int NOT NULL DEFAULT 0 COMMENT '第15天新增充值人数',
  `pay_16` int NOT NULL DEFAULT 0 COMMENT '第16天新增充值人数',
  `pay_17` int NOT NULL DEFAULT 0 COMMENT '第17天新增充值人数',
  `pay_18` int NOT NULL DEFAULT 0 COMMENT '第18天新增充值人数',
  `pay_19` int NOT NULL DEFAULT 0 COMMENT '第19天新增充值人数',
  `pay_20` int NOT NULL DEFAULT 0 COMMENT '第20天新增充值人数',
  `pay_21` int NOT NULL DEFAULT 0 COMMENT '第21天新增充值人数',
  `pay_22` int NOT NULL DEFAULT 0 COMMENT '第22天新增充值人数',
  `pay_23` int NOT NULL DEFAULT 0 COMMENT '第23天新增充值人数',
  `pay_24` int NOT NULL DEFAULT 0 COMMENT '第24天新增充值人数',
  `pay_25` int NOT NULL DEFAULT 0 COMMENT '第25天新增充值人数',
  `pay_26` int NOT NULL DEFAULT 0 COMMENT '第26天新增充值人数',
  `pay_27` int NOT NULL DEFAULT 0 COMMENT '第27天新增充值人数',
  `pay_28` int NOT NULL DEFAULT 0 COMMENT '第28天新增充值人数',
  `pay_29` int NOT NULL DEFAULT 0 COMMENT '第29天新增充值人数',
  `pay_30` int NOT NULL DEFAULT 0 COMMENT '第30天新增充值人数',
  `today_update_time` bigint NOT NULL COMMENT '今天基础更新时间',
  `daily_update_time` bigint NULL DEFAULT NULL COMMENT '相对今天留存更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `days_game_id_channel_area`(`date` ASC, `game_project_id` ASC, `game_id` ASC, `app_channel` ASC, `area_code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 633 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '连续充值人数' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for report_daily_pay_retain
-- ----------------------------
DROP TABLE IF EXISTS `report_daily_pay_retain`;
CREATE TABLE `report_daily_pay_retain`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `date` int NOT NULL COMMENT '日期',
  `game_project_id` int NOT NULL DEFAULT 0 COMMENT '项目',
  `game_id` int NOT NULL DEFAULT 0 COMMENT '游戏id',
  `app_channel` tinyint(1) NOT NULL DEFAULT 1 COMMENT '应用渠道',
  `area_code` varchar(10) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '地区',
  `retain_1` int NOT NULL DEFAULT 0 COMMENT '第1天新增充值人数',
  `retain_2` int NOT NULL DEFAULT 0 COMMENT '第2天新增充值登录人数',
  `retain_3` int NOT NULL DEFAULT 0 COMMENT '第3天新增充值登录人数',
  `retain_4` int NOT NULL DEFAULT 0 COMMENT '第4天新增充值登录人数',
  `retain_5` int NOT NULL DEFAULT 0 COMMENT '第5天新增充值登录人数',
  `retain_6` int NOT NULL DEFAULT 0 COMMENT '第6天新增充值登录人数',
  `retain_7` int NOT NULL DEFAULT 0 COMMENT '第7天新增充值登录人数',
  `retain_8` int NOT NULL DEFAULT 0 COMMENT '第8天新增充值登录人数',
  `retain_9` int NOT NULL DEFAULT 0 COMMENT '第9天新增充值登录人数',
  `retain_10` int NOT NULL DEFAULT 0 COMMENT '第10天新增充值登录人数',
  `retain_11` int NOT NULL DEFAULT 0 COMMENT '第11天新增充值登录人数',
  `retain_12` int NOT NULL DEFAULT 0 COMMENT '第12天新增充值登录人数',
  `retain_13` int NOT NULL DEFAULT 0 COMMENT '第13天新增充值登录人数',
  `retain_14` int NOT NULL DEFAULT 0 COMMENT '第14天新增充值登录人数',
  `retain_15` int NOT NULL DEFAULT 0 COMMENT '第15天新增充值登录人数',
  `retain_16` int NOT NULL DEFAULT 0 COMMENT '第16天新增充值登录人数',
  `retain_17` int NOT NULL DEFAULT 0 COMMENT '第17天新增充值登录人数',
  `retain_18` int NOT NULL DEFAULT 0 COMMENT '第18天新增充值登录人数',
  `retain_19` int NOT NULL DEFAULT 0 COMMENT '第19天新增充值登录人数',
  `retain_20` int NOT NULL DEFAULT 0 COMMENT '第20天新增充值登录人数',
  `retain_21` int NOT NULL DEFAULT 0 COMMENT '第21天新增充值登录人数',
  `retain_22` int NOT NULL DEFAULT 0 COMMENT '第22天新增充值登录人数',
  `retain_23` int NOT NULL DEFAULT 0 COMMENT '第23天新增充值登录人数',
  `retain_24` int NOT NULL DEFAULT 0 COMMENT '第24天新增充值登录人数',
  `retain_25` int NOT NULL DEFAULT 0 COMMENT '第25天新增充值登录人数',
  `retain_26` int NOT NULL DEFAULT 0 COMMENT '第26天新增充值登录人数',
  `retain_27` int NOT NULL DEFAULT 0 COMMENT '第27天新增充值登录人数',
  `retain_28` int NOT NULL DEFAULT 0 COMMENT '第28天新增充值登录人数',
  `retain_29` int NOT NULL DEFAULT 0 COMMENT '第29天新增充值登录人数',
  `retain_30` int NOT NULL DEFAULT 0 COMMENT '第30天新增充值登录人数',
  `retain_60` int NOT NULL DEFAULT 0 COMMENT '第60天新增充值登录人数',
  `retain_90` int NOT NULL DEFAULT 0 COMMENT '第90天新增充值登录人数',
  `update_time` bigint NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `days_game_id_channel_area`(`date` ASC, `game_project_id` ASC, `game_id` ASC, `app_channel` ASC, `area_code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 217 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '付费留存' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for report_daily_summary
-- ----------------------------
DROP TABLE IF EXISTS `report_daily_summary`;
CREATE TABLE `report_daily_summary`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `date` int NOT NULL COMMENT '日期 20230101 ',
  `game_project_id` int NOT NULL DEFAULT 0 COMMENT '项目',
  `game_id` int NOT NULL COMMENT '游戏',
  `app_channel` tinyint(1) NOT NULL COMMENT '应用渠道：1(官方)、2(安卓)、3(苹果)',
  `area_code` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地区',
  `dau` int NOT NULL DEFAULT 0 COMMENT 'DAU：登陆了SDK账号，就算1个DAU',
  `register` int NOT NULL DEFAULT 0 COMMENT '新账号注册数',
  `device` int NOT NULL DEFAULT 0 COMMENT '新设备数',
  `pay_money` int NOT NULL DEFAULT 0 COMMENT '总付费数',
  `pay_people` int NOT NULL DEFAULT 0 COMMENT '总付费人数',
  `pay_times` int NOT NULL DEFAULT 0 COMMENT '付费次数',
  `new_pay_money` int NOT NULL DEFAULT 0 COMMENT '新账号总付费数',
  `new_pay_people` int NOT NULL DEFAULT 0 COMMENT '新账号总付费人数',
  `new_pay_times` int NOT NULL DEFAULT 0 COMMENT '新账号付费次数',
  `first_pay_money` int NOT NULL DEFAULT 0 COMMENT '首次充值总额',
  `first_pay_times` int NOT NULL DEFAULT 0 COMMENT '首次充值充值人次',
  `first_pay_people` int NOT NULL DEFAULT 0 COMMENT '首次充值人数',
  `order_pay_money` int NOT NULL DEFAULT 0 COMMENT '总付费数-按订单',
  `order_pay_people` int NOT NULL DEFAULT 0 COMMENT '总付费人数-按订单',
  `order_pay_times` int NOT NULL DEFAULT 0 COMMENT '总付费人次-按订单',
  `order_new_pay_money` int NOT NULL DEFAULT 0 COMMENT '新账号总付费人数-按订单',
  `order_new_pay_people` int NOT NULL DEFAULT 0 COMMENT '新账号总付费人数-按订单',
  `order_new_pay_times` int NOT NULL DEFAULT 0 COMMENT '新账号总付费人次-按订单',
  `pcu` int NOT NULL DEFAULT 0 COMMENT 'PCU：当日最高同时在线人数',
  `retain_total` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '次留：前一天注册的人的登录数',
  `update_time` int NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `date_game_channel`(`date` ASC, `game_project_id` ASC, `game_id` ASC, `app_channel` ASC, `area_code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 263 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '每日综合数据统计表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for report_daily_user_retain
-- ----------------------------
DROP TABLE IF EXISTS `report_daily_user_retain`;
CREATE TABLE `report_daily_user_retain`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `date` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '日期',
  `game_project_id` int NOT NULL DEFAULT 0 COMMENT '项目',
  `game_id` int NOT NULL DEFAULT 0 COMMENT '游戏id',
  `app_channel` tinyint(1) NOT NULL DEFAULT 1 COMMENT '应用渠道',
  `area_code` varchar(10) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '地区',
  `retain_1` int NOT NULL DEFAULT 0 COMMENT '第1日留存',
  `retain_2` int NOT NULL DEFAULT 0 COMMENT '第2日留存',
  `retain_3` int NOT NULL DEFAULT 0 COMMENT '第3日留存',
  `retain_4` int NOT NULL DEFAULT 0 COMMENT '第4日留存',
  `retain_5` int NOT NULL DEFAULT 0 COMMENT '第5日留存',
  `retain_6` int NOT NULL DEFAULT 0 COMMENT '第6日留存',
  `retain_7` int NOT NULL DEFAULT 0 COMMENT '第7日留存',
  `retain_8` int NOT NULL DEFAULT 0 COMMENT '第8日留存',
  `retain_9` int NOT NULL DEFAULT 0 COMMENT '第9日留存',
  `retain_10` int NOT NULL DEFAULT 0 COMMENT '第10日留存',
  `retain_11` int NOT NULL DEFAULT 0 COMMENT '第11日留存',
  `retain_12` int NOT NULL DEFAULT 0 COMMENT '第12日留存',
  `retain_13` int NOT NULL DEFAULT 0 COMMENT '第13日留存',
  `retain_14` int NOT NULL DEFAULT 0 COMMENT '第14日留存',
  `retain_15` int NOT NULL DEFAULT 0 COMMENT '第15日留存',
  `retain_16` int NOT NULL DEFAULT 0 COMMENT '第16日留存',
  `retain_17` int NOT NULL DEFAULT 0 COMMENT '第17日留存',
  `retain_18` int NOT NULL DEFAULT 0 COMMENT '第18日留存',
  `retain_19` int NOT NULL DEFAULT 0 COMMENT '第19日留存',
  `retain_20` int NOT NULL DEFAULT 0 COMMENT '第20日留存',
  `retain_21` int NOT NULL DEFAULT 0 COMMENT '第21日留存',
  `retain_22` int NOT NULL DEFAULT 0 COMMENT '第22日留存',
  `retain_23` int NOT NULL DEFAULT 0 COMMENT '第23日留存',
  `retain_24` int NOT NULL DEFAULT 0 COMMENT '第24日留存',
  `retain_25` int NOT NULL DEFAULT 0 COMMENT '第25日留存',
  `retain_26` int NOT NULL DEFAULT 0 COMMENT '第26日留存',
  `retain_27` int NOT NULL DEFAULT 0 COMMENT '第27日留存',
  `retain_28` int NOT NULL DEFAULT 0 COMMENT '第28日留存',
  `retain_29` int NOT NULL DEFAULT 0 COMMENT '第29日留存',
  `retain_30` int NOT NULL DEFAULT 0 COMMENT '第30日留存',
  `retain_60` int NOT NULL DEFAULT 0 COMMENT '第60日留存',
  `retain_90` int NOT NULL DEFAULT 0 COMMENT '第90日留存',
  `update_time` bigint NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `days_game_id_channel_area`(`date` ASC, `game_project_id` ASC, `game_id` ASC, `app_channel` ASC, `area_code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 700 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '账号留存统计表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_depts
-- ----------------------------
DROP TABLE IF EXISTS `sys_depts`;
CREATE TABLE `sys_depts`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '部门名称',
  `parent_id` bigint NOT NULL DEFAULT 0 COMMENT '父部门id',
  `sort` int NOT NULL DEFAULT 0 COMMENT '显示顺序',
  `leader_user_id` bigint NULL DEFAULT 0 COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '联系电话',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '邮箱',
  `status` tinyint NOT NULL COMMENT '部门状态（0正常 1停用）',
  `creator` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '创建者',
  `created_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updater` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '更新者',
  `updated_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '部门表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `sort` int NOT NULL DEFAULT 0 COMMENT '字典排序',
  `label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '字典标签',
  `value` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '字典键值',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '字典类型',
  `status` tinyint NOT NULL DEFAULT 0 COMMENT '状态（0正常 1停用）',
  `color_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '颜色类型',
  `css_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT 'css 样式',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updater` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` int NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `dict_type`(`dict_type` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1364 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '字典数据表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '字典主键',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '字典名称',
  `type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '字典类型',
  `status` tinyint NOT NULL DEFAULT 0 COMMENT '状态（0正常 1停用）',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updater` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` int NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `dict_type`(`type` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 185 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '字典类型表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_menus`;
CREATE TABLE `sys_menus`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `menu_type` tinyint NOT NULL DEFAULT 0 COMMENT '0代表菜单、1代表iframe、2代表外链、3代表按钮',
  `parent_id` int UNSIGNED NOT NULL DEFAULT 0,
  `title` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '显示名称',
  `path` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '访问的URL',
  `redirect` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '重定向路径',
  `name` varchar(300) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '组件名字',
  `component` varchar(300) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '组件路径',
  `rank` int UNSIGNED NOT NULL DEFAULT 1 COMMENT ' 菜单排序，值越高排的越后（只针对顶级路由）',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '菜单状态（0正常 1停用）',
  `icon` varchar(300) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT 'icon ',
  `keep_alive` tinyint(1) NOT NULL DEFAULT 0 COMMENT '缓存页面（是否缓存该路由页面，开启后会保存该页面的整体状态，刷新后会清空状态）',
  `show_link` tinyint(1) NOT NULL DEFAULT 0 COMMENT '菜单（是否显示该菜单）',
  `hidden_tag` tinyint(1) NOT NULL DEFAULT 0 COMMENT '标签页（当前菜单名称或自定义信息禁止添加到标签页）',
  `fixed_tag` tinyint(1) NOT NULL DEFAULT 0 COMMENT '固定标签页（当前菜单名称是否固定显示在标签页且不可关闭）',
  `extra_icon` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '右侧图标',
  `enter_transition` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '进场动画（页面加载动画）',
  `leave_transition` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '离场动画（页面加载动画）',
  `active_path` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '菜单激活（将某个菜单激活，主要用于通过query或params传参的路由，当它们通过配置showLink: false后不在菜单中显示，就不会有任何菜单高亮，而通过设置activePath指定激活菜单即可获得高亮，activePath为指定激活菜单的path）',
  `auths` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '权限标识（按钮级别权限设置）',
  `frame_src` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '链接地址（需要内嵌的iframe链接地址）',
  `frame_loading` tinyint(1) NOT NULL DEFAULT 0 COMMENT '加载动画（内嵌的iframe页面是否开启首次加载动画）',
  `show_parent` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否显示父级菜单',
  `created_at` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted_at` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `module_sort`(`rank` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 142 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '系统菜单' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '岗位编码',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '岗位名称',
  `sort` int NOT NULL COMMENT '显示顺序',
  `status` tinyint NOT NULL COMMENT '状态（0正常 1停用）',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updater` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '岗位信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`  (
  `role_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '台子ID',
  `permission_id` int UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`role_id`, `permission_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_roles
-- ----------------------------
DROP TABLE IF EXISTS `sys_roles`;
CREATE TABLE `sys_roles`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `name` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '角色名',
  `status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态：1 正常 0 禁用',
  `code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色权限字符串',
  `sort` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '显示顺序',
  `data_scope` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '数据权限',
  `data_scope_menu_ids` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单权限',
  `type` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '角色类型',
  `remark` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `creator` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者',
  `created_at` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updater` int UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `updated_at` int NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted_at` int NOT NULL DEFAULT 0 COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `status`(`status` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_setting
-- ----------------------------
DROP TABLE IF EXISTS `sys_setting`;
CREATE TABLE `sys_setting`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `source_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '来源ID：platformID or gameID',
  `set_name` varchar(20) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '',
  `set_val` mediumtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `ss`(`set_name` ASC, `source_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '系统设置表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_user_logs
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_logs`;
CREATE TABLE `sys_user_logs`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `log_type` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '日志类型',
  `user_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `account` char(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '账号名',
  `status` int NOT NULL DEFAULT 0 COMMENT '操作结果',
  `descriptor` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '操作描述',
  `path` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '模块名',
  `ip` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT 'ip',
  `module` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '模块',
  `request` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '参数',
  `response` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '操作结果',
  `request_id` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '请求ID',
  `user_agent` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT 'UserAgent',
  `created_at` int NOT NULL DEFAULT 0 COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `account`(`account` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3102 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '记录系统日志' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',
  `role_id` varchar(20) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_id`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `account` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '帐户名',
  `password` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '登陆密码',
  `salt` varchar(16) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '盐值',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '-1:未激活，0：正常，1：禁用',
  `nickname` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `mobile` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '绑定的手机号',
  `full_name` varchar(12) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `reg_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '注册时间',
  `reg_ip` char(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '注册IP',
  `last_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后登陆时间',
  `last_ip` char(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '最后登陆IP',
  `last_login_way` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后登陆方式： 0(visitor)、1(email)、 2(facebook)、3(google)、4(apple)',
  `dept_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '部门ID',
  `role_ids` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '角色ids',
  `post_ids` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '岗位编号数组',
  `wechat` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '微信openId',
  `qq` varchar(15) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `email` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '用户邮箱',
  `sex` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '2:未知 1:男 0: 女',
  `avatar` varchar(400) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `login_times` mediumint UNSIGNED NOT NULL DEFAULT 0 COMMENT '登陆次数',
  `remark` varchar(800) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `opt_user` int NOT NULL DEFAULT 0 COMMENT '操作人',
  `created_at` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted_at` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `pu`(`account` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '管理员表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for vip_server_manage
-- ----------------------------
DROP TABLE IF EXISTS `vip_server_manage`;
CREATE TABLE `vip_server_manage`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `admin_uid` int NOT NULL DEFAULT 0 COMMENT '负责人',
  `game_id` int NOT NULL DEFAULT 0 COMMENT '游戏id',
  `vip_staff_id` int NOT NULL DEFAULT 0 COMMENT 'vip客服',
  `level` tinyint(1) NOT NULL DEFAULT 0 COMMENT '0初级 1中级 2高级',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '0禁用 1启用',
  `create_time` int NOT NULL DEFAULT 0,
  `create_user` int NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_time` int NOT NULL DEFAULT 0 COMMENT '更新时间',
  `update_user` int NOT NULL DEFAULT 0 COMMENT '更新人',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = 'VIP区服管理负责人' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for vip_server_manage_list
-- ----------------------------
DROP TABLE IF EXISTS `vip_server_manage_list`;
CREATE TABLE `vip_server_manage_list`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '0',
  `vip_server_manage_id` int NOT NULL DEFAULT 0 COMMENT 'VIP区服管理负责人表id',
  `game_server_id` int NOT NULL DEFAULT 0 COMMENT '游戏区服管理表id',
  `zone_id` int NOT NULL DEFAULT 0 COMMENT '游戏区服id',
  `game_id` int NOT NULL DEFAULT 0 COMMENT '游戏id',
  `deleted_at` timestamp NULL DEFAULT NULL,
  `create_time` int NOT NULL DEFAULT 0,
  `create_user` int NOT NULL DEFAULT 0,
  `update_time` int NOT NULL DEFAULT 0,
  `update_user` int NOT NULL DEFAULT 0 COMMENT '操作人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 104 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = 'VIP区服管理负责人管理区服列表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for vip_staff
-- ----------------------------
DROP TABLE IF EXISTS `vip_staff`;
CREATE TABLE `vip_staff`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `card` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '客服名片',
  `mail` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '企业邮箱',
  `name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '客服名称',
  `nickname` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '客服昵称',
  `pid` int NOT NULL DEFAULT 0 COMMENT 'vip组归属 0组长',
  `staff_id` varchar(45) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '0' COMMENT '七鱼客服id',
  `staff_icon` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '客服头像',
  `staff_level` tinyint(1) NOT NULL DEFAULT 0 COMMENT '客服等级 0初级 1中级 2高级',
  `staff_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '七鱼客服名称',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '1启用 0禁用',
  `wechat` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '微信号',
  `wechat_qr_code` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '图片',
  `create_time` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  `create_user` int NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_time` int NOT NULL DEFAULT 0 COMMENT '更新时间',
  `update_user` int NOT NULL DEFAULT 0 COMMENT '更新人',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
