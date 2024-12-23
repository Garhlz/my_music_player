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
    const { username, password, name, phone, email, sex } = req.body;

    if (!username || !password) {
        return res.status(400).json({ error: '用户名和密码不能为空' });
    }

    try {
        const [result] = await db.execute('SELECT * FROM user WHERE username = ?', [username]);
        const existingUser = result || []; // 确保 existingUser 是一个数组

        if (existingUser.length > 0) {
            return res.status(409).json({ error: '用户名已存在' });
        }

        const hashedPassword = await bcrypt.hash(password, 10); // 加密密码
        await db.execute('INSERT INTO user (username, password, name, phone, email, sex) VALUES (?, ?, ?, ?, ?, ?)', 
            [username, hashedPassword, name, phone, email, sex]); // 存储加密后的密码

        return res.status(201).json({ message: '注册成功' });
    } catch (error) {
        console.error('注册失败:', error);
        return res.status(500).json({ error: '注册过程中发生错误，请稍后再试' });
    }
});
//----------------------------------用户登录注册----------------------------------





//----------------------------------用户信息----------------------------------
// 获取用户信息接口
router.get('/user/{id}',verifyToken,  async (req, res) => {
    const userId = req.params.id;

    try {
        const [user] = await db.execute('SELECT id, username, name, phone, email, sex FROM user WHERE id = ?', [userId]);
        if (user.length > 0) {
            return res.status(200).json({ user: user[0] });
        } else {
            return res.status(404).json({ error: '用户未找到' });
        }
    } catch (error) {
        console.error('数据库查询失败:', error);
        return res.status(500).json({ error: '服务器错误' });
    }
});
//----------------------------------用户信息----------------------------------





//----------------------------------歌曲----------------------------------
// 获取歌曲列表接口
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

// 获取专辑信息
router.get('/albums', verifyToken, async (req, res) => {
  try {
    const { ids } = req.query;
    if (!ids) {
      return res.status(400).json({ message: false, error: '缺少专辑ID' });
    }

    // 处理数组格式的参数
    let albumIds = Array.isArray(ids) ? ids : [ids];
    
    // 确保所有ID都是数字
    albumIds = albumIds.map(id => parseInt(id)).filter(id => !isNaN(id));
    
    if (albumIds.length === 0) {
      return res.status(400).json({ message: false, error: '无效的专辑ID' });
    }

    // 构建查询语句
    const placeholders = albumIds.map(() => '?').join(',');
    const query = `
      SELECT id, name, cover, release_date, description
      FROM album
      WHERE id IN (${placeholders})
    `;
    
    const [albums] = await db.execute(query, albumIds);
    
    res.json({
      message: true,
      data: albums
    });
  } catch (error) {
    console.error('获取专辑信息失败:', error);
    res.status(500).json({ message: false, error: '获取专辑信息失败' });
  }
});
//----------------------------------歌曲----------------------------------




//----------------------------------喜欢的歌曲----------------------------------
//获取用户喜欢的歌曲
router.get('/liked-songs/:id', verifyToken, async (req, res) => {
    //注意两个接口的区别，一个是获取单个用户的所有喜欢的歌曲，一个是用户的喜欢的歌曲，分页查询  
    //注意一定要打上反双引号，否则保留字报错
  const userId = req.params.id;
  const songs = await db.execute('SELECT * FROM song WHERE id IN (SELECT song_id FROM `like` WHERE user_id = ?)', [userId]);
  res.json({ message: true, data: songs });
});

// 获取用户喜欢的歌曲，分页查询
router.get('/liked-songs', verifyToken, async (req, res) => {
  try {
    const userId = req.user.id;
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
        al.name as album_name,
        al.cover as album_cover
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
      url: song.url
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
//----------------------------------歌曲----------------------------------





//----------------------------------歌单----------------------------------
// 获取用户歌单列表
router.get('/my-playlists', verifyToken, async (req, res) => {
  try {
    const userId = req.user.id.toString(); // 转为字符串
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
//----------------------------------歌单----------------------------------

module.exports = router;
