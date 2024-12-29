INSERT INTO `artist` (`name`, `description`) VALUES
('The Beatles', '披头士是20世纪60年代最具影响力的英国摇滚乐队之一，他们的音乐风格对全球流行文化产生了深远影响。'),
('Bob Dylan', '鲍勃·迪伦是美国民谣歌手和词曲创作人，他的作品对现代音乐和社会运动有着重要的影响。');


INSERT INTO `album` (`name`, `cover`, `release_date`, `description`) VALUES
('The Beatles', '/assets/covers/the_beatles.jpg', '1968-11-22', '《The Beatles》是披头士的第九张专辑，也被称为“白色专辑”，以其多样化的音乐风格和创新性广受好评。'),
('Abbey Road', '/assets/covers/abbey_road.jpg', '1969-09-26', '《Abbey Road》是披头士的第十一张专辑，专辑中的“Come Together”和“Something”成为经典，展示了他们创作的巅峰。'),
('Blonde on Blonde', '/assets/covers/blonde_on_blonde.jpg', '1966-06-20', '《Blonde on Blonde》是鲍勃·迪伦的第七张专辑，被认为是迪伦音乐生涯的里程碑之一，融合了民谣与摇滚元素。'),
('Highway 61 Revisited', '/assets/covers/highway_61_revisited.jpg', '1965-08-30', '《Highway 61 Revisited》是鲍勃·迪伦的第六张专辑，开创了他转向电声音乐的风格，收录了经典的“Like a Rolling Stone”。');

-- 披头士的专辑
INSERT INTO `artist_album` (`artist_id`, `album_id`) VALUES
(1, 1),
(1, 2);

-- 鲍勃·迪伦的专辑
INSERT INTO `artist_album` (`artist_id`, `album_id`) VALUES
(2, 3),
(2, 4);

INSERT INTO `song` (`name`, `author_id`, `album_id`, `url`, `cover`, `duration`) VALUES
('Back in the U.S.S.R.', 1, 1, '/songs/the_beatles_back_in_the_ussr.mp3', '/assets/covers/back_in_the_ussr.jpg', 183),
('Dear Prudence', 1, 1, '/songs/the_beatles_dear_prudence.mp3', '/assets/covers/dear_prudence.jpg', 189),
('Glass Onion', 1, 1, '/songs/the_beatles_glass_onion.mp3', '/assets/covers/glass_onion.jpg', 154),
('Ob-La-Di, Ob-La-Da', 1, 1, '/songs/the_beatles_ob_la_di_ob_la_da.mp3', '/assets/covers/ob_la_di_ob_la_da.jpg', 113),
('Wild Honey Pie', 1, 1, '/songs/the_beatles_wild_honey_pie.mp3', '/assets/covers/wild_honey_pie.jpg', 60),
('The Continuing Story of Bungalow Bill', 1, 1, '/songs/the_beatles_bungalow_bill.mp3', '/assets/covers/bungalow_bill.jpg', 205),
('While My Guitar Gently Weeps', 1, 1, '/songs/the_beatles_while_my_guitar_gently_weeps.mp3', '/assets/covers/while_my_guitar_gently_weeps.jpg', 185),
('Happiness Is a Warm Gun', 1, 1, '/songs/the_beatles_happiness_is_a_warm_gun.mp3', '/assets/covers/happiness_is_a_warm_gun.jpg', 183),
('Martha My Dear', 1, 1, '/songs/the_beatles_martha_my_dear.mp3', '/assets/covers/martha_my_dear.jpg', 158),
('I’m So Tired', 1, 1, '/songs/the_beatles_im_so_tired.mp3', '/assets/covers/im_so_tired.jpg', 202),
('Blackbird', 1, 1, '/songs/the_beatles_blackbird.mp3', '/assets/covers/blackbird.jpg', 144),
('Piggies', 1, 1, '/songs/the_beatles_piggies.mp3', '/assets/covers/piggies.jpg', 160),
('Rocky Raccoon', 1, 1, '/songs/the_beatles_rocky_raccoon.mp3', '/assets/covers/rocky_raccoon.jpg', 222),
('Don’t Pass Me By', 1, 1, '/songs/the_beatles_dont_pass_me_by.mp3', '/assets/covers/dont_pass_me_by.jpg', 200),
('Why Don’t We Do It in the Road?', 1, 1, '/songs/the_beatles_why_dont_we_do_it_in_the_road.mp3', '/assets/covers/why_dont_we_do_it_in_the_road.jpg', 91),
('I Will', 1, 1, '/songs/the_beatles_i_will.mp3', '/assets/covers/i_will.jpg', 134),
('Junk', 1, 1, '/songs/the_beatles_junk.mp3', '/assets/covers/junk.jpg', 173),
('Yer Blues', 1, 1, '/songs/the_beatles_yer_blues.mp3', '/assets/covers/yer_blues.jpg', 215),
('Mother Nature’s Son', 1, 1, '/songs/the_beatles_mother_natures_son.mp3', '/assets/covers/mother_natures_son.jpg', 157),
('Everybody’s Got Something to Hide Except Me and My Monkey', 1, 1, '/songs/the_beatles_everybodys_got_something.mp3', '/assets/covers/everybodys_got_something.jpg', 152),
('Sexy Sadie', 1, 1, '/songs/the_beatles_sexy_sadie.mp3', '/assets/covers/sexy_sadie.jpg', 176),
('Helter Skelter', 1, 1, '/songs/the_beatles_helter_skelter.mp3', '/assets/covers/helter_skelter.jpg', 271),
('Long, Long, Long', 1, 1, '/songs/the_beatles_long_long_long.mp3', '/assets/covers/long_long_long.jpg', 157);


