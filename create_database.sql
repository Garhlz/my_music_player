CREATE DATABASE IF NOT EXISTS `music_player`
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci;

USE `music_player`;

-- 创建用户表
CREATE TABLE `user` (
    `id` INT PRIMARY KEY AUTO_INCREMENT COMMENT '用户ID,主键,自动递增',
    `username` VARCHAR(50) NOT NULL UNIQUE COMMENT '用户名,唯一且不能为空',
    `password` VARCHAR(100) NOT NULL COMMENT '用户密码,不能为空',
    `name` VARCHAR(50) DEFAULT NULL COMMENT '用户姓名,允许为空',
    `phone` VARCHAR(20) DEFAULT NULL COMMENT '用户电话号码,允许为空',
    `email` VARCHAR(100) DEFAULT NULL COMMENT '用户电子邮件,允许为空',
    `sex` VARCHAR(10) DEFAULT NULL COMMENT '用户性别,允许为空',
    `avatar` VARCHAR(255) DEFAULT '/assets/avatars/default-user.jpg' COMMENT '用户头像,允许为空',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '用户状态,默认为1',
    `created_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '账户创建时间,默认为当前时间',
    `updated_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '账户更新时间,默认为当前时间'
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='用户表,存储用户的基本信息';

-- 创建管理员表
CREATE TABLE `manager` (
    `id` INT PRIMARY KEY AUTO_INCREMENT COMMENT '管理员ID,主键,自动递增',
    `username` VARCHAR(50) NOT NULL UNIQUE COMMENT '管理员用户名,唯一且不能为空',
    `password` VARCHAR(100) NOT NULL COMMENT '管理员密码,不能为空'
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='管理员表,存储管理员的基本信息';

-- 创建艺术家表
CREATE TABLE `artist` (
    `id` INT PRIMARY KEY AUTO_INCREMENT COMMENT '艺术家ID,主键,自动递增',
    `name` VARCHAR(100) NOT NULL COMMENT '艺术家名称,不能为空',
    `sex` VARCHAR(10) DEFAULT NULL COMMENT '艺术家性别,允许为空',
    `avatar` VARCHAR(255) DEFAULT '/assets/avatars/default-artist.jpg' COMMENT '艺术家头像,允许为空',
    `description` TEXT DEFAULT NULL COMMENT '艺术家描述,允许为空',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间,默认为当前时间'
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='艺术家表,存储艺术家的基本信息';

-- 创建专辑表
CREATE TABLE `album` (
    `id` INT PRIMARY KEY AUTO_INCREMENT COMMENT '专辑ID,主键,自动递增',
    `name` VARCHAR(100) NOT NULL COMMENT '专辑名称,不能为空',
    `cover` VARCHAR(255) NOT NULL COMMENT '专辑封面,不能为空',
    `release_date` DATE DEFAULT NULL COMMENT '专辑发行日期,允许为空',
    `description` TEXT DEFAULT NULL COMMENT '专辑描述,允许为空',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间,默认为当前时间'
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='专辑表,存储专辑的基本信息';

-- 创建艺术家与专辑关联表 (多对多关系)
CREATE TABLE `artist_album` (
    `artist_id` INT NOT NULL COMMENT '艺术家ID,关联艺术家表的ID',
    `album_id` INT NOT NULL COMMENT '专辑ID,关联专辑表的ID',
    PRIMARY KEY (`artist_id`, `album_id`),
    FOREIGN KEY (`artist_id`) REFERENCES `artist` (`id`) ON DELETE CASCADE,
    FOREIGN KEY (`album_id`) REFERENCES `album` (`id`) ON DELETE CASCADE
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='艺术家与专辑关联表,存储艺术家与专辑的关系';

-- 创建歌曲表
CREATE TABLE `song` (
    `id` INT PRIMARY KEY AUTO_INCREMENT COMMENT '歌曲ID,主键,自动递增',
    `name` VARCHAR(100) NOT NULL COMMENT '歌曲名称,不能为空',
    `author_id` INT NOT NULL COMMENT '歌曲作者ID,不能为空,对应艺术家的ID',
    `album_id` INT NOT NULL COMMENT '专辑ID,不能为空,对应专辑的ID',
    `url` VARCHAR(255) NOT NULL COMMENT '歌曲存放地址,不能为空',
    `cover` VARCHAR(255) NOT NULL COMMENT '歌曲封面,不能为空',
    `duration` INT NOT NULL DEFAULT 180 COMMENT '歌曲时长,默认为180秒',
    `play_count` INT NOT NULL DEFAULT 0 COMMENT '播放次数,默认为0',
    `like_count` INT NOT NULL DEFAULT 0 COMMENT '点赞次数,默认为0',
    `is_public` BOOLEAN DEFAULT TRUE COMMENT '是否公开,默认为公开',
    `create_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '歌曲创建时间,默认为当前时间',
    FOREIGN KEY (`author_id`) REFERENCES `artist` (`id`),
    FOREIGN KEY (`album_id`) REFERENCES `album` (`id`)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='歌曲表,存储歌曲的基本信息';

-- 创建播放列表信息表
CREATE TABLE `playlist_info` (
    `id` INT PRIMARY KEY AUTO_INCREMENT COMMENT '播放列表ID,主键,自动递增',
    `name` VARCHAR(100) NOT NULL COMMENT '播放列表名称,不能为空',
    `user_id` INT NOT NULL COMMENT '用户ID,关联用户表的ID',
    `cover` VARCHAR(255) DEFAULT '/assets/covers/playlists/default.jpg' COMMENT '播放列表封面,允许为空',
    `description` TEXT DEFAULT NULL COMMENT '播放列表描述,允许为空',
    `is_public` BOOLEAN DEFAULT FALSE COMMENT '是否公开,默认为不公开',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间,默认为当前时间',
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间,默认为当前时间',
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='播放列表信息表,存储播放列表的基本信息';

-- 创建播放列表歌曲关联表
CREATE TABLE `playlist_songs` (
    `id` INT PRIMARY KEY AUTO_INCREMENT COMMENT '播放列表歌曲关联ID,主键,自动递增',
    `playlist_id` INT NOT NULL COMMENT '播放列表ID,关联播放列表表的ID',
    `song_id` INT NOT NULL COMMENT '歌曲ID,关联歌曲表的ID',
    `user_id` INT NOT NULL COMMENT '用户ID,关联用户表的ID',
    `added_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间,默认为当前时间',
    FOREIGN KEY (`playlist_id`) REFERENCES `playlist_info` (`id`),
    FOREIGN KEY (`song_id`) REFERENCES `song` (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='播放列表歌曲关联表,存储播放列表与歌曲的关系';

-- 创建点赞表
CREATE TABLE `like` (
    `id` INT PRIMARY KEY AUTO_INCREMENT COMMENT '点赞ID,主键,自动递增',
    `user_id` INT NOT NULL COMMENT '用户ID,关联用户表的ID',
    `song_id` INT NOT NULL COMMENT '歌曲ID,关联歌曲表的ID',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '点赞时间,默认为当前时间',
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
    FOREIGN KEY (`song_id`) REFERENCES `song` (`id`),
    UNIQUE KEY `unique_user_song` (`user_id`, `song_id`)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='点赞表,存储用户对歌曲的点赞记录';

-- 创建评论表
CREATE TABLE `comment` (
    `id` INT PRIMARY KEY AUTO_INCREMENT COMMENT '评论ID,主键,自动递增',
    `user_id` INT NOT NULL COMMENT '用户ID,关联用户表的ID',
    `song_id` INT NOT NULL COMMENT '歌曲ID,关联歌曲表的ID',
    `text` TEXT NOT NULL COMMENT '评论内容,不能为空',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '评论创建时间,默认为当前时间',
    `parent_id` INT DEFAULT NULL COMMENT '父评论ID,允许为空',
    `like_count` INT NOT NULL DEFAULT 0 COMMENT '点赞数,默认为0',
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '评论更新时间,默认为当前时间',
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
    FOREIGN KEY (`song_id`) REFERENCES `song` (`id`),
    FOREIGN KEY (`parent_id`) REFERENCES `comment` (`id`) ON DELETE SET NULL
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT='评论表,存储用户对歌曲的评论';

CREATE TABLE IF NOT EXISTS user_likes (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    song_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (song_id) REFERENCES songs(id),
    UNIQUE KEY unique_user_song (user_id, song_id)
);

-- 插入示例数据
-- 插入用户数据
INSERT INTO `user` (`username`, `password`, `name`, `phone`, `email`, `sex`) VALUES
('user1', 'password1', '张三', '13800000001', 'zhangsan@example.com', '男'),
('user2', 'password2', '李四', '13800000002', 'lisi@example.com', '女'),
('admin', 'admin', '管理员', '13800000003', 'admin@example.com', '男');

-- 插入管理员数据
INSERT INTO `manager` (`username`, `password`) VALUES
('admin1', 'adminpassword1'),
('admin2', 'adminpassword2');

-- 插入艺术家数据
INSERT INTO `artist` (`name`, `sex`, `description`) VALUES
('周杰伦', '男', '华语流行乐男歌手、音乐人'),
('林俊杰', '男', '新加坡华语流行乐男歌手');

-- 插入专辑数据
INSERT INTO `album` (`name`, `cover`, `release_date`, `description`) VALUES
('叶惠美', '/assets/covers/albums/yehuimei.jpg', '2003-07-31', '周杰伦第四张专辑'),
('JJ陆', '/assets/covers/albums/jjlu.jpg', '2016-12-12', '林俊杰第十二张专辑');

-- 插入艺术家与专辑关联数据
INSERT INTO `artist_album` (`artist_id`, `album_id`) VALUES
(1, 1),
(2, 2);

-- 插入歌曲数据
INSERT INTO `song` (`name`, `author_id`, `album_id`, `url`, `cover`, `duration`) VALUES
('晴天', 1, 1, '/assets/music/qingtian.mp3', '/assets/covers/qingtian.jpg', 269),
('不为谁而作的歌', 2, 2, '/assets/music/buweishui.mp3', '/assets/covers/buweishui.jpg', 264);

-- 插入播放列表信息
INSERT INTO `playlist_info` (`name`, `user_id`, `description`, `is_public`) VALUES
('我喜欢的音乐', 1, '我的个人收藏', TRUE),
('学习音乐', 2, '适合学习时听的音乐', FALSE);

-- 插入播放列表歌曲
INSERT INTO `playlist_songs` (`playlist_id`, `song_id`, `user_id`) VALUES
(1, 1, 1),
(1, 2, 1),
(2, 1, 2);

-- 插入点赞数据
INSERT INTO `like` (`user_id`, `song_id`) VALUES
(1, 1),
(2, 1),
(1, 2);

-- 插入评论数据
INSERT INTO `comment` (`user_id`, `song_id`, `text`) VALUES
(1, 1, '这首歌真好听！'),
(2, 1, '经典之作'),
(1, 2, '很有感觉的一首歌');