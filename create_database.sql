DROP DATABASE IF EXISTS `music_player`;

CREATE DATABASE IF NOT EXISTS `music_player`
    CHARACTER SET utf8mb4
    COLLATE utf8mb4_unicode_ci;

USE `music_player`;

-- 创建用户表
CREATE TABLE `user`
(
    `id`           INT PRIMARY KEY AUTO_INCREMENT COMMENT '用户ID,主键,自动递增',
    `username`     VARCHAR(50)  NOT NULL UNIQUE COMMENT '用户名,唯一且不能为空',
    `password`     VARCHAR(100) NOT NULL COMMENT '用户密码,不能为空',
    `name`         VARCHAR(50)           DEFAULT NULL COMMENT '用户姓名,允许为空',
    `phone`        VARCHAR(20)           DEFAULT NULL COMMENT '用户电话号码,允许为空',
    `email`        VARCHAR(100)          DEFAULT NULL COMMENT '用户电子邮件,允许为空',
    `sex`          VARCHAR(10)           DEFAULT NULL COMMENT '用户性别,允许为空',
    `avatar`       VARCHAR(255)          DEFAULT '/assets/avatars/default-user.jpg' COMMENT '用户头像,允许为空',
    bio            TEXT,                            -- 个人简介
    gender         TINYINT               DEFAULT 0, -- 性别：0-保密，1-男，2-女
    birthday       DATE,                            -- 生日
    location       VARCHAR(100),                    -- 地区
    `status`       TINYINT      NOT NULL DEFAULT 1 COMMENT '用户状态,默认为1',
    `created_time` TIMESTAMP             DEFAULT CURRENT_TIMESTAMP COMMENT '账户创建时间,默认为当前时间',
    `updated_time` TIMESTAMP             DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '账户更新时间,默认为当前时间'
) ENGINE = InnoDB
  row_format = Dynamic
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci COMMENT ='用户表,存储用户的基本信息';

-- 创建管理员表
CREATE TABLE `manager`
(
    `id`       INT PRIMARY KEY AUTO_INCREMENT COMMENT '管理员ID,主键,自动递增',
    `username` VARCHAR(50)  NOT NULL UNIQUE COMMENT '管理员用户名,唯一且不能为空',
    `password` VARCHAR(100) NOT NULL COMMENT '管理员密码,不能为空'
) ENGINE = MyISAM
  row_format = Compact
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci COMMENT ='管理员表,存储管理员的基本信息';

-- 创建艺术家表
CREATE TABLE `artist`
(
    `id`          INT PRIMARY KEY AUTO_INCREMENT COMMENT '艺术家ID,主键,自动递增',
    `name`        VARCHAR(100) NOT NULL COMMENT '艺术家名称,不能为空',
    `sex`         VARCHAR(10)  DEFAULT NULL COMMENT '艺术家性别,允许为空',
    `avatar`      VARCHAR(255) DEFAULT '/assets/avatars/default-artist.jpg' COMMENT '艺术家头像,允许为空',
    `description` TEXT         DEFAULT NULL COMMENT '艺术家描述,允许为空',
    `created_at`  TIMESTAMP    DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间,默认为当前时间'
) ENGINE = InnoDB
  row_format = Dynamic
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci COMMENT ='艺术家表,存储艺术家的基本信息';

-- 创建专辑表
CREATE TABLE `album`
(
    `id`           INT PRIMARY KEY AUTO_INCREMENT COMMENT '专辑ID,主键,自动递增',
    `name`         VARCHAR(100) NOT NULL COMMENT '专辑名称,不能为空',
    `cover`        VARCHAR(255) NOT NULL COMMENT '专辑封面,不能为空',
    `release_date` DATE      DEFAULT NULL COMMENT '专辑发行日期,允许为空',
    `description`  TEXT      DEFAULT NULL COMMENT '专辑描述,允许为空',
    `created_at`   TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间,默认为当前时间'
) ENGINE = InnoDB
  row_format = Dynamic
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci COMMENT ='专辑表,存储专辑的基本信息';

