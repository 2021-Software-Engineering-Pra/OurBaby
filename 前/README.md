 # 联想商城前端部分
 前端采用 [vue](https://cn.vuejs.org/) ，使用了一个现代化开源UI库-[View UI](https://www.iviewui.com/docs/introduce),最后使用 [webpack](https://www.webpackjs.com/) 打包，

 后端用 Go 编写，使用了 [Gin 框架](https://github.com/gin-gonic/gin)

 感谢以上开源框架的作者
## 说明
* package.json(依赖包以及运行脚本)
* webpack.config.js(webpack打包信息)
* index.html(用于dev-server预览)
* src<br>
  |-api(接口信息)<br>
  | &emsp;|-public.js(封装的请求方式，不用更改)<br>
  | &emsp;|-index.js&goods.js(请求详情)<br>
  |-component(组件，在多个页面中使用)<br>
  |-css(样式文件夹)<br>
  |-img(图片文件夹)<br>
  |-js(javascript文件夹)<br>
  |-page(页面，.vue文件)<br>
  |-router(路由，页面需在此处注册)<br>
  |-static(商品图片文件夹，以pid命名)<br>
  |-app.vue(vue渲染页面)<br>
  |-index.js(创建vue实体)<br>
  |index.html(主页面显示)
 ## 使用方法
 1. 安装node
 2. npm install
 3. npm run dev
 4. 即可在本地 8080 端口预览

