-- MySQL dump 10.13  Distrib 8.4.2, for Win64 (x86_64)
--
-- Host: localhost    Database: music_db1
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
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='评论表,存储用户对歌曲的评论';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
INSERT INTO `comment` VALUES (1,2,1,'back to UUSR comrade!!!','2024-12-28 11:37:40',NULL,NULL,1,'2025-08-07 16:05:04',1),(2,2,1,'I like this song','2024-12-28 11:38:01',1,2,0,'2024-12-28 11:38:01',0),(3,3,1,'it\'s just like the beach boy\'s california girls','2024-12-28 11:38:54',1,2,0,'2024-12-28 11:38:54',0),(4,1,1,'111','2024-12-28 15:47:36',1,2,0,'2024-12-28 15:47:36',0),(5,1,1,'111','2024-12-29 11:36:04',NULL,NULL,0,'2025-08-07 15:36:55',1),(6,1,1,'111','2024-12-29 11:36:11',1,3,0,'2024-12-29 11:36:11',0),(7,5,1,'喔哦喔哦喔哦~','2025-08-02 14:46:46',1,2,0,'2025-08-02 14:46:46',0),(8,5,2,'hello world test','2025-08-07 16:00:34',NULL,NULL,0,'2025-08-07 16:00:34',1),(9,5,2,'hello world test','2025-08-07 16:00:59',NULL,NULL,0,'2025-08-07 16:00:59',1);
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
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='评论点赞表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment_like`
--

LOCK TABLES `comment_like` WRITE;
/*!40000 ALTER TABLE `comment_like` DISABLE KEYS */;
INSERT INTO `comment_like` VALUES (1,2,1,'2024-12-28 15:47:29'),(2,3,1,'2024-12-28 15:47:30'),(3,4,5,'2025-08-02 14:46:05'),(4,6,5,'2025-08-02 14:46:07'),(7,7,5,'2025-08-02 14:46:49'),(9,1,5,'2025-08-07 16:05:04');
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
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='播放列表信息表,存储播放列表的基本信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `playlist_info`
--

LOCK TABLES `playlist_info` WRITE;
/*!40000 ALTER TABLE `playlist_info` DISABLE KEYS */;
INSERT INTO `playlist_info` VALUES (1,'night car songs',1,'/assets/covers/playlists/default.jpg','111',0,'2024-12-28 11:35:31','2024-12-29 11:36:44'),(3,'DOOMER SONG',2,'/assets/covers/playlists/default.jpg','awesrasfaesgesagswagtaewgtw',0,'2024-12-28 11:42:16','2024-12-28 11:42:34'),(4,'hello world',1,'/assets/covers/playlists/default.jpg',NULL,0,'2024-12-28 14:00:04','2024-12-28 15:46:28'),(7,'list1',5,'/assets/covers/playlists/default.jpg','list1',1,'2025-08-02 14:47:39','2025-08-02 14:47:50'),(8,'test2',5,'/assets/covers/playlists/default.jpg','test update',1,'2025-08-07 06:42:33','2025-08-07 06:44:42'),(9,'test2',5,'/assets/covers/playlists/default.jpg','my playlist test2',1,'2025-08-07 06:43:00','2025-08-07 06:43:00'),(12,'test3',5,'/assets/covers/playlists/default.jpg','my playlist test3',1,'2025-08-07 06:45:46','2025-08-07 06:45:46'),(13,'test3',5,'/assets/covers/playlists/default.jpg','my playlist test3',1,'2025-08-07 06:45:48','2025-08-07 06:45:48');
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
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='播放列表歌曲关联表,存储播放列表与歌曲的关系';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `playlist_songs`
--

