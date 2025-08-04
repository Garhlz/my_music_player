//所有后端逻辑放在这里
const express = require('express');
const router = express.Router();
const db = require('./db'); // 引入数据库连接
const bcrypt = require('bcryptjs'); // 引入 bcrypt 库
const jwt = require('jsonwebtoken'); // 引入 jwt 库
const dotenv = require('dotenv');


dotenv.config();
const JWT_SECRET = process.env.JWT_SECRET || 'your_jwt_secret_key'; // JWT 密钥
const {verifyToken} = require('./middleware/authMiddleware');



//----------------------------------用户登录注册----------------------------------
// 用户登录接口
router.post('/user/login', async (req, res) => {
    const { username, password } = req.body;
    console.log(username,password);
    if (!username || !password) {
        return res.status(400).json({ error: '用户名和密码不能为空' });
    }

    try {
        const [rows] = await db.execute('SELECT id, password FROM user WHERE username = ?', [username]);

        if (rows.length > 0) {
            const isMatch = await bcrypt.compare(password, rows[0].password);
            if (isMatch) {
                const token = jwt.sign({ id: rows[0].id, username }, JWT_SECRET, { expiresIn: '1h' });
                return res.status(200).json({ message: '登录成功', token, userId:rows[0].id });
            } else {
                return res.status(401).json({ error: '密码错误' });
            }
        } else {
            return res.status(404).json({ error: '用户未找到' });
        }
    } catch (error) {
        console.error('数据库查询失败:', error);
        return res.status(500).json({ error: '服务器错误' });
    }
});

// 用户注册接口
router.post('/user/register', async (req, res) => {
    const { 
        username, 
        password, 
        name, 
        phone, 
        email, 
        sex,
        bio,
        gender,
        birthday,
        location 
    } = req.body;

    if (!username || !password) {
        return res.status(400).json({ error: '用户名和密码不能为空' });
    }

    try {
        const [result] = await db.execute('SELECT * FROM user WHERE username = ?', [username]);
        const existingUser = result || [];

        if (existingUser.length > 0) {
            return res.status(409).json({ error: '用户名已存在' });
        }

        const hashedPassword = await bcrypt.hash(password, 10);
        await db.execute(`
            INSERT INTO user (
                username, password, name, phone, email, sex,
                bio, gender, birthday, location, avatar, status
            ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
        `, [
            username, 
            hashedPassword, 
            name || null, 
            phone || null, 
            email || null, 
            sex || null,
            bio || null,
            gender || 0,
            birthday || null,
            location || null,
            '/assets/avatars/default-user.jpg',
            1
        ]);

        return res.status(201).json({ message: '注册成功' });
    } catch (error) {
        console.error('注册失败:', error);
        return res.status(500).json({ error: '注册过程中发生错误，请稍后再试' });
    }
});
//----------------------------------用户登录注册----------------------------------





//----------------------------------用户信息----------------------------------
// 获取用户信息接口
router.get('/user/:id', async (req, res) => {
  try {
    const userId = req.params.id;
    
    // 获取基本用户信息
    const [user] = await db.execute(`
      SELECT 
        u.id,
        u.username,
        u.name,
        u.avatar,
        u.email,
        u.phone,
        u.gender,
        u.birthday,
        u.location,
        u.bio,
        COALESCE((SELECT COUNT(*) FROM follow WHERE follower_id = u.id), 0) as followings,
        COALESCE((SELECT COUNT(*) FROM follow WHERE following_id = u.id), 0) as followers
      FROM user u
      WHERE u.id = ?
    `, [userId]);

    if (!user[0]) {
      return res.status(404).json({
        message: false,
        data: null,
        error: '用户不存在'
      });
    }

    res.json({
      message: true,
      data: {
        data: user[0]
      }
    });
  } catch (error) {
    console.error('获取用户信息失败:', error);
    res.status(500).json({
      message: false,
      data: null,
      error: '获取用户信息失败'
    });
  }
});

// 更新用户信息
router.put('/user/:id', verifyToken, async (req, res) => {
  try {
    const userId = req.params.id;
    
    // 验证是否是当前用户
    if (userId.toString() !== req.user.id.toString()) {
      return res.status(403).json({
        message: false,
        data: null,
        error: '无权修改其他用户信息'
      });
    }

    const {
      name,
      avatar,
      email,
      phone,
      gender,
      birthday,
      location,
      bio
    } = req.body;

    // 更新用户信息
    const query = `
      UPDATE user 
      SET 
        name = ?,
        avatar = ?,
        email = ?,
        phone = ?,
        gender = ?,
        birthday = ?,
        location = ?,
        bio = ?
      WHERE id = ?
    `;

    await db.execute(query, [
      name,
      avatar,
      email,
      phone,
      gender,
      birthday,
      location,
      bio,
      userId
    ]);

    res.json({
      message: true,
      data: {
        data: { updated: true }
      }
    });
  } catch (error) {
    console.error('更新用户信息失败:', error);
    res.status(500).json({
      message: false,
      data: null,
      error: '更新用户信息失败'
    });
  }
});
//----------------------------------用户信息----------------------------------





