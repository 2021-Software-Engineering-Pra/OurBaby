<template>
  <div class="topbar">
    <Row type="flex" justify="start">
      <Col span="13">
        <nav class="main-nav float-left">
          <ul>
            <li v-for="(item,index) in topbar">
              <router-link :to="link1[index]">{{item}}</router-link>
            </li>
          </ul>
        </nav>
      </Col>
      <Col span="6" style="padding-top:24px;">
        <Input
          search
          enter-button
          v-model="searchkey"
          @on-search="search"
          placeholder="Enter something to search..."
        />
      </Col>
      <Col span="4">
        <nav class="main-nav float-right">
          <ul style="fixed">
            <li>
              <a v-if='login==false' href="javascript:;" @click="showlogin=true">{{navbar[0]}}</a>
              <Dropdown  v-else>
                <router-link style="padding-top:0" :to="link[0]">{{navbar[0]}}</router-link>
                <DropdownMenu v-if="login" slot="list">
                  <DropdownItem ><a href="javascript:logout;" @click="logout">{{droplist}}</a></DropdownItem>
                </DropdownMenu>
              </Dropdown>
              
            </li>
            <li>
              <!-- <a v-if='login==false' href="#login" @click="showregister=true">{{navbar[1]}}</a> -->
              <router-link  :to="link[1]">{{navbar[1]}}</router-link>
            </li>
          </ul>
        </nav>
      </Col>
    </Row>
    <Modal title="欢迎登陆" v-model="showlogin" >
      <Form>
            <FormItem label="用户名">
            <Input v-model="loginuname" placeholder="请输入用户名"  @on-focus="reseterror"></Input>
            <span style="color:red">{{unameerror}}</span>
            </FormItem>
            <FormItem label="密码">
            <Input v-model="loginpasswd" type="password" @on-focus="reseterrorp" @on-enter='getlogin' placeholder="请输入密码"></Input>
            <span style="color:red" >{{upasswderror}}</span>
            </FormItem>
        </Form>
        <div slot="footer">
              <Button size='large' long type="primary" :loading='loginloading' @click='getlogin'>登录</Button>
            </div>
      </Modal> 
  </div>
</template>
<script>
import { loginOut,userLogin } from "../api/index";
export default {
  data() {
    return {
      searchkey: "",
      navbar: ["登录", "注册"],
      link: ["/#login", "/register"],
      topbar: ["首页", "轻薄本", "游戏本", "Thinkpad"],
      link1: ["/", "/kind/light", "/kind/game", "/kind/Thinkpad"],
      login: false,
      droplist: "",
      showlogin:false,
      showregister:false,
      loginpasswd:'',
      loginuname:'',
      loginloading:false,
      unameerror:"",
      upasswderror:""
    };
  },
  methods: {
    search() {
      // this.$router.push('/search?key='+this.searchkey)
      if(this.searchkey==""){
        this.$Message.info("至少输点什么再搜吧~")
        return
      }
      this.$router.push({ path: "/search", query: { key: this.searchkey } });
      // this.$router.push({name:'/search',params:{key:this.searchkey}})
    },
    logout() {
      loginOut().then(res => {
        if (res.status == 1) {
           this.$Message["success"]({
            background: true,
            content: "注销成功"
          });
          this.login=false
          this.$router.push("/")
          this.navbar=["登录", "注册"]
          this.link=["/#login", "/register"]
        } else {
          this.$Message["error"]({
            background: true,
            content: res.msg
          });
        }
      });
    },
    getlogin(){
        console.log(this.loginuname,this.loginpasswd)
        if (this.loginuname==""){
          this.unameerror="请输入用户名"
          return
        }
        if (this.loginpasswd==""){
          this.upasswderror="请输入密码"
          return
        }
        this.loginloading=true
        userLogin({uname:this.loginuname,upasswd:this.loginpasswd}).then(
          res=>{
           if(res.status==1){
                        this.$Message['success']({
                            background:true,
                            content:"登录成功"
                        })
                        this.showlogin=false
                        this.$router.go(0)
                    }else{
                        this.$Message['error']({
                            background:true,
                            content:res.msg
                        })
                        this.loginloading=false
                    }
          }
        ).catch(e=>{
          this.$Message['error']({
                            background:true,
                            content:"服务器不给力，等会再试吧"
                        })
          this.loginloading=false
        })
    },
    reseterror(){
      this.unameerror=""
    },
    reseterrorp(){
      this.upasswderror=""
    }
  },
  created() {
    let uname = document.cookie.split("=");
    if (uname.length > 1) {
      this.login = true;
      this.navbar = [uname[1], "购物车"];
      this.link = ["/user/" + uname[1], "/cart"];
      this.droplist = "注销";
    }
  }
};
</script>
<style lang="css">
body{
  background: #fff;
  color: #444;
  /* font-family: "Open Sans", sans-serif; */
  font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;
}

a {
  color: #636399;
  transition: 0.5s;
}

a:hover,
a:active
a:focus {
  color: #0a98c0;
  outline: none;
  text-decoration: none;
}

p {
  font-size: 18px;
  padding: 0;
  margin: 0 0 30px 0;
}

h1,
h2,
h3,
h4,
h5,
h6 {
  font-family: "Montserrat", sans-serif;
  font-weight: 400;
  margin: 0 0 20px 0;
  padding: 0;
}
.topbar{
  height: 0.8%;
  width: 100%;
}

.float-left{float:left!important}

.float-right{float:right!important}
.main-nav,
.main-nav * {
  margin: auto;
  padding-top: 5px;
  list-style: none;
}

.main-nav > ul > li {
  position: relative;
  white-space: nowrap;
  float: left;
}
.main-nav a{
  display: block;
  position: relative;
  color: #413e66;
  padding: 10px 10px;
  transition: 0.3s;
  font-size: 20px;
  font-family: "Open Sans", sans-serif;
  text-transform: uppercase;
  font-weight: 600;
}

</style>