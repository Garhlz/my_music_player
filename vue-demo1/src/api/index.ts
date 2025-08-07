import {
  AlbumApi,
  ArtistApi,
  AuthApi,
  CommentApi,
  Configuration,
  LikeApi,
  PlaylistApi,
  SongApi,
  UserApi,
} from '@/api-client';
import axios from 'axios';

// 1. 创建一个自定义的 Axios 实例
const axiosInstance = axios.create({
  // 可以在这里设置一些通用的配置，比如超时时间
  // timeout: 10000,
});


// 2. 添加请求拦截器
axiosInstance.interceptors.request.use(
  (config) => {
    // 在发送请求之前做些什么
    const token = localStorage.getItem('token'); // 从 localStorage 获取 token
    if (token) {
      // 如果 token 存在，则为每个请求都加上 Authorization 头
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    // 对请求错误做些什么
    return Promise.reject(error);
  },
);

// (可选) 你还可以添加响应拦截器，用于统一处理错误，比如 token 过期后的自动跳转登录页
axiosInstance.interceptors.response.use(
  (response) => {
    // 对响应数据做点什么
    return response;
  },
  (error) => {
    // 对响应错误做点什么
    if (error.response && error.response.status === 401) {
      // 例如，如果收到 401 Unauthorized 错误，清除本地 token 并跳转到登录页
      localStorage.removeItem('token');
      // window.location.href = '/login'; // 实际项目中，你会使用 Vue Router 来跳转
      console.error('认证失败，请重新登录。');
    }
    return Promise.reject(error);
  },
);


// 3. 创建 API 配置对象
const apiConfig = new Configuration({
  // 注意：这里的 BasePath 应该与你的 Go 后端路由组一致
  // 我们在 router.go 中使用的是 /api/v1，所以这里不需要再加 http... 部分
  // Vite 的 proxy 会帮我们代理
  basePath: '/api/v1',
});

// 4. 创建所有 API 的实例，并将我们自定义的 axios 实例传递进去
// 第三个参数就是自定义的 axios 实例
export const authApi = new AuthApi(apiConfig, undefined, axiosInstance);
export const userApi = new UserApi(apiConfig, undefined, axiosInstance);
export const songApi = new SongApi(apiConfig, undefined, axiosInstance);
export const albumApi = new AlbumApi(apiConfig, undefined, axiosInstance);
export const artistApi = new ArtistApi(apiConfig, undefined, axiosInstance);
export const likeApi = new LikeApi(apiConfig, undefined, axiosInstance);
export const playlistApi = new PlaylistApi(apiConfig, undefined, axiosInstance);
export const commentApi = new CommentApi(apiConfig, undefined, axiosInstance);