-- MySQL dump 10.13  Distrib 8.4.2, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: music_db1
-- ------------------------------------------------------
-- Server version	8.4.2

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `album`
--

DROP TABLE IF EXISTS `album`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `album` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '专辑ID,主键,自动递增',
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '专辑名称,不能为空',
  `cover` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '专辑封面,不能为空',
  `release_date` date DEFAULT NULL COMMENT '专辑发行日期,允许为空',
  `description` text COLLATE utf8mb4_unicode_ci COMMENT '专辑描述,允许为空',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间,默认为当前时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='专辑表,存储专辑的基本信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `album`
--

LOCK TABLES `album` WRITE;
/*!40000 ALTER TABLE `album` DISABLE KEYS */;
INSERT INTO `album` VALUES (1,'The Beatles','/assets/covers/albums/the_beatles.jpg','1968-11-22','《The Beatles》是披头士的第九张专辑，也被称为“白色专辑”，以其多样化的音乐风格和创新性广受好评。','2024-12-28 11:20:05'),(2,'Abbey Road','/assets/covers/albums/abbey_road.jpg','1969-09-26','《Abbey Road》是披头士的第十一张专辑，专辑中的“Come Together”和“Something”成为经典，展示了他们创作的巅峰。','2024-12-28 11:20:05'),(3,'Blonde on Blonde','/assets/covers/albums/blonde_on_blonde.jpg','1966-06-20','《Blonde on Blonde》是鲍勃·迪伦的第七张专辑，被认为是迪伦音乐生涯的里程碑之一，融合了民谣与摇滚元素。','2024-12-28 11:20:05'),(4,'Highway 61 Revisited','/assets/covers/albums/highway_61_revisited.jpg','1965-08-30','《Highway 61 Revisited》是鲍勃·迪伦的第六张专辑，开创了他转向电声音乐的风格，收录了经典的“Like a Rolling Stone”。','2024-12-28 11:20:05');
/*!40000 ALTER TABLE `album` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `artist`
--

DROP TABLE IF EXISTS `artist`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `artist` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '艺术家ID,主键,自动递增',
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '艺术家名称,不能为空',
  `sex` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '艺术家性别,允许为空',
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '/assets/avatars/default-artist.jpg' COMMENT '艺术家头像,允许为空',
  `description` text COLLATE utf8mb4_unicode_ci COMMENT '艺术家描述,允许为空',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间,默认为当前时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='艺术家表,存储艺术家的基本信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `artist`
--

LOCK TABLES `artist` WRITE;
/*!40000 ALTER TABLE `artist` DISABLE KEYS */;
INSERT INTO `artist` VALUES (1,'The Beatles',NULL,'/assets/avatars/the_beatles.jpg','披头士是20世纪60年代最具影响力的英国摇滚乐队之一，他们的音乐风格对全球流行文化产生了深远影响。','2024-12-28 11:20:05'),(2,'Bob Dylan',NULL,'/assets/avatars/bob_dylan.jpg','鲍勃·迪伦是美国民谣歌手和词曲创作人，他的作品对现代音乐和社会运动有着重要的影响。','2024-12-28 11:20:05');
/*!40000 ALTER TABLE `artist` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `artist_album`
--

DROP TABLE IF EXISTS `artist_album`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `artist_album` (
  `artist_id` int NOT NULL COMMENT '艺术家ID,关联艺术家表的ID',
  `album_id` int NOT NULL COMMENT '专辑ID,关联专辑表的ID',
  PRIMARY KEY (`artist_id`,`album_id`),
  KEY `album_id` (`album_id`),
  CONSTRAINT `artist_album_ibfk_1` FOREIGN KEY (`artist_id`) REFERENCES `artist` (`id`) ON DELETE CASCADE,
  CONSTRAINT `artist_album_ibfk_2` FOREIGN KEY (`album_id`) REFERENCES `album` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='艺术家与专辑关联表,存储艺术家与专辑的关系';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `artist_album`
--

LOCK TABLES `artist_album` WRITE;
/*!40000 ALTER TABLE `artist_album` DISABLE KEYS */;
INSERT INTO `artist_album` VALUES (1,1),(1,2),(2,3),(2,4);
/*!40000 ALTER TABLE `artist_album` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment`
--

DROP TABLE IF EXISTS `comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comment` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '评论ID,主键,自动递增',
  `user_id` int NOT NULL COMMENT '用户ID,关联用户表的ID',
  `song_id` int NOT NULL COMMENT '歌曲ID,关联歌曲表的ID',
  `text` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '评论内容,不能为空',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论创建时间,默认为当前时间',
  `parent_id` int DEFAULT NULL COMMENT '父评论ID,允许为空',
  `reply_to_user_id` int DEFAULT NULL COMMENT '回复目标用户ID',
  `like_count` int NOT NULL DEFAULT '0' COMMENT '点赞数,默认为0',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '评论更新时间,默认为当前时间',
  `is_root` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否为根评论 (TRUE: 根评论, FALSE: 回复评论)',
  PRIMARY KEY (`id`),
  KEY `reply_to_user_id` (`reply_to_user_id`),
  KEY `idx_comment_song_id` (`song_id`),
  KEY `idx_comment_user_id` (`user_id`),
  KEY `idx_comment_parent_id` (`parent_id`),
  CONSTRAINT `comment_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `comment_ibfk_2` FOREIGN KEY (`song_id`) REFERENCES `song` (`id`),
  CONSTRAINT `comment_ibfk_3` FOREIGN KEY (`parent_id`) REFERENCES `comment` (`id`) ON DELETE SET NULL,
  CONSTRAINT `comment_ibfk_4` FOREIGN KEY (`reply_to_user_id`) REFERENCES `user` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='评论表,存储用户对歌曲的评论';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
INSERT INTO `comment` VALUES (1,2,1,'back to UUSR comrade!!!','2024-12-28 11:37:40',NULL,NULL,1,'2025-08-07 16:05:04',1),(2,2,1,'I like this song','2024-12-28 11:38:01',1,2,0,'2024-12-28 11:38:01',0),(3,3,1,'it\'s just like the beach boy\'s california girls','2024-12-28 11:38:54',1,2,0,'2024-12-28 11:38:54',0),(4,1,1,'111','2024-12-28 15:47:36',1,2,1,'2025-08-14 15:39:06',0),(5,1,1,'111','2024-12-29 11:36:04',NULL,NULL,0,'2025-08-07 15:36:55',1),(6,1,1,'111','2024-12-29 11:36:11',1,3,0,'2025-08-14 07:11:53',0),(7,5,1,'喔哦喔哦喔哦~','2025-08-02 14:46:46',1,2,0,'2025-08-02 14:46:46',0),(8,5,2,'hello world test','2025-08-07 16:00:34',NULL,NULL,1,'2025-08-15 06:32:11',1),(9,5,2,'hello world test','2025-08-07 16:00:59',NULL,NULL,0,'2025-08-07 16:00:59',1),(10,5,1,'测试测试测试\n','2025-08-14 07:12:21',NULL,NULL,0,'2025-08-14 07:45:28',1),(12,5,1,'喔哦喔哦喔哦~','2025-08-14 07:22:46',NULL,NULL,0,'2025-08-14 07:22:46',1),(13,5,1,'1','2025-08-14 07:23:14',NULL,NULL,0,'2025-08-14 07:23:14',1),(14,5,1,'2','2025-08-14 07:23:17',NULL,NULL,0,'2025-08-14 07:23:17',1),(15,5,1,'3','2025-08-14 07:23:19',NULL,NULL,0,'2025-08-14 07:23:19',1),(16,5,1,'4','2025-08-14 07:23:33',NULL,NULL,0,'2025-08-14 07:23:33',1),(17,5,1,'5','2025-08-14 07:23:34',NULL,NULL,0,'2025-08-14 07:23:34',1),(18,5,1,'6','2025-08-14 07:23:36',NULL,NULL,0,'2025-08-14 07:23:36',1),(19,5,1,'7','2025-08-14 07:23:42',NULL,NULL,0,'2025-08-14 07:23:42',1),(20,5,1,'8','2025-08-14 07:23:44',NULL,NULL,0,'2025-08-14 07:23:44',1),(21,5,1,'9','2025-08-14 07:23:46',NULL,NULL,0,'2025-08-14 07:23:46',1),(22,5,1,'10','2025-08-14 07:23:52',NULL,NULL,0,'2025-08-14 07:23:52',1),(23,15,2,'hello','2025-08-15 06:30:36',9,5,0,'2025-08-15 06:30:36',0);
/*!40000 ALTER TABLE `comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment_like`
--

DROP TABLE IF EXISTS `comment_like`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comment_like` (
  `id` int NOT NULL AUTO_INCREMENT,
  `comment_id` int NOT NULL COMMENT '评论ID',
  `user_id` int NOT NULL COMMENT '用户ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '点赞时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_comment_user` (`comment_id`,`user_id`) COMMENT '确保用户对同一评论只能点赞一次',
  KEY `fk_comment_like_user` (`user_id`),
  CONSTRAINT `fk_comment_like_comment` FOREIGN KEY (`comment_id`) REFERENCES `comment` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_comment_like_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='评论点赞表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment_like`
--

LOCK TABLES `comment_like` WRITE;
/*!40000 ALTER TABLE `comment_like` DISABLE KEYS */;
INSERT INTO `comment_like` VALUES (1,2,1,'2024-12-28 15:47:29'),(2,3,1,'2024-12-28 15:47:30'),(7,7,5,'2025-08-02 14:46:49'),(9,1,5,'2025-08-07 16:05:04'),(20,4,15,'2025-08-14 15:39:06'),(23,8,15,'2025-08-15 06:32:11');
/*!40000 ALTER TABLE `comment_like` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `follow`
--

DROP TABLE IF EXISTS `follow`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `follow` (
  `id` int NOT NULL AUTO_INCREMENT,
  `follower_id` int NOT NULL,
  `following_id` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_follow` (`follower_id`,`following_id`),
  KEY `following_id` (`following_id`),
  CONSTRAINT `follow_ibfk_1` FOREIGN KEY (`follower_id`) REFERENCES `user` (`id`) ON DELETE CASCADE,
  CONSTRAINT `follow_ibfk_2` FOREIGN KEY (`following_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `follow`
--

LOCK TABLES `follow` WRITE;
/*!40000 ALTER TABLE `follow` DISABLE KEYS */;
/*!40000 ALTER TABLE `follow` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `manager`
--

DROP TABLE IF EXISTS `manager`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `manager` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '管理员ID,主键,自动递增',
  `username` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '管理员用户名,唯一且不能为空',
  `password` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '管理员密码,不能为空',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='管理员表,存储管理员的基本信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `manager`
--

LOCK TABLES `manager` WRITE;
/*!40000 ALTER TABLE `manager` DISABLE KEYS */;
/*!40000 ALTER TABLE `manager` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `playlist_info`
--

DROP TABLE IF EXISTS `playlist_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `playlist_info` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '播放列表ID,主键,自动递增',
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '播放列表名称,不能为空',
  `user_id` int NOT NULL COMMENT '用户ID,关联用户表的ID',
  `cover` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '/assets/covers/playlists/default.jpg' COMMENT '播放列表封面,允许为空',
  `description` text COLLATE utf8mb4_unicode_ci COMMENT '播放列表描述,允许为空',
  `is_public` tinyint(1) DEFAULT '0' COMMENT '是否公开,默认为不公开',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间,默认为当前时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间,默认为当前时间',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `playlist_info_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='播放列表信息表,存储播放列表的基本信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `playlist_info`
--

LOCK TABLES `playlist_info` WRITE;
/*!40000 ALTER TABLE `playlist_info` DISABLE KEYS */;
INSERT INTO `playlist_info` VALUES (1,'night car songs',1,'/assets/covers/playlists/default.jpg','111',0,'2024-12-28 11:35:31','2024-12-29 11:36:44'),(3,'DOOMER SONG',2,'/assets/covers/playlists/default.jpg','awesrasfaesgesagswagtaewgtw',0,'2024-12-28 11:42:16','2024-12-28 11:42:34'),(4,'hello world',1,'/assets/covers/playlists/default.jpg',NULL,0,'2024-12-28 14:00:04','2024-12-28 15:46:28'),(7,'test1.0',5,'/assets/covers/playlists/default.jpg','这是一个不可爱的测试歌单\n',1,'2025-08-02 14:47:39','2025-08-14 02:32:46'),(8,'test2',5,'/assets/covers/playlists/default.jpg','test update',1,'2025-08-07 06:42:33','2025-08-07 06:44:42'),(12,'test3',5,'/assets/covers/playlists/default.jpg','my playlist test3',1,'2025-08-07 06:45:46','2025-08-12 09:40:10'),(14,'test4',5,'/assets/covers/playlists/default.jpg','这是第四个测试歌单',0,'2025-08-12 09:38:48','2025-08-12 09:38:48'),(15,'test5',5,'/assets/covers/playlists/default.jpg','',0,'2025-08-12 13:21:00','2025-08-12 13:21:00'),(16,'test6',5,'/assets/covers/playlists/default.jpg','test6',0,'2025-08-12 14:45:49','2025-08-12 14:57:53'),(18,'test7',5,'/assets/covers/playlists/default.jpg','这是一个不可爱的测试歌单7',1,'2025-08-14 03:19:50','2025-08-14 03:19:50'),(19,'test8',5,'/assets/covers/playlists/default.jpg','测试测试测试',0,'2025-08-14 14:47:20','2025-08-14 14:47:20'),(20,'test1',15,'/assets/covers/playlists/default.jpg','测试用户1的测试歌单1',1,'2025-08-15 06:35:38','2025-08-15 06:35:38');
/*!40000 ALTER TABLE `playlist_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `playlist_songs`
--

DROP TABLE IF EXISTS `playlist_songs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `playlist_songs` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '播放列表歌曲关联ID,主键,自动递增',
  `playlist_id` int NOT NULL COMMENT '播放列表ID,关联播放列表表的ID',
  `song_id` int NOT NULL COMMENT '歌曲ID,关联歌曲表的ID',
  `user_id` int NOT NULL COMMENT '用户ID,关联用户表的ID',
  `added_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间,默认为当前时间',
  PRIMARY KEY (`id`),
  KEY `playlist_id` (`playlist_id`),
  KEY `song_id` (`song_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `playlist_songs_ibfk_1` FOREIGN KEY (`playlist_id`) REFERENCES `playlist_info` (`id`),
  CONSTRAINT `playlist_songs_ibfk_2` FOREIGN KEY (`song_id`) REFERENCES `song` (`id`),
  CONSTRAINT `playlist_songs_ibfk_3` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='播放列表歌曲关联表,存储播放列表与歌曲的关系';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `playlist_songs`
--

LOCK TABLES `playlist_songs` WRITE;
/*!40000 ALTER TABLE `playlist_songs` DISABLE KEYS */;
INSERT INTO `playlist_songs` VALUES (2,1,36,1,'2024-12-28 11:36:08'),(3,3,44,2,'2024-12-28 11:42:34'),(4,4,1,1,'2024-12-28 15:46:28'),(6,7,1,5,'2025-08-02 14:47:50'),(8,8,3,5,'2025-08-07 06:48:27'),(10,7,2,5,'2025-08-11 14:59:20'),(11,7,3,5,'2025-08-11 14:59:23'),(12,12,3,5,'2025-08-12 05:03:16'),(13,15,1,5,'2025-08-12 13:21:02'),(14,16,1,5,'2025-08-12 14:45:51'),(15,7,37,5,'2025-08-13 03:17:23'),(16,18,1,5,'2025-08-14 03:30:32'),(17,18,2,5,'2025-08-14 03:34:12'),(18,15,42,5,'2025-08-14 03:34:16'),(19,14,2,5,'2025-08-14 06:45:06'),(20,19,42,5,'2025-08-15 04:48:58'),(21,19,12,5,'2025-08-15 04:49:02');
/*!40000 ALTER TABLE `playlist_songs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `song`
--

DROP TABLE IF EXISTS `song`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `song` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '歌曲ID,主键,自动递增',
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '歌曲名称,不能为空',
  `author_id` int NOT NULL COMMENT '歌曲作者ID,不能为空,对应艺术家的ID',
  `album_id` int NOT NULL COMMENT '专辑ID,不能为空,对应专辑的ID',
  `url` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '歌曲存放地址,不能为空',
  `cover` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '歌曲封面,不能为空',
  `duration` int NOT NULL DEFAULT '180' COMMENT '歌曲时长,默认为180秒',
  `play_count` int NOT NULL DEFAULT '0' COMMENT '播放次数,默认为0',
  `like_count` int NOT NULL DEFAULT '0' COMMENT '点赞次数,默认为0',
  `is_public` tinyint(1) DEFAULT '1' COMMENT '是否公开,默认为公开',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '歌曲创建时间,默认为当前时间',
  PRIMARY KEY (`id`),
  KEY `idx_song_name` (`name`),
  KEY `idx_song_album_id` (`album_id`),
  KEY `idx_song_author_id` (`author_id`),
  CONSTRAINT `song_ibfk_1` FOREIGN KEY (`author_id`) REFERENCES `artist` (`id`),
  CONSTRAINT `song_ibfk_2` FOREIGN KEY (`album_id`) REFERENCES `album` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=60 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='歌曲表,存储歌曲的基本信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `song`
--

LOCK TABLES `song` WRITE;
/*!40000 ALTER TABLE `song` DISABLE KEYS */;
INSERT INTO `song` VALUES (1,'Back in the U.S.S.R.',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',183,0,0,1,'2024-12-28 11:20:05'),(2,'Dear Prudence',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',189,0,0,1,'2024-12-28 11:20:05'),(3,'Glass Onion',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',154,0,0,1,'2024-12-28 11:20:05'),(4,'Ob-La-Di, Ob-La-Da',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',113,0,0,1,'2024-12-28 11:20:05'),(5,'Wild Honey Pie',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',60,0,0,1,'2024-12-28 11:20:05'),(6,'The Continuing Story of Bungalow Bill',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',205,0,0,1,'2024-12-28 11:20:05'),(7,'While My Guitar Gently Weeps',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',185,0,0,1,'2024-12-28 11:20:05'),(8,'Happiness Is a Warm Gun',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',183,0,0,1,'2024-12-28 11:20:05'),(9,'Martha My Dear',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',158,0,0,1,'2024-12-28 11:20:05'),(10,'I’m So Tired',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',202,0,0,1,'2024-12-28 11:20:05'),(11,'Blackbird',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',144,0,0,1,'2024-12-28 11:20:05'),(12,'Piggies',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',160,0,0,1,'2024-12-28 11:20:05'),(13,'Rocky Raccoon',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',222,0,0,1,'2024-12-28 11:20:05'),(14,'Don’t Pass Me By',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',200,0,0,1,'2024-12-28 11:20:05'),(15,'Why Don’t We Do It in the Road?',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',91,0,0,1,'2024-12-28 11:20:05'),(16,'I Will',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',134,0,0,1,'2024-12-28 11:20:05'),(17,'Junk',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',173,0,0,1,'2024-12-28 11:20:05'),(18,'Yer Blues',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',215,0,0,1,'2024-12-28 11:20:05'),(19,'Mother Nature’s Son',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',157,0,0,1,'2024-12-28 11:20:05'),(20,'Everybody’s Got Something to Hide Except Me and My Monkey',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',152,0,0,1,'2024-12-28 11:20:05'),(21,'Sexy Sadie',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',176,0,0,1,'2024-12-28 11:20:05'),(22,'Helter Skelter',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',271,0,0,1,'2024-12-28 11:20:05'),(23,'Long, Long, Long',1,1,'/assets/music/default.mp3','/assets/covers/albums/the_beatles.jpg',157,0,0,1,'2024-12-28 11:20:05'),(24,'Come Together',1,2,'/assets/music/default.mp3','/assets/covers/albums/abbey_road.jpg',259,0,0,1,'2024-12-28 11:20:05'),(25,'Something',1,2,'/assets/music/default.mp3','/assets/covers/albums/abbey_road.jpg',182,0,0,1,'2024-12-28 11:20:05'),(26,'Maxwell’s Silver Hammer',1,2,'/assets/music/default.mp3','/assets/covers/albums/abbey_road.jpg',207,0,0,1,'2024-12-28 11:20:05'),(27,'Oh! Darling',1,2,'/assets/music/default.mp3','/assets/covers/albums/abbey_road.jpg',187,0,0,1,'2024-12-28 11:20:05'),(28,'Octopus’s Garden',1,2,'/assets/music/default.mp3','/assets/covers/albums/abbey_road.jpg',174,0,0,1,'2024-12-28 11:20:05'),(29,'I Want You (She’s So Heavy)',1,2,'/assets/music/default.mp3','/assets/covers/albums/abbey_road.jpg',467,0,0,1,'2024-12-28 11:20:05'),(30,'Here Comes the Sun',1,2,'/assets/music/default.mp3','/assets/covers/albums/abbey_road.jpg',185,0,0,1,'2024-12-28 11:20:05'),(31,'Because',1,2,'/assets/music/default.mp3','/assets/covers/albums/abbey_road.jpg',137,0,0,1,'2024-12-28 11:20:05'),(32,'You Never Give Me Your Money',1,2,'/assets/music/default.mp3','/assets/covers/albums/abbey_road.jpg',210,0,0,1,'2024-12-28 11:20:05'),(33,'Sun King',1,2,'/assets/music/default.mp3','/assets/covers/albums/abbey_road.jpg',143,0,0,1,'2024-12-28 11:20:05'),(34,'Mean Mr. Mustard',1,2,'/assets/music/default.mp3','/assets/covers/albums/abbey_road.jpg',99,0,0,1,'2024-12-28 11:20:05'),(35,'Polythene Pam',1,2,'/assets/music/default.mp3','/assets/covers/albums/abbey_road.jpg',102,0,0,1,'2024-12-28 11:20:05'),(36,'She Came in Through the Bathroom Window',1,2,'/assets/music/default.mp3','/assets/covers/albums/abbey_road.jpg',111,0,0,1,'2024-12-28 11:20:05'),(37,'Golden Slumbers',1,2,'/assets/music/default.mp3','/assets/covers/albums/abbey_road.jpg',144,0,0,1,'2024-12-28 11:20:05'),(38,'Carry That Weight',1,2,'/assets/music/default.mp3','/assets/covers/albums/abbey_road.jpg',137,0,0,1,'2024-12-28 11:20:05'),(39,'The End',1,2,'/assets/music/default.mp3','/assets/covers/albums/abbey_road.jpg',139,0,0,1,'2024-12-28 11:20:05'),(40,'Her Majesty',1,2,'/assets/music/default.mp3','/assets/covers/albums/abbey_road.jpg',23,0,0,1,'2024-12-28 11:20:05'),(41,'Rainy Day Women #12 & 35',2,3,'/assets/music/default.mp3','/assets/covers/albums/blonde_on_blonde.jpg',227,0,0,1,'2024-12-28 11:20:05'),(42,'Pledging My Time',2,3,'/assets/music/default.mp3','/assets/covers/albums/blonde_on_blonde.jpg',311,0,0,1,'2024-12-28 11:20:05'),(43,'Visions of Johanna',2,3,'/assets/music/default.mp3','/assets/covers/albums/blonde_on_blonde.jpg',420,0,0,1,'2024-12-28 11:20:05'),(44,'Just Like a Woman',2,3,'/assets/music/default.mp3','/assets/covers/albums/blonde_on_blonde.jpg',262,0,0,1,'2024-12-28 11:20:05'),(45,'Obviously 5 Believers',2,3,'/assets/music/default.mp3','/assets/covers/albums/blonde_on_blonde.jpg',211,0,0,1,'2024-12-28 11:20:05'),(46,'Sad Eyed Lady of the Lowlands',2,3,'/assets/music/default.mp3','/assets/covers/albums/blonde_on_blonde.jpg',696,0,0,1,'2024-12-28 11:20:05'),(47,'4th Time Around',2,3,'/assets/music/default.mp3','/assets/covers/albums/blonde_on_blonde.jpg',188,0,0,1,'2024-12-28 11:20:05'),(48,'I Want You',2,3,'/assets/music/default.mp3','/assets/covers/albums/blonde_on_blonde.jpg',213,0,0,1,'2024-12-28 11:20:05'),(49,'Stuck Inside of Mobile with the Memphis Blues Again',2,3,'/assets/music/default.mp3','/assets/covers/albums/blonde_on_blonde.jpg',359,0,0,1,'2024-12-28 11:20:05'),(50,'Temporary Like Achilles',2,3,'/assets/music/default.mp3','/assets/covers/albums/blonde_on_blonde.jpg',265,0,0,1,'2024-12-28 11:20:05'),(51,'Like a Rolling Stone',2,4,'/assets/music/default.mp3','/assets/covers/albums/highway_61_revisited.jpg',367,0,0,1,'2024-12-28 11:20:05'),(52,'Tombstone Blues',2,4,'/assets/music/default.mp3','/assets/covers/albums/highway_61_revisited.jpg',295,0,0,1,'2024-12-28 11:20:05'),(53,'It Takes a Lot to Laugh, It Takes a Train to Cry',2,4,'/assets/music/default.mp3','/assets/covers/albums/highway_61_revisited.jpg',314,0,0,1,'2024-12-28 11:20:05'),(54,'Ballad of a Thin Man',2,4,'/assets/music/default.mp3','/assets/covers/albums/highway_61_revisited.jpg',245,0,0,1,'2024-12-28 11:20:05'),(55,'From a Buick 6',2,4,'/assets/music/default.mp3','/assets/covers/albums/highway_61_revisited.jpg',215,0,0,1,'2024-12-28 11:20:05'),(56,'Queen Jane Approximately',2,4,'/assets/music/default.mp3','/assets/covers/albums/highway_61_revisited.jpg',339,0,0,1,'2024-12-28 11:20:05'),(57,'Highway 61 Revisited',2,4,'/assets/music/default.mp3','/assets/covers/albums/highway_61_revisited.jpg',280,0,0,1,'2024-12-28 11:20:05'),(58,'Just Like Tom Thumb’s Blues',2,4,'/assets/music/default.mp3','/assets/covers/albums/highway_61_revisited.jpg',330,0,0,1,'2024-12-28 11:20:05'),(59,'Desolation Row',2,4,'/assets/music/default.mp3','/assets/covers/albums/highway_61_revisited.jpg',719,0,0,1,'2024-12-28 11:20:05');
/*!40000 ALTER TABLE `song` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户ID,主键,自动递增',
  `username` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名,唯一且不能为空',
  `password` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户密码,不能为空',
  `name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户姓名,允许为空',
  `phone` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户电话号码,允许为空',
  `email` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户电子邮件,允许为空',
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '/assets/avatars/default-user.jpg' COMMENT '用户头像,允许为空',
  `bio` text COLLATE utf8mb4_unicode_ci,
  `gender` tinyint DEFAULT '0',
  `birthday` date DEFAULT NULL,
  `location` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '用户状态,默认为1',
  `created_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '账户创建时间,默认为当前时间',
  `updated_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '账户更新时间,默认为当前时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_user_phone` (`phone`),
  KEY `idx_user_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户表,存储用户的基本信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'admin','$2a$10$XQ3lTxi12672s7o401OCRe8X/uvhXPZMPPtf08fQgYs48wR1dyx2S','管理员','18971548334','111@111.com','/assets/avatars/default-user.jpg','admin',1,NULL,'111',1,'2024-12-28 11:22:31','2024-12-29 11:37:40'),(2,'user1','$2a$10$VsmbRSJ041kGQ4fbhzmjvO7YM2sLWDUTSUEgo0OrY1XSM.BhYxqi2','doomer no.1','18971548334','111@111.com','/assets/avatars/default-user.jpg','I like music, you too?',1,NULL,'111',1,'2024-12-28 11:37:01','2024-12-28 11:41:24'),(3,'user2','$2a$10$wTp8u5h5iWLE2m5ezAqgnuyovh35yseaezWjGhLM8OSToxLuOcelm','user2','18971548334','111@111.com','/assets/avatars/default-user.jpg',NULL,2,'2024-11-24','111',1,'2024-12-28 11:38:18','2024-12-28 11:38:18'),(4,'Elaine','$2a$10$kaerMaf6Md9dUgshdPr1XOEo0qHKmJn2emKsX8auxJIAZWKPjI7ta','Eglantine','13111111111','example@example.com','/assets/avatars/default-user.jpg','阳光小橘',0,'2025-01-07','nowhere',1,'2025-06-09 14:45:25','2025-06-09 14:45:25'),(5,'hkwxsl','$2a$10$KiEv79t43A/sYNyWAleeNO2V0HbUe54G/20t31VGvP5r0MhB2UmBC','好困我先睡了喵喵','10011111112','example@example.com','/assets/avatars/default-user.jpg','我真的是一只猫, 超级巨大~\n我有多巨大?',2,'2005-03-30','小青家里',1,'2025-08-02 14:42:50','2025-08-14 14:46:53'),(6,'test1','$2a$10$HTR/8eLCUlnt6vsmPRsONe9iYgyjkLRbevdgw/BWZifGUgqk9hl6.',NULL,NULL,NULL,'/assets/avatars/default-user.jpg',NULL,0,NULL,NULL,1,'2025-08-03 09:00:43','2025-08-03 09:00:43'),(7,'test2','$2a$10$21kAXcns94wW1dFX34H0zewycN8a03cCMHtX47y3R3zvChp31slnG',NULL,NULL,NULL,'/assets/avatars/default-user.jpg',NULL,0,NULL,NULL,1,'2025-08-03 18:08:39','2025-08-03 18:08:39'),(8,'test3','$2a$10$xZhQcOWckwe/ptj0DiqfHOEC7qnwkr0VaL4NkCR/UID7KZfsvye2K',NULL,NULL,NULL,'/assets/avatars/default-user.jpg',NULL,0,NULL,NULL,1,'2025-08-03 18:39:56','2025-08-03 18:39:56'),(9,'test4','$2a$10$1QflnuTCAgzT5sG41jzjleos4F4Z1CXq9lECM9ozhQZEDU0b1ftby',NULL,NULL,NULL,'/assets/avatars/default-user.jpg',NULL,0,NULL,NULL,1,'2025-08-04 02:59:27','2025-08-04 02:59:27'),(10,'test5','$2a$10$kYF73nDO2VmPJd9j5RHxPOT//FF4vCanx9/7SeBY6GbeTPhfQFF5.',NULL,NULL,NULL,'/assets/avatars/default-user.jpg',NULL,0,NULL,NULL,1,'2025-08-06 03:36:31','2025-08-06 03:36:31'),(11,'test6','$2a$10$eYG8hTmGTvmQ9ec3i2lE..bT9gg3n2tSOujsx3KQf9k8cV0fxOq3a','test6',NULL,NULL,'/assets/avatars/default-user.jpg',NULL,0,NULL,NULL,1,'2025-08-08 02:15:43','2025-08-08 02:15:43'),(12,'test7','$2a$10$mIFZaew/MbvtFgnj.26HRuUPO2T2SmDqzRr/Pv7d7IB94dIycSfRS','test7',NULL,NULL,'/assets/avatars/default-user.jpg',NULL,0,NULL,NULL,1,'2025-08-08 02:49:58','2025-08-08 02:49:58'),(13,'test8','$2a$10$ddRojhKQc6GyS8Rw20QzAe.vOYTf.hBZApewLxz1PxMwE1kGfsGpK','test8',NULL,NULL,'/assets/avatars/default-user.jpg',NULL,0,NULL,NULL,1,'2025-08-08 02:51:46','2025-08-08 02:51:46'),(14,'test9','$2a$10$W9ykZji7sW0pkmQf5S.aceRv32sBwBM07yfmYVDPAjoD6KclsPl1y','test9','13111111111','example@example.com','/assets/avatars/default-user.jpg','真的是一只猫, 我只会说喵喵~',2,NULL,'广东, 深圳',1,'2025-08-08 02:52:56','2025-08-08 02:52:56'),(15,'testuser','$2a$10$wJh5/SsHKOBWdrTlMTaV3.ahmOwzXh2tR5s4XTY93zT0eoG7GRXe2','测试用户1',NULL,NULL,'/assets/avatars/default-user.jpg',NULL,0,NULL,NULL,1,'2025-08-14 14:47:45','2025-08-14 14:47:45');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_likes`
--

DROP TABLE IF EXISTS `user_likes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_likes` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `song_id` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_user_song` (`user_id`,`song_id`),
  KEY `song_id` (`song_id`),
  CONSTRAINT `user_likes_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `user_likes_ibfk_2` FOREIGN KEY (`song_id`) REFERENCES `song` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=95 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_likes`
--

LOCK TABLES `user_likes` WRITE;
/*!40000 ALTER TABLE `user_likes` DISABLE KEYS */;
INSERT INTO `user_likes` VALUES (2,1,36,'2024-12-28 15:46:58'),(3,1,51,'2024-12-28 15:47:00'),(4,2,1,'2024-12-28 11:41:39'),(5,2,8,'2024-12-28 11:41:44'),(6,1,21,'2024-12-28 13:57:04'),(15,1,8,'2024-12-28 15:48:27'),(21,1,2,'2024-12-29 11:34:45'),(22,1,4,'2024-12-29 11:34:48'),(31,5,10,'2025-08-10 12:17:35'),(32,5,8,'2025-08-10 12:17:36'),(35,5,5,'2025-08-10 12:17:39'),(37,5,12,'2025-08-10 12:17:43'),(39,5,15,'2025-08-10 12:17:45'),(40,5,16,'2025-08-10 12:17:45'),(41,5,17,'2025-08-10 12:17:46'),(48,5,37,'2025-08-13 03:17:17'),(49,5,50,'2025-08-13 08:19:55'),(50,5,49,'2025-08-13 08:19:56'),(51,5,48,'2025-08-13 08:19:57'),(52,5,47,'2025-08-13 08:19:58'),(53,5,46,'2025-08-13 08:19:59'),(55,5,44,'2025-08-13 08:20:00'),(58,5,41,'2025-08-13 08:20:03'),(59,5,42,'2025-08-13 08:58:20'),(64,5,2,'2025-08-14 03:34:56'),(92,5,43,'2025-08-15 04:48:05'),(93,5,23,'2025-08-15 04:50:35'),(94,15,2,'2025-08-15 07:42:13');
/*!40000 ALTER TABLE `user_likes` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-08-16 16:16:31
