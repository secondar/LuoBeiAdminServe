/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : luobeiadmin

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 22/08/2022 09:42:46
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for lb_admin
-- ----------------------------
DROP TABLE IF EXISTS `lb_admin`;
CREATE TABLE `lb_admin`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `account` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '账户',
  `password` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `interfere` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '干扰码',
  `role` int(11) NOT NULL COMMENT '角色',
  `state` tinyint(1) NOT NULL COMMENT '状态，0=禁用,1=启用',
  `addtime` datetime(0) NULL DEFAULT NULL COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lb_admin
-- ----------------------------
INSERT INTO `lb_admin` VALUES (1, 'admin', '281ab141a12f67f5238719cd876ce96e', 'e10adc3949ba59abbe56e057f20f883e', 1, 1, '2021-10-17 13:22:23');

-- ----------------------------
-- Table structure for lb_admin_log
-- ----------------------------
DROP TABLE IF EXISTS `lb_admin_log`;
CREATE TABLE `lb_admin_log`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `aid` int(11) NOT NULL COMMENT '管理员ID',
  `type` tinyint(1) NOT NULL COMMENT '类型，1=登录日志',
  `content` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '内容',
  `addtime` datetime(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lb_admin_log
-- ----------------------------

-- ----------------------------
-- Table structure for lb_admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `lb_admin_menu`;
CREATE TABLE `lb_admin_menu`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL COMMENT '父级ID',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '菜单名称',
  `type` tinyint(1) NOT NULL COMMENT '类型1-目录，2=菜单，3=按钮',
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '图标',
  `show` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否显示0=隐藏，1=显示',
  `link` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否为外链0-否，1=是',
  `api_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '接口地址',
  `characteristic` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '权限标识',
  `router` varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '路由地址，如果是外链的话这里会变成链接地址',
  `sort` int(11) NULL DEFAULT NULL COMMENT '排序，越小越靠前',
  `component` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '组件名称',
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '组件路径',
  `addtime` datetime(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lb_admin_menu
-- ----------------------------
INSERT INTO `lb_admin_menu` VALUES (1, 0, '系统管理', 1, 'el-icon-setting', 1, 0, NULL, 'System', 'system', 1, '', '', '2021-10-21 02:06:40');
INSERT INTO `lb_admin_menu` VALUES (2, 1, '菜单管理', 2, '', 1, 0, '/admin/api/v1/menu/list', 'MenuManagement', 'menus', 3, 'menus', '/system/menus', '2021-10-21 02:09:31');
INSERT INTO `lb_admin_menu` VALUES (3, 1, '用户管理', 2, '', 1, 0, '/admin/api/v1/admin/getlist', 'UserManagement', 'user', 1, 'user', '/system/user', '2021-10-21 14:07:38');
INSERT INTO `lb_admin_menu` VALUES (4, 1, '角色管理', 2, '', 1, 0, '/admin/api/v1/role/list', 'RoleManagement', 'role', 2, 'role', '/system/role', '2021-10-22 13:22:12');
INSERT INTO `lb_admin_menu` VALUES (5, 3, '新增用户', 3, '', 1, 0, '/admin/api/v1/admin/add', 'AddUser', '', 999, '', '', '2021-10-29 08:31:30');
INSERT INTO `lb_admin_menu` VALUES (6, 3, '编辑用户', 3, '', 1, 0, '/admin/api/v1/admin/edit', 'EditUser', 'user', 999, '', '', '2021-10-29 08:31:52');
INSERT INTO `lb_admin_menu` VALUES (7, 3, '删除用户', 3, '', 1, 0, '/admin/api/v1/admin/delete', 'DeleteUser', 'user', 999, '', '', '2021-10-29 08:32:11');
INSERT INTO `lb_admin_menu` VALUES (8, 4, '新增角色', 3, '', 1, 0, '/admin/api/v1/role/add', 'AddRole', 'role', 999, '', '', '2021-10-29 08:33:51');
INSERT INTO `lb_admin_menu` VALUES (9, 4, '编辑角色', 3, '', 1, 0, '/admin/api/v1/role/edit', 'EditRole', 'user', 999, '', '', '2021-10-29 08:34:32');
INSERT INTO `lb_admin_menu` VALUES (10, 4, '删除角色', 3, '', 1, 0, '/admin/api/v1/role/delete', 'DeleteRole', 'user', 999, '', '', '2021-10-29 08:34:57');
INSERT INTO `lb_admin_menu` VALUES (11, 2, '新增菜单', 3, '', 1, 0, '/admin/api/v1/menu/add', 'AddMenu', 'user', 999, '', '', '2021-10-29 08:35:38');
INSERT INTO `lb_admin_menu` VALUES (12, 2, '编辑菜单', 3, '', 1, 0, '/admin/api/v1/menu/edit', 'EditMenu', 'user', 999, '', '', '2021-10-29 08:36:00');
INSERT INTO `lb_admin_menu` VALUES (13, 2, '删除菜单', 3, '', 1, 0, '/admin/api/v1/menu/delete', 'DeleteMenu', 'user', 999, '', '', '2021-10-29 08:36:23');
INSERT INTO `lb_admin_menu` VALUES (14, 4, '设置角色路由', 3, '', 1, 0, '/admin/api/v1/router/setting', 'SettingRole', 'user', 999, '', '', '2021-10-29 08:37:06');
INSERT INTO `lb_admin_menu` VALUES (15, 4, '获取角色路由', 3, '', 1, 0, '/v1/router/getrouter', 'GetUserRouter', 'user', 999, '', '', '2021-10-29 08:37:26');
INSERT INTO `lb_admin_menu` VALUES (16, 0, '文章管理', 1, 'el-icon-notebook-1', 1, 0, NULL, 'Article', 'article', 2, '', '', '2021-11-05 03:32:34');
INSERT INTO `lb_admin_menu` VALUES (17, 16, '分类管理', 2, '', 1, 0, '/admin/api/v1/article/sort/getlist', 'ArticleSortList', 'sort', 3, 'sort', '/article/sort', '2021-11-05 03:33:54');
INSERT INTO `lb_admin_menu` VALUES (18, 16, '文章管理', 2, '', 1, 0, '/admin/api/v1/article/getlist', 'ArticleList', 'articlelist', 2, 'articlelist', '/article/list', '2021-11-07 10:57:20');
INSERT INTO `lb_admin_menu` VALUES (19, 16, '创建文章', 2, '', 1, 0, '/admin/api/v1/article/add', 'ArticleAdd', 'articleadd', 1, 'articleadd', '/article/add', '2021-11-08 13:13:35');
INSERT INTO `lb_admin_menu` VALUES (20, 16, '编辑文章', 2, '', 0, 0, '/admin/api/v1/article/edit', 'ArticleEdit', 'articleedit', 3, 'articleedit', '/article/edit', '2021-11-09 05:55:35');
INSERT INTO `lb_admin_menu` VALUES (21, 17, '添加分类', 3, '', 1, 0, '/admin/api/v1/article/sort/add', 'ArticleSortAdd', 'sort', 999, '', '', '2021-11-09 08:28:47');
INSERT INTO `lb_admin_menu` VALUES (22, 17, '删除分类', 3, '', 1, 0, '/admin/api/v1/article/sort/delete', 'ArticleSortDelete', 'sort', 999, '', '', '2021-11-09 08:29:03');
INSERT INTO `lb_admin_menu` VALUES (23, 17, '编辑分类', 3, '', 1, 0, '/admin/api/v1/article/sort/edit', 'ArticleSortEdit', 'sort', 999, '', '', '2021-11-09 08:29:41');
INSERT INTO `lb_admin_menu` VALUES (24, 16, '删除文章', 3, '', 1, 0, '/admin/api/v1/article/delete', 'ArticleDelete', '', 999, '', '', '2021-11-09 12:30:30');
INSERT INTO `lb_admin_menu` VALUES (25, 16, '查看文章详情', 3, '', 1, 0, '/admin/api/v1/article/details', 'ArticleDetails', '', 999, '', '', '2021-11-09 12:31:16');
INSERT INTO `lb_admin_menu` VALUES (26, 1, '系统配置', 2, '', 1, 0, NULL, 'ConfManagement', 'system', 999, 'system', '/system/system', '2021-11-09 12:51:57');
INSERT INTO `lb_admin_menu` VALUES (27, 26, '更改配置', 3, '', 1, 0, '/v1/system/save', 'SaveConf', '', 999, '', '', '2021-11-09 13:20:55');

-- ----------------------------
-- Table structure for lb_admin_on_line
-- ----------------------------
DROP TABLE IF EXISTS `lb_admin_on_line`;
CREATE TABLE `lb_admin_on_line`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `aid` int(11) NOT NULL COMMENT 'adminid',
  `account` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'account',
  `token` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'token',
  `expiration_time` datetime(0) NOT NULL COMMENT '失效时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lb_admin_on_line
-- ----------------------------

-- ----------------------------
-- Table structure for lb_admin_role
-- ----------------------------
DROP TABLE IF EXISTS `lb_admin_role`;
CREATE TABLE `lb_admin_role`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '权限名称',
  `remarks` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `state` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否启用,1=启用',
  `addtime` datetime(0) NOT NULL COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lb_admin_role
-- ----------------------------
INSERT INTO `lb_admin_role` VALUES (1, '超级管理员', NULL, 1, '2021-10-24 04:40:07');
INSERT INTO `lb_admin_role` VALUES (2, '管理员', NULL, 1, '2021-10-27 09:28:41');

-- ----------------------------
-- Table structure for lb_admin_router
-- ----------------------------
DROP TABLE IF EXISTS `lb_admin_router`;
CREATE TABLE `lb_admin_router`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `role` int(11) NOT NULL COMMENT '角色ID',
  `menu` int(11) NOT NULL COMMENT '菜单ID',
  `is_pid` tinyint(1) NULL DEFAULT NULL COMMENT '是否是PID(非用户提交而是通过代码找出来的)',
  `addtime` datetime(0) NOT NULL COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 22 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lb_admin_router
-- ----------------------------
INSERT INTO `lb_admin_router` VALUES (16, 2, 17, 1, '2021-11-09 08:58:18');
INSERT INTO `lb_admin_router` VALUES (17, 2, 21, 0, '2021-11-09 08:58:19');
INSERT INTO `lb_admin_router` VALUES (18, 2, 22, 0, '2021-11-09 08:58:19');
INSERT INTO `lb_admin_router` VALUES (19, 2, 23, 0, '2021-11-09 08:58:19');
INSERT INTO `lb_admin_router` VALUES (20, 2, 18, 0, '2021-11-09 08:58:19');
INSERT INTO `lb_admin_router` VALUES (21, 2, 16, 1, '2021-11-09 08:58:19');

-- ----------------------------
-- Table structure for lb_article
-- ----------------------------
DROP TABLE IF EXISTS `lb_article`;
CREATE TABLE `lb_article`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `sort` int(11) NOT NULL COMMENT '分类',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标题',
  `thumbnail` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '缩略图',
  `content` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '内容',
  `describe` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '描述',
  `hot` int(11) NOT NULL DEFAULT 0 COMMENT '点击数',
  `addtime` datetime(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lb_article
-- ----------------------------
INSERT INTO `lb_article` VALUES (1, 2, '测试', '', '<p>内容</p>', '摘要', 0, '2021-11-08 13:26:24');
INSERT INTO `lb_article` VALUES (2, 2, '测试', '', '<p>内容</p>', '摘要', 0, '2021-11-08 13:26:24');
INSERT INTO `lb_article` VALUES (3, 2, '测试', '', '<p>内容</p>', '摘要', 0, '2021-11-08 13:26:24');
INSERT INTO `lb_article` VALUES (4, 2, '测试', '', '<p>内容</p>', '摘要', 0, '2021-11-08 13:26:24');
INSERT INTO `lb_article` VALUES (5, 2, '测试', '', '<p>内容</p>', '摘要', 0, '2021-11-08 13:26:24');
INSERT INTO `lb_article` VALUES (6, 2, '测试', '', '<p>内容</p>', '摘要', 0, '2021-11-08 13:26:24');
INSERT INTO `lb_article` VALUES (7, 2, '测试', '', '<p>内容</p>', '摘要', 0, '2021-11-08 13:26:24');
INSERT INTO `lb_article` VALUES (8, 2, '测试', '', '<p>内容</p>', '摘要', 0, '2021-11-08 13:26:24');
INSERT INTO `lb_article` VALUES (9, 2, '测试', '', '<p>内容</p>', '摘要', 0, '2021-11-08 13:26:24');

-- ----------------------------
-- Table structure for lb_article_sort
-- ----------------------------
DROP TABLE IF EXISTS `lb_article_sort`;
CREATE TABLE `lb_article_sort`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级分类',
  `title` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标题',
  `state` tinyint(1) NOT NULL COMMENT '状态，1=启用',
  `sort` int(11) NOT NULL DEFAULT 1 COMMENT '排序，越低越前',
  `addtime` datetime(0) NOT NULL COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lb_article_sort
-- ----------------------------
INSERT INTO `lb_article_sort` VALUES (1, 0, ' 新闻中心', 1, 1, '2021-11-06 13:23:59');
INSERT INTO `lb_article_sort` VALUES (2, 1, '新闻资讯', 1, 1, '2021-11-06 13:31:16');

-- ----------------------------
-- Table structure for lb_system
-- ----------------------------
DROP TABLE IF EXISTS `lb_system`;
CREATE TABLE `lb_system`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '网站标题',
  `tail` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '网站小尾巴',
  `keyword` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '网站关键词',
  `describe` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '网站描述',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lb_system
-- ----------------------------
INSERT INTO `lb_system` VALUES (1, '萝北后台管理系统-LuoBeiAdmin', '萝北后台管理系统-LuoBeiAdmin', '萝北后台管理系统，LuoBeiAdmin，beego后台管理系统，vue后台管理系统，前后端分离管理系统，go，beego，vuelementadmin', '萝北后台管理系统，LuoBeiAdmin，beego后台管理系统，vue后台管理系统，前后端分离管理系统，go，beego，vuelementadmin');

SET FOREIGN_KEY_CHECKS = 1;
