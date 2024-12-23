const jwt = require('jsonwebtoken');
const dotenv = require('dotenv');

// 加载环境变量
dotenv.config();

// 获取 JWT 密钥
const JWT_SECRET = process.env.JWT_SECRET || 'your_jwt_secret_key';

/**
 * 验证 JWT 是否有效的中间件
 */
const verifyToken = (req, res, next) => {
  // 从请求头中获取 token
  const token = req.header('Authorization') && req.header('Authorization').split(' ')[1];
  
  if (!token) {
    return res.status(401).json({ message: 'Access denied. No token provided.' });
  }

  try {
    // 验证 token
    const decoded = jwt.verify(token, JWT_SECRET);
    
    // 将解码后的用户信息存到请求对象中，以便后续使用
    req.user = decoded;

    // 继续处理请求
    next();
  } catch (err) {
    // 如果 token 无效或过期，返回 401 错误
    return res.status(401).json({ message: 'Invalid or expired token.' });
  }
};

module.exports = { verifyToken };