INSERT INTO `song` (`name`, `author_id`, `album_id`, `url`, `cover`, `duration`) VALUES
('Come Together', 1, 2, '/songs/abbey_road_come_together.mp3', '/assets/covers/come_together.jpg', 259),
('Something', 1, 2, '/songs/abbey_road_something.mp3', '/assets/covers/something.jpg', 182),
('Maxwell’s Silver Hammer', 1, 2, '/songs/abbey_road_maxwells_silver_hammer.mp3', '/assets/covers/maxwells_silver_hammer.jpg', 207),
('Oh! Darling', 1, 2, '/songs/abbey_road_oh_darling.mp3', '/assets/covers/oh_darling.jpg', 187),
('Octopus’s Garden', 1, 2, '/songs/abbey_road_octopuss_garden.mp3', '/assets/covers/octopuss_garden.jpg', 174),
('I Want You (She’s So Heavy)', 1, 2, '/songs/abbey_road_i_want_you.mp3', '/assets/covers/i_want_you.jpg', 467),
('Here Comes the Sun', 1, 2, '/songs/abbey_road_here_comes_the_sun.mp3', '/assets/covers/here_comes_the_sun.jpg', 185),
('Because', 1, 2, '/songs/abbey_road_because.mp3', '/assets/covers/because.jpg', 137),
('You Never Give Me Your Money', 1, 2, '/songs/abbey_road_you_never_give_me.mp3', '/assets/covers/you_never_give_me.jpg', 210),
('Sun King', 1, 2, '/songs/abbey_road_sun_king.mp3', '/assets/covers/sun_king.jpg', 143),
('Mean Mr. Mustard', 1, 2, '/songs/abbey_road_mean_mr_mustard.mp3', '/assets/covers/mean_mr_mustard.jpg', 99),
('Polythene Pam', 1, 2, '/songs/abbey_road_polythene_pam.mp3', '/assets/covers/polythene_pam.jpg', 102),
('She Came in Through the Bathroom Window', 1, 2, '/songs/abbey_road_she_came_in_through_the_bathroom_window.mp3', '/assets/covers/she_came_in_through_the_bathroom_window.jpg', 111),
('Golden Slumbers', 1, 2, '/songs/abbey_road_golden_slumbers.mp3', '/assets/covers/golden_slumbers.jpg', 144),
('Carry That Weight', 1, 2, '/songs/abbey_road_carry_that_weight.mp3', '/assets/covers/carry_that_weight.jpg', 137),
('The End', 1, 2, '/songs/abbey_road_the_end.mp3', '/assets/covers/the_end.jpg', 139),
('Her Majesty', 1, 2, '/songs/abbey_road_her_majesty.mp3', '/assets/covers/her_majesty.jpg', 23);