//----------------------------------歌曲----------------------------------
// 获取歌曲列表接口
router.get('/song/:id', verifyToken, async (req, res) => {
    try {
        const songId = req.params.id;
        const query = `
            SELECT 
                s.*,
                a.name as artist_name,
                a.id as artist_id,
                al.name as album_name,
                al.id as album_id,
                al.cover as album_cover
            FROM song s
            LEFT JOIN artist a ON s.author_id = a.id
            LEFT JOIN album al ON s.album_id = al.id
            WHERE s.id = ?
        `;

        const [song] = await db.execute(query, [songId]);

        if (song.length === 0) {
            return res.status(404).json({
                message: false,
                error: '歌曲不存在'
            });
        }

        // 格式化返回数据
        const formattedSong = {
            id: song[0].id,
            name: song[0].name,
            artist_name: song[0].artist_name || '未知歌手',
            artist_id: song[0].artist_id,
            album_name: song[0].album_name || '未知专辑',
            album_id: song[0].album_id,
            cover: song[0].album_cover || '/assets/default-cover.jpg',
            duration: song[0].duration,
            url: song[0].url,
            create_time: song[0].create_time,
            play_count: song[0].play_count || 0,
            like_count: song[0].like_count || 0
        };

        res.json({
            message: true,
            data: formattedSong
        });

    } catch (error) {
        console.error('获取歌曲详情失败:', error);
        res.status(500).json({
            message: false,
            error: '获取歌曲详情失败'
        });
    }
});


router.get('/songs', verifyToken, async (req, res) => {
  try {
    // 转换并验证参数
    const page = parseInt(req.query.page) || 1;
    const pageSize = parseInt(req.query.pageSize) || 10;
    const search = req.query.search || '';
    const sortBy = req.query.sortBy || 'id';
    const offset = (page - 1) * pageSize;
    
    // 定义允许的排序字段
    const allowedSortFields = {
      'id': 'id',
      'name': 'name',
      'play_count': 'play_count',
      'like_count': 'like_count',
      'create_time': 'create_time',
      'latest': 'create_time'
    };

    // 验证排序字段
    const validSortField = allowedSortFields[sortBy] || 'id';
    
    let query = `
      SELECT s.*, a.name as artist_name, al.name as album_name, al.cover as album_cover
      FROM song s
      LEFT JOIN artist a ON s.author_id = a.id
      LEFT JOIN album al ON s.album_id = al.id
      WHERE s.is_public = TRUE
    `;
    
    let countQuery = `
      SELECT COUNT(*) as total
      FROM song s
      WHERE s.is_public = TRUE
    `;

    let queryParams = [];
    let countParams = [];

    // 添加搜索条件
    if (search) {
      query += ` AND (s.name LIKE '%${search}%' OR a.name LIKE '%${search}%')`;
      countQuery += ` AND (s.name LIKE '%${search}%' OR a.name LIKE '%${search}%')`;
    }

    // 添加排序和分页
    query += ` ORDER BY s.${validSortField} DESC LIMIT ${pageSize} OFFSET ${offset}`;

    // 执行查询
    const [songs] = await db.execute(query);
    const [totalResult] = await db.execute(countQuery);

    res.json({
      message: true,
      data: {
        total: totalResult[0].total,
        list: songs
      }
    });
  } catch (error) {
    console.error('获取歌曲列表失败:', error);
    res.status(500).json({ 
      message: false, 
      error: '获取歌曲列表失败',
      details: process.env.NODE_ENV === 'development' ? error.message : undefined
    });
  }
});
// 获取歌手信息
router.get('/artists', verifyToken, async (req, res) => {
  try {
    const { ids } = req.query;
    if (!ids) {
      return res.status(400).json({ message: false, error: '缺少歌手ID' });
    }

    // 处理数组格式的参数
    let artistIds = Array.isArray(ids) ? ids : [ids];
    
    // 确保所有ID都是数字
    artistIds = artistIds.map(id => parseInt(id)).filter(id => !isNaN(id));
    
    if (artistIds.length === 0) {
      return res.status(400).json({ message: false, error: '无效的歌手ID' });
    }

    // 构建查询语句
    const placeholders = artistIds.map(() => '?').join(',');
    const query = `
      SELECT id, name, sex, avatar, description
      FROM artist
      WHERE id IN (${placeholders})
    `;
    
    const [artists] = await db.execute(query, artistIds);
    
    res.json({
      message: true,
      data: artists
    });
  } catch (error) {
    console.error('获取歌手信息失败:', error);
    res.status(500).json({ message: false, error: '获取歌手信息失败' });
  }
});
//----------------------------------歌曲----------------------------------




