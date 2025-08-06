//我直接把所有前端的请求集成在这个文件中统一修改，对于我的项目来说更加合适
import {api} from '@/utils/api';

//----------------------------------用户登录注册----------------------------------
//---用户信息相关接口---
export const  login = async (username, password) => {
    const resp = await api.post('/user/login',{
      username,
      password
    });
    return resp;
};

export const register = async (userData) => {
  const resp = await api.post('/user/register', userData);
  return resp;
};
//----------------------------------用户登录注册----------------------------------



//----------------------------------用户相关接口----------------------------------
// 获取用户信息
export const getUserInfo = async (userId) => {
  const resp = await api.get(`/user/${userId}`);
  return resp;
};

// 更新用户信息
export const updateUserInfo = async (userId, data) => {
  const resp = await api.put(`/user/${userId}`, data);
  return resp;
};
//----------------------------------用户相关接口----------------------------------



//----------------------------------歌曲相关接口----------------------------------
export const getSongs = async (params) => {
  const resp = await api.get('/songs', { params });
  return resp;
};

export const getSongById = async (id) => {
  const resp = await api.get(`/song/${id}`);
  return resp;
};
//----------------------------------歌曲相关接口----------------------------------





//----------------------------------喜欢的歌曲相关接口----------------------------------
//---喜欢的歌曲相关接口---
export const getLikedSongsById = async (id) => {
  const response = await api.get(`/liked-songs/${id}`);
  return response;
};

//获取喜欢的歌曲
export const getLikedSongs = async (params) => {
  const response = await api.get('/liked-songs', { params });
  return response;
};

// 添加喜欢的歌曲
export const addLikedSong = async (songId) => {
    await api.post(`/like/${songId}`);
};
  
  // 取消喜欢的歌曲
export const removeLikedSong = async (songId) => {
  const response = await api.delete(`/like/${songId}`);
  return response;
};
//----------------------------------喜欢的歌曲相关接口----------------------------------





//----------------------------------歌单相关接口----------------------------------
// 获取用户歌单列表
export const getMyPlaylists = async (params = {}) => {
  const resp = await api.get('/my-playlists', { params })
  return resp
}

// 创建歌单
export const createPlaylist = async (data) => {
  const resp = await api.post('/playlist', data)
  return resp
}

// 删除歌单
export const deletePlaylist = async (playlistId) => {
  const resp = await api.delete(`/playlist/${playlistId}`)
  return resp
}

// 编辑歌单
export const updatePlaylist = async (playlistId, data) => {
  const resp = await api.put(`/playlist/${playlistId}`, data)
  return resp
}   

// 添加歌曲到歌单
export const addSongToPlaylist = async (playlistId, songId) => {
  const resp = await api.post(`/playlist/${playlistId}/songs`, { songId })
  return resp
}

// 从歌单中获取歌曲
export const getSongFromPlaylistById = async (playlistId, params) => {
    const resp = await api.get(`/playlist/${playlistId}`,{ params });
    return resp
}

// 从歌单中移除歌曲
export const removeSongFromPlaylist = async (playlistId, songId) => {
    const resp = await api.delete(`/playlist/${playlistId}/songs/${songId}`);
    return resp
}
//----------------------------------歌单相关接口----------------------------------

//----------------------------------专辑相关接口----------------------------------
// 获取专辑列表
export const getAlbums = async (params) => {
    const resp = await api.get('/albums', { params });
    return resp;
}

// 获取专辑详情及其歌曲
export const getAlbumById = async (albumId, params) => {
    const resp = await api.get(`/album/${albumId}`, { params });
    return resp;
}
//----------------------------------专辑相关接口----------------------------------



//----------------------------------歌手相关接口----------------------------------
// 获取歌手列表
export const getArtists = async (params) => {
    const resp = await api.get('/artists', { params });
    return resp;
}

// 获取歌手详情
export const getArtistById = async (artistId, params) => {
    const resp = await api.get(`/artist/${artistId}`, { params });
    return resp;
}
//----------------------------------歌手相关接口----------------------------------



//----------------------------------评论相关接口----------------------------------
// 获取歌曲评论列表
export const getComments = async (songId, params) => {
  const resp = await api.get(`/comments/${songId}`, { params });
  return resp;
};

// 获取评论回复列表
export const getCommentReplies = async (commentId, params) => {
  const resp = await api.get(`/comments/${commentId}/replies`, { params });
  return resp;
};

// 发表评论
export const createComment = async (data) => {
  const resp = await api.post('/comments', {
    songId: data.songId,
    text: data.text,
    parentId: data.parentId,
    replyToUserId: data.replyToUserId
  });
  return resp;
};

// 点赞评论
export const likeComment = async (commentId) => {
  const resp = await api.post(`/comments/${commentId}/like`);
  return resp;
};

// 删除评论
export const deleteComment = async (commentId) => {
  const resp = await api.delete(`/comments/${commentId}`);
  return resp;
};
//----------------------------------评论相关接口----------------------------------