INSERT INTO `song` (`name`, `author_id`, `album_id`, `url`, `cover`, `duration`) VALUES
('Rainy Day Women #12 & 35', 2, 3, '/songs/bob_dylan_rainy_day_women.mp3', '/assets/covers/rainy_day_women.jpg', 227),
('Pledging My Time', 2, 3, '/songs/bob_dylan_pledging_my_time.mp3', '/assets/covers/pledging_my_time.jpg', 311),
('Visions of Johanna', 2, 3, '/songs/bob_dylan_visions_of_johanna.mp3', '/assets/covers/visions_of_johanna.jpg', 420),
('Just Like a Woman', 2, 3, '/songs/bob_dylan_just_like_a_woman.mp3', '/assets/covers/just_like_a_woman.jpg', 262),
('Obviously 5 Believers', 2, 3, '/songs/bob_dylan_obviously_5_believers.mp3', '/assets/covers/obviously_5_believers.jpg', 211),
('Sad Eyed Lady of the Lowlands', 2, 3, '/songs/bob_dylan_sad_eyed_lady.mp3', '/assets/covers/sad_eyed_lady.jpg', 696),
('4th Time Around', 2, 3, '/songs/bob_dylan_4th_time_around.mp3', '/assets/covers/4th_time_around.jpg', 188),
('I Want You', 2, 3, '/songs/bob_dylan_i_want_you.mp3', '/assets/covers/i_want_you.jpg', 213),
('Stuck Inside of Mobile with the Memphis Blues Again', 2, 3, '/songs/bob_dylan_stuck_inside_mobile.mp3', '/assets/covers/stuck_inside_mobile.jpg', 359),
('Temporary Like Achilles', 2, 3, '/songs/bob_dylan_temporary_like_achilles.mp3', '/assets/covers/temporary_like_achilles.jpg', 265);


INSERT INTO `song` (`name`, `author_id`, `album_id`, `url`, `cover`, `duration`) VALUES
('Like a Rolling Stone', 2, 4, '/songs/bob_dylan_like_a_rolling_stone.mp3', '/assets/covers/like_a_rolling_stone.jpg', 367),
('Tombstone Blues', 2, 4, '/songs/bob_dylan_tombstone_blues.mp3', '/assets/covers/tombstone_blues.jpg', 295),
('It Takes a Lot to Laugh, It Takes a Train to Cry', 2, 4, '/songs/bob_dylan_it_takes_a_lot.mp3', '/assets/covers/it_takes_a_lot.jpg', 314),
('Ballad of a Thin Man', 2, 4, '/songs/bob_dylan_ballad_of_a_thin_man.mp3', '/assets/covers/ballad_of_a_thin_man.jpg', 245),
('From a Buick 6', 2, 4, '/songs/bob_dylan_from_a_buick_6.mp3', '/assets/covers/from_a_buick_6.jpg', 215),
('Queen Jane Approximately', 2, 4, '/songs/bob_dylan_queen_jane_approximately.mp3', '/assets/covers/queen_jane_approximately.jpg', 339),
('Highway 61 Revisited', 2, 4, '/songs/bob_dylan_highway_61_revisited.mp3', '/assets/covers/highway_61_revisited.jpg', 280),
('Just Like Tom Thumb’s Blues', 2, 4, '/songs/bob_dylan_just_like_tom_thumbs_blues.mp3', '/assets/covers/just_like_tom_thumbs_blues.jpg', 330),
('Desolation Row', 2, 4, '/songs/bob_dylan_desolation_row.mp3', '/assets/covers/desolation_row.jpg', 719);


