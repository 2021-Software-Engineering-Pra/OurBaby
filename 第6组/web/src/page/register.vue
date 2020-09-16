<template>
<div id='register' style="max-width:600px;margin:auto;padding-top:25px">
    <Card title="欢迎注册">
        <div>
        <label for="">用户名</label><br>
    <Input v-model="uname" @input="checkuname" @on-blur="ishavaname" placeholder="请输入用户名 5-16 位" style="max-width: 300px" />
    <span id='unameerror' style="color:red">{{unameerror}}{{used}}</span>
    </div>
    <div>
        <label for="">密码</label><br>
    <Input v-model="passwd" @input="cpasswd()" type="password" placeholder="请输入密码 7-18 位" style="max-width: 300px" />
    <span id='pa' style="color:red">{{passwderror}}</span>
    </div>
    <div>
        <label for="">确认密码</label><br>
    <Input v-model="npasswd" @input="cnpasswd()" type="password" placeholder="请再次输入密码" style="max-width: 300px" />
    <span id='npassederror' style="color:red">{{npasswderror}}</span>
    </div>
    <div>
        <label for="">邮箱</label><br>
    <Input v-model="email" @input="cemail" placeholder="请输入邮箱,最多支持 20 位" style="max-width: 300px" />
    <span id='emailerror' style="color:red">{{emailerror}}</span>
    </div>
    <div>
        <label for="">手机号</label><br>
    <Input type="tel" v-model="phone" @input="cphone" placeholder="请输入手机号" style="max-width: 300px" />
    <span id='phoneerror' style="color:red">{{phoneerror}}</span>
    </div>
    <div style='padding-top:20px'>
       <Radio v-model="tongyi"></Radio> <span>已阅读并同意<a href="#" @click="modal=true">用户注册协议</a></span> 
    </div>
     
     <div style='padding-top:20px'>
          <Button type="primary" @click="zhuce" size='large' :loading='loading'>立即注册</Button>
     </div>
   
    <Modal v-model="modal" :title="modaltitle" @on-ok="ok" @on-cancel='cancel'>
        <p>欢迎注册</p>
    </Modal>
    </Card>
    
</div>

</template>
<script>
import http from '../api/public'
import {isUname} from '../api/index'
export default {
    data(){
        return{
            modaltitle:"欢迎注册",
            modal:false,
            uname:'',
            passwd:'',
            npasswd:'',
            email:'',
            phone:'',
            unameerror:'',
            used:'',
            passwderror:'',
            npasswderror:'',
            tongyi:false,
            emailerror:'',
            phoneerror:'',
            loading:false,
            unameok:false,
        }
    },
    methods:{
        checkuname(){
            if (this.uname.length<5||this.uname.length>16){
                this.unameerror="用户名长度不符"
                return false
            }else{
                this.unameerror=""
                return true
            }
        },
        ishavaname(){
            if(this.checkuname()){
                isUname({uname:this.uname}).then(res=>{
                    if(res.status==0){
                        this.used=res.msg
                        this.unameok=false
                    }else{
                        this.used=""
                        this.unameok=true
                    }
                    }
                )
            }
            
        },
        cpasswd(){
            if (this.passwd.length<7||this.passwd.length>18){
                this.passwderror="密码长度不符"
                return false
            }else{
                this.passwderror=""
                return true
            }
        },
        cnpasswd(){
             if (this.npasswd!=this.passwd){
                this.npasswderror="两次密码不同"
                return false
            }else{
                this.npasswderror=''
                return true
            }
        },
        cemail(){
             if (this.email.match(/^\w+@\w+\.\w+/)==null){
                
                this.emailerror="邮箱格式不对"
                return false
            }
            else if(this.email.length>20){
                this.emailerror="邮箱过长，请更换邮箱"
                return false
            }
            else{
                this.emailerror=''
                return true
            }
        },
        cphone(){
            
             if (this.phone.length!=11){
                this.phoneerror="手机号长度不对"
                return false
            }
            else if(this.phone.match('^1[0-9]+$')==null){
                this.phoneerror="手机号格式不对"
                return false
            }
            else{
                this.phoneerror=''
                return true
            }
        },
        ok(){
            this.tongyi=true
        },
        cancel(){
            this.tongyi=false
        },
        zhuce(){
            if(this.checkuname()&&this.unameok&&this.cpasswd()&&this.cnpasswd()&&this.cemail()&&this.cphone()&&this.tongyi){
                this.loading=true
                http.fetchPost('api/zhuce',{'uname':this.uname,'passwd':this.passwd,'email':this.email,'phone':this.phone}).then(
                    res=>{
                        if (res.status=1){
                            this.$Message.success({
                                background:true,
                                content:"注册成功，即将跳转到之前页面"})
                            this.$router.go(-2)
                        }else{
                            alert(res.msg)
                            this.loading=false
                        }
                    }
                ).catch(err=>{
                    this.$Message.warning({
                        background:true,
                        content:"请求超时，请重试"})
                    this.loading=false
                })
                    
                
            }
            
        }

    }
}
</script>
<style>
    #register {
        max-width: 60%;
        margin: auto;
        
    }
   
</style>