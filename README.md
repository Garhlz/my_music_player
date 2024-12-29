# [English Version](#project-introduction-online-music-streaming-system)

# 项目介绍：在线音乐播放系统

本项目是一个现代化的在线音乐播放系统，采用前后端分离架构。前端使用 Vue 3 框架，后端基于 Express 和 MySQL 开发，旨在为用户提供全面的音乐播放、管理和社交互动等功能。系统的设计和实现参考了国内外知名的音乐平台，如网易云音乐和 Spotify，力求提供一个既富有互动性又简洁流畅的音乐体验。

# 如何启动

目前项目仍处于开发阶段，尚未部署到服务器，因此只能以开发模式启动。

## 环境准备

1. 确保已安装 [Node.js](https://nodejs.org/)和npm
2. 运行 `/create_database.sql` 创建数据库
3. 运行 `/data.sql` 导入演示数据（包括披头士和鲍勃·迪伦的歌曲信息）

## 配置后端

1. 在 `/express-demo1/.env` 文件中修改数据库连接配置。请注意，`password` 存储为密文，建议手动添加一个管理员账号。
2. 后端项目使用 [Express Generator](https://expressjs.com/en/starter/generator.html) 生成项目结构。

### 启动后端服务

1. 开启一个终端，切换到 `/express-demo1` 目录。
2. 输入以下命令以下载依赖：
   ```bash
   npm install
   ```
3. 输入以下命令以启动后端服务：
   ```bash
   npm start
   ```

## 配置前端

前端使用 [Vite](https://vitejs.dev/) 和 Vue 3 全家桶。

### 启动前端服务

1. 开启另一个终端，切换到 `/vue-demo1` 目录。
2. 输入以下命令以下载依赖：
   ```bash
   npm install
   ```
3. 输入以下命令以启动前端服务：
   ```bash
   npm run dev
   ```
## 默认账户

默认账号为admin,密码为111111


## 资源位置

部分演示音乐和图片已存放在 `/vue-demo1/assets/` 目录中。

# 重要信息

- 后端接口全部定义在 `/express-demo1/routeFile.js` 中。
- 前端接口全部定义在 `/vue-demo1/src/api/axiosFile.js` 中。
- 在 `/vue-demo1/README.md` 中说明了部分开发过程以及 AI 生成的接口文档。

## 核心功能

### 1. 私人音乐库
- 创建个性化歌单
- 收藏喜爱的歌曲
- 支持歌单公开/私密设置
- 灵活的歌单管理（增、删、改、查）

### 2. 智能音乐发现
- 系统提供公共音乐库，用户可以通过歌名、歌手、专辑等多维度进行搜索
- 支持按热度、时间等方式排序
- 专辑详情页浏览

### 3. 歌手专区
- 歌手详情页展示
- 歌手作品集
- 相关专辑展示

### 4. 评论系统与用户互动
- 用户可以在歌曲下发表评论并与其他用户互动（如点赞和回复）
- 支持分页加载评论
- 用户信息管理与个人主页展示

### 5. 音乐播放与收藏功能
- 提供稳定的在线音乐播放功能
- 支持播放列表管理
- 用户可以收藏喜爱的歌曲

## 技术架构与亮点

### 1. 前后端分离架构
- **前端**：Vue 3 + Element Plus
- **后端**：Express + MySQL
- **认证机制**：JWT

### 2. 响应式设计与无感知 Token 认证
- 系统采用响应式设计，确保在不同设备上流畅运行
- Token 认证机制无感知地确保了每次用户操作的安全性

### 3. 数据安全与性能优化
- 密码加密存储、接口权限控制、数据验证机制确保用户数据安全
- 分页、缓存和数据索引优化了系统的性能

## 技术实现

### 1. 数据库设计
采用 MySQL 数据库，设计了多个核心实体，如用户、歌曲、歌手、专辑、歌单、评论等，并通过外键确保数据的一致性。数据库表结构优化了查询效率和存储空间，选择了 InnoDB 存储引擎以支持事务和外键约束，适应复杂查询和高并发负载。

### 2. 后端开发
后端使用 Express 框架，快速构建 RESTful 风格的 API，前后端分离架构使得开发更为模块化，易于维护。用户身份通过 JWT 进行认证，确保系统的安全性与数据交换的一致性。

### 3. 前端开发
前端使用 Vue 3 框架，结合 Vue Router 和 Pinia 进行路由和状态管理，确保了页面的快速响应与良好的用户体验。使用 Element Plus 组件库快速构建 UI 界面，提供响应式设计，确保在不同设备上流畅运行。

### 4. API 接口与数据交互
项目的 API 设计遵循 RESTful 原则，支持常见的 HTTP 请求方法（GET、POST、PUT、DELETE）。所有需要认证的接口均采用 JWT 验证，保证了数据的安全传输。

## 未来计划与优化

### 近期优化
- 歌单搜索功能的优化
- 评论系统 BUG 修复
- 个人主页 UI 美化

### 未来功能规划(暑假)
- **部署到云服务器**
  - 采用 nginx 进行反向代理和负载均衡
  - 尝试使用 docker 进行容器化和快速部署

- **后端重写**
  - 切换为 Java 或者 Go 语言
  - 加入 Redis 和消息队列

- **音乐播放增强**
  - 修复音乐播放功能
  - 跨组件播放控制
  - 播放队列管理
  - 音频可视化效果
  - 可以加入歌词功能

- **内容管理**
  - 修复搜索和排序的 bug
  - 音乐文件上传和下载
  - 考虑采用云服务器的对象存储服务
  - 封面图片管理
  - 音乐标签系统

- **社交功能**
  - 用户关注系统
  - 歌单分享功能
  - 音乐动态 feed 流

- **运营管理**
  - 管理员后台
  - 数据统计分析
  - 内容审核系统

## 项目亮点
1. 使用最新的 Vue 3 组合式 API，提升代码的可维护性与复用性
2. 完备的用户认证机制和权限控制，确保用户数据安全
3. 丰富的音乐社交功能，增强用户互动体验
4. 灵活的歌单管理系统，支持个性化的音乐管理
5. 优化的用户体验，响应式设计和无感知 Token 认证带来流畅体验

---

# Project Introduction: Online Music Streaming System

This project is a modern online music streaming system that adopts a front-end and back-end separation architecture. The front end uses the Vue 3 framework, while the back end is developed based on Express and MySQL. The system aims to provide users with comprehensive music playback, management, and social interaction features. The design and implementation of the system reference well-known music platforms both domestically and internationally, such as NetEase Cloud Music and Spotify, striving to offer an interactive yet simple and smooth music experience.

# How to Start

The project is currently in development and has not yet been deployed to a server, so it can only be started in development mode.

## Environment Preparation

1. Ensure that [Node.js](https://nodejs.org/) and npm is installed.
2. Run `/create_database.sql` to create the database.
3. Run `/data.sql` to import demo data (including songs by The Beatles and Bob Dylan).

## Configure the Backend

1. Modify the database connection configuration in the `/express-demo1/.env` file. Note that the `password` is stored as an encrypted string, and it is recommended to manually add an admin account.
2. The backend project uses [Express Generator](https://expressjs.com/en/starter/generator.html) to generate the project structure.

### Start the Backend Service

1. Open a terminal and navigate to the `/express-demo1` directory.
2. Enter the following command to download dependencies:
   ```bash
   npm install
   ```
3. Enter the following command to start the backend service:
   ```bash
   npm start
   ```

## Configure the Frontend

The frontend uses [Vite](https://vitejs.dev/) and the Vue 3 ecosystem.

### Start the Frontend Service

1. Open another terminal and navigate to the `/vue-demo1` directory.
2. Enter the following command to download dependencies:
   ```bash
   npm install
   ```
3. Enter the following command to start the frontend service:
   ```bash
   npm run dev
   ```
## Default account

The default user name and password are admin and 111111 respectively

## Resource Location

Some demo music and images are stored in the `/vue-demo1/assets/` directory.

# Important Information

- All backend interfaces are defined in `/express-demo1/routeFile.js`.
- All frontend interfaces are defined in `/vue-demo1/src/api/axiosFile.js`.
- The `/vue-demo1/README.md` explains part of the development process and the AI-generated API documentation.

## Core Features

### 1. Personal Music Library
- Create personalized playlists
- Favorite songs
- Support for public/private playlist settings
- Flexible playlist management (add, delete, modify, query)

### 2. Intelligent Music Discovery
- The system provides a public music library, allowing users to search by song name, artist, album, etc.
- Support sorting by popularity, time, etc.
- Browse album detail pages

### 3. Artist Section
- Display artist detail pages
- Artist discography
- Related album display

### 4. Comment System and User Interaction
- Users can comment on songs and interact with other users (like and reply)
- Support for paginated comment loading
- User information management and personal homepage display

### 5. Music Playback and Favorite Features
- Provide stable online music playback functionality
- Support playlist management
- Users can favorite songs

## Technical Architecture and Highlights

### 1. Frontend and Backend Separation Architecture
- **Frontend**: Vue 3 + Element Plus
- **Backend**: Express + MySQL
- **Authentication Mechanism**: JWT

### 2. Responsive Design and Invisible Token Authentication
- The system adopts a responsive design to ensure smooth operation on different devices
- The token authentication mechanism ensures the security of each user operation invisibly

### 3. Data Security and Performance Optimization
- Passwords are stored encrypted, interface permission control, and data validation mechanisms ensure user data security
- Pagination, caching, and data indexing optimize system performance

## Technical Implementation

### 1. Database Design
Using MySQL, multiple core entities such as users, songs, artists, albums, playlists, and comments are designed, ensuring data consistency through foreign keys. The database table structure optimizes query efficiency and storage space, choosing the InnoDB storage engine to support transactions and foreign key constraints, adapting to complex queries and high concurrency loads.

### 2. Backend Development
The backend uses the Express framework to quickly build RESTful-style APIs, and the front-end and back-end separation architecture makes development more modular and easier to maintain. User identity is authenticated through JWT, ensuring system security and data exchange consistency.

### 3. Frontend Development
The frontend uses the Vue 3 framework, combined with Vue Router and Pinia for routing and state management, ensuring fast page response and a good user experience. The Element Plus component library is used to quickly build the UI interface, providing a responsive design to ensure smooth operation on different devices.

### 4. API Interfaces and Data Interaction
The project's API design follows RESTful principles, supporting common HTTP request methods (GET, POST, PUT, DELETE). All interfaces requiring authentication use JWT verification to ensure secure data transmission.

## Future Plans and Optimizations

### Recent Optimizations
- Optimization of playlist search functionality
- Bug fixes in the comment system
- UI beautification of personal homepage

### Future Feature Planning (Summer)
- **Deployment to Cloud Server**
  - Use Nginx for reverse proxy and load balancing
  - Try using Docker for containerization and rapid deployment

- **Backend Rewrite**
  - Switch to Java or Go language
  - Add Redis and message queues

- **Enhanced Music Playback**
  - Fix music playback functionality
  - Cross-component playback control
  - Playlist management
  - Audio visualization effects
  - Consider adding lyrics functionality

- **Content Management**
  - Fix bugs in search and sorting
  - Music file upload and download
  - Consider using cloud server object storage services
  - Cover image management
  - Music tagging system

- **Social Features**
  - User follow system
  - Playlist sharing functionality
  - Music dynamic feed

- **Operational Management**
  - Admin backend
  - Data statistical analysis
  - Content review system

## Project Highlights
1. Utilizes the latest Vue 3 Composition API, enhancing code maintainability and reusability
2. Comprehensive user authentication mechanisms and permission controls ensure user data security
3. Rich music social features enhance user interaction experience
4. Flexible playlist management system supports personalized music management
5. Optimized user experience with responsive design and invisible token authentication for a smooth experience