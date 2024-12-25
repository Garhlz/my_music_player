# 音乐播放系统产品说明文档

## 产品概述
这是一个现代化的在线音乐播放系统，为用户提供全方位的音乐服务体验。系统采用Vue3 + Express技术栈开发，提供了丰富的音乐管理和社交功能。

## 核心功能

### 🎵 个性化音乐体验
1. **私人音乐库**
   - 创建个性化歌单
   - 收藏喜欢的音乐
   - 支持歌单公开/私密设置
   - 灵活的歌单管理（增删改查）

2. **智能音乐发现**
   - 公共音乐库浏览
   - 多维度音乐搜索（歌名、歌手、专辑）
   - 支持按热度、时间等多种方式排序
   - 专辑详情页浏览

3. **歌手专区**
   - 歌手详情页展示
   - 歌手作品集合
   - 相关专辑展示

### 💬 社交互动
1. **评论系统**
   - 发表音乐评论
   - 评论互动（回复、点赞）
   - 评论管理（删除）
   - 分页加载评论

2. **用户互动**
   - 个人主页展示
   - 用户信息管理
   - 歌单分享（开发中）

### 🎧 音乐播放
- 在线音乐播放
- 播放列表管理
- 音乐收藏功能

## 技术特点
1. **前后端分离架构**
   - 前端：Vue3 + Element Plus
   - 后端：Express + MySQL
   - JWT认证机制

2. **用户体验优化**
   - 响应式设计
   - 无感知Token认证
   - 统一的错误处理

3. **数据安全**
   - 密码加密存储
   - 接口权限控制
   - 数据验证机制

## API接口清单

### 用户认证接口
| 接口 | 方法 | 路径 | 描述 |
|-----|------|------|------|
| 用户注册 | POST | /api/user/register | 新用户注册 |
| 用户登录 | POST | /api/user/login | 用户登录获取token |

### 用户信息接口
| 接口 | 方法 | 路径 | 描述 |
|-----|------|------|------|
| 获取用户信息 | GET | /api/user/:id | 获取指定用户信息 |
| 更新用户信息 | PUT | /api/user/:id | 更新用户个人信息 |

### 音乐相关接口
| 接口 | 方法 | 路径 | 描述 |
|-----|------|------|------|
| 获取公共歌曲 | GET | /api/songs | 获取公共歌曲列表 |
| 获取单曲详情 | GET | /api/song/:id | 获取单个歌曲详情 |
| 获取专辑列表 | GET | /api/albums | 获取专辑列表 |
| 获取专辑详情 | GET | /api/album/:id | 获取专辑及其歌曲 |
| 获取歌手详情 | GET | /api/artist/:id | 获取歌手及其歌曲 |

### 个人音乐管理接口
| 接口 | 方法 | 路径 | 描述 |
|-----|------|------|------|
| 获取喜欢的歌曲 | GET | /api/liked-songs | 获取用户喜欢的歌曲 |
| 添加喜欢的歌曲 | POST | /api/like/:songId | 添加歌曲到喜欢列表 |
| 取消喜欢歌曲 | DELETE | /api/like/:songId | 从喜欢列表移除歌曲 |
| 获取用户歌单 | GET | /api/my-playlists | 获取用户创建的歌单 |
| 创建歌单 | POST | /api/playlist | 创建新歌单 |
| 删除歌单 | DELETE | /api/playlist/:id | 删除指定歌单 |
| 编辑歌单 | PUT | /api/playlist/:id | 更新歌单信息 |
| 获取歌单详情 | GET | /api/playlist/:id | 获取歌单及其歌曲 |
| 添加歌曲到歌单 | POST | /api/playlist/:playlistId/songs | 向歌单添加歌曲 |
| 从歌单移除歌曲 | DELETE | /api/playlist/:playlistId/songs/:songId | 从歌单移除歌曲 |

### 评论相关接口
| 接口 | 方法 | 路径 | 描述 |
|-----|------|------|------|
| 获取评论列表 | GET | /api/comments/:songId | 获取歌曲评论列表 |
| 获取评论回复 | GET | /api/comments/:commentId/replies | 获取评论的回复列表 |
| 发表评论 | POST | /api/comments | 发表评论或回复 |
| 点赞评论 | POST | /api/comments/:commentId/like | 点赞或取消点赞评论 |
| 删除评论 | DELETE | /api/comments/:commentId | 删除评论或回复 |

## 产品迭代计划

### 近期优化项目
1. **功能修复**
   - 歌单搜索功能优化
   - 评论系统bug���复
   - 个人主页UI美化

2. **性能优化**
   - 接口响应速度优化
   - 数据缓存机制
   - 前端组件复用优化

### 未来规划功能
1. **音乐播放增强**
   - 跨组件播放控制
   - 播放队列管理
   - 音频可视化效果

2. **内容管理**
   - 音乐文件上传
   - 封面图片管理
   - 音乐标签系统

3. **社交功能**
   - 用户关注系统
   - 歌单分享功能
   - 音乐动态feed流

4. **运营管理**
   - 管理员后台
   - 数据统计分析
   - 内容审核系统

## 技术支持
- 接口统一采用RESTful风格
- 所有需要认证的接口都需要携带 `Authorization: Bearer token` 请求头
- 分页接口支持 page、pageSize、search、sortBy 等通用参数
- 统一的响应格式确保了接口调用的一致性

## 项目亮点
1. 采用最新的Vue3组合式API，提供更好的代码组织和复用
2. 完善的权限控制和用户认证机制
3. 丰富的音乐社交功能
4. 灵活的歌单管理系统
5. 优秀的用户交互体验

## 联系与支持
如有任何问题或建议，请通过以下方式联系我们：
- 项目仓库：[GitHub地址]
- 问题反馈：[Issues页面]
- 技术支持：[开发者邮箱] 