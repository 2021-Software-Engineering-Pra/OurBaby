<template>
  <div>
      <Card title="确认订单">
        <div style="padding-left:30px">
            <h2>收货地址</h2>
     <div style="display:inline-block">
        <div style="width:300px;padding-left:15px;float:left;padding-top:10px" v-for='(i,index) in address' :key='index'>
          <a href="javascript:;" style="color:black" @click="changechosenid(index)">
          <Card :class="{'chosened':index==chosenaid}" style="height:150px">
            <h1>{{i.recvname}}</h1>
            <hr/>
            <h2>{{i.phone}}</h2>
            <p>{{i.detailed}}</p>
          </Card>
        </a>
      </div>
      <div style="float:left;width:170px;padding-left:15px;padding-top:10px">
        <a href="javascript:;" style="color:black" @click="addressmodal=true">
          <Card style="height:150px">
            <p>添加收货地址</p>
            <Icon style="margin:auto" type="ios-add-circle-outline" color="rgb(77,90,108)" size=100 />
          </Card>
        </a>
      </div>
     </div>

      <br>
      <h2>订单信息</h2>
        <div v-for='(i,index) in product' :key='index' style="padding-top:10px">
          <Card>
          <div style="padding-left:20px;display:inline-block">
            <div style="float:left">
            <img :src="'/src/static/'+i.linkid+'/main.jpg'" style="max-width:100px" alt srcset />
          </div>
          <div style="padding-left:120px">
            <h1>{{i.pname}}</h1>
            <p>{{i.peizhi}}</p>
            <p>单价：{{i.price}} 元</p>
            <p>数量：{{countnum[index]}}</p>
          </div>
        </div>
        </Card>      
       
        </div>
        <div style="padding-left:60%;padding-top:10px">
          <p>订单金额：{{total}} 元</p>
          <div >
          <Button type="primary" size='large' @click="submitorder">提交订单</Button>
          </div>
        </div>
          <div>
          <Button @click="goback">返回上一页</Button>
        </div>
        </div>
    </Card>
    <!-- 添加收货地址面板 -->
     <Modal
        v-model="addressmodal"
        title="添加收获地址"
        @on-ok="oktoadd"
        >
        <Form>
            <FormItem label="收货人">
            <Input v-model="recv" placeholder="请输入收货人"></Input>
            </FormItem>
            <FormItem label="手机">
            <Input v-model="phone" type="tel" placeholder="请输入手机号"></Input>
            </FormItem>
            <FormItem label="详细地址">
            <Input v-model="detail" type="textarea" :autosize="{minRows: 2,maxRows: 5}" placeholder="请输入详细地址"/>
        </FormItem>
        </Form>
    </Modal>
    <div>
        <Modal
        draggable
        v-model="showpay"
        :title="modaltitle"
        @on-visible-change='modalstatus'
        >
        <div>
            <Input v-model="passwd" type='password' @on-enter='oktopay' ref="foc" placeholder="请输入6位数字或字符"></Input>
        </div>
            
        <div slot="footer">
            <Button type="primary" @click="oktopay" long :loading='payload' :disabled="checkok">{{oktext}}</Button>
        </div>
    </Modal>
    </div>
  </div>
