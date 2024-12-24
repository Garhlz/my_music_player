# 说明
这是我的数据库设计期末大作业，是一个音乐播放网站，采用了vue+express前后端分离开发 
# 如何启动
目前处于开发过程中，尚未部署到服务器，只能用开发模式启动 

确保安装了nodejs 

通过/create_database.sql创建数据库，在/express-demo1/.env中修改数据库连接，数据库的数据可以自己（找ai）生成，注意password存储的是密文，建议手动添加一个管理员账号 

后端使用了express-generator生成项目结构 
切换到/express-demo1目录，终端输入npm install下载依赖 

输入npm start启动后端服务 

前端用的是vite和vue3全家桶，切换到vue-demo1目录，终端输入npm install下载依赖 

终端输入npm run dev启动前端服务 

部分演示音乐和图片已经存放在/vue-demo1/assets/中

# 重要信息
后端接口全部定义在/express-demo1/routeFile.js中，前端则全部定义在/vue-demo1/src/api/axiosFile.js中 

这可能不是最佳实践，但是经过权衡，对我目前来说最为方便 

在/vue-demo1/README.md中说明了部分开发过程以及ai生成的接口文档
