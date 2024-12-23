-- 用户表 (user)
INSERT INTO `user` (`username`, `password`, `name`, `phone`, `email`, `sex`, `created_time`) VALUES
('user1', 'password1', '张三', '13800000001', 'zhangsan@example.com', '男', NOW()),
('user2', 'password2', '李四', '13800000002', 'lisi@example.com', '女', NOW()),
('user3', 'password3', '王五', '13800000003', 'wangwu@example.com', '男', NOW()),
('user4', 'password4', '赵六', '13800000004', 'zhaoliu@example.com', '女', NOW()),
('user5', 'password5', '孙七', '13800000005', 'sunqi@example.com', '男', NOW()),
('user6', 'password6', '周八', '13800000006', 'zhouba@example.com', '女', NOW()),
('user7', 'password7', '吴九', '13800000007', 'wujiu@example.com', '男', NOW()),
('user8', 'password8', '郑十', '13800000008', 'zhengshi@example.com', '女', NOW()),
('user9', 'password9', '钱十一', '13800000009', 'qianshiyi@example.com', '男', NOW()),
('user10', 'password10', '何十二', '13800000010', 'heshier@example.com', '女', NOW()),
('user11', 'password11', '林十三', '13800000011', 'linshisan@example.com', '男', NOW()),
('user12', 'password12', '郭十四', '13800000012', 'guoshishi@example.com', '女', NOW()),
('user13', 'password13', '冯十五', '13800000013', 'fengshiwu@example.com', '男', NOW()),
('user14', 'password14', '陈十六', '13800000014', 'chenshiliu@example.com', '女', NOW()),
('user15', 'password15', '褚十七', '13800000015', 'chushiqi@example.com', '男', NOW()),
('user16', 'password16', '卫十八', '13800000016', 'weishiba@example.com', '女', NOW()),
('user17', 'password17', '蒋十九', '13800000017', 'jiangshijiu@example.com', '男', NOW()),
('user18', 'password18', '沈二十', '13800000018', 'shenershi@example.com', '女', NOW()),
('user19', 'password19', '韩二十一', '13800000019', 'hanshiyi@example.com', '男', NOW()),
('user20', 'password20', '杨二十二', '13800000020', 'yangershi@example.com', '女', NOW());

-- 管理员表 (manager)
INSERT INTO `manager` (`username`, `password`) VALUES
('admin1', 'adminpassword1'),
('admin2', 'adminpassword2'),
('admin3', 'adminpassword3'),
('admin4', 'adminpassword4'),
('admin5', 'adminpassword5'),
('admin6', 'adminpassword6'),
('admin7', 'adminpassword7'),
('admin8', 'adminpassword8'),
('admin9', 'adminpassword9'),
('admin10', 'adminpassword10'),
('admin11', 'adminpassword11'),
('admin12', 'adminpassword12'),
('admin13', 'adminpassword13'),
('admin14', 'adminpassword14'),
('admin15', 'adminpassword15'),
('admin16', 'adminpassword16'),
('admin17', 'adminpassword17'),
('admin18', 'adminpassword18'),
('admin19', 'adminpassword19'),
('admin20', 'adminpassword20');

-- 歌曲表 (song)
INSERT INTO `song` (`name`, `author_id`, `album_id`, `is_public`, `loc`, `pic`, `create_time`) VALUES
('歌曲1', 1, 1, TRUE, '/songs/1.mp3', '/images/song1.jpg', NOW()),
('歌曲2', 2, 2, TRUE, '/songs/2.mp3', '/images/song2.jpg', NOW()),
('歌曲3', 3, 3, TRUE, '/songs/3.mp3', '/images/song3.jpg', NOW()),
('歌曲4', 4, 4, TRUE, '/songs/4.mp3', '/images/song4.jpg', NOW()),
('歌曲5', 5, 5, TRUE, '/songs/5.mp3', '/images/song5.jpg', NOW()),
('歌曲6', 6, 6, TRUE, '/songs/6.mp3', '/images/song6.jpg', NOW()),
('歌曲7', 7, 7, TRUE, '/songs/7.mp3', '/images/song7.jpg', NOW()),
('歌曲8', 8, 8, TRUE, '/songs/8.mp3', '/images/song8.jpg', NOW()),
('歌曲9', 9, 9, TRUE, '/songs/9.mp3', '/images/song9.jpg', NOW()),
('歌曲10', 10, 10, TRUE, '/songs/10.mp3', '/images/song10.jpg', NOW()),
('歌曲11', 11, 11, TRUE, '/songs/11.mp3', '/images/song11.jpg', NOW()),
('歌曲12', 12, 12, TRUE, '/songs/12.mp3', '/images/song12.jpg', NOW()),
('歌曲13', 13, 13, TRUE, '/songs/13.mp3', '/images/song13.jpg', NOW()),
('歌曲14', 14, 14, TRUE, '/songs/14.mp3', '/images/song14.jpg', NOW()),
('歌曲15', 15, 15, TRUE, '/songs/15.mp3', '/images/song15.jpg', NOW()),
('歌曲16', 16, 16, TRUE, '/songs/16.mp3', '/images/song16.jpg', NOW()),
('歌曲17', 17, 17, TRUE, '/songs/17.mp3', '/images/song17.jpg', NOW()),
('歌曲18', 18, 18, TRUE, '/songs/18.mp3', '/images/song18.jpg', NOW()),
('歌曲19', 19, 19, TRUE, '/songs/19.mp3', '/images/song19.jpg', NOW()),
('歌曲20', 20, 20, TRUE, '/songs/20.mp3', '/images/song20.jpg', NOW());