//----------------------------------喜欢的歌曲----------------------------------
//获取用户喜欢的歌曲
router.get('/liked-songs/:id', verifyToken, async (req, res) => {
  const userId = req.params.id;
  const songs = await db.execute('SELECT * FROM song WHERE id IN (SELECT song_id FROM `like` WHERE user_id = ?)', [userId]);
  res.json({ message: true, data: songs });
  //注意两个接口的区别，一个是获取单个用户的所有喜欢的歌曲，一个是用户的喜欢的歌曲，分页查询  
  //注意一定要打上反双引号，否则保留字报错
});
// 获取用户喜欢的歌曲，分页查询
router.get('/liked-songs', verifyToken, async (req, res) => {
  try {
    const userId = req.query.id;
    // 确保所有参数都被正确转换
    const page = parseInt(req.query.page) || 1;
    const pageSize = parseInt(req.query.pageSize) || 10;
    const search = String(req.query.search || '');
    const sortBy = String(req.query.sortBy || 'latest');
    const offset = (page - 1) * pageSize;

    // 构建基础查询
    let query = `
      SELECT 
        s.*,
        a.name as artist_name,
        a.id as artist_id,
        al.name as album_name,
        al.cover as album_cover,
        al.id as album_id
      FROM \`like\` l
      JOIN song s ON l.song_id = s.id
      LEFT JOIN artist a ON s.author_id = a.id
      LEFT JOIN album al ON s.album_id = al.id
      WHERE l.user_id = ${userId}
    `;

    // 构建计数查询
    let countQuery = `
      SELECT COUNT(*) as total
      FROM \`like\` l
      JOIN song s ON l.song_id = s.id
      WHERE l.user_id = ${userId}
    `;

    // 添加搜索条件
    if (search) {
      query += ` AND (s.name LIKE '%${search}%' OR a.name LIKE '%${search}%')`;
      countQuery += ` AND (s.name LIKE '%${search}%' OR a.name LIKE '%${search}%')`;
    }

    // 添加排序
    const sortMapping = {
      'latest': 'l.created_at DESC',
      'oldest': 'l.created_at ASC',
      'name': 's.name ASC',
      'duration': 's.duration ASC'
    };
    
    query += ` ORDER BY ${sortMapping[sortBy] || 'l.created_at DESC'}`;

    // 添加分页
    query += ` LIMIT ${pageSize} OFFSET ${offset}`;

    // 执行查询
    const [songs] = await db.execute(query);
    const [totalResult] = await db.execute(countQuery);

    // 格式化返回数据
    const formattedSongs = songs.map(song => ({
      id: song.id,
      name: song.name,
      artist: song.artist_name || '未知歌手',
      album: song.album_name || '未知专辑',
      cover: song.album_cover || '/assets/default-cover.jpg',
      duration: song.duration || 0,
      artist_id: song.artist_id,
      album_id: song.album_id,
      url: song.url
      //注意这里也需要修改，返回了格式化数据，前后端方便对接
    }));

    res.json({
      message: true,
      data: {
        total: totalResult[0].total,
        list: formattedSongs
      }
    });

  } catch (error) {
    console.error('获取喜欢的歌曲失败:', error);
    res.status(500).json({
      message: false,
      error: '获取喜欢的歌曲失败'
    });
  }
});
// 添加喜欢的歌曲
router.post('/like/:songId', verifyToken, async (req, res) => {
  try {
    const userId = req.user.id;
    const songId = req.params.songId;

    // 检查歌曲是否存在
    const [songExists] = await db.execute(
      'SELECT id FROM song WHERE id = ?',
      [songId]
    );

    if (songExists.length === 0) {
      return res.status(404).json({
        message: false,
        error: '歌曲不存在'
      });
    }

    // 添加到喜欢列表
    await db.execute(
      'INSERT INTO `like` (user_id, song_id) VALUES (?, ?) ON DUPLICATE KEY UPDATE created_at = CURRENT_TIMESTAMP',
      [userId, songId]
    );

    res.json({
      message: true,
      data: '添加成功'
    });

  } catch (error) {
    console.error('添加喜欢的歌曲失败:', error);
    res.status(500).json({
      message: false,
      error: '添加喜欢的歌曲失败'
    });
  }
});
// 取消喜欢的歌曲
router.delete('/like/:songId', verifyToken, async (req, res) => {
  try {
    const userId = req.user.id;
    const songId = req.params.songId;

    // 删除喜欢记录
    const [result] = await db.execute(
      'DELETE FROM `like` WHERE user_id = ? AND song_id = ?',
      [userId, songId]
    );

    if (result.affectedRows === 0) {
      return res.status(404).json({
        message: false,
        error: '该歌曲不在喜欢列表中'
      });
    }

    res.json({
      message: true,
      data: '取消成功'
    });

  } catch (error) {
    console.error('取消喜欢的歌曲失败:', error);
    res.status(500).json({
      message: false,
      error: '取消喜欢的歌曲失败'
    });
  }
});
//----------------------------------喜欢的歌曲----------------------------------





//----------------------------------歌单----------------------------------
// 获取用户歌单列表
router.get('/my-playlists', verifyToken, async (req, res) => {
  try {
    // const userId = req.user.id.toString(); // 转为字符串
    //这里修改，不能从拦截器添加的当前用户id处获取，逻辑错误
    const userId = req.query.id;
    const page = parseInt(req.query.page) || 1;
    const pageSize = parseInt(req.query.pageSize) || 10;
    const offset = (page - 1) * pageSize; // 保持为数字
    const limit = pageSize; // 保持为数字

    const [playlists] = await db.query(
      `
      SELECT 
        pi.*,
        u.username as creator_name,
        (SELECT COUNT(*) FROM playlist_songs ps WHERE ps.playlist_id = pi.id) as song_count
      FROM playlist_info pi
      LEFT JOIN user u ON pi.user_id = u.id
      WHERE pi.user_id = ?
      ORDER BY pi.created_at DESC LIMIT ? OFFSET ?`,
      [userId, limit, offset] // userId 保持字符串，limit 和 offset 保持数字
    );

    // 获取总数
    const [total] = await db.query(
      'SELECT COUNT(*) as total FROM playlist_info WHERE user_id = ?',
      [userId]
    );

    res.json({
      message: true,
      data: {
        playlists,
        total: total[0].total
      }
    });
  } catch (error) {
    console.error('获取歌单列表失败:', error);
    res.json({
      message: false,
      error: '获取歌单列表失败'
    });
  }
});

