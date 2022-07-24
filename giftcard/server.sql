
-- MySQL dump 10.13  Distrib 5.7.34, for osx10.15 (x86_64)
--
-- Host: 127.0.0.1    Database: giftcard
-- ------------------------------------------------------
-- Server version	5.7.35

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `gorp_migrations`
--

DROP TABLE IF EXISTS `gorp_migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `gorp_migrations` (
  `id` varchar(255) NOT NULL,
  `applied_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `gorp_migrations`
--

LOCK TABLES `gorp_migrations` WRITE;
/*!40000 ALTER TABLE `gorp_migrations` DISABLE KEYS */;
INSERT INTO `gorp_migrations` VALUES ('example_20220422173400.sql','2022-07-17 23:30:55');
/*!40000 ALTER TABLE `gorp_migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `migrations`
--

DROP TABLE IF EXISTS `migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `migrations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `batch` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `migrations`
--

LOCK TABLES `migrations` WRITE;
/*!40000 ALTER TABLE `migrations` DISABLE KEYS */;
INSERT INTO `migrations` VALUES (1,'2021_04_12_160526_create_system_user_table',1),(2,'2021_04_18_215320_create_system_menu_table',2),(3,'2021_04_18_215515_create_system_role_table',2),(4,'2021_04_18_215521_create_system_user_role_table',2),(5,'2021_04_18_222634_create_system_role_menu_table',2),(6,'2021_04_18_224723_create_system_dict_data_table',2),(7,'2021_04_18_224727_create_system_dict_type_table',2),(8,'2021_04_18_224817_create_system_dept_table',2),(9,'2021_04_18_224835_create_system_post_table',2),(10,'2021_04_18_224912_create_system_login_log_table',2),(11,'2021_04_18_224938_create_system_oper_log_table',2),(12,'2021_04_26_141249_create_system_user_post_table',2),(13,'2021_05_07_215228_create_system_role_dept_table',2),(14,'2021_06_24_111216_create_system_uploadfile_table',2),(15,'2021_11_03_223648_change_router_length',3),(16,'2021_11_07_231534_user_add_backend_setting',3),(17,'2021_11_11_140915_create_system_app_table',3),(18,'2021_11_11_140935_create_system_app_group_table',3),(19,'2021_11_11_141720_create_system_api_table',3),(20,'2021_11_11_141724_create_system_api_group_table',3),(21,'2021_11_11_141849_create_system_api_column_table',3),(22,'2021_11_11_151525_create_system_app_api_table',3),(23,'2021_11_26_163202_create_system_api_log_table',3),(24,'2021_11_26_163818_create_system_notice_table',3),(25,'2021_11_26_164006_create_system_queue_message_table',3),(26,'2021_12_20_163839_table_struct_version050',3),(27,'2021_12_25_163880_create_system_queue_log_table',3),(28,'2021_12_25_163890_create_system_queue_message_receive_table',3),(29,'2021_04_18_224626_create_setting_config_table',4),(30,'2021_04_18_225055_create_setting_crontab_table',4),(31,'2021_04_18_225100_create_setting_crontab_log_table',4),(32,'2021_04_18_225100_create_setting_generate_columns_table',4),(33,'2021_04_18_225100_create_setting_generate_tables_table',4),(34,'2022_02_20_223768_table_struct_version061',5),(35,'2022_04_09_223768_table_struct_version063',5),(36,'2022_04_15_223768_table_struct_version064',5);
/*!40000 ALTER TABLE `migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `setting_config`
--

DROP TABLE IF EXISTS `setting_config`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `setting_config` (
  `key` varchar(32) NOT NULL COMMENT '配置键名',
  `value` varchar(255) DEFAULT NULL COMMENT '配置值',
  `name` varchar(255) DEFAULT NULL COMMENT '配置名称',
  `group_name` varchar(100) DEFAULT NULL COMMENT '组名称',
  `sort` tinyint(3) unsigned DEFAULT '0' COMMENT '排序',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`),
  KEY `setting_config_group_name_index` (`group_name`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='参数配置信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `setting_config`
--

LOCK TABLES `setting_config` WRITE;
/*!40000 ALTER TABLE `setting_config` DISABLE KEYS */;
INSERT INTO `setting_config` VALUES ('site_copyright',NULL,'版权信息','system',96,NULL,1),('site_desc',NULL,'网站描述','system',97,NULL,2),('site_keywords',NULL,'网站关键字','system',98,NULL,3),('site_name',NULL,'网站名称','system',99,NULL,4),('site_record_number',NULL,'网站备案号','system',95,NULL,5),('site_storage_mode','local','上传存储模式','system',93,NULL,6),('web_close','1','网站是否关闭','extend',0,'关闭网站后将无法访问',7),('web_login_verify','0','后台登录验证码方式','extend',0,'0 前端验证，1 后端验证',8);
/*!40000 ALTER TABLE `setting_config` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `setting_crontab`
--

DROP TABLE IF EXISTS `setting_crontab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `setting_crontab` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `name` varchar(100) DEFAULT NULL COMMENT '任务名称',
  `type` char(1) DEFAULT '4' COMMENT '任务类型 (1 command, 2 class, 3 url, 4 eval)',
  `target` varchar(500) DEFAULT NULL COMMENT '调用任务字符串',
  `parameter` varchar(1000) DEFAULT NULL COMMENT '调用任务参数',
  `rule` varchar(32) DEFAULT NULL COMMENT '任务执行表达式',
  `singleton` char(1) DEFAULT '0' COMMENT '是否单次执行 (0 是 1 不是)',
  `status` char(1) DEFAULT '0' COMMENT '状态 (0正常 1停用)',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='定时任务信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `setting_crontab`
--

LOCK TABLES `setting_crontab` WRITE;
/*!40000 ALTER TABLE `setting_crontab` DISABLE KEYS */;
INSERT INTO `setting_crontab` VALUES (3890924964000,'urlCrontab','3','http://127.0.0.1:9501/','','59 */1 * * * *','1','1',NULL,NULL,'2021-08-07 15:28:28','2021-08-07 15:44:55','请求127.0.0.1'),(10794906676384,'每月1号清理所有日志','2','AppSystemCrontabClearLogCrontab','','0 0 0 1 * *','1','1',NULL,NULL,'2022-04-11 03:15:48','2022-04-11 03:15:48','');
/*!40000 ALTER TABLE `setting_crontab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `setting_crontab_log`
--

DROP TABLE IF EXISTS `setting_crontab_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `setting_crontab_log` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `crontab_id` bigint(20) unsigned NOT NULL COMMENT '任务ID',
  `name` varchar(255) DEFAULT NULL COMMENT '任务名称',
  `target` varchar(500) DEFAULT NULL COMMENT '任务调用目标字符串',
  `parameter` varchar(1000) DEFAULT NULL COMMENT '任务调用参数',
  `exception_info` varchar(2000) DEFAULT NULL COMMENT '异常信息',
  `status` char(1) DEFAULT '0' COMMENT '执行状态 (0成功 1失败)',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='定时任务执行日志表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `setting_crontab_log`
--

LOCK TABLES `setting_crontab_log` WRITE;
/*!40000 ALTER TABLE `setting_crontab_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `setting_crontab_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `setting_generate_columns`
--

DROP TABLE IF EXISTS `setting_generate_columns`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `setting_generate_columns` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `table_id` bigint(20) unsigned NOT NULL COMMENT '所属表ID',
  `column_name` varchar(200) DEFAULT NULL COMMENT '字段名称',
  `column_comment` varchar(255) DEFAULT NULL COMMENT '字段注释',
  `column_type` varchar(50) DEFAULT NULL COMMENT '字段类型',
  `is_pk` char(1) DEFAULT '0' COMMENT '0 非主键 1 主键',
  `is_required` char(1) DEFAULT '0' COMMENT '0 非必填 1 必填',
  `is_insert` char(1) DEFAULT '0' COMMENT '0 非插入字段 1 插入字段',
  `is_edit` char(1) DEFAULT '0' COMMENT '0 非编辑字段 1 编辑字段',
  `is_list` char(1) DEFAULT '0' COMMENT '0 非列表显示字段 1 列表显示字段',
  `is_query` char(1) DEFAULT '0' COMMENT '0 非查询字段 1 查询字段',
  `query_type` varchar(100) DEFAULT 'eq' COMMENT '查询方式 eq 等于, neq 不等于, gt 大于, lt 小于, like 范围',
  `view_type` varchar(100) DEFAULT 'text' COMMENT '页面控件，text, textarea, password, select, checkbox, radio, date, upload, ma-upload（封装的上传控件）',
  `dict_type` varchar(200) DEFAULT NULL COMMENT '字典类型',
  `allow_roles` varchar(255) DEFAULT NULL COMMENT '允许查看该字段的角色',
  `options` varchar(1000) DEFAULT NULL COMMENT '字段其他设置',
  `sort` tinyint(3) unsigned DEFAULT '0' COMMENT '排序',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='代码生成业务字段信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `setting_generate_columns`
--

LOCK TABLES `setting_generate_columns` WRITE;
/*!40000 ALTER TABLE `setting_generate_columns` DISABLE KEYS */;
/*!40000 ALTER TABLE `setting_generate_columns` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `setting_generate_tables`
--

DROP TABLE IF EXISTS `setting_generate_tables`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `setting_generate_tables` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `table_name` varchar(200) DEFAULT NULL COMMENT '表名称',
  `table_comment` varchar(500) DEFAULT NULL COMMENT '表注释',
  `module_name` varchar(100) DEFAULT NULL COMMENT '所属模块',
  `namespace` varchar(255) DEFAULT NULL COMMENT '命名空间',
  `menu_name` varchar(100) DEFAULT NULL COMMENT '生成菜单名',
  `belong_menu_id` bigint(20) DEFAULT NULL COMMENT '所属菜单',
  `package_name` varchar(100) DEFAULT NULL COMMENT '控制器包名',
  `type` varchar(100) DEFAULT NULL COMMENT '生成类型，single 单表CRUD，tree 树表CRUD，parent_sub父子表CRUD',
  `generate_type` char(1) DEFAULT '0' COMMENT '0 压缩包下载 1 生成到模块',
  `generate_menus` varchar(255) DEFAULT NULL COMMENT '生成菜单列表',
  `build_menu` char(1) DEFAULT '0' COMMENT '是否构建菜单',
  `component_type` char(1) DEFAULT '0' COMMENT '组件显示方式',
  `options` varchar(1500) DEFAULT NULL COMMENT '其他业务选项',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='代码生成业务信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `setting_generate_tables`
--

LOCK TABLES `setting_generate_tables` WRITE;
/*!40000 ALTER TABLE `setting_generate_tables` DISABLE KEYS */;
/*!40000 ALTER TABLE `setting_generate_tables` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_api`
--

DROP TABLE IF EXISTS `system_api`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_api` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `group_id` bigint(20) unsigned NOT NULL COMMENT '接口组ID',
  `name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '接口名称',
  `access_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '接口访问名称',
  `class_name` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '类命名空间',
  `method_name` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '方法名',
  `auth_mode` char(1) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '认证模式 (0简易 1复杂)',
  `request_mode` char(1) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'A' COMMENT '请求模式 (A 所有 P POST G GET)',
  `description` text COLLATE utf8mb4_unicode_ci COMMENT '接口说明介绍',
  `response` text COLLATE utf8mb4_unicode_ci COMMENT '返回内容示例',
  `status` char(1) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '状态 (0正常 1停用)',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `system_api_group_id_index` (`group_id`),
  KEY `system_api_access_name_index` (`access_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_api`
--

LOCK TABLES `system_api` WRITE;
/*!40000 ALTER TABLE `system_api` DISABLE KEYS */;
/*!40000 ALTER TABLE `system_api` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_api_column`
--

DROP TABLE IF EXISTS `system_api_column`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_api_column` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `api_id` bigint(20) unsigned NOT NULL COMMENT '接口主键',
  `name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '字段名称',
  `type` char(1) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '字段类型 0 请求 1 返回',
  `data_type` varchar(16) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '数据类型',
  `is_required` char(1) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '是否必填 0 非必填 1 必填',
  `default_value` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '默认值',
  `status` char(1) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '状态 (0正常 1停用)',
  `description` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '字段说明',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `system_api_column_api_id_type_status_index` (`api_id`,`type`,`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_api_column`
--

LOCK TABLES `system_api_column` WRITE;
/*!40000 ALTER TABLE `system_api_column` DISABLE KEYS */;
/*!40000 ALTER TABLE `system_api_column` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_api_group`
--

DROP TABLE IF EXISTS `system_api_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_api_group` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '接口组名称',
  `status` char(1) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '状态 (0正常 1停用)',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_api_group`
--

LOCK TABLES `system_api_group` WRITE;
/*!40000 ALTER TABLE `system_api_group` DISABLE KEYS */;
/*!40000 ALTER TABLE `system_api_group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_api_log`
--

DROP TABLE IF EXISTS `system_api_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_api_log` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `api_id` bigint(20) unsigned NOT NULL COMMENT 'api ID',
  `api_name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '接口名称',
  `access_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '接口访问名称',
  `request_data` text COLLATE utf8mb4_unicode_ci COMMENT '请求数据',
  `response_code` varchar(5) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '响应状态码',
  `response_data` text COLLATE utf8mb4_unicode_ci COMMENT '响应数据',
  `ip` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '访问IP地址',
  `ip_location` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'IP所属地',
  `access_time` timestamp NULL DEFAULT NULL COMMENT '访问时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `system_api_log_api_id_index` (`api_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_api_log`
--

LOCK TABLES `system_api_log` WRITE;
/*!40000 ALTER TABLE `system_api_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `system_api_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_app`
--

DROP TABLE IF EXISTS `system_app`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_app` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `group_id` bigint(20) unsigned NOT NULL COMMENT '应用组ID',
  `app_name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '应用名称',
  `app_id` varchar(16) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '应用ID',
  `app_secret` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '应用密钥',
  `status` char(1) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '状态 (0正常 1停用)',
  `description` text COLLATE utf8mb4_unicode_ci COMMENT '应用介绍',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `system_app_group_id_app_id_app_secret_index` (`group_id`,`app_id`,`app_secret`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_app`
--

LOCK TABLES `system_app` WRITE;
/*!40000 ALTER TABLE `system_app` DISABLE KEYS */;
/*!40000 ALTER TABLE `system_app` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_app_api`
--

DROP TABLE IF EXISTS `system_app_api`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_app_api` (
  `app_id` bigint(20) unsigned NOT NULL COMMENT '应用ID',
  `api_id` bigint(20) unsigned NOT NULL COMMENT 'API—ID',
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_app_api`
--

LOCK TABLES `system_app_api` WRITE;
/*!40000 ALTER TABLE `system_app_api` DISABLE KEYS */;
/*!40000 ALTER TABLE `system_app_api` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_app_group`
--

DROP TABLE IF EXISTS `system_app_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_app_group` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '应用组名称',
  `status` char(1) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '状态 (0正常 1停用)',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_app_group`
--

LOCK TABLES `system_app_group` WRITE;
/*!40000 ALTER TABLE `system_app_group` DISABLE KEYS */;
/*!40000 ALTER TABLE `system_app_group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_dept`
--

DROP TABLE IF EXISTS `system_dept`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_dept` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `parent_id` bigint(20) unsigned NOT NULL COMMENT '父ID',
  `level` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '组级集合',
  `name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '部门名称',
  `leader` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '联系电话',
  `status` char(1) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '状态 (0正常 1停用)',
  `sort` tinyint(3) unsigned DEFAULT '0' COMMENT '排序',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_dept`
--

LOCK TABLES `system_dept` WRITE;
/*!40000 ALTER TABLE `system_dept` DISABLE KEYS */;
INSERT INTO `system_dept` VALUES (1,0,'0','曼艺科技','曼艺','16888888888','0',0,19151769377440,NULL,'2022-07-18 07:07:41','2022-07-18 07:07:41',NULL,NULL);
/*!40000 ALTER TABLE `system_dept` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_dict_data`
--

DROP TABLE IF EXISTS `system_dict_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_dict_data` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `type_id` bigint(20) unsigned NOT NULL COMMENT '字典类型ID',
  `label` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '字典标签',
  `value` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '字典值',
  `code` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '字典标示',
  `sort` tinyint(3) unsigned DEFAULT '0' COMMENT '排序',
  `status` char(1) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '状态 (0正常 1停用)',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `system_dict_data_type_id_index` (`type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_dict_data`
--

LOCK TABLES `system_dict_data` WRITE;
/*!40000 ALTER TABLE `system_dict_data` DISABLE KEYS */;
INSERT INTO `system_dict_data` VALUES (2035090111136,2035075124896,'InnoDB','InnoDB','table_engine',0,'0',NULL,NULL,'2021-06-26 16:37:11','2021-06-27 05:33:29',NULL,NULL),(2035095441568,2035075124896,'MyISAM','MyISAM','table_engine',0,'0',NULL,NULL,'2021-06-26 16:37:21','2021-06-27 05:33:29',NULL,NULL),(2058945281696,2058928973472,'本地存储','local','upload_mode',99,'0',NULL,NULL,'2021-06-27 05:33:43','2021-06-27 05:33:43',NULL,NULL),(2058951566496,2058928973472,'七牛云','qiniu','upload_mode',98,'0',NULL,NULL,'2021-06-27 05:33:55','2021-06-27 05:33:55',NULL,NULL),(2058957471392,2058928973472,'阿里云OSS','oss','upload_mode',97,'0',NULL,NULL,'2021-06-27 05:34:07','2021-06-27 05:34:07',NULL,NULL),(2058963571360,2058928973472,'腾讯云COS','cos','upload_mode',96,'0',NULL,NULL,'2021-06-27 05:34:19','2021-06-27 05:34:19',NULL,NULL),(2059041780384,2059023428768,'正常','0','data_status',0,'0',NULL,NULL,'2021-06-27 05:36:51','2021-06-27 05:37:01',NULL,'0为正常'),(2059051223200,2059023428768,'停用','1','data_status',0,'0',NULL,NULL,'2021-06-27 05:37:10','2021-06-27 05:37:10',NULL,'1为停用'),(3959904132768,3959885616288,'统计页面','statistics','dashboard',0,'0',NULL,NULL,'2021-08-09 04:53:53','2021-08-09 04:53:53',NULL,'管理员用'),(3959916887200,3959885616288,'工作台','work','dashboard',0,'0',NULL,NULL,'2021-08-09 04:54:18','2021-08-09 04:54:18',NULL,'员工使用'),(3959938423968,3959928216736,'男','0','sex',0,'0',NULL,NULL,'2021-08-09 04:55:00','2021-08-09 04:55:00',NULL,NULL),(3959942491808,3959928216736,'女','1','sex',0,'0',NULL,NULL,'2021-08-09 04:55:08','2021-08-09 04:55:08',NULL,NULL),(3959946307744,3959928216736,'未知','2','sex',0,'0',NULL,NULL,'2021-08-09 04:55:16','2021-08-09 04:55:16',NULL,NULL),(8126628581024,8126617434272,'通知','1','backend_notice_type',2,'0',NULL,NULL,'2021-11-11 09:29:27','2021-11-11 09:30:51',NULL,NULL),(8126697847456,8126617434272,'公告','2','backend_notice_type',1,'0',NULL,NULL,'2021-11-11 09:31:42','2021-11-11 09:31:42',NULL,NULL),(8259153751712,8259136986784,'所有','A','request_mode',5,'0',NULL,NULL,'2021-11-14 09:23:25','2021-11-14 09:23:25',NULL,NULL),(8259160266912,8259136986784,'POST','P','request_mode',3,'0',NULL,NULL,'2021-11-14 09:23:38','2021-11-14 09:23:38',NULL,NULL),(8259163719840,8259136986784,'GET','G','request_mode',4,'0',NULL,NULL,'2021-11-14 09:23:45','2021-11-14 09:23:45',NULL,NULL),(8259163719842,8259136986784,'PUT','U','request_mode',2,'0',NULL,NULL,'2021-11-14 09:23:45','2021-11-14 09:23:45',NULL,NULL),(8259163719843,8259136986784,'DELETE','D','request_mode',1,'0',NULL,NULL,'2021-11-14 09:23:45','2021-11-14 09:23:45',NULL,NULL),(8645168257696,8619580222112,'String','1','api_data_type',7,'0',NULL,NULL,'2021-11-23 02:49:00','2021-11-23 02:49:00',NULL,NULL),(8645183473312,8619580222112,'Integer','2','api_data_type',6,'0',NULL,NULL,'2021-11-23 02:49:29','2021-11-23 02:49:29',NULL,NULL),(8645187911328,8619580222112,'Array','3','api_data_type',5,'0',NULL,NULL,'2021-11-23 02:49:38','2021-11-23 02:49:38',NULL,NULL),(8645192008352,8619580222112,'Float','4','api_data_type',4,'0',NULL,NULL,'2021-11-23 02:49:46','2021-11-23 02:49:46',NULL,NULL),(8645196093088,8619580222112,'Boolean','5','api_data_type',3,'0',NULL,NULL,'2021-11-23 02:49:54','2021-11-23 02:49:54',NULL,NULL),(8645207885984,8619580222112,'Enum','6','api_data_type',2,'0',NULL,NULL,'2021-11-23 02:50:17','2021-11-23 02:50:17',NULL,NULL),(8645212278944,8619580222112,'Object','7','api_data_type',1,'0',NULL,NULL,'2021-11-23 02:50:26','2021-11-23 02:50:26',NULL,NULL),(8645212278946,8619580222112,'File','8','api_data_type',0,'0',NULL,NULL,'2021-12-25 10:32:48','2021-12-25 10:32:48',NULL,NULL),(10074768599200,10074681620640,'未生产','0','queue_produce_status',5,'0',NULL,NULL,'2021-12-25 10:25:28','2021-12-25 10:25:28',NULL,NULL),(10074773797536,10074681620640,'生产中','1','queue_produce_status',4,'0',NULL,NULL,'2021-12-25 10:25:38','2021-12-25 10:25:38',NULL,NULL),(10074780033696,10074681620640,'生产成功','2','queue_produce_status',3,'0',NULL,NULL,'2021-12-25 10:25:50','2021-12-25 10:25:50',NULL,NULL),(10074792360608,10074681620640,'生产失败','3','queue_produce_status',2,'0',NULL,NULL,'2021-12-25 10:26:14','2021-12-25 10:26:14',NULL,NULL),(10074800282272,10074681620640,'生产重复','4','queue_produce_status',1,'0',NULL,NULL,'2021-12-25 10:26:30','2021-12-25 10:26:30',NULL,NULL),(10074814135968,10074702532768,'未消费','0','queue_consume_status',5,'0',NULL,NULL,'2021-12-25 10:26:57','2021-12-25 10:26:57',NULL,NULL),(10074819077792,10074702532768,'消费中','1','queue_consume_status',4,'0',NULL,NULL,'2021-12-25 10:27:07','2021-12-25 10:27:07',NULL,NULL),(10074823951520,10074702532768,'消费成功','2','queue_consume_status',3,'0',NULL,NULL,'2021-12-25 10:27:16','2021-12-25 10:27:16',NULL,NULL),(10074828146848,10074702532768,'消费失败','3','queue_consume_status',2,'0',NULL,NULL,'2021-12-25 10:27:24','2021-12-25 10:27:24',NULL,NULL),(10074833623200,10074702532768,'消费重复','4','queue_consume_status',1,'0',NULL,NULL,'2021-12-25 10:27:35','2021-12-25 10:27:35',NULL,NULL),(10074923844256,10074866778272,'通知','notice','queue_msg_type',1,'0',NULL,NULL,'2021-12-25 10:30:31','2021-12-25 10:30:31',NULL,NULL),(10074938784928,10074866778272,'公告','announcement','queue_msg_type',2,'0',NULL,NULL,'2021-12-25 10:31:00','2021-12-25 10:31:00',NULL,NULL),(10074951884448,10074866778272,'待办','todo','queue_msg_type',3,'0',NULL,NULL,'2021-12-25 10:31:26','2021-12-25 10:31:26',NULL,NULL),(10074951884459,10074866778272,'抄送我的','carbon_copy_mine','queue_msg_type',4,'0',NULL,NULL,'2021-12-25 10:31:26','2021-12-25 10:31:26',NULL,NULL),(10074951884666,10074866778272,'私信','private_message','queue_msg_type',5,'0',NULL,NULL,'2021-12-25 10:31:26','2021-12-25 10:31:26',NULL,NULL),(10695566593184,10774866789575,'图片','image','attachment_type',10,'0',NULL,NULL,'2022-03-17 06:49:59','2022-03-17 06:49:59',NULL,NULL),(10695577459872,10774866789575,'文档','text','attachment_type',9,'0',NULL,NULL,'2022-03-17 06:50:20','2022-03-17 06:50:49',NULL,NULL),(10695585910944,10774866789575,'音频','audio','attachment_type',8,'0',NULL,NULL,'2022-03-17 06:50:37','2022-03-17 06:50:52',NULL,NULL),(10695590140064,10774866789575,'视频','video','attachment_type',7,'0',NULL,NULL,'2022-03-17 06:50:45','2022-03-17 06:50:57',NULL,NULL);
/*!40000 ALTER TABLE `system_dict_data` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_dict_type`
--

DROP TABLE IF EXISTS `system_dict_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_dict_type` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '字典名称',
  `code` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '字典标示',
  `status` char(1) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '状态 (0正常 1停用)',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_dict_type`
--

LOCK TABLES `system_dict_type` WRITE;
/*!40000 ALTER TABLE `system_dict_type` DISABLE KEYS */;
INSERT INTO `system_dict_type` VALUES (2035075124896,'数据表引擎','table_engine','0',NULL,NULL,'2021-06-26 16:36:42','2021-06-27 05:33:29',NULL,'数据表引擎字典'),(2058928973472,'存储模式','upload_mode','0',NULL,NULL,'2021-06-27 05:33:11','2021-06-27 05:33:11',NULL,'上传文件存储模式'),(2059023428768,'数据状态','data_status','0',NULL,NULL,'2021-06-27 05:36:16','2021-06-27 05:36:34',NULL,'通用数据状态'),(3959885616288,'后台首页','dashboard','0',NULL,NULL,'2021-08-09 04:53:17','2021-08-09 04:53:17',NULL,NULL),(3959928216736,'性别','sex','0',NULL,NULL,'2021-08-09 04:54:40','2021-08-09 04:54:40',NULL,NULL),(8126617434272,'后台公告类型','backend_notice_type','0',NULL,NULL,'2021-11-11 09:29:05','2021-11-11 09:29:14',NULL,NULL),(8259136986784,'请求方式','request_mode','0',NULL,NULL,'2021-11-14 09:22:52','2021-11-14 09:22:52',NULL,NULL),(8619580222112,'接口数据类型','api_data_type','0',NULL,NULL,'2021-11-22 12:56:03','2021-11-22 12:56:03',NULL,NULL),(10074681620640,'队列生产状态','queue_produce_status','0',NULL,NULL,'2021-12-25 10:22:38','2021-12-25 10:22:38',NULL,NULL),(10074702532768,'队列消费状态','queue_consume_status','0',NULL,NULL,'2021-12-25 10:23:19','2021-12-25 10:23:19',NULL,NULL),(10074866778272,'队列消息类型','queue_msg_type','0',NULL,NULL,'2021-12-25 10:28:40','2021-12-25 10:28:40',NULL,NULL),(10074866789575,'附件类型','attachment_type','0',NULL,NULL,'2022-03-17 06:49:23','2022-03-17 06:49:23',NULL,NULL);
/*!40000 ALTER TABLE `system_dict_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_login_log`
--

DROP TABLE IF EXISTS `system_login_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_login_log` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `username` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `ip` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '登录IP地址',
  `ip_location` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'IP所属地',
  `os` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '操作系统',
  `browser` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '浏览器',
  `status` char(1) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '登录状态 (0成功 1失败)',
  `message` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '提示消息',
  `login_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '登录时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_login_log`
--

LOCK TABLES `system_login_log` WRITE;
/*!40000 ALTER TABLE `system_login_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `system_login_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_menu`
--

DROP TABLE IF EXISTS `system_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_menu` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `parent_id` bigint(20) unsigned NOT NULL COMMENT '父ID',
  `level` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '组级集合',
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜单名称',
  `code` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜单标识代码',
  `icon` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '菜单图标',
  `route` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '路由地址',
  `component` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '组件路径',
  `redirect` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '跳转地址',
  `is_hidden` char(1) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '1' COMMENT '是否隐藏 (0是 1否)',
  `type` char(1) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单类型, (M菜单 B按钮 L链接 I iframe)',
  `status` char(1) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '状态 (0正常 1停用)',
  `sort` tinyint(3) unsigned DEFAULT '0' COMMENT '排序',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_menu`
--

LOCK TABLES `system_menu` WRITE;
/*!40000 ALTER TABLE `system_menu` DISABLE KEYS */;
INSERT INTO `system_menu` VALUES (1000,0,'0','权限管理','permission','ma-icon-permission','permission','',NULL,'1','M','0',99,NULL,NULL,'2021-07-25 10:48:47','2021-07-25 10:48:47',NULL,NULL),(1100,1000,'0,1000','用户管理','system:user','ma-icon-user','user','system/user/index',NULL,'1','M','0',99,NULL,NULL,'2021-07-25 10:50:15','2021-07-25 10:50:15',NULL,NULL),(1101,1100,'0,1000,1100','用户列表','system:user:index',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 10:50:15','2021-07-25 10:50:15',NULL,NULL),(1102,1100,'0,1000,1100','用户回收站列表','system:user:recycle',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 10:50:15','2021-07-25 10:50:15',NULL,NULL),(1103,1100,'0,1000,1100','用户保存','system:user:save',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 10:50:15','2021-07-25 10:50:15',NULL,NULL),(1104,1100,'0,1000,1100','用户更新','system:user:update',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 10:50:15','2021-07-25 10:50:15',NULL,NULL),(1105,1100,'0,1000,1100','用户删除','system:user:delete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 10:50:15','2021-07-25 10:50:15',NULL,NULL),(1106,1100,'0,1000,1100','用户读取','system:user:read',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 10:50:15','2021-07-25 10:50:15',NULL,NULL),(1107,1100,'0,1000,1100','用户恢复','system:user:recovery',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 10:50:15','2021-07-25 10:50:15',NULL,NULL),(1108,1100,'0,1000,1100','用户真实删除','system:user:realDelete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 10:50:15','2021-07-25 10:50:15',NULL,NULL),(1109,1100,'0,1000,1100','用户导入','system:user:import',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 10:50:15','2021-07-25 10:50:15',NULL,NULL),(1110,1100,'0,1000,1100','用户导出','system:user:export',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 10:50:15','2021-07-25 10:50:15',NULL,NULL),(1111,1100,'0,1000,1100','用户状态改变','system:user:changeStatus','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-07-25 10:53:02','2021-07-25 10:53:02',NULL,NULL),(1112,1100,'0,1000,1100','用户初始化密码','system:user:initUserPassword','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-07-25 10:55:55','2021-07-25 10:55:55',NULL,NULL),(1113,1100,'0,1000,1100','更新用户缓存','system:user:cache','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-08-08 10:30:57','2021-08-08 10:30:57',NULL,NULL),(1114,1100,'0,1000,1100','设置用户首页','system:user:homePage','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-08-08 10:34:30','2021-08-08 10:34:30',NULL,NULL),(1200,1000,'0,1000','菜单管理','system:menu','el-icon-menu','menu','system/menu/index',NULL,'1','M','0',96,NULL,NULL,'2021-07-25 11:01:47','2021-07-25 11:01:47',NULL,NULL),(1201,1200,'0,1000,1200','菜单列表','system:menu:index',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:01:47','2021-07-25 11:01:47',NULL,NULL),(1202,1200,'0,1000,1200','菜单回收站','system:menu:recycle',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:01:47','2021-07-25 11:01:47',NULL,NULL),(1203,1200,'0,1000,1200','菜单保存','system:menu:save',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:01:47','2021-07-25 11:01:47',NULL,NULL),(1204,1200,'0,1000,1200','菜单更新','system:menu:update',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:01:47','2021-07-25 11:01:47',NULL,NULL),(1205,1200,'0,1000,1200','菜单删除','system:menu:delete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:01:47','2021-07-25 11:01:47',NULL,NULL),(1206,1200,'0,1000,1200','菜单读取','system:menu:read',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:01:47','2021-07-25 11:01:47',NULL,NULL),(1207,1200,'0,1000,1200','菜单恢复','system:menu:recovery',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:01:47','2021-07-25 11:01:47',NULL,NULL),(1208,1200,'0,1000,1200','菜单真实删除','system:menu:realDelete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:01:47','2021-07-25 11:01:47',NULL,NULL),(1209,1200,'0,1000,1200','菜单导入','system:menu:import',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:01:47','2021-07-25 11:01:47',NULL,NULL),(1210,1200,'0,1000,1200','菜单导出','system:menu:export',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:01:47','2021-07-25 11:01:47',NULL,NULL),(1300,1000,'0,1000','部门管理','system:dept','ma-icon-dept','dept','system/dept/index',NULL,'1','M','0',97,NULL,NULL,'2021-07-25 11:16:33','2021-07-25 11:16:33',NULL,NULL),(1301,1300,'0,1000,1300','部门列表','system:dept:index',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:16:33','2021-07-25 11:16:33',NULL,NULL),(1302,1300,'0,1000,1300','部门回收站','system:dept:recycle',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:16:33','2021-07-25 11:16:33',NULL,NULL),(1303,1300,'0,1000,1300','部门保存','system:dept:save',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:16:33','2021-07-25 11:16:33',NULL,NULL),(1304,1300,'0,1000,1300','部门更新','system:dept:update',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:16:33','2021-07-25 11:16:33',NULL,NULL),(1305,1300,'0,1000,1300','部门删除','system:dept:delete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:16:33','2021-07-25 11:16:33',NULL,NULL),(1306,1300,'0,1000,1300','部门读取','system:dept:read',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:16:33','2021-07-25 11:16:33',NULL,NULL),(1307,1300,'0,1000,1300','部门恢复','system:dept:recovery',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:16:33','2021-07-25 11:16:33',NULL,NULL),(1308,1300,'0,1000,1300','部门真实删除','system:dept:realDelete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:16:33','2021-07-25 11:16:33',NULL,NULL),(1309,1300,'0,1000,1300','部门导入','system:dept:import',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:16:33','2021-07-25 11:16:33',NULL,NULL),(1310,1300,'0,1000,1300','部门导出','system:dept:export',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:16:33','2021-07-25 11:16:33',NULL,NULL),(1311,1300,'0,1000,1300','部门状态改变','system:dept:changeStatus','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-11-09 10:26:15','2021-11-09 10:26:15',NULL,NULL),(1400,1000,'0,1000','角色管理','system:role','ma-icon-role','role','system/role/index',NULL,'1','M','0',98,NULL,NULL,'2021-07-25 11:17:37','2021-07-25 11:17:37',NULL,NULL),(1401,1400,'0,1000,1400','角色列表','system:role:index',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:17:37','2021-07-25 11:17:37',NULL,NULL),(1402,1400,'0,1000,1400','角色回收站','system:role:recycle',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:17:38','2021-07-25 11:17:38',NULL,NULL),(1403,1400,'0,1000,1400','角色保存','system:role:save',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:17:38','2021-07-25 11:17:38',NULL,NULL),(1404,1400,'0,1000,1400','角色更新','system:role:update',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:17:38','2021-07-25 11:17:38',NULL,NULL),(1405,1400,'0,1000,1400','角色删除','system:role:delete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:17:38','2021-07-25 11:17:38',NULL,NULL),(1406,1400,'0,1000,1400','角色读取','system:role:read',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:17:38','2021-07-25 11:17:38',NULL,NULL),(1407,1400,'0,1000,1400','角色恢复','system:role:recovery',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:17:38','2021-07-25 11:17:38',NULL,NULL),(1408,1400,'0,1000,1400','角色真实删除','system:role:realDelete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:17:38','2021-07-25 11:17:38',NULL,NULL),(1409,1400,'0,1000,1400','角色导入','system:role:import',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:17:38','2021-07-25 11:17:38',NULL,NULL),(1410,1400,'0,1000,1400','角色导出','system:role:export',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 11:17:38','2021-07-25 11:17:38',NULL,NULL),(1411,1400,'0,1000,1400','角色状态改变','system:role:changeStatus','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-07-30 03:21:24','2021-07-30 03:21:24',NULL,NULL),(1412,1400,'0,1000,1400','更新菜单权限','system:role:menuPermission','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-08-09 03:52:33','2021-08-09 03:52:33',NULL,NULL),(1413,1400,'0,1000,1400','更新数据权限','system:role:dataPermission','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-08-09 03:52:52','2021-08-09 03:52:52',NULL,NULL),(1500,1000,'0,1000','岗位管理','system:post','ma-icon-post','post','system/post/index',NULL,'1','M','0',0,NULL,NULL,'2021-07-25 12:46:38','2021-07-25 12:46:38',NULL,NULL),(1501,1500,'0,1000,1500','岗位列表','system:post:index',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 12:46:38','2021-07-25 12:46:38',NULL,NULL),(1502,1500,'0,1000,1500','岗位回收站','system:post:recycle',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 12:46:38','2021-07-25 12:46:38',NULL,NULL),(1503,1500,'0,1000,1500','岗位保存','system:post:save',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 12:46:38','2021-07-25 12:46:38',NULL,NULL),(1504,1500,'0,1000,1500','岗位更新','system:post:update',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 12:46:38','2021-07-25 12:46:38',NULL,NULL),(1505,1500,'0,1000,1500','岗位删除','system:post:delete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 12:46:38','2021-07-25 12:46:38',NULL,NULL),(1506,1500,'0,1000,1500','岗位读取','system:post:read',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 12:46:38','2021-07-25 12:46:38',NULL,NULL),(1507,1500,'0,1000,1500','岗位恢复','system:post:recovery',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 12:46:38','2021-07-25 12:46:38',NULL,NULL),(1508,1500,'0,1000,1500','岗位真实删除','system:post:realDelete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 12:46:38','2021-07-25 12:46:38',NULL,NULL),(1509,1500,'0,1000,1500','岗位导入','system:post:import',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 12:46:38','2021-07-25 12:46:38',NULL,NULL),(1510,1500,'0,1000,1500','岗位导出','system:post:export',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-25 12:46:38','2021-07-25 12:46:38',NULL,NULL),(1511,1500,'0,1000,1500','岗位状态改变','system:post:changeStatus','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-11-09 10:26:15','2021-11-09 10:26:15',NULL,NULL),(2000,0,'0','数据中心','dataCenter','el-icon-coin','dataCenter','',NULL,'1','M','0',98,NULL,NULL,'2021-07-31 09:17:03','2021-07-31 09:17:03',NULL,NULL),(2100,2000,'0,2000','数据字典','system:dataDict','ma-icon-dict','dataDict','system/dataDict/index',NULL,'1','M','0',99,NULL,NULL,'2021-07-31 09:23:58','2021-07-31 09:23:58',NULL,NULL),(2101,2100,'0,2000,2100','数据字典列表','system:dataDict:index',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 09:23:58','2021-07-31 09:23:58',NULL,NULL),(2102,2100,'0,2000,2100','数据字典回收站','system:dataDict:recycle',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 09:23:58','2021-07-31 09:23:58',NULL,NULL),(2103,2100,'0,2000,2100','数据字典保存','system:dataDict:save',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 09:23:58','2021-07-31 09:23:58',NULL,NULL),(2104,2100,'0,2000,2100','数据字典更新','system:dataDict:update',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 09:23:58','2021-07-31 09:23:58',NULL,NULL),(2105,2100,'0,2000,2100','数据字典删除','system:dataDict:delete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 09:23:58','2021-07-31 09:23:58',NULL,NULL),(2106,2100,'0,2000,2100','数据字典读取','system:dataDict:read',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 09:23:58','2021-07-31 09:23:58',NULL,NULL),(2107,2100,'0,2000,2100','数据字典恢复','system:dataDict:recovery',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 09:23:58','2021-07-31 09:23:58',NULL,NULL),(2108,2100,'0,2000,2100','数据字典真实删除','system:dataDict:realDelete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 09:23:58','2021-07-31 09:23:58',NULL,NULL),(2109,2100,'0,2000,2100','数据字典导入','system:dataDict:import',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 09:23:58','2021-07-31 09:23:58',NULL,NULL),(2110,2100,'0,2000,2100','数据字典导出','system:dataDict:export',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 09:23:58','2021-07-31 09:23:58',NULL,NULL),(2111,2100,'0,2000,2100','清除字典缓存','system:dataDict:clearCache','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-08-10 05:14:32','2021-08-10 05:14:32',NULL,NULL),(2112,2100,'0,2000,2100','字典状态改变','system:dataDict:changeStatus','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-11-09 10:26:15','2021-11-09 10:26:15',NULL,NULL),(2200,2000,'0,2000','字典类型','system:dictType','','dictType','',NULL,'0','M','0',0,NULL,NULL,'2021-07-31 10:33:45','2021-07-31 10:33:45',NULL,NULL),(2201,2200,'0,2000,2200','字典类型列表','system:dictType:index',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 10:33:45','2021-07-31 10:33:45',NULL,NULL),(2202,2200,'0,2000,2200','字典类型回收站','system:dictType:recycle',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 10:33:45','2021-07-31 10:33:45',NULL,NULL),(2203,2200,'0,2000,2200','字典类型保存','system:dictType:save',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 10:33:45','2021-07-31 10:33:45',NULL,NULL),(2204,2200,'0,2000,2200','字典类型更新','system:dictType:update',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 10:33:45','2021-07-31 10:33:45',NULL,NULL),(2205,2200,'0,2000,2200','字典类型删除','system:dictType:delete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 10:33:45','2021-07-31 10:33:45',NULL,NULL),(2206,2200,'0,2000,2200','字典类型读取','system:dictType:read',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 10:33:46','2021-07-31 10:33:46',NULL,NULL),(2207,2200,'0,2000,2200','字典类型恢复','system:dictType:recovery',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 10:33:46','2021-07-31 10:33:46',NULL,NULL),(2208,2200,'0,2000,2200','字典类型真实删除','system:dictType:realDelete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 10:33:46','2021-07-31 10:33:46',NULL,NULL),(2209,2200,'0,2000,2200','字典类型导入','system:dictType:import',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 10:33:46','2021-07-31 10:33:46',NULL,NULL),(2210,2200,'0,2000,2200','字典类型导出','system:dictType:export',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 10:33:46','2021-07-31 10:33:46',NULL,NULL),(2300,2000,'0,2000','附件管理','system:attachment','ma-icon-attach','attachment','system/attachment/index',NULL,'1','M','0',98,NULL,NULL,'2021-07-31 10:36:34','2021-07-31 10:36:34',NULL,NULL),(2301,2300,'0,2000,2300','附件删除','system:attachment:delete','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-07-31 10:37:20','2021-07-31 10:37:20',NULL,NULL),(2302,2300,'0,2000,2300','附件列表','system:attachment:index','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-07-31 10:38:05','2021-07-31 10:38:05',NULL,NULL),(2303,2300,'0,2000,2300','附件回收站','system:attachment:recycle','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-07-31 10:38:57','2021-07-31 10:38:57',NULL,NULL),(2304,2300,'0,2000,2300','附件真实删除','system:attachment:realDelete','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-07-31 10:40:44','2021-07-31 10:40:44',NULL,NULL),(2400,2000,'0,2000','数据表维护','system:dataMaintain','ma-icon-db','dataMaintain','system/dataMaintain/index',NULL,'1','M','0',95,NULL,NULL,'2021-07-31 10:43:20','2021-07-31 10:43:20',NULL,NULL),(2401,2400,'0,2000,2400','数据表列表','system:dataMaintain:index','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-07-31 10:44:03','2021-07-31 10:44:03',NULL,NULL),(2402,2400,'0,2000,2400','数据表详细','system:dataMaintain:detailed','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-07-31 10:45:17','2021-07-31 10:45:17',NULL,NULL),(2403,2400,'0,2000,2400','数据表清理碎片','system:dataMaintain:fragment','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-07-31 10:46:04','2021-07-31 10:46:04',NULL,NULL),(2404,2400,'0,2000,2400','数据表优化','system:dataMaintain:optimize','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-07-31 10:46:31','2021-07-31 10:46:31',NULL,NULL),(2500,2000,'0,2000','应用中心','apps','el-icon-goods','apps','',NULL,'1','M','0',90,NULL,NULL,'2021-11-11 13:16:47','2021-11-11 13:16:47',NULL,NULL),(2510,2500,'0,2000,2500','应用分组','system:appGroup','ma-icon-group','appGroup','system/appGroup/index',NULL,'1','M','0',0,NULL,NULL,'2021-11-11 13:31:28','2021-11-11 13:31:28',NULL,NULL),(2511,2510,'0,2000,2500,2510','应用分组列表','system:appGroup:index',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:31:28','2021-11-11 13:31:28',NULL,NULL),(2512,2510,'0,2000,2500,2510','应用分组回收站','system:appGroup:recycle',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:31:28','2021-11-11 13:31:28',NULL,NULL),(2513,2510,'0,2000,2500,2510','应用分组保存','system:appGroup:save',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:31:28','2021-11-11 13:31:28',NULL,NULL),(2514,2510,'0,2000,2500,2510','应用分组更新','system:appGroup:update',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:31:28','2021-11-11 13:31:28',NULL,NULL),(2515,2510,'0,2000,2500,2510','应用分组删除','system:appGroup:delete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:31:28','2021-11-11 13:31:28',NULL,NULL),(2516,2510,'0,2000,2500,2510','应用分组读取','system:appGroup:read',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:31:28','2021-11-11 13:31:28',NULL,NULL),(2517,2510,'0,2000,2500,2510','应用分组恢复','system:appGroup:recovery',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:31:28','2021-11-11 13:31:28',NULL,NULL),(2518,2510,'0,2000,2500,2510','应用分组真实删除','system:appGroup:realDelete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:31:28','2021-11-11 13:31:28',NULL,NULL),(2519,2510,'0,2000,2500,2510','应用分组导入','system:appGroup:import',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:31:28','2021-11-11 13:31:28',NULL,NULL),(2520,2510,'0,2000,2500,2510','应用分组导出','system:appGroup:export',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:31:28','2021-11-11 13:31:28',NULL,NULL),(2530,2500,'0,2000,2500','应用管理','system:app','el-icon-goods','app','system/app/index',NULL,'1','M','0',0,NULL,NULL,'2021-11-11 13:35:35','2021-11-11 13:35:35',NULL,NULL),(2531,2530,'0,2000,2500,2530','应用列表','system:app:index',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:35:35','2021-11-11 13:35:35',NULL,NULL),(2532,2530,'0,2000,2500,2530','应用回收站','system:app:recycle',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:35:35','2021-11-11 13:35:35',NULL,NULL),(2533,2530,'0,2000,2500,2530','应用保存','system:app:save',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:35:35','2021-11-11 13:35:35',NULL,NULL),(2534,2530,'0,2000,2500,2530','应用更新','system:app:update',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:35:35','2021-11-11 13:35:35',NULL,NULL),(2535,2530,'0,2000,2500,2530','应用删除','system:app:delete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:35:35','2021-11-11 13:35:35',NULL,NULL),(2536,2530,'0,2000,2500,2530','应用读取','system:app:read',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:35:35','2021-11-11 13:35:35',NULL,NULL),(2537,2530,'0,2000,2500,2530','应用恢复','system:app:recovery',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:35:35','2021-11-11 13:35:35',NULL,NULL),(2538,2530,'0,2000,2500,2530','应用真实删除','system:app:realDelete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:35:35','2021-11-11 13:35:35',NULL,NULL),(2539,2530,'0,2000,2500,2530','应用导入','system:app:import',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:35:35','2021-11-11 13:35:35',NULL,NULL),(2540,2530,'0,2000,2500,2530','应用导出','system:app:export',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:35:35','2021-11-11 13:35:35',NULL,NULL),(2541,2530,'0,2000,2500,2530','应用绑定接口','system:app:bind','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:44:24','2021-11-11 13:44:24',NULL,NULL),(2600,2000,'0,2000','应用接口','apis','el-icon-set-up','apis','',NULL,'1','M','0',80,NULL,NULL,'2021-11-11 13:23:57','2021-11-11 13:23:57',NULL,NULL),(2610,2600,'0,2000,2600','接口分组','system:apiGroup','ma-icon-group','apiGroup','system/apiGroup/index',NULL,'1','M','0',0,NULL,NULL,'2021-11-11 13:37:54','2021-11-11 13:37:54',NULL,NULL),(2611,2610,'0,2000,2600,2610','接口分组列表','system:apiGroup:index',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:37:54','2021-11-11 13:37:54',NULL,NULL),(2612,2610,'0,2000,2600,2610','接口分组回收站','system:apiGroup:recycle',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:37:54','2021-11-11 13:37:54',NULL,NULL),(2613,2610,'0,2000,2600,2610','接口分组保存','system:apiGroup:save',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:37:54','2021-11-11 13:37:54',NULL,NULL),(2614,2610,'0,2000,2600,2610','接口分组更新','system:apiGroup:update',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:37:54','2021-11-11 13:37:54',NULL,NULL),(2615,2610,'0,2000,2600,2610','接口分组删除','system:apiGroup:delete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:37:54','2021-11-11 13:37:54',NULL,NULL),(2616,2610,'0,2000,2600,2610','接口分组读取','system:apiGroup:read',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:37:54','2021-11-11 13:37:54',NULL,NULL),(2617,2610,'0,2000,2600,2610','接口分组恢复','system:apiGroup:recovery',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:37:54','2021-11-11 13:37:54',NULL,NULL),(2618,2610,'0,2000,2600,2610','接口分组真实删除','system:apiGroup:realDelete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:37:54','2021-11-11 13:37:54',NULL,NULL),(2619,2610,'0,2000,2600,2610','接口分组导入','system:apiGroup:import',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:37:54','2021-11-11 13:37:54',NULL,NULL),(2620,2610,'0,2000,2600,2610','接口分组导出','system:apiGroup:export',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:37:54','2021-11-11 13:37:54',NULL,NULL),(2630,2600,'0,2000,2600','接口管理','system:api','el-icon-files','api','system/api/index',NULL,'1','M','0',0,NULL,NULL,'2021-11-11 13:44:24','2021-11-11 13:44:24',NULL,NULL),(2631,2630,'0,2000,2600,2630','接口列表','system:api:index',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:44:24','2021-11-11 13:44:24',NULL,NULL),(2632,2630,'0,2000,2600,2630','接口回收站','system:api:recycle',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:44:24','2021-11-11 13:44:24',NULL,NULL),(2633,2630,'0,2000,2600,2630','接口保存','system:api:save',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:44:24','2021-11-11 13:44:24',NULL,NULL),(2634,2630,'0,2000,2600,2630','接口更新','system:api:update',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:44:24','2021-11-11 13:44:24',NULL,NULL),(2635,2630,'0,2000,2600,2630','接口删除','system:api:delete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:44:24','2021-11-11 13:44:24',NULL,NULL),(2636,2630,'0,2000,2600,2630','接口读取','system:api:read',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:44:24','2021-11-11 13:44:24',NULL,NULL),(2637,2630,'0,2000,2600,2630','接口恢复','system:api:recovery',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:44:24','2021-11-11 13:44:24',NULL,NULL),(2638,2630,'0,2000,2600,2630','接口真实删除','system:api:realDelete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:44:24','2021-11-11 13:44:24',NULL,NULL),(2639,2630,'0,2000,2600,2630','接口导入','system:api:import',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:44:24','2021-11-11 13:44:24',NULL,NULL),(2640,2630,'0,2000,2600,2630','接口导出','system:api:export',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:44:24','2021-11-11 13:44:24',NULL,NULL),(2650,2600,'0,2000,2600','接口字段管理','system:apiColumn','el-icon-menu','apiColumn','system/apiColumn/index',NULL,'0','M','0',0,NULL,NULL,'2021-11-11 13:46:25','2021-11-11 13:46:25',NULL,NULL),(2651,2650,'0,2000,2600,2650','接口字段列表','system:apiColumn:index',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:46:25','2021-11-11 13:46:25',NULL,NULL),(2652,2650,'0,2000,2600,2650','接口字段回收站','system:apiColumn:recycle',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:46:25','2021-11-11 13:46:25',NULL,NULL),(2653,2650,'0,2000,2600,2650','接口字段保存','system:apiColumn:save',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:46:25','2021-11-11 13:46:25',NULL,NULL),(2654,2650,'0,2000,2600,2650','接口字段更新','system:apiColumn:update',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:46:25','2021-11-11 13:46:25',NULL,NULL),(2655,2650,'0,2000,2600,2650','接口字段删除','system:apiColumn:delete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:46:25','2021-11-11 13:46:25',NULL,NULL),(2656,2650,'0,2000,2600,2650','接口字段读取','system:apiColumn:read',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:46:25','2021-11-11 13:46:25',NULL,NULL),(2657,2650,'0,2000,2600,2650','接口字段恢复','system:apiColumn:recovery',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:46:25','2021-11-11 13:46:25',NULL,NULL),(2658,2650,'0,2000,2600,2650','接口字段真实删除','system:apiColumn:realDelete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:46:25','2021-11-11 13:46:25',NULL,NULL),(2659,2650,'0,2000,2600,2650','接口字段导入','system:apiColumn:import',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:46:25','2021-11-11 13:46:25',NULL,NULL),(2660,2650,'0,2000,2600,2650','接口字段导出','system:apiColumn:export',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-11-11 13:46:25','2021-11-11 13:46:25',NULL,NULL),(2700,2000,'0,2000','系统公告','system:notice','el-icon-flag','notice','system/notice/index',NULL,'1','M','0',94,NULL,NULL,'2021-12-25 10:10:20','2021-12-25 10:10:20',NULL,NULL),(2701,2700,'0,2000,2700','系统公告列表','system:notice:index',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-12-25 10:10:20','2021-12-25 10:10:20',NULL,NULL),(2702,2700,'0,2000,2700','系统公告回收站','system:notice:recycle',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-12-25 10:10:20','2021-12-25 10:10:20',NULL,NULL),(2703,2700,'0,2000,2700','系统公告保存','system:notice:save',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-12-25 10:10:20','2021-12-25 10:10:20',NULL,NULL),(2704,2700,'0,2000,2700','系统公告更新','system:notice:update',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-12-25 10:10:20','2021-12-25 10:10:20',NULL,NULL),(2705,2700,'0,2000,2700','系统公告删除','system:notice:delete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-12-25 10:10:20','2021-12-25 10:10:20',NULL,NULL),(2706,2700,'0,2000,2700','系统公告读取','system:notice:read',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-12-25 10:10:20','2021-12-25 10:10:20',NULL,NULL),(2707,2700,'0,2000,2700','系统公告恢复','system:notice:recovery',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-12-25 10:10:20','2021-12-25 10:10:20',NULL,NULL),(2708,2700,'0,2000,2700','系统公告真实删除','system:notice:realDelete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-12-25 10:10:20','2021-12-25 10:10:20',NULL,NULL),(2709,2700,'0,2000,2700','系统公告导入','system:notice:import',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-12-25 10:10:20','2021-12-25 10:10:20',NULL,NULL),(2710,2700,'0,2000,2700','系统公告导出','system:notice:export',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-12-25 10:10:20','2021-12-25 10:10:20',NULL,NULL),(3000,0,'0','系统监控','monitor','el-icon-aim','monitor','',NULL,'1','M','0',97,NULL,NULL,'2021-07-31 10:49:24','2021-07-31 10:49:24',NULL,NULL),(3100,3000,'0,3000','依赖监控','system:monitor:rely','ma-icon-rely','rely','system/monitor/rely/index',NULL,'1','M','0',97,NULL,NULL,'2021-07-31 10:51:48','2021-07-31 10:51:48',NULL,NULL),(3101,3100,'0,3000,3100','依赖包详细','system:monitor:relyDetail','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-08-09 03:44:20','2021-08-09 03:44:20',NULL,NULL),(3200,3000,'0,3000','服务监控','system:monitor:server','el-icon-umbrella','server','system/monitor/server/index',NULL,'1','M','0',99,NULL,NULL,'2021-07-31 10:52:46','2021-07-31 10:52:46',NULL,NULL),(3300,3000,'0,3000','日志监控','logs','el-icon-calendar','logs','',NULL,'1','M','0',95,NULL,NULL,'2021-07-31 10:54:01','2021-07-31 10:54:01',NULL,NULL),(3400,3300,'0,3000,3200','登录日志','system:loginLog','el-icon-reading','loginLog','system/loginLog/index',NULL,'1','M','0',0,NULL,NULL,'2021-07-31 10:54:55','2021-07-31 10:54:55',NULL,NULL),(3401,3400,'0,3000,3200,3300','登录日志删除','system:loginLog:delete','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-07-31 10:56:43','2021-07-31 10:56:43',NULL,NULL),(3500,3300,'0,3000,3200','操作日志','system:operLog','el-icon-document','operLog','system/operLog/index',NULL,'1','M','0',0,NULL,NULL,'2021-07-31 10:55:40','2021-07-31 10:55:40',NULL,NULL),(3501,3500,'0,3000,3200,3400','操作日志删除','system:operLog:delete','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-07-31 10:56:19','2021-07-31 10:56:19',NULL,NULL),(3600,3000,'0,3000','在线用户','system:onlineUser','ma-icon-online','onlineUser','system/monitor/onlineUser/index',NULL,'1','M','0',98,NULL,NULL,'2021-08-08 10:26:31','2021-08-08 10:26:31',NULL,NULL),(3601,3500,'0,3000,3500','在线用户列表','system:onlineUser:index','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-08-08 10:26:55','2021-08-08 10:26:55',NULL,NULL),(3602,3500,'0,3000,3500','强退用户','system:onlineUser:kick','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-08-08 10:27:21','2021-08-08 10:27:21',NULL,NULL),(3700,3000,'0,3000','缓存监控','system:cache','el-icon-odometer','cache','system/monitor/cache/index',NULL,'1','M','0',98,NULL,NULL,'2021-10-26 12:50:31','2021-10-26 12:50:31',NULL,NULL),(3701,3700,'0,3000,3700','获取Redis信息','system:cache:monitor','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-10-26 12:50:31','2021-10-26 12:50:31',NULL,NULL),(3702,3700,'0,3000,3700','删除一个缓存','system:cache:delete','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-10-26 12:50:31','2021-10-26 12:50:31',NULL,NULL),(3703,3700,'0,3000,3700','清空所有缓存','system:cache:clear','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-10-26 12:50:31','2021-10-26 12:50:31',NULL,NULL),(3800,3300,'0,3000,3300','接口日志','system:apiLog','el-icon-calendar','apiLog','system/apiLog/index',NULL,'1','M','0',0,NULL,NULL,'2021-12-06 13:50:05','2021-12-06 13:50:05',NULL,NULL),(3801,3800,'0,3000,3300,3800','接口日志删除','system:apiLog:delete','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-12-06 13:50:40','2021-12-06 13:50:40',NULL,NULL),(3850,3300,'0,3000,3300','队列日志','system:queueLog','el-icon-guide','queueLog','system/queueLog/index',NULL,'1','M','0',0,NULL,NULL,'2021-12-25 08:41:14','2021-12-25 08:41:14',NULL,NULL),(3851,3850,'0,3000,3300,3850','删除队列日志','system:queueLog:delete','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-12-25 08:42:42','2021-12-25 08:42:42',NULL,NULL),(3852,3850,'0,3000,3300,3850','更新队列状态','system:queueLog:updateStatus','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-12-25 08:45:03','2021-12-25 08:47:16',NULL,NULL),(4000,0,'0','开发工具','devTools','ma-icon-tool','devTools','',NULL,'1','M','0',95,NULL,NULL,'2021-07-31 11:03:32','2021-07-31 11:03:32',NULL,NULL),(4100,4000,'0,4000','模块管理','setting:module','el-icon-brush','module','setting/module/index',NULL,'1','M','0',99,NULL,NULL,'2021-07-31 11:33:49','2021-07-31 11:33:49',NULL,NULL),(4101,4100,'0,4000,4100','新增模块','setting:module:save','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-07-31 11:45:29','2021-07-31 11:45:29',NULL,NULL),(4102,4100,'0,4000,4100','模块删除','setting:module:delete','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-07-31 11:34:15','2021-07-31 11:34:15',NULL,NULL),(4103,4100,'0,4000,4100','模块列表','setting:module:index','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-08-09 03:48:27','2021-08-09 03:48:27',NULL,NULL),(4104,4100,'0,4000,4100','模块启停用','setting:module:status',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2022-02-13 10:10:20','2022-02-13 10:10:20',NULL,NULL),(4105,4100,'0,4000,4100','安装模块','setting:module:install',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2022-02-13 10:10:20','2022-02-13 10:10:20',NULL,NULL),(4200,4000,'0,4000','代码生成器','setting:code','ma-icon-code','code','setting/code/index',NULL,'1','M','0',98,NULL,NULL,'2021-07-31 11:36:17','2021-07-31 11:36:17',NULL,NULL),(4201,4200,'0,4000,4200','预览代码','setting:code:preview','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-07-31 11:36:44','2021-07-31 11:36:44',NULL,NULL),(4202,4200,'0,4000,4200','装载数据表','setting:code:loadTable','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-07-31 11:38:03','2021-07-31 11:38:03',NULL,NULL),(4203,4200,'0,4000,4200','删除业务表','setting:code:delete','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-07-31 11:38:53','2021-07-31 11:38:53',NULL,NULL),(4204,4200,'0,4000,4200','同步业务表','setting:code:sync','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-07-31 11:39:18','2021-07-31 11:39:18',NULL,NULL),(4205,4200,'0,4000,4200','生成代码','setting:code:generate','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-07-31 11:39:48','2021-07-31 11:39:48',NULL,NULL),(4206,4200,'0,4000,4200','更新业务表','setting:code:update','','codeEdit','setting/code/edit',NULL,'0','M','0',0,6798055202976,6798055202976,'2021-07-31 11:43:12','2022-04-15 08:21:32',NULL,NULL),(4300,4000,'0,4000','数据表设计器','setting:table','el-icon-magic-stick','table','setting/table/index',NULL,'1','M','0',0,NULL,NULL,'2021-07-31 11:44:08','2021-07-31 11:44:08',NULL,NULL),(4301,4300,'0,4000,4300','保存数据表','setting:table:save','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-07-31 11:44:36','2021-07-31 11:44:36',NULL,NULL),(4400,4000,'0,4000','定时任务','setting:crontab','el-icon-alarm-clock','crontab','setting/crontab/index','','1','M','0',0,NULL,NULL,'2021-07-31 11:47:49','2021-07-31 11:47:49',NULL,NULL),(4401,4400,'0,4000,4400','定时任务列表','setting:crontab:index',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 11:47:49','2021-07-31 11:47:49',NULL,NULL),(4402,4400,'0,4000,4400','定时任务保存','setting:crontab:save',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 11:47:49','2021-07-31 11:47:49',NULL,NULL),(4403,4400,'0,4000,4400','定时任务更新','setting:crontab:update',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 11:47:49','2021-07-31 11:47:49',NULL,NULL),(4404,4400,'0,4000,4400','定时任务删除','setting:crontab:delete',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 11:47:49','2021-07-31 11:47:49',NULL,NULL),(4405,4400,'0,4000,4400','定时任务读取','setting:crontab:read',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 11:47:49','2021-07-31 11:47:49',NULL,NULL),(4406,4400,'0,4000,4400','定时任务导入','setting:crontab:import',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 11:47:49','2021-07-31 11:47:49',NULL,NULL),(4407,4400,'0,4000,4400','定时任务导出','setting:crontab:export',NULL,NULL,NULL,NULL,'1','B','0',0,NULL,NULL,'2021-07-31 11:47:49','2021-07-31 11:47:49',NULL,NULL),(4408,4400,'0,4000,4400','定时任务执行','setting:crontab:run','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-08-07 15:44:06','2021-08-07 15:44:06',NULL,NULL),(4409,4400,'0,4000,4400','定时任务日志删除','setting:crontab:deleteLog','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-12-06 14:06:12','2021-12-06 14:06:12',NULL,NULL),(4500,0,'0','系统设置','setting:config','','setting','','','0','M','0',0,NULL,NULL,'2021-07-31 11:58:57','2021-07-31 11:58:57',NULL,NULL),(4501,4500,'0,4500','获取系统组配置','setting:config:getSystemConfig','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-08-10 05:08:49','2021-08-10 05:08:49',NULL,NULL),(4502,4500,'0,4500','获取扩展组配置','setting:config:getExtendConfig','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-08-10 05:09:18','2021-08-10 05:09:18',NULL,NULL),(4503,4500,'0,4500','保存系统组配置','setting:config:saveSystemConfig','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-08-10 05:11:25','2021-08-10 05:11:25',NULL,NULL),(4504,4500,'0,4500','新增配置 ','setting:config:save','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-08-10 05:11:56','2021-08-10 05:11:56',NULL,NULL),(4505,4500,'0,4500','更新配置','setting:config:update','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-08-10 05:12:25','2021-08-10 05:12:25',NULL,NULL),(4506,4500,'0,4500','删除配置','setting:config:delete','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-08-10 05:13:33','2021-08-10 05:13:33',NULL,NULL),(4507,4500,'0,4500','清除配置缓存','setting:config:clearCache','',NULL,'',NULL,'1','B','0',0,NULL,NULL,'2021-08-10 05:13:59','2021-08-10 05:13:59',NULL,NULL),(4600,4000,'0,4000','系统接口','systemInterface','el-icon-histogram','systemInterface','setting/systemInterface/index',NULL,'1','M','0',0,NULL,NULL,'2022-03-30 02:00:28','2022-03-30 02:00:28',NULL,NULL),(4601,4000,'0,4000','表单设计器','formDesigner','el-icon-set-up','formDesigner','setting/formDesigner/index',NULL,'1','M','0',0,NULL,NULL,'2022-03-30 02:13:37','2022-03-30 02:32:36',NULL,NULL);
/*!40000 ALTER TABLE `system_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_notice`
--

DROP TABLE IF EXISTS `system_notice`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_notice` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `message_id` bigint(20) DEFAULT NULL COMMENT '消息ID',
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '标题',
  `type` char(1) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '公告类型（1通知 2公告）',
  `content` text COLLATE utf8mb4_unicode_ci COMMENT '公告内容',
  `click_num` int(11) DEFAULT '0' COMMENT '浏览次数',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_notice`
--

LOCK TABLES `system_notice` WRITE;
/*!40000 ALTER TABLE `system_notice` DISABLE KEYS */;
/*!40000 ALTER TABLE `system_notice` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_oper_log`
--

DROP TABLE IF EXISTS `system_oper_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_oper_log` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `username` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `method` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '请求方式',
  `router` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '请求路由',
  `service_name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '业务名称',
  `ip` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求IP地址',
  `ip_location` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'IP所属地',
  `request_data` text COLLATE utf8mb4_unicode_ci COMMENT '请求数据',
  `response_code` varchar(5) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '响应状态码',
  `response_data` text COLLATE utf8mb4_unicode_ci COMMENT '响应数据',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_oper_log`
--

LOCK TABLES `system_oper_log` WRITE;
/*!40000 ALTER TABLE `system_oper_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `system_oper_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_post`
--

DROP TABLE IF EXISTS `system_post`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_post` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '岗位名称',
  `code` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '岗位代码',
  `sort` tinyint(3) unsigned DEFAULT '0' COMMENT '排序',
  `status` char(1) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '状态 (0正常 1停用)',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_post`
--

LOCK TABLES `system_post` WRITE;
/*!40000 ALTER TABLE `system_post` DISABLE KEYS */;
INSERT INTO `system_post` VALUES (1,'首席执行官','CEO',1,'0',1,1,NULL,NULL,NULL,NULL);
/*!40000 ALTER TABLE `system_post` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_queue_log`
--

DROP TABLE IF EXISTS `system_queue_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_queue_log` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `exchange_name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '交换机名称',
  `routing_key_name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '路由名称',
  `queue_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '队列名称',
  `queue_content` longtext COLLATE utf8mb4_unicode_ci COMMENT '队列数据',
  `log_content` text COLLATE utf8mb4_unicode_ci COMMENT '队列日志',
  `produce_status` char(1) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '生产状态 0:未生产 1:生产中 2:生产成功 3:生产失败 4:生产重复',
  `consume_status` char(1) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '消费状态 0:未消费 1:消费中 2:消费成功 3:消费失败 4:消费重复',
  `delay_time` int(10) unsigned NOT NULL COMMENT '延迟时间（秒）',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_queue_log`
--

LOCK TABLES `system_queue_log` WRITE;
/*!40000 ALTER TABLE `system_queue_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `system_queue_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_queue_message`
--

DROP TABLE IF EXISTS `system_queue_message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_queue_message` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `content_type` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '内容类型',
  `title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '消息标题',
  `content` longtext COLLATE utf8mb4_unicode_ci COMMENT '消息内容',
  `send_by` bigint(20) unsigned NOT NULL COMMENT '发送人',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `system_queue_message_content_type_index` (`content_type`),
  KEY `system_queue_message_send_by_index` (`send_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_queue_message`
--

LOCK TABLES `system_queue_message` WRITE;
/*!40000 ALTER TABLE `system_queue_message` DISABLE KEYS */;
/*!40000 ALTER TABLE `system_queue_message` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_queue_message_receive`
--

DROP TABLE IF EXISTS `system_queue_message_receive`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_queue_message_receive` (
  `message_id` bigint(20) unsigned NOT NULL COMMENT '队列消息主键',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '接收用户主键',
  `read_status` char(1) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '已读状态 (0未读 1已读)',
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_queue_message_receive`
--

LOCK TABLES `system_queue_message_receive` WRITE;
/*!40000 ALTER TABLE `system_queue_message_receive` DISABLE KEYS */;
/*!40000 ALTER TABLE `system_queue_message_receive` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_role`
--

DROP TABLE IF EXISTS `system_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_role` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，角色ID',
  `name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色名称',
  `code` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色代码',
  `data_scope` char(1) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '数据范围（0：全部数据权限 1：自定义数据权限 2：本部门数据权限 3：本部门及以下数据权限 4：本人数据权限）',
  `status` char(1) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '状态 (0正常 1停用)',
  `sort` tinyint(3) unsigned DEFAULT '0' COMMENT '排序',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_role`
--

LOCK TABLES `system_role` WRITE;
/*!40000 ALTER TABLE `system_role` DISABLE KEYS */;
INSERT INTO `system_role` VALUES (1,'超级管理员（创始人）','superAdmin','0','0',0,1,0,'2022-07-18 07:07:41','2022-07-18 07:07:41',NULL,'系统内置角色，不可删除'),(2,'高级管理员','admin','0','0',3,1,1,'2022-07-23 17:45:21','2022-07-23 18:21:06',NULL,'高级管理员111');
/*!40000 ALTER TABLE `system_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_role_dept`
--

DROP TABLE IF EXISTS `system_role_dept`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_role_dept` (
  `role_id` bigint(20) unsigned NOT NULL COMMENT '角色主键',
  `dept_id` bigint(20) unsigned NOT NULL COMMENT '部门主键',
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_role_dept`
--

LOCK TABLES `system_role_dept` WRITE;
/*!40000 ALTER TABLE `system_role_dept` DISABLE KEYS */;
/*!40000 ALTER TABLE `system_role_dept` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_role_menu`
--

DROP TABLE IF EXISTS `system_role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_role_menu` (
  `role_id` bigint(20) unsigned NOT NULL COMMENT '角色主键',
  `menu_id` bigint(20) unsigned NOT NULL COMMENT '菜单主键',
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_role_menu`
--

LOCK TABLES `system_role_menu` WRITE;
/*!40000 ALTER TABLE `system_role_menu` DISABLE KEYS */;
/*!40000 ALTER TABLE `system_role_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_uploadfile`
--

DROP TABLE IF EXISTS `system_uploadfile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_uploadfile` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `storage_mode` char(1) COLLATE utf8mb4_unicode_ci DEFAULT '1' COMMENT '状态 (1 本地 2 阿里云 3 七牛云 4 腾讯云)',
  `origin_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '原文件名',
  `object_name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '新文件名',
  `mime_type` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '资源类型',
  `storage_path` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '存储目录',
  `suffix` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件后缀',
  `size_byte` bigint(20) DEFAULT NULL COMMENT '字节数',
  `size_info` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件大小',
  `url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'url地址',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `system_uploadfile_storage_path_index` (`storage_path`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_uploadfile`
--

LOCK TABLES `system_uploadfile` WRITE;
/*!40000 ALTER TABLE `system_uploadfile` DISABLE KEYS */;
/*!40000 ALTER TABLE `system_uploadfile` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_user`
--

DROP TABLE IF EXISTS `system_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID，主键',
  `username` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `password` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `user_type` varchar(3) COLLATE utf8mb4_unicode_ci DEFAULT '100' COMMENT '用户类型：(100系统用户)',
  `nickname` varchar(30) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户昵称',
  `phone` varchar(11) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '手机',
  `email` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户邮箱',
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户头像',
  `signed` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '个人签名',
  `dashboard` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '后台首页类型',
  `dept_id` bigint(20) unsigned DEFAULT NULL COMMENT '部门ID',
  `status` char(1) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '状态 (0正常 1停用)',
  `login_ip` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '最后登陆IP',
  `login_time` timestamp NULL DEFAULT NULL COMMENT '最后登陆时间',
  `backend_setting` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '后台设置数据',
  `created_by` bigint(20) DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint(20) DEFAULT NULL COMMENT '更新者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `system_user_username_unique` (`username`),
  KEY `system_user_dept_id_index` (`dept_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_user`
--

LOCK TABLES `system_user` WRITE;
/*!40000 ALTER TABLE `system_user` DISABLE KEYS */;
INSERT INTO `system_user` VALUES (1,'superAdmin','0192023a7bbd73250516f069df18b500','100','创始人','16858888988','admin@adminmine.com',NULL,NULL,'statistics',1,'0',NULL,NULL,NULL,0,0,'2022-07-18 07:07:41','2022-07-18 07:07:41',NULL,NULL);
/*!40000 ALTER TABLE `system_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_user_post`
--

DROP TABLE IF EXISTS `system_user_post`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_user_post` (
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户主键',
  `post_id` bigint(20) unsigned NOT NULL COMMENT '岗位主键',
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_user_post`
--

LOCK TABLES `system_user_post` WRITE;
/*!40000 ALTER TABLE `system_user_post` DISABLE KEYS */;
/*!40000 ALTER TABLE `system_user_post` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_user_role`
--

DROP TABLE IF EXISTS `system_user_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_user_role` (
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户主键',
  `role_id` bigint(20) unsigned NOT NULL COMMENT '角色主键',
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_user_role`
--

LOCK TABLES `system_user_role` WRITE;
/*!40000 ALTER TABLE `system_user_role` DISABLE KEYS */;
INSERT INTO `system_user_role` VALUES (1,1,1);
/*!40000 ALTER TABLE `system_user_role` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-07-24 21:38:59

