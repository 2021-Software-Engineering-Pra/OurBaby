
import vue from 'vue';
import app from "./app.vue"
import ViewUI from 'view-design' //引入view UI
import 'view-design/dist/styles/iview.css'
import router from './router/router'
vue.use(ViewUI)
router.beforeEach((to, from, next) => {
    /* 路由发生变化修改页面title */
    if (to.meta.title) {
      document.title = to.meta.title;
    }
    next();
  })
new vue({
    el:'#app',
    router,//router:router
    render:h=>h(app)
})