// 创建歌单
router.post('/playlist', verifyToken, async (req, res) => {
  try {
    const userId = req.user.id;
    const { name, description, isPublic = false } = req.body;

    if (!name) {
      return res.status(400).json({
        message: false,
        error: '歌单名称不能为空'
      });
    }

    const query = `
      INSERT INTO playlist_info 
      (name, description, user_id, is_public, cover) 
      VALUES (?, ?, ?, ?, ?)
    `;

    const [result] = await db.execute(query, [
      name,
      description || null,
      userId,
      isPublic,
      '/assets/covers/playlists/default.jpg'
    ]);

    res.json({
      message: true,
      data: {
        id: result.insertId,
        name,
        description,
        isPublic
      }
    });

  } catch (error) {
    console.error('创建歌单失败:', error);
    res.status(500).json({
      message: false,
      error: '创建歌单失败'
    });
  }
});

// 删除歌单
router.delete('/playlist/:id', verifyToken, async (req, res) => {
  try {
    const userId = req.user.id;
    const playlistId = req.params.id;
    console.log(userId);
    console.log(playlistId);
    // 验证歌单所有权
    const [playlist] = await db.execute(
      'SELECT user_id FROM playlist_info WHERE id = ?',
      [playlistId]
    );
    console.log(playlist);
    console.log(userId);
    if (!playlist.length || playlist[0].user_id !== userId) {
      return res.status(403).json({
        message: false,
        error: '没有权限删除此歌单'
      });
    }

    // 首先删除歌单中的所有歌曲
    await db.execute(
      'DELETE FROM playlist_songs WHERE playlist_id = ?',
      [playlistId]
    );

    // 然后删除歌单
    await db.execute(
      'DELETE FROM playlist_info WHERE id = ?',
      [playlistId]
    );

    res.json({
      message: true,
      data: { id: playlistId }
    });

  } catch (error) {
    console.error('删除歌单失败:', error);
    res.status(500).json({
      message: false,
      error: '删除歌单失败'
    });
  }
});

// 编辑歌单
router.put('/playlist/:id', verifyToken, async (req, res) => {
  try {
    const userId = req.user.id;
    const playlistId = req.params.id;
    const { name, description, isPublic, cover } = req.body;

    // 验证歌单所有权
    const [playlist] = await db.execute(
      'SELECT user_id FROM playlist_info WHERE id = ?',
      [playlistId]
    );

    if (!playlist.length || playlist[0].user_id !== userId) {
      return res.status(403).json({
        message: false,
        error: '没有权限编辑此歌单'
      });
    }

    // 构建更新查询
    const updates = [];
    const params = [];

    if (name) {
      updates.push('name = ?');
      params.push(name);
    }
    if (description !== undefined) {
      updates.push('description = ?');
      params.push(description);
    }
    if (isPublic !== undefined) {
      updates.push('is_public = ?');
      params.push(isPublic);
    }
    if (cover) {
      updates.push('cover = ?');
      params.push(cover);
    }

    if (updates.length === 0) {
      return res.status(400).json({
        message: false,
        error: '没有提供要更新的内容'
      });
    }

    params.push(playlistId);

    await db.execute(
      `UPDATE playlist_info SET ${updates.join(', ')} WHERE id = ?`,
      params
    );

    res.json({
      message: true,
      data: {
        id: playlistId,
        name,
        description,
        isPublic,
        cover
      }
    });

  } catch (error) {
    console.error('编辑歌单失败:', error);
    res.status(500).json({
      message: false,
      error: '编辑歌单失败'
    });
  }
});

// 添加歌曲到歌单
router.post('/playlist/:playlistId/songs', verifyToken, async (req, res) => {
  try {
    const playlistId = req.params.playlistId.toString();
    const songId = req.body.songId.toString();
    const userId = req.user.id.toString();

    // 验证歌单所有权
    const [playlist] = await db.query(
      'SELECT * FROM playlist_info WHERE id = ? AND user_id = ?',
      [playlistId, userId]
    );

    if (playlist.length === 0) {
      return res.json({
        message: false,
        error: '无权操作此歌单'
      });
    }

    // 检查歌曲是否已在歌单中
    const [existingSong] = await db.query(
      'SELECT * FROM playlist_songs WHERE playlist_id = ? AND song_id = ?',
      [playlistId, songId]
    );

    if (existingSong.length > 0) {
      return res.json({
        message: false,
        error: '歌曲已在歌单中'
      });
    }

    // 添加歌曲到歌单 (添加 user_id 字段)
    await db.query(
      'INSERT INTO playlist_songs (playlist_id, song_id, user_id) VALUES (?, ?, ?)',
      [playlistId, songId, userId]
    );

    // 更新歌单的更新时间
    await db.query(
      'UPDATE playlist_info SET updated_at = NOW() WHERE id = ?',
      [playlistId]
    );

    res.json({
      message: true,
      data: '添加成功'
    });
  } catch (error) {
    console.error('添加歌曲到歌单失败:', error);
    res.json({
      message: false,
      error: '添加歌曲失败'
    });
  }
});

