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



//----------------------------------歌曲相关接口----------------------------------
export const getSongs = async (params) => {
  const resp = await api.get('/songs', { params });
  return resp;
};

//---艺术家相关接口---
export const getArtists = async (ids) => {
  const resp = await api.get('/artists', { params: { ids } });
  return resp;
};

//---专辑相关接口---
export const getAlbums = async (ids) => {
  const resp = await api.get('/albums', { params: { ids } });
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
//----------------------------------歌单相关接口----------------------------------