CREATE DATABASE IF NOT EXISTS `my_music_player`
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci;

USE `my_music_player`;

CREATE TABLE `user` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT COMMENT '用户ID,主键,自动递增',
  `username` varchar(32) UNIQUE NOT NULL COMMENT '用户名,唯一且不能为空',
  `password` varchar(64) NOT NULL COMMENT '用户密码,不能为空',
  `name` varchar(32) DEFAULT NULL COMMENT '用户姓名,允许为空',
  `phone` varchar(11) DEFAULT NULL COMMENT '用户电话号码,允许为空',
  `email` VARCHAR(255) DEFAULT NULL COMMENT '用户电子邮件,允许为空',
  `sex` varchar(2) DEFAULT NULL COMMENT '用户性别,允许为空',
  `created_time` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '账户创建时间,默认为当前时间'
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='用户表,存储用户的基本信息';

CREATE TABLE `manager` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT COMMENT '管理员ID,主键,自动递增',
  `username` varchar(32) UNIQUE NOT NULL COMMENT '管理员用户名,唯一且不能为空',
  `password` varchar(64) NOT NULL COMMENT '管理员密码,不能为空'
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='管理员表,存储管理员的基本信息';

CREATE TABLE `song` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT COMMENT '歌曲ID,主键,自动递增',
  `name` varchar(32) NOT NULL COMMENT '歌曲名称,不能为空',
  `author_id` bigint NOT NULL COMMENT '歌曲作者ID,不能为空,对应艺术家的ID',
  `album_id` bigint NOT NULL COMMENT '专辑ID,不能为空,对应专辑的ID',
  `is_public` bool DEFAULT TRUE COMMENT '是否公开,默认为公开',
  `loc` varchar(255) DEFAULT NULL COMMENT '歌曲存放位置,允许为空',
  `pic` varchar(255) DEFAULT NULL COMMENT '歌曲封面图片,允许为空',
  `create_time` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '歌曲创建时间,默认为当前时间'
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='歌曲表,存储歌曲的基本信息';

CREATE TABLE `like` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT COMMENT '点赞ID,主键,自动递增',
  `user_id` bigint NOT NULL COMMENT '用户ID,关联用户表的ID',
  `song_id` bigint NOT NULL COMMENT '歌曲ID,关联歌曲表的ID'
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='用户对歌曲的点赞记录,存储用户和歌曲的关联';

CREATE TABLE `comment` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT COMMENT '评论ID,主键,自动递增',
  `user_id` bigint NOT NULL COMMENT '用户ID,关联用户表的ID',
  `song_id` bigint NOT NULL COMMENT '歌曲ID,关联歌曲表的ID',
  `create_time` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '评论创建时间,默认为当前时间',
  `text` text NOT NULL COMMENT '评论内容,不能为空'
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='评论表,存储用户对歌曲的评论';

CREATE TABLE `playlist` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT COMMENT '歌单ID,主键,自动递增',
  `user_id` bigint NOT NULL COMMENT '用户ID,关联用户表的ID',
  `song_id` bigint NOT NULL COMMENT '歌曲ID,关联歌曲表的ID'
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='歌单表,存储用户创建的歌单和歌曲的关联';

CREATE TABLE `artist` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT COMMENT '艺术家ID,主键,自动递增',
  `name` varchar(32) NOT NULL COMMENT '艺术家名称,不能为空',
  `sex` varchar(2) DEFAULT NULL COMMENT '艺术家性别,允许为空'
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='艺术家表,存储艺术家的基本信息';

CREATE TABLE `album` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT COMMENT '专辑ID,主键,自动递增',
  `name` varchar(32) NOT NULL COMMENT '专辑名称,不能为空',
  `cover` varchar(255) DEFAULT NULL COMMENT '专辑封面,允许为空'
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='专辑表,存储专辑的基本信息';

CREATE TABLE `album_list` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT COMMENT '专辑列表ID,主键,自动递增',
  `album_id` bigint NOT NULL COMMENT '专辑ID,关联专辑表的ID',
  `user_id` bigint NOT NULL COMMENT '用户ID,关联用户表的ID'
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='专辑列表表,存储用户与专辑的关联';

CREATE TABLE `album_artist` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT COMMENT '专辑艺术家关联ID,主键,自动递增',
  `artist_id` bigint NOT NULL COMMENT '艺术家ID,关联艺术家表的ID',
  `album_id` bigint NOT NULL COMMENT '专辑ID,关联专辑表的ID'
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='专辑与艺术家关联表,存储艺术家与专辑的关系';