// 从歌单中获取歌曲
router.get('/playlist/:id', async (req, res) => {
  try {
    const playlistId = req.params.id.toString();
    const page = parseInt(req.query.page) || 1;
    const pageSize = parseInt(req.query.pageSize) || 10;
    const search = req.query.search || '';
    const sortBy = req.query.sortBy || 'latest';
    const offset = (page - 1) * pageSize;

    // 定义允许的排序字段
    const allowedSortFields = {
      'id': 's.id',
      'name': 's.name',
      'latest': 'ps.added_at',
      'oldest': 'ps.added_at ASC',
      'duration': 's.duration'
    };

    // 验证排序字段
    const sortField = allowedSortFields[sortBy] || 'ps.added_at';
    const sortOrder = sortBy === 'oldest' ? '' : 'DESC';
    
    // 获取歌单信息
    const [playlist] = await db.query(
      `SELECT pi.*, u.username as creator_name 
      FROM playlist_info pi 
      LEFT JOIN user u ON pi.user_id = u.id 
      WHERE pi.id = ?`,
      [playlistId]
    );

    if (playlist.length === 0) {
      return res.json({
        message: false,
        error: '歌单不存在'
      });
    }

    // 修改歌曲查询，通过 playlist_songs 表关联
    let query = `
      SELECT 
        s.*,
        al.name as album_name,
        al.cover as album_cover,
        al.id as album_id,
        ar.name as artist_name,
        ar.id as artist_id,
        ps.added_at
      FROM playlist_songs ps
      JOIN song s ON ps.song_id = s.id
      LEFT JOIN album al ON s.album_id = al.id
      LEFT JOIN artist ar ON s.author_id = ar.id
      WHERE ps.playlist_id = ?
    `;

    // 修改计数查询
    let countQuery = `
      SELECT COUNT(*) as total
      FROM playlist_songs ps
      JOIN song s ON ps.song_id = s.id
      WHERE ps.playlist_id = ?
    `;

    let queryParams = [playlistId];
    let countParams = [playlistId];

    // 添加搜索条件
    if (search) {
      query += ` AND (s.name LIKE ? OR ar.name LIKE ?)`;
      countQuery += ` AND (s.name LIKE ? OR ar.name LIKE ?)`;
      const searchPattern = `%${search}%`;
      queryParams.push(searchPattern, searchPattern);
      countParams.push(searchPattern, searchPattern);
    }

    // 添加排序和分页
    query += ` ORDER BY ${sortField} ${sortOrder} LIMIT ? OFFSET ?`;
    queryParams.push(pageSize, offset);

    // 执行查询
    const [songs] = await db.query(query, queryParams);
    const [totalResult] = await db.query(countQuery, countParams);

    res.json({
      message: true,
      data: {
        playlist: playlist[0],
        total: totalResult[0].total,
        songs: songs
      }
    });
  } catch (error) {
    console.error('获取歌单歌曲失败:', error);
    res.json({
      message: false,
      error: '获取歌单歌曲失败'
    });
  }
});

// 从歌单中移除歌曲
router.delete('/playlist/:playlistId/songs/:songId', async (req, res) => {
  try {
    const playlistId = req.params.playlistId.toString();
    const songId = req.params.songId.toString();
    
    // 检查歌单是否存在
    const [playlist] = await db.query(
      'SELECT * FROM playlist_info WHERE id = ?',
      [playlistId]
    );

    if (playlist.length === 0) {
      return res.json({
        message: false,
        error: '歌单不存在'
      });
    }

    // 检查歌曲是否在歌单中
    const [playlistSong] = await db.query(
      'SELECT * FROM playlist_songs WHERE playlist_id = ? AND song_id = ?',
      [playlistId, songId]
    );

    if (playlistSong.length === 0) {
      return res.json({
        message: false,
        error: '歌曲不在歌单中'
      });
    }

    // 从歌单中移除歌曲
    await db.query(
      'DELETE FROM playlist_songs WHERE playlist_id = ? AND song_id = ?',
      [playlistId, songId]
    );

    res.json({
      message: true,
      data: '歌曲已从歌单中移除'
    });
  } catch (error) {
    console.error('移除歌曲失败:', error);
    res.json({
      message: false,
      error: '移除歌曲失败'
    });
  }
});
//----------------------------------歌单----------------------------------



//----------------------------------专辑----------------------------------

