/*
 Navicat Premium Data Transfer

 Source Server         : docker-mysql-master
 Source Server Type    : MySQL
 Source Server Version : 80032
 Source Host           : 192.168.122.247:3306
 Source Schema         : xiaozhu_log

 Target Server Type    : MySQL
 Target Server Version : 80032
 File Encoding         : 65001

 Date: 17/04/2025 14:39:40
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for log_access
-- ----------------------------
DROP TABLE IF EXISTS `log_access`;
CREATE TABLE `log_access`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `trace_id` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '跟踪ID',
  `request_domain` varchar(300) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT 'domain',
  `request_path` varchar(300) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT 'path',
  `request_method` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '请求方法名',
  `request_header` varchar(600) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '请求头',
  `request_query` varchar(1024) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '请求query',
  `request_body` varchar(8192) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '请求体',
  `request_ip` char(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT 'IP ',
  `response_code` varchar(10) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '’‘' COMMENT '响应码',
  `response_body` varchar(1024) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '响应',
  `duration` decimal(10, 4) UNSIGNED NOT NULL DEFAULT 0.0000 COMMENT '运行时长,单位秒',
  `ts` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '事件时间',
  `ua` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '浏览器',
  `level` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '级别',
  `area` varchar(300) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区',
  `version` varchar(45) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '接口版本',
  `game_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `device_id` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '设备号',
  `user_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `server_id` varchar(15) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '服务器ID',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `trace_id`(`trace_id` ASC) USING BTREE,
  INDEX `ts`(`ts` ASC) USING BTREE,
  INDEX `response_code`(`response_code` ASC) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `request_header`(`request_header` ASC) USING BTREE,
  INDEX `device_id`(`device_id` ASC) USING BTREE,
  INDEX `request_ip`(`request_ip` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = 'API日志' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for log_active
-- ----------------------------
DROP TABLE IF EXISTS `log_active`;
CREATE TABLE `log_active`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `promote_code` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '推广码',
  `adid` varchar(64) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '广告id',
  `app_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏项目ID',
  `game_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `app_channel` int UNSIGNED NOT NULL DEFAULT 1 COMMENT '应用渠道：1(官方)、2(安卓)、3(苹果)',
  `os` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '操作系统：android、ios',
  `cause` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '广告归因依据',
  `device_id` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '设备号',
  `ip` char(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT 'IP ',
  `area_code` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区码',
  `area` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区',
  `ts` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '事件时间，即激活时间',
  `created_at` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  `days` int NULL DEFAULT 0 COMMENT '日期',
  `request_id` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '生成记录的请求日志',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `device_id_ts`(`device_id` ASC, `ts` ASC) USING BTREE,
  INDEX `add_time`(`created_at` ASC) USING BTREE,
  INDEX `days`(`days` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15715 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '激活日志' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for log_adjust_event
-- ----------------------------
DROP TABLE IF EXISTS `log_adjust_event`;
CREATE TABLE `log_adjust_event`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  `game_project_id` int NOT NULL DEFAULT 0 COMMENT '项目id',
  `game_id` int NOT NULL DEFAULT 99999 COMMENT '后台应用游戏id',
  `days` int NOT NULL DEFAULT 0 COMMENT '时间发生天，created_at 时间戳转换',
  `channel_id` int NOT NULL DEFAULT 0 COMMENT '媒体渠道',
  `app_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'adjust平台应用id',
  `app_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '应用名称',
  `app_version` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '应用版本',
  `store` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '点击目标商店',
  `user_id` bigint NULL DEFAULT 0 COMMENT '平台user_id',
  `adid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '设备广告标识码，所有平台都适用',
  `device_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT 'gps_adid|idfa',
  `gps_adid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT 'Google Play 商店广告 ID',
  `android_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '安卓id',
  `idfa` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '广告商 ID（仅限 iOS）',
  `idfv` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '供应商的大写 iOS ID',
  `trace_id` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '',
  `tracker` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '6 或 7 个字符 调整跟踪器标记',
  `tracker_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '同 Adjust 控制面板中所定义的当前跟踪链接名称',
  `activity_kind` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '事件类型：展示、点击、安装、被拒安装、被拒再归因、会话、再归因、卸载、重装、再归因重装、归因更新、应用内事件或广告支出',
  `event` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '事件标识符',
  `event_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '事件名称',
  `created_at` int NOT NULL DEFAULT 0 COMMENT '事件发生时间',
  `click_time` int NULL DEFAULT 0 COMMENT '点击时间',
  `sk_campaign_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '广告展示id',
  `campaign_id` bigint NULL DEFAULT 0 COMMENT '广告id,  google_ads_campaign_id 谷歌，fb_deeplink_campaign_id  FB',
  `campaign_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '广告名称',
  `adgroup_id` bigint NULL DEFAULT 0 COMMENT '广告组id',
  `adgroup_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '广告组名称',
  `creative_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0',
  `creative_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '广告创意名称',
  `cost_type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '消耗类型，仅支持广告',
  `cost_amount` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '用户参与成本（仅适用于广告支出跟踪）',
  `cost_currency` char(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '成本数据的 ISO 4217 货币代码（仅适用于广告支出跟踪）',
  `reporting_cost` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '转换为应用的报告货币并在调整仪表板中报告的用户参与成本（仅适用于广告支出跟踪）',
  `rejection_reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '拒绝归因原因',
  `conversion_duration` int NULL DEFAULT 0 COMMENT '点击和安装或再归因之间的时间（以秒为单位）',
  `network_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '网络类型',
  `ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT 'ip地址',
  `mac` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT 'mac取值顺序 md5, sha1',
  `match_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '归因方法',
  `mcc` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '移动国家/地区码',
  `mnc` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '移动网络代码',
  `country` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '设备两字符国家代码',
  `country_subdivision` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '国家/地区的设备细分，例如州',
  `city` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '城市',
  `language` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '设备语言',
  `device_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '设备类型',
  `device_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '设备名称',
  `device_model` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '设备型号',
  `os` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '操作系统',
  `os_version` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '操作系统版本',
  `sdk_version` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT 'sdk版本',
  `random` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '随机数（每个回调唯一）',
  `random_user_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '随机用户 ID（每个设备每个应用程序）',
  `timezone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '设备时区',
  `last_time_spent` int NULL DEFAULT 0 COMMENT '用户上次会话的长度（以秒为单位）',
  `time_spent` int NULL DEFAULT 0 COMMENT '用户当前会话的时长（以秒为单位）',
  `session_count` int NULL DEFAULT 0 COMMENT '当前Adjust SDK记录的会话数',
  `lifetime_session_count` int NULL DEFAULT 0 COMMENT '整个用户生命周期记录的会话数量',
  `is_reattributed` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '1 如果用户从较早的来源至少被重新归因一次；如果用户从未被重新归因过，则为 0',
  `revenue_float` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '收入（以整个货币单位计算）',
  `revenue_cny` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '收入，单位 CNY（人民币）',
  `revenue` int NULL DEFAULT 0 COMMENT '收入，以美分为单位',
  `currency` char(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '原始 ISO 4217 货币代码',
  `environment` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '环境，生产环境 production, 沙箱 sandbox',
  `reporting_revenue` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '调整仪表板中报告的收入（以整个货币单位表示）',
  `reporting_currency` char(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '报告 ISO 4217 货币代码的仪表板',
  `push_token` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '推送通知令牌，即注册令牌（Android）、设备令牌（iOS）',
  `publisher_parameters` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '由Adjust SDK收集的自定义发布商参数（从不显示在Adjust Dashboard中）',
  `add_time` int NOT NULL DEFAULT 0 COMMENT '本地记录时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `random`(`random` ASC) USING BTREE,
  INDEX `created_at`(`created_at` ASC) USING BTREE,
  INDEX `event`(`activity_kind` ASC, `event_name` ASC, `event` ASC) USING BTREE,
  INDEX `days_adid`(`days` ASC, `adid` ASC) USING BTREE,
  INDEX `tracker`(`game_id` ASC, `tracker` ASC) USING BTREE,
  INDEX `device_id_game_id_adid`(`device_id` ASC, `game_id` ASC, `adid` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for log_adjust_event_only
-- ----------------------------
DROP TABLE IF EXISTS `log_adjust_event_only`;
CREATE TABLE `log_adjust_event_only`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
  `game_project_id` int NOT NULL DEFAULT 0 COMMENT '项目id',
  `game_id` int NOT NULL DEFAULT 99999 COMMENT '后台应用游戏id',
  `days` int NOT NULL DEFAULT 0 COMMENT '时间发生天，created_at 时间戳转换',
  `channel_id` int NOT NULL DEFAULT 0 COMMENT '媒体渠道',
  `app_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'adjust平台应用id',
  `app_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '应用名称',
  `app_version` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '应用版本',
  `store` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '点击目标商店',
  `user_id` bigint NULL DEFAULT 0 COMMENT '平台user_id',
  `adid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '设备广告标识码，所有平台都适用',
  `device_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT 'gps_adid|idfa',
  `gps_adid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT 'Google Play 商店广告 ID',
  `android_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '安卓id',
  `idfa` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '广告商 ID（仅限 iOS）',
  `idfv` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '供应商的大写 iOS ID',
  `trace_id` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '',
  `tracker` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '6 或 7 个字符 调整跟踪器标记',
  `tracker_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '同 Adjust 控制面板中所定义的当前跟踪链接名称',
  `activity_kind` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '事件类型：展示、点击、安装、被拒安装、被拒再归因、会话、再归因、卸载、重装、再归因重装、归因更新、应用内事件或广告支出',
  `event` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '事件标识符',
  `event_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '事件名称',
  `created_at` int NOT NULL DEFAULT 0 COMMENT '事件发生时间',
  `click_time` int NULL DEFAULT 0 COMMENT '点击时间',
  `sk_campaign_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '广告展示id',
  `campaign_id` bigint NULL DEFAULT 0 COMMENT '广告id,  google_ads_campaign_id 谷歌，fb_deeplink_campaign_id  FB',
  `campaign_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '广告名称',
  `adgroup_id` bigint NULL DEFAULT 0 COMMENT '广告组id',
  `adgroup_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '广告组名称',
  `creative_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0',
  `creative_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '广告创意名称',
  `cost_type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '消耗类型，仅支持广告',
  `cost_amount` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '用户参与成本（仅适用于广告支出跟踪）',
  `cost_currency` char(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '成本数据的 ISO 4217 货币代码（仅适用于广告支出跟踪）',
  `reporting_cost` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '转换为应用的报告货币并在调整仪表板中报告的用户参与成本（仅适用于广告支出跟踪）',
  `rejection_reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '拒绝归因原因',
  `conversion_duration` int NULL DEFAULT 0 COMMENT '点击和安装或再归因之间的时间（以秒为单位）',
  `network_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '网络类型',
  `ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT 'ip地址',
  `mac` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT 'mac取值顺序 md5, sha1',
  `match_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '归因方法',
  `mcc` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '移动国家/地区码',
  `mnc` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '移动网络代码',
  `country` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '设备两字符国家代码',
  `country_subdivision` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '国家/地区的设备细分，例如州',
  `city` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '城市',
  `language` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '设备语言',
  `device_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '设备类型',
  `device_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '设备名称',
  `device_model` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '设备型号',
  `os` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '操作系统',
  `os_version` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '操作系统版本',
  `sdk_version` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT 'sdk版本',
  `random` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '随机数（每个回调唯一）',
  `random_user_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '随机用户 ID（每个设备每个应用程序）',
  `timezone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '设备时区',
  `last_time_spent` int NULL DEFAULT 0 COMMENT '用户上次会话的长度（以秒为单位）',
  `time_spent` int NULL DEFAULT 0 COMMENT '用户当前会话的时长（以秒为单位）',
  `session_count` int NULL DEFAULT 0 COMMENT '当前Adjust SDK记录的会话数',
  `lifetime_session_count` int NULL DEFAULT 0 COMMENT '整个用户生命周期记录的会话数量',
  `is_reattributed` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '1 如果用户从较早的来源至少被重新归因一次；如果用户从未被重新归因过，则为 0',
  `revenue_float` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '收入（以整个货币单位计算）',
  `revenue_cny` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '收入，单位 CNY（人民币）',
  `revenue` int NULL DEFAULT 0 COMMENT '收入，以美分为单位',
  `currency` char(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '原始 ISO 4217 货币代码',
  `environment` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '环境，生产环境 production, 沙箱 sandbox',
  `reporting_revenue` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '调整仪表板中报告的收入（以整个货币单位表示）',
  `reporting_currency` char(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '报告 ISO 4217 货币代码的仪表板',
  `push_token` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '推送通知令牌，即注册令牌（Android）、设备令牌（iOS）',
  `publisher_parameters` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '由Adjust SDK收集的自定义发布商参数（从不显示在Adjust Dashboard中）',
  `add_time` int NOT NULL DEFAULT 0 COMMENT '本地记录时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `random`(`days` ASC, `game_project_id` ASC, `game_id` ASC, `activity_kind` ASC, `event_name` ASC, `random` ASC) USING BTREE,
  INDEX `created_at`(`created_at` ASC) USING BTREE,
  INDEX `event`(`activity_kind` ASC, `event_name` ASC, `event` ASC) USING BTREE,
  INDEX `tracker`(`game_id` ASC, `tracker` ASC) USING BTREE,
  INDEX `days_project_adid`(`days` ASC, `game_project_id` ASC, `adid` ASC) USING BTREE,
  INDEX `device_id_game_id_adid`(`device_id` ASC, `game_id` ASC, `adid` ASC) USING BTREE,
  INDEX `game_project_id_ip`(`game_project_id` ASC, `ip` ASC) USING BTREE,
  INDEX `project_adid`(`game_project_id` ASC, `adid` ASC) USING BTREE,
  INDEX `game_id_ip`(`game_id` ASC, `ip` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for log_device
-- ----------------------------
DROP TABLE IF EXISTS `log_device`;
CREATE TABLE `log_device`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `promote_code` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '推广码',
  `adid` varchar(64) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '广告id',
  `app_id` int NOT NULL DEFAULT 0 COMMENT '游戏项目ID',
  `game_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `app_channel` int NOT NULL DEFAULT 1 COMMENT '应用渠道：1(官方)、2(安卓)、3(苹果)',
  `area_code` varchar(10) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT 'CN' COMMENT '地区',
  `area` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区',
  `os` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '操作系统：android、ios',
  `cause` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '广告归因依据',
  `device_id` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '设备号',
  `ip` char(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT 'IP ',
  `ts` bigint NOT NULL COMMENT '更新时间时间',
  `days` int NOT NULL DEFAULT 0 COMMENT '日期',
  `created_at` int NOT NULL DEFAULT 0 COMMENT 'log_active的创建日期',
  `channel_id` int NOT NULL DEFAULT 0 COMMENT '媒体渠道',
  `campaign_id` bigint NOT NULL DEFAULT 0 COMMENT '广告计划',
  `adgroup_id` bigint NOT NULL DEFAULT 0 COMMENT '广告组',
  `creative_id` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '创意',
  `request_id` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '生成记录的请求日志',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `game_device`(`game_id` ASC, `device_id` ASC) USING BTREE,
  INDEX `days`(`days` ASC) USING BTREE,
  INDEX `device_id`(`device_id` ASC) USING BTREE,
  INDEX `create_time`(`created_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13517 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '游戏设备表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for log_exchange
-- ----------------------------
DROP TABLE IF EXISTS `log_exchange`;
CREATE TABLE `log_exchange`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `os` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '操作系统：android、ios',
  `device_id` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '设备号',
  `ip` char(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT 'IP ',
  `area_code` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区码',
  `area` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区',
  `create_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '事件时间，即激活时间',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `promote_code` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '推广码',
  `adid` varchar(64) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '广告id',
  `user_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `account` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL DEFAULT '' COMMENT '帐户名',
  `order_num` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '订单号',
  `pay_money` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付金额',
  `app_channel` int UNSIGNED NOT NULL DEFAULT 1 COMMENT '应用渠道',
  `pay_currency` tinyint NOT NULL DEFAULT 1 COMMENT '支付货币类型:1(积分)',
  `game_project_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏项目ID',
  `game_id` mediumint UNSIGNED NOT NULL DEFAULT 0,
  `game_name` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '游戏名',
  `zone_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '0' COMMENT '区服',
  `zone_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '区服名',
  `game_role_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '0',
  `game_role_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名',
  `game_role_level` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '0' COMMENT '角色级别',
  `days` date NULL DEFAULT NULL COMMENT '日期',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `order_num`(`order_num` ASC) USING BTREE,
  INDEX `ts`(`create_time` ASC) USING BTREE,
  INDEX `game_zone_role`(`game_id` ASC, `zone_id` ASC, `game_role_id` ASC) USING BTREE,
  INDEX `days`(`days` ASC) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '兑换商品日志' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for log_login
-- ----------------------------
DROP TABLE IF EXISTS `log_login`;
CREATE TABLE `log_login`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `game_project_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏项目ID',
  `game_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `app_channel` int UNSIGNED NOT NULL DEFAULT 1 COMMENT '应用渠道：1(官方)、2(安卓)、3(苹果)',
  `os` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '操作系统：android、ios',
  `device_id` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '设备号',
  `ip` char(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT 'IP ',
  `create_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '事件时间，即登陆时间',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `user_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `account` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL DEFAULT '' COMMENT '帐户名',
  `account_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '帐号类型：0  游客; 1 普通账号 ; 2 手机账号; 3  微信账号;4 机器人',
  `promote_code` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '推广码',
  `adid` varchar(64) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '广告id',
  `area_code` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区码',
  `area` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区',
  `login_way` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '登陆方式： 0(visitor)、1(email)、 2(facebook)、3(google)、4(apple)',
  `days` date NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `member_id_ts`(`user_id` ASC, `create_time` ASC) USING BTREE,
  INDEX `days`(`days` ASC) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `ip`(`ip` ASC) USING BTREE,
  INDEX `device_id`(`device_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '玩家登录日志' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for log_pay
-- ----------------------------
DROP TABLE IF EXISTS `log_pay`;
CREATE TABLE `log_pay`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `os` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '操作系统：android、ios',
  `device_id` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '设备号',
  `ip` char(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT 'IP ',
  `area_code` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区码',
  `area` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区',
  `create_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '事件时间，即激活时间',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `promote_code` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '推广码',
  `adid` varchar(64) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '广告id',
  `user_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `account` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL DEFAULT '' COMMENT '帐户名',
  `merchant_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '商户ID',
  `pay_channel` varchar(20) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '支付编号',
  `order_num` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '订单号',
  `pay_money` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付金额',
  `currency` varchar(45) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT 'USD' COMMENT '货币代码:USD(美元)、HKG(港元) 、MAC(澳门元) 、TWD(新台币) ...',
  `app_channel` int UNSIGNED NOT NULL DEFAULT 1 COMMENT '应用渠道：1(官方)、2(安卓)、3(苹果)',
  `game_project_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏项目ID',
  `game_id` mediumint UNSIGNED NOT NULL DEFAULT 0,
  `game_name` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '游戏名',
  `zone_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '0' COMMENT '区服',
  `zone_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '区服名',
  `game_role_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '0',
  `game_role_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名',
  `game_role_level` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '0' COMMENT '角色级别',
  `is_first` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '首充：是(1)、否(0)',
  `first_game_id` mediumint UNSIGNED NOT NULL DEFAULT 0 COMMENT '首次进游戏ID',
  `days` date NULL DEFAULT NULL COMMENT '日期',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `order_num`(`order_num` ASC) USING BTREE,
  INDEX `ts`(`create_time` ASC) USING BTREE,
  INDEX `is_first`(`is_first` ASC) USING BTREE,
  INDEX `days`(`days` ASC) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `game_zone_role`(`game_id` ASC, `zone_id` ASC, `game_role_id` ASC) USING BTREE,
  INDEX `pay_money`(`pay_money` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '充值日志' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for log_player
-- ----------------------------
DROP TABLE IF EXISTS `log_player`;
CREATE TABLE `log_player`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `promote_code` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '推广码',
  `adid` varchar(64) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '广告id',
  `os` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '操作系统：android、ios',
  `device_id` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '设备号',
  `ip` char(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT 'IP ',
  `area_code` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区码',
  `area` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区',
  `create_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '事件时间，即创建角色时间',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `user_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `account` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL DEFAULT '' COMMENT '帐户名',
  `app_channel` int UNSIGNED NOT NULL DEFAULT 1 COMMENT '应用渠道：1(官方)、2(安卓)、3(苹果)',
  `game_project_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏项目ID',
  `game_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `game_name` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '游戏名',
  `zone_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '区服',
  `zone_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '区服名',
  `game_role_id` bigint NOT NULL DEFAULT 0,
  `game_role_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名',
  `game_role_level` int NOT NULL DEFAULT 0 COMMENT '角色级别',
  `online` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '在线时长单位为秒',
  `total_pay_money` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '支付金额',
  `days` date NULL DEFAULT NULL COMMENT '日期',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `ugzg`(`user_id` ASC, `game_id` ASC, `zone_id` ASC, `game_role_id` ASC) USING BTREE,
  INDEX `days`(`days` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '角色创建日志' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for log_player_login
-- ----------------------------
DROP TABLE IF EXISTS `log_player_login`;
CREATE TABLE `log_player_login`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `promote_code` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '推广码',
  `adid` varchar(64) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '广告id',
  `os` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '操作系统：android、ios',
  `device_id` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '设备号',
  `ip` char(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT 'IP ',
  `area_code` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区码',
  `area` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区',
  `create_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '事件时间，即激活时间',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `user_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `account` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL DEFAULT '' COMMENT '帐户名',
  `app_channel` int UNSIGNED NOT NULL DEFAULT 1 COMMENT '应用渠道：1(官方)、2(安卓)、3(苹果)',
  `game_project_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏项目ID',
  `game_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `game_name` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '游戏名',
  `zone_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '区服',
  `zone_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '区服名',
  `game_role_id` bigint NOT NULL DEFAULT 0,
  `game_role_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名',
  `game_role_level` int NOT NULL DEFAULT 0 COMMENT '角色级别',
  `days` date NULL DEFAULT NULL COMMENT '日期',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `mgzr`(`game_role_id` ASC, `create_time` ASC, `game_id` ASC, `zone_id` ASC) USING BTREE,
  INDEX `days`(`days` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '角色登录日志' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for log_player_online
-- ----------------------------
DROP TABLE IF EXISTS `log_player_online`;
CREATE TABLE `log_player_online`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `promote_code` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '推广码',
  `adid` varchar(64) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '广告id',
  `os` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '操作系统：android、ios',
  `device_id` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '设备号',
  `ip` char(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT 'IP ',
  `area_code` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区码',
  `area` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区',
  `create_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '事件时间，即激活时间',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `user_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `account` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL DEFAULT '' COMMENT '帐户名',
  `app_channel` int UNSIGNED NOT NULL DEFAULT 1 COMMENT '应用渠道：1(官方)、2(安卓)、3(苹果)',
  `game_project_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏项目ID',
  `game_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `game_name` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `zone_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '区服',
  `zone_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '区服名',
  `game_role_id` bigint NOT NULL DEFAULT 0,
  `game_role_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名',
  `game_role_level` int NOT NULL DEFAULT 0 COMMENT '角色级别',
  `online` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '在线时长单位为秒',
  `days` date NULL DEFAULT NULL COMMENT '日期',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `ugzgd`(`user_id` ASC, `zone_id` ASC, `game_role_id` ASC, `game_id` ASC, `days` ASC) USING BTREE,
  INDEX `days`(`days` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '角色在线时长日志' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for log_player_upgrade
-- ----------------------------
DROP TABLE IF EXISTS `log_player_upgrade`;
CREATE TABLE `log_player_upgrade`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `promote_code` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '推广码',
  `adid` varchar(64) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '广告id',
  `os` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '操作系统：android、ios',
  `device_id` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '设备号',
  `ip` char(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT 'IP ',
  `area_code` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区码',
  `area` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区',
  `create_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '事件时间，即激活时间',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `user_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `account` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL DEFAULT '' COMMENT '帐户名',
  `app_channel` int UNSIGNED NOT NULL DEFAULT 1 COMMENT '应用渠道：1(官方)、2(安卓)、3(苹果)',
  `game_project_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏项目ID',
  `game_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `game_name` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `zone_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '区服',
  `zone_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '区服名',
  `game_role_id` bigint NOT NULL DEFAULT 0,
  `game_role_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名',
  `game_role_level` int NOT NULL DEFAULT 0 COMMENT '角色级别',
  `days` date NULL DEFAULT NULL COMMENT '日期',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `ugzgd`(`user_id` ASC, `zone_id` ASC, `game_role_id` ASC, `game_id` ASC, `days` ASC) USING BTREE,
  INDEX `days`(`days` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '角色升级日志' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for log_register
-- ----------------------------
DROP TABLE IF EXISTS `log_register`;
CREATE TABLE `log_register`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `game_project_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏项目ID',
  `game_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `app_channel` int UNSIGNED NOT NULL DEFAULT 1 COMMENT '应用渠道：1(官方)、2(安卓)、3(苹果)',
  `os` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '操作系统：android、ios',
  `device_id` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '设备号',
  `ip` char(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT 'IP ',
  `area_code` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '城市',
  `area` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地区',
  `create_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '事件时间，即注册时间',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `media` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '媒体',
  `cause` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '广告归因依据',
  `cause_data` varchar(800) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '归因数据',
  `user_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `account` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL DEFAULT '' COMMENT '帐户名',
  `account_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '帐号类型：0  游客; 1 普通账号 ; 2 手机账号; 3  微信账号;4 机器人',
  `channel_id` int NOT NULL DEFAULT 0 COMMENT '媒体渠道',
  `campaign_id` bigint NOT NULL DEFAULT 0 COMMENT '广告计划',
  `adgroup_id` bigint NOT NULL DEFAULT 0 COMMENT '广告组',
  `creative_id` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '创意',
  `promote_code` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '推广码',
  `adid` varchar(64) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '广告id',
  `days` date NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `ug`(`user_id` ASC, `game_project_id` ASC) USING BTREE,
  INDEX `cause`(`cause` ASC) USING BTREE,
  INDEX `ts`(`create_time` ASC) USING BTREE,
  INDEX `days`(`days` ASC) USING BTREE,
  INDEX `ip`(`ip` ASC) USING BTREE,
  INDEX `device_id`(`device_id` ASC) USING BTREE,
  INDEX `prohect_adid`(`game_project_id` ASC, `adid` ASC) USING BTREE,
  INDEX `project_adid`(`game_project_id` ASC, `adid` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '玩家注册日志' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for log_role_pay_day
-- ----------------------------
DROP TABLE IF EXISTS `log_role_pay_day`;
CREATE TABLE `log_role_pay_day`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `adid` varchar(64) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '广告id',
  `user_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `account` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL DEFAULT '' COMMENT '帐户名',
  `pay_channel` varchar(20) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '支付编号',
  `pay_money` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付金额',
  `currency` varchar(45) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT 'USD' COMMENT '货币代码:USD(美元)、HKG(港元) 、MAC(澳门元) 、TWD(新台币) ...',
  `app_channel` int UNSIGNED NOT NULL DEFAULT 1 COMMENT '应用渠道：1(官方)、2(安卓)、3(苹果)',
  `game_project_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏项目ID',
  `game_id` mediumint UNSIGNED NOT NULL DEFAULT 0,
  `game_name` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '游戏名',
  `zone_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '0' COMMENT '区服',
  `zone_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '区服名',
  `game_role_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '0',
  `game_role_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名',
  `game_role_level` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '0' COMMENT '角色级别',
  `days` date NULL DEFAULT NULL COMMENT '日期',
  `reg_date` date NULL DEFAULT NULL COMMENT '注册日期',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `upggd`(`game_role_id` ASC, `game_id` ASC, `user_id` ASC, `days` ASC, `pay_channel` ASC) USING BTREE,
  INDEX `days`(`days` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '角色充值按天' ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