INSERT INTO `comment` (`id`, `user_id`, `song_id`, `text`, `created_at`, `parent_id`, `reply_to_user_id`, `like_count`, `updated_at`) VALUES (1,2,1,'back to UUSR comrade!!!','2024-12-28 11:37:40',NULL,NULL,0,'2024-12-28 11:37:40'),(2,2,1,'I like this song','2024-12-28 11:38:01',1,2,0,'2024-12-28 11:38:01'),(3,3,1,'it\'s just like the beach boy\'s california girls','2024-12-28 11:38:54',1,2,0,'2024-12-28 11:38:54'),(4,1,1,'111','2024-12-28 15:47:36',1,2,0,'2024-12-28 15:47:36');

INSERT INTO `comment_like` (`id`, `comment_id`, `user_id`, `created_at`) VALUES (1,2,1,'2024-12-28 15:47:29'),(2,3,1,'2024-12-28 15:47:30');

INSERT INTO `like` (`id`, `user_id`, `song_id`, `created_at`) VALUES (2,1,36,'2024-12-28 15:46:58'),(3,1,51,'2024-12-28 15:47:00'),(4,2,1,'2024-12-28 11:41:39'),(5,2,8,'2024-12-28 11:41:44'),(6,1,21,'2024-12-28 13:57:04'),(10,1,1,'2024-12-28 15:48:47'),(15,1,8,'2024-12-28 15:48:27');

INSERT INTO `playlist_info` (`id`, `name`, `user_id`, `cover`, `description`, `is_public`, `created_at`, `updated_at`) VALUES (1,'night car songs',1,'/assets/covers/playlists/default.jpg','11',0,'2024-12-28 11:35:31','2024-12-28 15:46:49'),(3,'DOOMER SONG',2,'/assets/covers/playlists/default.jpg','awesrasfaesgesagswagtaewgtw',0,'2024-12-28 11:42:16','2024-12-28 11:42:34'),(4,'hello world',1,'/assets/covers/playlists/default.jpg',NULL,0,'2024-12-28 14:00:04','2024-12-28 15:46:28');

INSERT INTO `playlist_songs` (`id`, `playlist_id`, `song_id`, `user_id`, `added_at`) VALUES (2,1,36,1,'2024-12-28 11:36:08'),(3,3,44,2,'2024-12-28 11:42:34'),(4,4,1,1,'2024-12-28 15:46:28');

INSERT INTO `user` (`id`, `username`, `password`, `name`, `phone`, `email`, `sex`, `avatar`, `bio`, `gender`, `birthday`, `location`, `status`, `created_time`, `updated_time`) VALUES (1,'admin','$2a$10$XQ3lTxi12672s7o401OCRe8X/uvhXPZMPPtf08fQgYs48wR1dyx2S','管理员','18971548334','111@111.com',NULL,'/assets/avatars/default-user.jpg','11',1,NULL,'111',1,'2024-12-28 11:22:31','2024-12-28 15:47:57'),(2,'user1','$2a$10$VsmbRSJ041kGQ4fbhzmjvO7YM2sLWDUTSUEgo0OrY1XSM.BhYxqi2','doomer no.1','18971548334','111@111.com',NULL,'/assets/avatars/default-user.jpg','I like music, you too?',1,NULL,'111',1,'2024-12-28 11:37:01','2024-12-28 11:41:24'),(3,'user2','$2a$10$wTp8u5h5iWLE2m5ezAqgnuyovh35yseaezWjGhLM8OSToxLuOcelm','user2','18971548334','111@111.com',NULL,'/assets/avatars/default-user.jpg',NULL,2,'2024-11-24','111',1,'2024-12-28 11:38:18','2024-12-28 11:38:18');