// 获取专辑列表
router.get('/albums', async (req, res) => {
  try {
    const page = parseInt(req.query.page) || 1;
    const pageSize = parseInt(req.query.pageSize) || 10;
    const search = req.query.search || '';
    const sortBy = req.query.sortBy || 'latest';
    const offset = (page - 1) * pageSize;

    // 定义允许的排序字段
    const allowedSortFields = {
      'id': 'a.id',
      'name': 'a.name',
      'latest': 'a.created_at',
      'oldest': 'a.created_at ASC',
      'release_date': 'a.release_date'
    };

    // 验证排序字段
    const sortField = allowedSortFields[sortBy] || 'a.created_at';
    const sortOrder = sortBy === 'oldest' ? '' : 'DESC';

    // 构建查询
    let query = `
      SELECT DISTINCT
        a.*,
        GROUP_CONCAT(DISTINCT ar.name) as artist_names,
        COUNT(DISTINCT s.id) as song_count
      FROM album a
      LEFT JOIN artist_album aa ON a.id = aa.album_id
      LEFT JOIN artist ar ON aa.artist_id = ar.id
      LEFT JOIN song s ON a.id = s.album_id
    `;

    let countQuery = `
      SELECT COUNT(DISTINCT a.id) as total
      FROM album a
      LEFT JOIN artist_album aa ON a.id = aa.album_id
      LEFT JOIN artist ar ON aa.artist_id = ar.id
    `;

    let queryParams = [];
    let countParams = [];

    // 添加搜索条件
    if (search) {
      query += ` WHERE a.name LIKE ? OR ar.name LIKE ?`;
      countQuery += ` WHERE a.name LIKE ? OR ar.name LIKE ?`;
      const searchPattern = `%${search}%`;
      queryParams.push(searchPattern, searchPattern);
      countParams.push(searchPattern, searchPattern);
    }

    query += ` GROUP BY a.id ORDER BY ${sortField} ${sortOrder} LIMIT ? OFFSET ?`;
    queryParams.push(pageSize, offset);

    // 执行查询
    const [albums] = await db.query(query, queryParams);
    const [totalResult] = await db.query(countQuery, countParams);

    res.json({
      message: true,
      data: {
        total: totalResult[0].total,
        albums: albums
      }
    });
  } catch (error) {
    console.error('获取专辑列表失败:', error);
    res.json({
      message: false,
      error: '获取专辑列表失败'
    });
  }
});

// 获取专辑详情及其歌曲
router.get('/album/:id', async (req, res) => {
  try {
    const albumId = req.params.id.toString();
    const page = parseInt(req.query.page) || 1;
    const pageSize = parseInt(req.query.pageSize) || 10;
    const search = req.query.search || '';
    const sortBy = req.query.sortBy || 'latest';
    const offset = (page - 1) * pageSize;

    // 定义允许的排序字段
    const allowedSortFields = {
      'id': 's.id',
      'name': 's.name',
      'latest': 's.create_time',
      'oldest': 's.create_time ASC',
      'duration': 's.duration',
      'play_count': 's.play_count'
    };

    // 验证排序字段
    const sortField = allowedSortFields[sortBy] || 's.create_time';
    const sortOrder = sortBy === 'oldest' ? '' : 'DESC';

    // 获取专辑信息
    const [album] = await db.query(
      `SELECT a.*, 
        GROUP_CONCAT(DISTINCT ar.name) as artist_names,
        GROUP_CONCAT(DISTINCT ar.id) as artist_ids
      FROM album a
      LEFT JOIN artist_album aa ON a.id = aa.album_id
      LEFT JOIN artist ar ON aa.artist_id = ar.id
      WHERE a.id = ?
      GROUP BY a.id`,
      [albumId]
    );

    if (album.length === 0) {
      return res.json({
        message: false,
        error: '专辑不存在'
      });
    }

    // 构建歌曲查询
    let query = `
      SELECT 
        s.*,
        a.name as artist_name,
        al.cover as album_cover,
        al.name as album_name
      FROM song s
      LEFT JOIN artist a ON s.author_id = a.id
      LEFT JOIN album al ON s.album_id = al.id
      WHERE s.album_id = ?
    `;

    // 构建计数查询
    let countQuery = `
      SELECT COUNT(*) as total
      FROM song s
      LEFT JOIN artist a ON s.author_id = a.id
      WHERE s.album_id = ?
    `;

    let queryParams = [albumId];
    let countParams = [albumId];

    // 添加搜索条件
    if (search) {
      query += ` AND (s.name LIKE ? OR a.name LIKE ?)`;
      countQuery += ` AND (s.name LIKE ? OR a.name LIKE ?)`;
      const searchPattern = `%${search}%`;
      queryParams.push(searchPattern, searchPattern);
      countParams.push(searchPattern, searchPattern);
    }

    // 添加排序和分页
    query += ` ORDER BY ${sortField} ${sortOrder} LIMIT ? OFFSET ?`;
    queryParams.push(pageSize, offset);

    // 执行查询
    const [songs] = await db.query(query, queryParams);
    const [totalResult] = await db.query(countQuery, countParams);

    res.json({
      message: true,
      data: {
        album: album[0],
        total: totalResult[0].total,
        songs: songs
      }
    });
  } catch (error) {
    console.error('获取专辑详情失败:', error);
    res.json({
      message: false,
      error: '获取专辑详情失败'
    });
  }
});
//----------------------------------专辑----------------------------------



