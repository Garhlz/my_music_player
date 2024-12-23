const express = require('express');           // Express 框架
const path = require('path');                 // 路径操作
const cookieParser = require('cookie-parser');  // 用于解析 cookie
const logger = require('morgan');             // 用于记录日志
const cors = require('cors');                // 用于解决跨域问题
const bodyParser = require('body-parser');
const dotenv = require('dotenv');

dotenv.config();
// 创建 Express 应用实例
const app = express();

app.use(bodyParser.json());

// 使用 CORS 中间件来支持跨域请求
app.use(cors());


// 使用中间件来记录请求日志
app.use(logger('dev'));  // 记录日志

// 使用中间件来处理请求体
app.use(express.json());  // 解析 JSON 格式的请求体
app.use(express.urlencoded({ extended: false }));  // 解析 URL 编码格式的请求体
app.use(cookieParser());  // 解析 cookie

// Serve static files from the "public" directory
app.use('/assets', express.static(path.join(__dirname, 'public/assets')));

// ------------------- 以下是路由导入和路径配置部分 -------------------
//之后只保留一个路由文件，对我来说反而方便管理
const routeFile = require('./routeFile');
app.use('/', routeFile);
// ------------------- 路由部分结束 -------------------

// 没有使用模板引擎渲染html页面，而是直接返回即可
// 404 错误处理：所有路由都不匹配时，返回 404 错误
app.use((req, res, next) => {
  res.status(404).json({ error: 'Not Found' });
});

// 全局错误处理器
app.use((err, req, res, next) => {
  console.error(err);  // 打印错误信息
  res.status(500).json({ error: 'Something went wrong!' });
});

module.exports = app;  // 导出 Express 应用实例，供其他模块使用
