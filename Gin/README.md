 # 联想商城后端部分
 前端采用 [vue](https://cn.vuejs.org/) ，使用 [webpack](https://www.webpackjs.com/) 打包

 后端用 Go 编写，使用了 [Gin 框架](https://github.com/gin-gonic/gin)
## 说明
  go mod download 下载依赖包
## 注意
测试前请先配置数据库相关
   
   1. 请确保本机已经安装 Mysql 并运行在默认的 3306 端口
   2. 请新建用户 test 密码同为 test 。
   3. 然后查看并执行本目录下的 database.sql 语句，请注意检查 sql 内的内容，以免因为重名导致数据损失，并赋予 test用户操作数据库 shop 的权限