//----------------------------------歌手----------------------------------
// 获取艺术家详情及其歌曲
router.get('/artist/:id', async (req, res) => {
  try {
    const artistId = req.params.id.toString();
    const page = parseInt(req.query.page) || 1;
    const pageSize = parseInt(req.query.pageSize) || 10;
    const search = req.query.search || '';
    const sortBy = req.query.sortBy || 'latest';
    const offset = (page - 1) * pageSize;

    // 定义允许的排序字段
    const allowedSortFields = {
      'id': 's.id',
      'name': 's.name',
      'latest': 's.create_time',
      'oldest': 's.create_time ASC',
      'duration': 's.duration',
      'play_count': 's.play_count'
    };

    // 验证排序字段
    const sortField = allowedSortFields[sortBy] || 's.create_time';
    const sortOrder = sortBy === 'oldest' ? '' : 'DESC';

    // 获取艺术家信息
    const [artist] = await db.query(
      `SELECT ar.*, 
        GROUP_CONCAT(DISTINCT al.name) as album_names,
        GROUP_CONCAT(DISTINCT al.id) as album_ids
      FROM artist ar
      LEFT JOIN artist_album aa ON ar.id = aa.artist_id
      LEFT JOIN album al ON aa.album_id = al.id
      WHERE ar.id = ?
      GROUP BY ar.id`,
      [artistId]
    );
    //最好还是把所有信息都left_join出来

    if (artist.length === 0) {
      return res.json({
        message: false,
        error: '艺术家不存在'
      });
    }

    // 构建歌曲查询
    let query = `
      SELECT 
        s.*,
        al.name as album_name,
        al.cover as album_cover,
        al.id as album_id,
        ar.name as artist_name,
        ar.id as artist_id
      FROM song s
      LEFT JOIN album al ON s.album_id = al.id
      LEFT JOIN artist ar ON s.author_id = ar.id
      WHERE s.author_id = ?
    `;

    // 构建计数查询
    let countQuery = `
      SELECT COUNT(*) as total
      FROM song s
      WHERE s.author_id = ?
    `;

    let queryParams = [artistId];
    let countParams = [artistId];

    // 添加搜索条件
    if (search) {
      query += ` AND (s.name LIKE ? OR al.name LIKE ?)`;
      countQuery += ` AND s.name LIKE ?`;
      const searchPattern = `%${search}%`;
      queryParams.push(searchPattern, searchPattern);
      countParams.push(searchPattern);
    }

    // 添加排序和分页
    query += ` ORDER BY ${sortField} ${sortOrder} LIMIT ? OFFSET ?`;
    queryParams.push(pageSize, offset);

    // 执行查询
    const [songs] = await db.query(query, queryParams);
    const [totalResult] = await db.query(countQuery, countParams);

    // 处理专辑ID和名称字符串
    if (artist[0].album_ids) {
      artist[0].album_ids = artist[0].album_ids.split(',').map(id => parseInt(id));
      artist[0].album_names = artist[0].album_names.split(',');
    } else {
      artist[0].album_ids = [];
      artist[0].album_names = [];
    }

    res.json({
      message: true,
      data: {
        artist: artist[0],
        total: totalResult[0].total,
        songs: songs
      }
    });
  } catch (error) {
    console.error('获取艺术家详情失败:', error);
    res.json({
      message: false,
      error: '获取艺术家详情失败'
    });
  }
});
//----------------------------------歌手----------------------------------





//----------------------------------评论相关接口----------------------------------
// 获取歌曲评论列表
router.get('/comments/:songId', async (req, res) => {
  try {
    const songId = req.params.songId.toString(); // 转换为字符串
    const page = parseInt(req.query.page) || 1;
    const pageSize = parseInt(req.query.pageSize) || 10;
    const offset = (page - 1) * pageSize;

    // 获取主��论
    const query = `
      SELECT 
        c.*,
        u.username,
        u.avatar,
        (SELECT COUNT(*) FROM comment_like WHERE comment_id = c.id) as like_count
      FROM comment c
      LEFT JOIN user u ON c.user_id = u.id
      WHERE c.song_id = ? AND c.parent_id IS NULL
      ORDER BY c.created_at DESC
      LIMIT ?, ?
    `;

    const countQuery = `
      SELECT COUNT(*) as total
      FROM comment
      WHERE song_id = ?
    `;

    // 使用数组传递参数，确保参数类型正确
    const [comments] = await db.execute(query, [songId, offset.toString(), pageSize.toString()]);
    const [totalResult] = await db.execute(countQuery, [songId]);

    res.json({
      message: true,
      data: {
        comments,
        total: totalResult[0].total
      }
    });
  } catch (error) {
    console.error('获取评论失败:', error);
    res.status(500).json({
      message: false,
      error: '获取评论失败'
    });
  }
});