-- 点赞表 (like)
INSERT INTO `like` (`user_id`, `song_id`) VALUES
(1, 1), (1, 2), (2, 2), (2, 3), (3, 3), (3, 4),
(4, 4), (4, 5), (5, 5), (5, 6), (6, 6), (6, 7),
(7, 7), (7, 8), (8, 8), (8, 9), (9, 9), (9, 10), 
(10, 10), (10, 11);

-- 评论表 (comment)
INSERT INTO `comment` (`user_id`, `song_id`, `text`) VALUES
(1, 1, '非常好听的歌曲！'),
(2, 2, '喜欢这首歌的旋律。'),
(3, 3, '歌词很有深意。'),
(4, 4, '这首歌很动听。'),
(5, 5, '旋律很特别。'),
(6, 6, '这首歌让我想起了很多回忆。'),
(7, 7, '歌手的声音很好听。'),
(8, 8, '这首歌节奏很棒。'),
(9, 9, '这首歌非常振奋人心。'),
(10, 10, '一首值得反复听的好歌。'),
(11, 11, '歌词非常感人。'),
(12, 12, '这首歌非常有感染力。'),
(13, 13, '旋律和歌词都很好。'),
(14, 14, '一首很有情感的歌。'),
(15, 15, '非常喜欢这首歌的气氛。'),
(16, 16, '歌手的声音非常有力量。'),
(17, 17, '这首歌很适合放松时听。'),
(18, 18, '特别喜欢这首歌的编曲。'),
(19, 19, '歌词深刻，值得思考。'),
(20, 20, '这首歌让我很放松。');

-- 歌单表 (playlist)
INSERT INTO `playlist` (`user_id`, `song_id`) VALUES
(1, 1), (1, 2), (2, 3), (2, 4), (3, 5), (3, 6),
(4, 7), (4, 8), (5, 9), (5, 10), (6, 11), (6, 12),
(7, 13), (7, 14), (8, 15), (8, 16), (9, 17), (9, 18),
(10, 19), (10, 20);

-- 艺术家表 (artist)
INSERT INTO `artist` (`name`, `sex`) VALUES
('艺术家1', '男'),
('艺术家2', '女'),
('艺术家3', '男'),
('艺术家4', '女'),
('艺术家5', '男'),
('艺术家6', '女'),
('艺术家7', '男'),
('艺术家8', '女'),
('艺术家9', '男'),
('艺术家10', '女'),
('艺术家11', '男'),
('艺术家12', '女'),
('艺术家13', '男'),
('艺术家14', '女'),
('艺术家15', '男'),
('艺术家16', '女'),
('艺术家17', '男'),
('艺术家18', '女'),
('艺术家19', '男'),
('艺术家20', '女');

-- 专辑表 (album)
INSERT INTO `album` (`name`, `cover`) VALUES
('专辑1', '/images/album1.jpg'),
('专辑2', '/images/album2.jpg'),
('专辑3', '/images/album3.jpg'),
('专辑4', '/images/album4.jpg'),
('专辑5', '/images/album5.jpg'),
('专辑6', '/images/album6.jpg'),
('专辑7', '/images/album7.jpg'),
('专辑8', '/images/album8.jpg'),
('专辑9', '/images/album9.jpg'),
('专辑10', '/images/album10.jpg'),
('专辑11', '/images/album11.jpg'),
('专辑12', '/images/album12.jpg'),
('专辑13', '/images/album13.jpg'),
('专辑14', '/images/album14.jpg'),
('专辑15', '/images/album15.jpg'),
('专辑16', '/images/album16.jpg'),
('专辑17', '/images/album17.jpg'),
('专辑18', '/images/album18.jpg'),
('专辑19', '/images/album19.jpg'),
('专辑20', '/images/album20.jpg');

-- 专辑与艺术家关联表 (album_artist)
INSERT INTO `album_artist` (`artist_id`, `album_id`) VALUES
(1, 1), (2, 2), (3, 3), (4, 4), (5, 5), (6, 6),
(7, 7), (8, 8), (9, 9), (10, 10), (11, 11), (12, 12),
(13, 13), (14, 14), (15, 15), (16, 16), (17, 17), (18, 18),
(19, 19), (20, 20);

-- 专辑列表表 (album_list)
INSERT INTO `album_list` (`album_id`, `user_id`) VALUES
(1, 1), (2, 2), (3, 3), (4, 4), (5, 5), (6, 6),
(7, 7), (8, 8), (9, 9), (10, 10), (11, 11), (12, 12),
(13, 13), (14, 14), (15, 15), (16, 16), (17, 17), (18, 18),
(19, 19), (20, 20);