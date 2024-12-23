//我直接把所有前端的请求集成在这个文件中统一修改，对于我的项目来说更加合适
import {api} from '@/utils/api';



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


//---歌曲相关接口---
export const getSongs = async (params) => {
  const resp = await api.get('/songs', { params });
  return resp;
};

export const getLikedSongs = async () => {
  const resp = await api.get('/liked-songs');
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