// 获取评论的回复列表
router.get('/comments/:commentId/replies', async (req, res) => {
  try {
    const commentId = req.params.commentId.toString(); // 转换为字符串
    
    const query = `
      SELECT 
        c.*,
        u.username,
        u.avatar,
        ru.username as reply_to_username,
        (SELECT COUNT(*) FROM comment_like WHERE comment_id = c.id) as like_count
      FROM comment c
      LEFT JOIN user u ON c.user_id = u.id
      LEFT JOIN user ru ON c.reply_to_user_id = ru.id
      WHERE c.parent_id = ?
      ORDER BY c.created_at ASC
    `;

    const [replies] = await db.execute(query, [commentId]);

    res.json({
      message: true,
      data: {
        data: {
          replies
        }
      }
    });
  } catch (error) {
    console.error('获取回复失败:', error);
    res.status(500).json({
      message: false,
      data: null,
      error: '获取回复失败'
    });
  }
});

// 点赞评论
router.post('/comments/:commentId/like', verifyToken, async (req, res) => {
  try {
    const commentId = req.params.commentId.toString(); // 确保 commentId 是字符串
    const userId = req.user.id.toString(); // 确保 userId 是字符串

    // 检查评论是否存在
    const [comment] = await db.execute(
      'SELECT id FROM comment WHERE id = ?',
      [commentId]
    );

    if (comment.length === 0) {
      return res.status(404).json({
        message: false,
        data: null,
        error: '评论不存在'
      });
    }

    // 检查是否已经点赞
    const [existingLike] = await db.execute(
      'SELECT * FROM comment_like WHERE comment_id = ? AND user_id = ?',
      [commentId, userId]
    );

    if (existingLike.length > 0) {
      // 如果已经点赞，则取消点赞
      await db.execute(
        'DELETE FROM comment_like WHERE comment_id = ? AND user_id = ?',
        [commentId, userId]
      );
    } else {
      // 如果未点赞，则添加点赞
      await db.execute(
        'INSERT INTO comment_like (comment_id, user_id) VALUES (?, ?)',
        [commentId, userId]
      );
    }

    // 获取最新点赞数
    const [likeCount] = await db.execute(
      'SELECT COUNT(*) as count FROM comment_like WHERE comment_id = ?',
      [commentId]
    );

    res.json({
      message: true,
      data: {
        data: {
          likeCount: likeCount[0].count,
          isLiked: existingLike.length === 0
        }
      }
    });
  } catch (error) {
    console.error('点赞操作失败:', error);
    res.status(500).json({
      message: false,
      data: null,
      error: '点赞操作失败'
    });
  }
});

// 删除评论
router.delete('/comments/:commentId', verifyToken, async (req, res) => {
  try {
    const commentId = req.params.commentId;
    const userId = req.user.id;

    // 检查是否是评论作者
    const [comment] = await db.query(
      'SELECT user_id FROM comment WHERE id = ?',
      [commentId]
    );

    if (comment.length === 0) {
      return res.status(404).json({
        message: false,
        error: '评论不存在'
      });
    }

    if (comment[0].user_id !== userId) {
      return res.status(403).json({
        message: false,
        error: '无权删除该评论'
      });
    }

    // 删除评论及其所有回复
    await db.execute('DELETE FROM comment WHERE id = ? OR parent_id = ?', [commentId, commentId]);

    res.json({
      message: true,
      data: '删除成功'
    });
  } catch (error) {
    console.error('删除评论失败:', error);
    res.status(500).json({
      message: false,
      error: '删除评论失败'
    });
  }
});

// 发表评论
router.post('/comments', verifyToken, async (req, res) => {
  try {
    const { songId, text, parentId, replyToUserId } = req.body;
    const userId = req.user.id;

    // 验证必要参数
    if (!songId || !text) {
      return res.status(400).json({
        message: false,
        data: null,
        error: '缺少必要参数'
      });
    }

    // 检查歌曲是否存在
    const [song] = await db.execute('SELECT id FROM song WHERE id = ?', [songId]);
    if (song.length === 0) {
      return res.status(404).json({
        message: false,
        data: null,
        error: '歌曲不存在'
      });
    }

    // 如果是回复评论，检查父评论是否存在
    if (parentId) {
      const [parentComment] = await db.execute(
        'SELECT id FROM comment WHERE id = ?',
        [parentId]
      );
      if (parentComment.length === 0) {
        return res.status(404).json({
          message: false,
          data: null,
          error: '父评论不存在'
        });
      }
    }

    // 插入评论
    const query = `
      INSERT INTO comment 
        (song_id, user_id, text, parent_id, reply_to_user_id, created_at) 
      VALUES 
        (?, ?, ?, ?, ?, NOW())
    `;

    const [result] = await db.execute(query, [
      songId,
      userId,
      text,
      parentId || null,
      replyToUserId || null
    ]);

    // 获取新创建的评论详情
    const [newComment] = await db.execute(`
      SELECT 
        c.*,
        u.username,
        u.avatar,
        ru.username as reply_to_username
      FROM comment c
      LEFT JOIN user u ON c.user_id = u.id
      LEFT JOIN user ru ON c.reply_to_user_id = ru.id
      WHERE c.id = ?
    `, [result.insertId]);

    res.json({
      message: true,
      data: {
        data: newComment[0]
      }
    });

  } catch (error) {
    console.error('发表评论失败:', error);
    res.status(500).json({
      message: false,
      data: null,
      error: '发表评论失败'
    });
  }
});


//----------------------------------评论相关接口----------------------------------



module.exports = router;