-- 创建艺术家与专辑关联表 (多对多关系)
CREATE TABLE `artist_album`
(
    `artist_id` INT NOT NULL COMMENT '艺术家ID,关联艺术家表的ID',
    `album_id`  INT NOT NULL COMMENT '专辑ID,关联专辑表的ID',
    PRIMARY KEY (`artist_id`, `album_id`),
    FOREIGN KEY (`artist_id`) REFERENCES `artist` (`id`) ON DELETE CASCADE,
    FOREIGN KEY (`album_id`) REFERENCES `album` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB
  row_format = Compact
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci COMMENT ='艺术家与专辑关联表,存储艺术家与专辑的关系';

-- 创建歌曲表
CREATE TABLE `song`
(
    `id`          INT PRIMARY KEY AUTO_INCREMENT COMMENT '歌曲ID,主键,自动递增',
    `name`        VARCHAR(100) NOT NULL COMMENT '歌曲名称,不能为空',
    `author_id`   INT          NOT NULL COMMENT '歌曲作者ID,不能为空,对应艺术家的ID',
    `album_id`    INT          NOT NULL COMMENT '专辑ID,不能为空,对应专辑的ID',
    `url`         VARCHAR(255) NOT NULL COMMENT '歌曲存放地址,不能为空',
    `cover`       VARCHAR(255) NOT NULL COMMENT '歌曲封面,不能为空',
    `duration`    INT          NOT NULL DEFAULT 180 COMMENT '歌曲时长,默认为180秒',
    `play_count`  INT          NOT NULL DEFAULT 0 COMMENT '播放次数,默认为0',
    `like_count`  INT          NOT NULL DEFAULT 0 COMMENT '点赞次数,默认为0',
    `is_public`   BOOLEAN               DEFAULT TRUE COMMENT '是否公开,默认为公开',
    `create_time` TIMESTAMP             DEFAULT CURRENT_TIMESTAMP COMMENT '歌曲创建时间,默认为当前时间',
    FOREIGN KEY (`author_id`) REFERENCES `artist` (`id`),
    FOREIGN KEY (`album_id`) REFERENCES `album` (`id`)
) ENGINE = InnoDB
  row_format = Compact
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci COMMENT ='歌曲表,存储歌曲的基本信息';

-- 创建播放列表信息表
CREATE TABLE `playlist_info`
(
    `id`          INT PRIMARY KEY AUTO_INCREMENT COMMENT '播放列表ID,主键,自动递增',
    `name`        VARCHAR(100) NOT NULL COMMENT '播放列表名称,不能为空',
    `user_id`     INT          NOT NULL COMMENT '用户ID,关联用户表的ID',
    `cover`       VARCHAR(255) DEFAULT '/assets/covers/playlists/default.jpg' COMMENT '播放列表封面,允许为空',
    `description` TEXT         DEFAULT NULL COMMENT '播放列表描述,允许为空',
    `is_public`   BOOLEAN      DEFAULT FALSE COMMENT '是否公开,默认为不公开',
    `created_at`  TIMESTAMP    DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间,默认为当前时间',
    `updated_at`  TIMESTAMP    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间,默认为当前时间',
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE = InnoDB
  row_format = Dynamic
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci COMMENT ='播放列表信息表,存储播放列表的基本信息';

-- 创建播放列表歌曲关联表
CREATE TABLE `playlist_songs`
(
    `id`          INT PRIMARY KEY AUTO_INCREMENT COMMENT '播放列表歌曲关联ID,主键,自动递增',
    `playlist_id` INT NOT NULL COMMENT '播放列表ID,关联播放列表表的ID',
    `song_id`     INT NOT NULL COMMENT '歌曲ID,关联歌曲表的ID',
    `user_id`     INT NOT NULL COMMENT '用户ID,关联用户表的ID',
    `added_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间,默认为当前时间',
    FOREIGN KEY (`playlist_id`) REFERENCES `playlist_info` (`id`),
    FOREIGN KEY (`song_id`) REFERENCES `song` (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE = InnoDB
  row_format = Compact
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci COMMENT ='播放列表歌曲关联表,存储播放列表与歌曲的关系';

-- 创建点赞表
CREATE TABLE `like`
(
    `id`         INT PRIMARY KEY AUTO_INCREMENT COMMENT '点赞ID,主键,自动递增',
    `user_id`    INT NOT NULL COMMENT '用户ID,关联用户表的ID',
    `song_id`    INT NOT NULL COMMENT '歌曲ID,关联歌曲表的ID',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '点赞时间,默认为当前时间',
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
    FOREIGN KEY (`song_id`) REFERENCES `song` (`id`),
    UNIQUE KEY `unique_user_song` (`user_id`, `song_id`)
) ENGINE = InnoDB
  row_format = Compact
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci COMMENT ='点赞表,存储用户对歌曲的点赞记录';

-- 创建评论表
CREATE TABLE `comment`
(
    `id`               INT PRIMARY KEY AUTO_INCREMENT COMMENT '评论ID,主键,自动递增',
    `user_id`          INT  NOT NULL COMMENT '用户ID,关联用户表的ID',
    `song_id`          INT  NOT NULL COMMENT '歌曲ID,关联歌曲表的ID',
    `text`             TEXT NOT NULL COMMENT '评论内容,不能为空',
    `created_at`       TIMESTAMP     DEFAULT CURRENT_TIMESTAMP COMMENT '评论创建时间,默认为当前时间',
    `parent_id`        INT           DEFAULT NULL COMMENT '父评论ID,允许为空',
    `reply_to_user_id` int           DEFAULT NULL COMMENT '回复目标用户ID',
    `like_count`       INT  NOT NULL DEFAULT 0 COMMENT '点赞数,默认为0',
    `updated_at`       TIMESTAMP     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '评论更新时间,默认为当前时间',
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
    FOREIGN KEY (`song_id`) REFERENCES `song` (`id`),
    FOREIGN KEY (`parent_id`) REFERENCES `comment` (`id`) ON DELETE SET NULL,
    FOREIGN KEY (`reply_to_user_id`) REFERENCES `user` (`id`) ON DELETE SET NULL
) ENGINE = InnoDB
  row_format = Dynamic
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci COMMENT ='评论表,存储用户对歌曲的评论';

CREATE TABLE IF NOT EXISTS user_likes
(
    id         INT PRIMARY KEY AUTO_INCREMENT,
    user_id    INT NOT NULL,
    song_id    INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user (id),
    FOREIGN KEY (song_id) REFERENCES song (id),
    UNIQUE KEY unique_user_song (user_id, song_id)
) ENGINE = InnoDB
  row_format = Compact
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci;
-- 评论点赞表
CREATE TABLE `comment_like`
(
    `id`         int       NOT NULL AUTO_INCREMENT,
    `comment_id` int       NOT NULL COMMENT '评论ID',
    `user_id`    int       NOT NULL COMMENT '用户ID',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '点赞时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_comment_user` (`comment_id`, `user_id`) COMMENT '确保用户对同一评论只能点赞一次',
    CONSTRAINT `fk_comment_like_comment` FOREIGN KEY (`comment_id`) REFERENCES `comment` (`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_comment_like_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB
  row_format = Compact
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='评论点赞表';

-- 关注关系表
CREATE TABLE follow
(
    id           INT PRIMARY KEY AUTO_INCREMENT,
    follower_id  INT NOT NULL,                           -- 关注者ID
    following_id INT NOT NULL,                           -- 被关注者ID
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (follower_id) REFERENCES user (id) ON DELETE CASCADE,
    FOREIGN KEY (following_id) REFERENCES user (id) ON DELETE CASCADE,
    UNIQUE KEY unique_follow (follower_id, following_id) -- 防止重复关注
) ENGINE = InnoDB
  row_format = Compact
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci;

-- 针对 user 表
CREATE INDEX idx_user_phone ON user(phone);
CREATE INDEX idx_user_email ON user(email);

-- 针对 song 表
CREATE INDEX idx_song_name ON song(name);
CREATE INDEX idx_song_album_id ON song(album_id);
CREATE INDEX idx_song_author_id ON song(author_id);

-- 针对 comment 表
CREATE INDEX idx_comment_song_id ON comment(song_id);
CREATE INDEX idx_comment_user_id ON comment(user_id);
CREATE INDEX idx_comment_parent_id ON comment(parent_id);

-- 展示表
# SHOW TABLES;

-- 删除表，注意顺序要先删除有外键依赖的表
# DROP TABLE IF EXISTS comment_like;
# DROP TABLE IF EXISTS playlist_songs;
# DROP TABLE IF EXISTS playlist_info;
# DROP TABLE IF EXISTS `like`;
# DROP TABLE IF EXISTS comment;
# DROP TABLE IF EXISTS user_likes;
# DROP TABLE IF EXISTS song;
# DROP TABLE IF EXISTS artist_album;
# DROP TABLE IF EXISTS album;
# DROP TABLE IF EXISTS artist;
# DROP TABLE IF EXISTS manager;
# DROP TABLE IF EXISTS user;
# DROP TABLE IF EXISTS follow;
