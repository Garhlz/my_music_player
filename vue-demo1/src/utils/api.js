import { ElMessage, ElMessageBox } from 'element-plus';
import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8080', // 后端 API 根地址
  timeout: 10000, // 超时时间
  headers: {
    'Content-Type': 'application/json',  // 设置请求头中的 Content-Type 为 JSON
  },
});

api.interceptors.request.use(
  (config) => {
    if (config.url === '/user/login' || config.url === '/user/register') {
      return config;
    }

    const token = localStorage.getItem('token');
    if (!token) {
      ElMessage.error('请先登录');
      setTimeout(() => {
        window.location.href = '/login';  // 跳转到登录页
      }, 1000);
      return Promise.reject('No token');
    }
    config.headers['Authorization'] = `Bearer ${token}`;
    return config;
  },
  (error) => Promise.reject(error)
);

api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response && error.response.status === 401) {
      localStorage.removeItem('token');
      ElMessage.error('身份验证失败，请重新登录');
      setTimeout(() => {
        window.location.href = '/login';  // 跳转到登录页
      }, 1000);
      return Promise.resolve(null);
    }
    return Promise.reject(error);
  }
);

export { api };
