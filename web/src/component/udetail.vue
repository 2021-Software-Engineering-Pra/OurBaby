<template>
  <div>
    <h1>{{info[0].uname}}</h1>
    <h1>余额：{{info[0].yue}}</h1>
    <Form>
            <FormItem label="邮箱">
             <Input  v-model="info[0].email" placeholder="请输入邮箱,最多支持 20 位"/><span>在上面输入你想要更改的新邮箱，然后</span>
            <Button  size='small' @click="upemail">点这里更改</Button>
           
            </FormItem>
            <FormItem label="手机">
            <Input v-model="info[0].phone" placeholder="请输入手机号"/><span>在上面输入你想要更改的新手机号，然后</span>
            <Button size='small' @click="upphone">点这里更改</Button>
            </FormItem>
            <Button size='large' type="primary" @click="change">更改登录密码</Button>
            <Button size='large' type="error" @click="changepay">更改支付密码</Button>
        </Form>
        <Modal
        v-model="modal"
        :title="modaltitle"
        @on-visible-change='modalchange'
        >
        <Form>
            <span>{{attention}}</span>
            <FormItem label="旧密码">
            <Input v-model="opass" type="password" placeholder="请输入原密码"></Input>
            </FormItem>
            <FormItem label="新密码">
            <Input v-model="npass" type="password" :placeholder="payhold"></Input>
            </FormItem>
        </Form>
        <div slot="footer">
            <Button long type="primary" @click="queren" :disabled='ischecked'>确认更改</Button>
        </div>
        </Modal>
  </div>
</template>
<script>
import {userInfo,updatePasswd,updateEmail,updatePhone,updatePaypass} from '../api/index'
export default {
  data() {
    return {
      chosenid: null,
      info: [
      ],
      opass:'',
      npass:'',
      modal:false,
      payhold:'',
      modaltitle:"更改登录密码",
      attention:"",
      changemode:0

    }
  },
  methods:{
      change(){
          this.modaltitle="更改登录密码"
          this.changemode=0
          this.attention=""
          this.payhold="请输入 7-18 位数字或字符"
          this.modal=true
      },
      modalchange(){
          if(this.modal){
              this.npass=""
              this.opass=""
          }
      },
      cemail(){
             if (this.info[0].email.match(/^\w+@\w+\.\w+/)==null){
                
                this.$Message.error("邮箱格式不对")
                return false
            }
            else if(this.info[0].email.length>20){
                this.$Message.error("邮箱过长，请更换邮箱")
                return false
            }
            else{
                return true
            }
        },
        cphone(){
            
             if (this.info[0].phone.length!=11){
                this.$Message.error("手机号长度不对")
                return false
            }
            else if(this.info[0].phone.match('^1[0-9]+')==null){
                this.$Message.error("手机号格式不对")
                return false
            }
            else{
                return true
            }
        },
      queren(){
          if(this.changemode==1){
              updatePaypass({opaypass:this.opass,npaypass:this.npass}).then(res=>{
                   if(res.status==1){
                    this.$Message['success']({
                            background:true,
                            content:"更改成功"
                        })
                        this.opass=""
                        this.npass=""
                    }else{
                        this.$Message['error']({
                                    background:true,
                                    content:res.msg
                                })
                                this.opass=""
                                this.npass=""
              }
              
              }).catch(e=>{
                  this.$Message.error("服务器不给力，等会再试吧")
              })
              this.modal=false
              this.ischecked=true
              return
          }
          if(this.opass==""){
              return
          }
          updatePasswd({opasswd:this.opass,npasswd:this.npass}).then(res=>{
              if(res.status==1){
                  this.$Message['success']({
                            background:true,
                            content:"更改成功"
                        })
                        this.opass=""
                        this.npass=""
              }else{
                  this.$Message['error']({
                            background:true,
                            content:res.msg
                        })
                        this.opass=""
                        this.npass=""
              }
          }).catch(e=>{
                  this.$Message.error("服务器不给力，等会再试吧")
              })
              this.modal=false
              this.ischecked=true
      },
      getuser(){
          userInfo().then(res=>{
              if(res.status==1){
                  this.info=res.data
              }else{
                  this.$Message['error']({
                            background:true,
                            content:res.msg
                        })
              }
          })
      },
      upemail(){
          if(!this.cemail()){
              return
          }
          updateEmail({email:this.info[0].email}).then(res=>{
              if(res.status==1){
                  this.$Message['success']({
                            background:true,
                            content:"更改成功"
                        })
              }else{
                  this.$Message['error']({
                            background:true,
                            content:res.msg
                        })
                        
              }
          })
      },
      upphone(){
          if(!this.cphone()){
              return
          }
          updatePhone({phone:this.info[0].phone}).then(res=>{
              if(res.status==1){
                  this.$Message['success']({
                            background:true,
                            content:"更改成功"
                        })
              }else{
                  this.$Message['error']({
                            background:true,
                            content:res.msg
                        })
              }
          })
      },
      changepay(){
          this.modaltitle="更改支付密码"
           this.payhold="请输入 6 位数字或字符"
            this.attention="如果你从未设置过支付密码，旧密码选项请留空白"
            this.modal=true
            this.changemode=1
      }
  },
  computed:{
      ischecked(){
          if(this.changemode){
              if(this.npass.length==6){
                  return false
              }else{
                  return true
              }
          }else{
              if(this.npass.length>6&&this.npass.length<19){
                 return false
              }else{
                  return true
              }
          }
      }
  },
  mounted(){
      this.getuser()
  }
};
</script>