</template>
<script>
import {Products, addOrder,toPay} from '../api/goods'
import {getAddress,addAddress,NothavePaypass,updatePaypass} from '../api/index'
export default {
  data() {
    return {
      pid: this.$route.query.pid,
      countnum: this.$route.query.count,
      chosenaid: 0,
      addressmodal:false,
      showpay:false,
      payload:false,
      havapaypass:false,
      oktext:"",
      orid:'',
      modaltitle:'',
      recv:'',
      phone:'',
      detail:'',
      product: [
        
      ],
      address: [
       
      ],
      orid:"",
      passwd:''
    };
  },
  methods:{
      initproduct(){
     if(typeof(this.pid)=="string"){
          let a=[],b=[]
          a.push(this.pid)
          this.pid=a
          b.push(Number(this.countnum))
          this.countnum=b
        }
        Products({pid:this.pid}).then(res=>{
            if(res.status==1){
              if(res.data[0]==null){
                this.$Message.error("请重新下单")
                this.$router.go(-1)
              }
                this.product=res.data[0]
                
            }else{
                this.$Message.info(res.msg)
            }
        }).catch(e=>{
            this.$Message.error("请重新下单")
            this.$router.go(-1)
        })
    },
    initAddress(){
      getAddress().then(res=>{
        if(res.status=1){
          this.address=res.data
        
        }else{
          this.$Message.info(res.msg)
        }
        
      }).catch(e=>{
        this.$Message.error("服务器不给力，等会再试吧")
      })
    },
    submitorder(){
      
      if(this.address==null){
        this.addressmodal=true
        return
      }
      console.log(this.address[this.chosenaid].aid,this.pid,this.countnum)
  
      addOrder({pid:this.pid,count:this.countnum,aid:this.address[this.chosenaid].aid}).then(res=>{
          if(res.status==1){
            this.orid=res.msg
           NothavePaypass().then(res=>{
              console.log(res)
              if(res.status==1){
                this.havapaypass=false
                this.$Message.info(res.msg)
                this.modaltitle="请设置支付密码"
                this.oktext="确认设置"
              }else{
                this.havapaypass=true
                this.modaltitle="请输入支付密码"
                this.oktext="立即支付"
              }
              this.showpay=true
              //自动聚焦
                this.$nextTick(()=>{
                  this.$refs.foc.focus()
                })
            }).catch(e=>{
              this.$Message.error("服务器不给力，等会再试吧")
            })
          }else{
            this.$Message.error(res.msg)
          }
         
      }).catch(e=>{
        this.$Message.error("服务器不给力，等会再试吧")
      })
       
    },
    changechosenid(index){
      this.chosenaid=index
    },
    goback(){
      this.$router.go(-1)
    },
    goorder(){
      this.$router.push({path:'/orderdetail',query:{orid:this.orid}})
    },
    cphone(){
            if(this.recv==""||this.phone==""||this.detail==""){
              return false
            }
             if (this.phone.length!=11){
                this.$Message.error("手机号长度不对")
                return false
            }
            if(this.phone.match('^1[0-9]+$')==null){
                this.$Message.error("手机号格式不对")
                return false
            }
            else{
                return true
            }
        },
    oktoadd(){
      if(!this.cphone()){
        return
      }
      addAddress({recvname:this.recv,phone:this.phone,detailed:this.detail}).then(res=>{
                    console.log(res)
                    if(res.status==1){
                        this.$Message['success']({
                            background:true,
                            content:"添加成功"
                        })
                        this.initAddress()
                    }else{
                        this.$Message['error']({
                            background:true,
                            content:res.msg
                        })
                    }
                    
                })
    },
    oktopay(){
      if(this.passwd.length!=6){
        return
      }
      this.payload=true
      if(this.havapaypass){
        toPay({orid:this.orid,paypass:this.passwd}).then(res=>{
         if(res.status==1){
           this.$Message.success("支付成功")
           this.$router.push({path:'/orderdetail',query:{orid:this.orid}})
         }else{
           this.$Message.error(res.msg)
          //  this.showpay=false
          this.passwd=""
          this.$nextTick(()=>{
          this.$refs.foc.focus()
          })
           this.payload=false
         }
      }).catch(e=>{
        console.log(e)
        this.$Message.error("服务器不给力，等会再试吧")
           this.payload=false
      })
      }else{
        updatePaypass({npaypass:this.passwd}).then(res=>{
          console.log(res)
          if(res.status==1){
            this.payload=false
            this.$Message.success("设置成功")
            this.passwd=""
            this.havapaypass=true
            this.showpay=false
          }else{
            this.$Message.error(res.msg)
          }
        }).catch(e=>{
          this.$Message.error("服务器不给力，等会再试吧")
        })
      }
      this.payload=false
    },
    modalstatus(){
      if(this.showpay==false){
        this.goorder()
      }
    }
  },
  mounted(){
    this.initproduct()
    this.initAddress()
    
  },
  computed:{
    total(){
      let tem=0
      this.product.forEach((item,index)=>{
        tem+=item.price*this.countnum[index]
      })
      return tem
    },
    checkok(){
      if(this.passwd.length==6){
        return false
      }else{
        return true
      }
    }
  }
};
</script>
<style lang="css" scoped>
  .chosened{
    border: 1px red solid;
  }
</style>