/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50553
 Source Host           : localhost:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 50553
 File Encoding         : 65001

 Date: 21/08/2019 17:00:46
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_blog_admin
-- ----------------------------
DROP TABLE IF EXISTS `t_blog_admin`;
CREATE TABLE `t_blog_admin`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(6) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '盐值',
  `created_at` int(11) NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(11) NULL DEFAULT 0 COMMENT '更新时间',
  `isdel` tinyint(4) NOT NULL DEFAULT 1 COMMENT '是否删除 0 删除 1 正常',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '后台用户表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for t_blog_article
-- ----------------------------
DROP TABLE IF EXISTS `t_blog_article`;
CREATE TABLE `t_blog_article`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uid` int(11) NULL DEFAULT 0 COMMENT '创建人id',
  `type` tinyint(4) NOT NULL DEFAULT 1 COMMENT '类型 1 php 2 go 3 算法 4 linux 5 js',
  `title` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `introduce` varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '说明',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '内容',
  `count` int(11) NOT NULL DEFAULT 0 COMMENT '阅读次数',
  `check_status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '状态 0 编辑中 1 待审核 2 审核通过 3 审核失败',
  `publish_time` int(11) NULL DEFAULT 0 COMMENT '发布时间',
  `created_at` timestamp not null DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp not null DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `isdel` tinyint(4) NOT NULL DEFAULT 1 COMMENT '是否删除 0 删除 1 正常',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '文章表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for t_blog_comment
-- ----------------------------
DROP TABLE IF EXISTS `t_blog_comment`;
CREATE TABLE `t_blog_comment`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `article_id` int(11) NOT NULL DEFAULT 0 COMMENT '文章id',
  `fid` int(11) NOT NULL DEFAULT 0 COMMENT '父级id',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '内容',
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `isdel` tinyint(4) NOT NULL DEFAULT 1 COMMENT '是否删除 0 删除 1 正常',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '评论表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for t_blog_user
-- ----------------------------
DROP TABLE IF EXISTS `t_blog_user`;
CREATE TABLE `t_blog_user`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `nickname` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `email` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `avatar` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(6) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '盐值',
  `active` tinyint(4) NOT NULL DEFAULT 0 COMMENT '账户状态 0 未激活 1 已激活',
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `isdel` tinyint(4) NOT NULL DEFAULT 1 COMMENT '是否删除 0 删除 1 正常',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户表' ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