LOCK TABLES `playlist_songs` WRITE;
/*!40000 ALTER TABLE `playlist_songs` DISABLE KEYS */;
INSERT INTO `playlist_songs` VALUES (2,1,36,1,'2024-12-28 11:36:08'),(3,3,44,2,'2024-12-28 11:42:34'),(4,4,1,1,'2024-12-28 15:46:28'),(6,7,1,5,'2025-08-02 14:47:50'),(8,8,3,5,'2025-08-07 06:48:27');
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
INSERT INTO `song` VALUES (1,'Back in the U.S.S.R.',1,1,'/songs/the_beatles_back_in_the_ussr.mp3','/assets/covers/back_in_the_ussr.jpg',183,0,0,1,'2024-12-28 11:20:05'),(2,'Dear Prudence',1,1,'/songs/the_beatles_dear_prudence.mp3','/assets/covers/dear_prudence.jpg',189,0,0,1,'2024-12-28 11:20:05'),(3,'Glass Onion',1,1,'/songs/the_beatles_glass_onion.mp3','/assets/covers/glass_onion.jpg',154,0,0,1,'2024-12-28 11:20:05'),(4,'Ob-La-Di, Ob-La-Da',1,1,'/songs/the_beatles_ob_la_di_ob_la_da.mp3','/assets/covers/ob_la_di_ob_la_da.jpg',113,0,0,1,'2024-12-28 11:20:05'),(5,'Wild Honey Pie',1,1,'/songs/the_beatles_wild_honey_pie.mp3','/assets/covers/wild_honey_pie.jpg',60,0,0,1,'2024-12-28 11:20:05'),(6,'The Continuing Story of Bungalow Bill',1,1,'/songs/the_beatles_bungalow_bill.mp3','/assets/covers/bungalow_bill.jpg',205,0,0,1,'2024-12-28 11:20:05'),(7,'While My Guitar Gently Weeps',1,1,'/songs/the_beatles_while_my_guitar_gently_weeps.mp3','/assets/covers/while_my_guitar_gently_weeps.jpg',185,0,0,1,'2024-12-28 11:20:05'),(8,'Happiness Is a Warm Gun',1,1,'/songs/the_beatles_happiness_is_a_warm_gun.mp3','/assets/covers/happiness_is_a_warm_gun.jpg',183,0,0,1,'2024-12-28 11:20:05'),(9,'Martha My Dear',1,1,'/songs/the_beatles_martha_my_dear.mp3','/assets/covers/martha_my_dear.jpg',158,0,0,1,'2024-12-28 11:20:05'),(10,'I’m So Tired',1,1,'/songs/the_beatles_im_so_tired.mp3','/assets/covers/im_so_tired.jpg',202,0,0,1,'2024-12-28 11:20:05'),(11,'Blackbird',1,1,'/songs/the_beatles_blackbird.mp3','/assets/covers/blackbird.jpg',144,0,0,1,'2024-12-28 11:20:05'),(12,'Piggies',1,1,'/songs/the_beatles_piggies.mp3','/assets/covers/piggies.jpg',160,0,0,1,'2024-12-28 11:20:05'),(13,'Rocky Raccoon',1,1,'/songs/the_beatles_rocky_raccoon.mp3','/assets/covers/rocky_raccoon.jpg',222,0,0,1,'2024-12-28 11:20:05'),(14,'Don’t Pass Me By',1,1,'/songs/the_beatles_dont_pass_me_by.mp3','/assets/covers/dont_pass_me_by.jpg',200,0,0,1,'2024-12-28 11:20:05'),(15,'Why Don’t We Do It in the Road?',1,1,'/songs/the_beatles_why_dont_we_do_it_in_the_road.mp3','/assets/covers/why_dont_we_do_it_in_the_road.jpg',91,0,0,1,'2024-12-28 11:20:05'),(16,'I Will',1,1,'/songs/the_beatles_i_will.mp3','/assets/covers/i_will.jpg',134,0,0,1,'2024-12-28 11:20:05'),(17,'Junk',1,1,'/songs/the_beatles_junk.mp3','/assets/covers/junk.jpg',173,0,0,1,'2024-12-28 11:20:05'),(18,'Yer Blues',1,1,'/songs/the_beatles_yer_blues.mp3','/assets/covers/yer_blues.jpg',215,0,0,1,'2024-12-28 11:20:05'),(19,'Mother Nature’s Son',1,1,'/songs/the_beatles_mother_natures_son.mp3','/assets/covers/mother_natures_son.jpg',157,0,0,1,'2024-12-28 11:20:05'),(20,'Everybody’s Got Something to Hide Except Me and My Monkey',1,1,'/songs/the_beatles_everybodys_got_something.mp3','/assets/covers/everybodys_got_something.jpg',152,0,0,1,'2024-12-28 11:20:05'),(21,'Sexy Sadie',1,1,'/songs/the_beatles_sexy_sadie.mp3','/assets/covers/sexy_sadie.jpg',176,0,0,1,'2024-12-28 11:20:05'),(22,'Helter Skelter',1,1,'/songs/the_beatles_helter_skelter.mp3','/assets/covers/helter_skelter.jpg',271,0,0,1,'2024-12-28 11:20:05'),(23,'Long, Long, Long',1,1,'/songs/the_beatles_long_long_long.mp3','/assets/covers/long_long_long.jpg',157,0,0,1,'2024-12-28 11:20:05'),(24,'Come Together',1,2,'/songs/abbey_road_come_together.mp3','/assets/covers/come_together.jpg',259,0,0,1,'2024-12-28 11:20:05'),(25,'Something',1,2,'/songs/abbey_road_something.mp3','/assets/covers/something.jpg',182,0,0,1,'2024-12-28 11:20:05'),(26,'Maxwell’s Silver Hammer',1,2,'/songs/abbey_road_maxwells_silver_hammer.mp3','/assets/covers/maxwells_silver_hammer.jpg',207,0,0,1,'2024-12-28 11:20:05'),(27,'Oh! Darling',1,2,'/songs/abbey_road_oh_darling.mp3','/assets/covers/oh_darling.jpg',187,0,0,1,'2024-12-28 11:20:05'),(28,'Octopus’s Garden',1,2,'/songs/abbey_road_octopuss_garden.mp3','/assets/covers/octopuss_garden.jpg',174,0,0,1,'2024-12-28 11:20:05'),(29,'I Want You (She’s So Heavy)',1,2,'/songs/abbey_road_i_want_you.mp3','/assets/covers/i_want_you.jpg',467,0,0,1,'2024-12-28 11:20:05'),(30,'Here Comes the Sun',1,2,'/songs/abbey_road_here_comes_the_sun.mp3','/assets/covers/here_comes_the_sun.jpg',185,0,0,1,'2024-12-28 11:20:05'),(31,'Because',1,2,'/songs/abbey_road_because.mp3','/assets/covers/because.jpg',137,0,0,1,'2024-12-28 11:20:05'),(32,'You Never Give Me Your Money',1,2,'/songs/abbey_road_you_never_give_me.mp3','/assets/covers/you_never_give_me.jpg',210,0,0,1,'2024-12-28 11:20:05'),(33,'Sun King',1,2,'/songs/abbey_road_sun_king.mp3','/assets/covers/sun_king.jpg',143,0,0,1,'2024-12-28 11:20:05'),(34,'Mean Mr. Mustard',1,2,'/songs/abbey_road_mean_mr_mustard.mp3','/assets/covers/mean_mr_mustard.jpg',99,0,0,1,'2024-12-28 11:20:05'),(35,'Polythene Pam',1,2,'/songs/abbey_road_polythene_pam.mp3','/assets/covers/polythene_pam.jpg',102,0,0,1,'2024-12-28 11:20:05'),(36,'She Came in Through the Bathroom Window',1,2,'/songs/abbey_road_she_came_in_through_the_bathroom_window.mp3','/assets/covers/she_came_in_through_the_bathroom_window.jpg',111,0,0,1,'2024-12-28 11:20:05'),(37,'Golden Slumbers',1,2,'/songs/abbey_road_golden_slumbers.mp3','/assets/covers/golden_slumbers.jpg',144,0,0,1,'2024-12-28 11:20:05'),(38,'Carry That Weight',1,2,'/songs/abbey_road_carry_that_weight.mp3','/assets/covers/carry_that_weight.jpg',137,0,0,1,'2024-12-28 11:20:05'),(39,'The End',1,2,'/songs/abbey_road_the_end.mp3','/assets/covers/the_end.jpg',139,0,0,1,'2024-12-28 11:20:05'),(40,'Her Majesty',1,2,'/songs/abbey_road_her_majesty.mp3','/assets/covers/her_majesty.jpg',23,0,0,1,'2024-12-28 11:20:05'),(41,'Rainy Day Women #12 & 35',2,3,'/songs/bob_dylan_rainy_day_women.mp3','/assets/covers/rainy_day_women.jpg',227,0,0,1,'2024-12-28 11:20:05'),(42,'Pledging My Time',2,3,'/songs/bob_dylan_pledging_my_time.mp3','/assets/covers/pledging_my_time.jpg',311,0,0,1,'2024-12-28 11:20:05'),(43,'Visions of Johanna',2,3,'/songs/bob_dylan_visions_of_johanna.mp3','/assets/covers/visions_of_johanna.jpg',420,0,0,1,'2024-12-28 11:20:05'),(44,'Just Like a Woman',2,3,'/songs/bob_dylan_just_like_a_woman.mp3','/assets/covers/just_like_a_woman.jpg',262,0,0,1,'2024-12-28 11:20:05'),(45,'Obviously 5 Believers',2,3,'/songs/bob_dylan_obviously_5_believers.mp3','/assets/covers/obviously_5_believers.jpg',211,0,0,1,'2024-12-28 11:20:05'),(46,'Sad Eyed Lady of the Lowlands',2,3,'/songs/bob_dylan_sad_eyed_lady.mp3','/assets/covers/sad_eyed_lady.jpg',696,0,0,1,'2024-12-28 11:20:05'),(47,'4th Time Around',2,3,'/songs/bob_dylan_4th_time_around.mp3','/assets/covers/4th_time_around.jpg',188,0,0,1,'2024-12-28 11:20:05'),(48,'I Want You',2,3,'/songs/bob_dylan_i_want_you.mp3','/assets/covers/i_want_you.jpg',213,0,0,1,'2024-12-28 11:20:05'),(49,'Stuck Inside of Mobile with the Memphis Blues Again',2,3,'/songs/bob_dylan_stuck_inside_mobile.mp3','/assets/covers/stuck_inside_mobile.jpg',359,0,0,1,'2024-12-28 11:20:05'),(50,'Temporary Like Achilles',2,3,'/songs/bob_dylan_temporary_like_achilles.mp3','/assets/covers/temporary_like_achilles.jpg',265,0,0,1,'2024-12-28 11:20:05'),(51,'Like a Rolling Stone',2,4,'/songs/bob_dylan_like_a_rolling_stone.mp3','/assets/covers/like_a_rolling_stone.jpg',367,0,0,1,'2024-12-28 11:20:05'),(52,'Tombstone Blues',2,4,'/songs/bob_dylan_tombstone_blues.mp3','/assets/covers/tombstone_blues.jpg',295,0,0,1,'2024-12-28 11:20:05'),(53,'It Takes a Lot to Laugh, It Takes a Train to Cry',2,4,'/songs/bob_dylan_it_takes_a_lot.mp3','/assets/covers/it_takes_a_lot.jpg',314,0,0,1,'2024-12-28 11:20:05'),(54,'Ballad of a Thin Man',2,4,'/songs/bob_dylan_ballad_of_a_thin_man.mp3','/assets/covers/ballad_of_a_thin_man.jpg',245,0,0,1,'2024-12-28 11:20:05'),(55,'From a Buick 6',2,4,'/songs/bob_dylan_from_a_buick_6.mp3','/assets/covers/from_a_buick_6.jpg',215,0,0,1,'2024-12-28 11:20:05'),(56,'Queen Jane Approximately',2,4,'/songs/bob_dylan_queen_jane_approximately.mp3','/assets/covers/queen_jane_approximately.jpg',339,0,0,1,'2024-12-28 11:20:05'),(57,'Highway 61 Revisited',2,4,'/songs/bob_dylan_highway_61_revisited.mp3','/assets/covers/highway_61_revisited.jpg',280,0,0,1,'2024-12-28 11:20:05'),(58,'Just Like Tom Thumb’s Blues',2,4,'/songs/bob_dylan_just_like_tom_thumbs_blues.mp3','/assets/covers/just_like_tom_thumbs_blues.jpg',330,0,0,1,'2024-12-28 11:20:05'),(59,'Desolation Row',2,4,'/songs/bob_dylan_desolation_row.mp3','/assets/covers/desolation_row.jpg',719,0,0,1,'2024-12-28 11:20:05');
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
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户表,存储用户的基本信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'admin','$2a$10$XQ3lTxi12672s7o401OCRe8X/uvhXPZMPPtf08fQgYs48wR1dyx2S','管理员','18971548334','111@111.com','/assets/avatars/default-user.jpg','admin',1,NULL,'111',1,'2024-12-28 11:22:31','2024-12-29 11:37:40'),(2,'user1','$2a$10$VsmbRSJ041kGQ4fbhzmjvO7YM2sLWDUTSUEgo0OrY1XSM.BhYxqi2','doomer no.1','18971548334','111@111.com','/assets/avatars/default-user.jpg','I like music, you too?',1,NULL,'111',1,'2024-12-28 11:37:01','2024-12-28 11:41:24'),(3,'user2','$2a$10$wTp8u5h5iWLE2m5ezAqgnuyovh35yseaezWjGhLM8OSToxLuOcelm','user2','18971548334','111@111.com','/assets/avatars/default-user.jpg',NULL,2,'2024-11-24','111',1,'2024-12-28 11:38:18','2024-12-28 11:38:18'),(4,'Elaine','$2a$10$kaerMaf6Md9dUgshdPr1XOEo0qHKmJn2emKsX8auxJIAZWKPjI7ta','Eglantine','13111111111','example@example.com','/assets/avatars/default-user.jpg','阳光小橘',0,'2025-01-07','nowhere',1,'2025-06-09 14:45:25','2025-06-09 14:45:25'),(5,'hkwxsl','$2a$10$KiEv79t43A/sYNyWAleeNO2V0HbUe54G/20t31VGvP5r0MhB2UmBC','好困我先睡了喵','10011111112','example@example.com','/assets/avatars/default-user.jpg','我真的是一只猫,超级巨大',2,'2025-07-31','123',1,'2025-08-02 14:42:50','2025-08-06 03:37:35'),(6,'test1','$2a$10$HTR/8eLCUlnt6vsmPRsONe9iYgyjkLRbevdgw/BWZifGUgqk9hl6.',NULL,NULL,NULL,'/assets/avatars/default-user.jpg',NULL,0,NULL,NULL,1,'2025-08-03 09:00:43','2025-08-03 09:00:43'),(7,'test2','$2a$10$21kAXcns94wW1dFX34H0zewycN8a03cCMHtX47y3R3zvChp31slnG',NULL,NULL,NULL,'/assets/avatars/default-user.jpg',NULL,0,NULL,NULL,1,'2025-08-03 18:08:39','2025-08-03 18:08:39'),(8,'test3','$2a$10$xZhQcOWckwe/ptj0DiqfHOEC7qnwkr0VaL4NkCR/UID7KZfsvye2K',NULL,NULL,NULL,'/assets/avatars/default-user.jpg',NULL,0,NULL,NULL,1,'2025-08-03 18:39:56','2025-08-03 18:39:56'),(9,'test4','$2a$10$1QflnuTCAgzT5sG41jzjleos4F4Z1CXq9lECM9ozhQZEDU0b1ftby',NULL,NULL,NULL,'/assets/avatars/default-user.jpg',NULL,0,NULL,NULL,1,'2025-08-04 02:59:27','2025-08-04 02:59:27'),(10,'test5','$2a$10$kYF73nDO2VmPJd9j5RHxPOT//FF4vCanx9/7SeBY6GbeTPhfQFF5.',NULL,NULL,NULL,'/assets/avatars/default-user.jpg',NULL,0,NULL,NULL,1,'2025-08-06 03:36:31','2025-08-06 03:36:31'),(11,'test6','$2a$10$eYG8hTmGTvmQ9ec3i2lE..bT9gg3n2tSOujsx3KQf9k8cV0fxOq3a','test6',NULL,NULL,'/assets/avatars/default-user.jpg',NULL,0,NULL,NULL,1,'2025-08-08 02:15:43','2025-08-08 02:15:43'),(12,'test7','$2a$10$mIFZaew/MbvtFgnj.26HRuUPO2T2SmDqzRr/Pv7d7IB94dIycSfRS','test7',NULL,NULL,'/assets/avatars/default-user.jpg',NULL,0,NULL,NULL,1,'2025-08-08 02:49:58','2025-08-08 02:49:58'),(13,'test8','$2a$10$ddRojhKQc6GyS8Rw20QzAe.vOYTf.hBZApewLxz1PxMwE1kGfsGpK','test8',NULL,NULL,'/assets/avatars/default-user.jpg',NULL,0,NULL,NULL,1,'2025-08-08 02:51:46','2025-08-08 02:51:46'),(14,'test9','$2a$10$W9ykZji7sW0pkmQf5S.aceRv32sBwBM07yfmYVDPAjoD6KclsPl1y','test9','13111111111','example@example.com','/assets/avatars/default-user.jpg','真的是一只猫, 我只会说喵喵~',2,NULL,'广东, 深圳',1,'2025-08-08 02:52:56','2025-08-08 02:52:56');
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
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_likes`
--

LOCK TABLES `user_likes` WRITE;
/*!40000 ALTER TABLE `user_likes` DISABLE KEYS */;
INSERT INTO `user_likes` VALUES (2,1,36,'2024-12-28 15:46:58'),(3,1,51,'2024-12-28 15:47:00'),(4,2,1,'2024-12-28 11:41:39'),(5,2,8,'2024-12-28 11:41:44'),(6,1,21,'2024-12-28 13:57:04'),(15,1,8,'2024-12-28 15:48:27'),(21,1,2,'2024-12-29 11:34:45'),(22,1,4,'2024-12-29 11:34:48'),(23,5,1,'2025-08-02 14:45:45'),(24,5,2,'2025-08-07 02:33:40');
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

-- Dump completed on 2025-08-08 19:40:17
