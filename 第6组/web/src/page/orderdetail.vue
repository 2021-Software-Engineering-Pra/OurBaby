<template>
  <div>
    <br>
    <div v-if="currentstep==3" style="max-width:80%;margin:auto">
      <Steps current="1" status="error">
        <Step title="未支付" content="您已拍下商品，但还未进行支付，请尽快完成支付"></Step>
        <Step title="已取消" content="您已取消本次购物，期待下次为您服务"></Step>
    </Steps>
    </div>
     <div v-else style="max-width:80%;margin:auto">
       <Steps :current="currentstep">
        <Step title="未支付" content="您已拍下商品，但还未进行支付，请尽快完成支付"></Step>
        <Step title="已支付" content="我们已经收到您的货款，正在为您发货"></Step>
        <Step title="已完成" content="您已完成本次购物，期待您的下次光临"></Step>
    </Steps>
     </div>
    <br>
    <Card title="订单详情">
      <div style="padding-left:30px">
        <h2>收货地址</h2>
      <div style="max-width:500px;">
        <Card >
        <h1>{{order.recvname}}</h1>
        <hr/>
        <h2>{{order.phone}}</h2>
        <p>{{order.detailed}}</p>
      </Card>
      </div>
      <br>
      <h2>订单信息</h2>
        <div>
          <div v-for="(i,index) in products" :key='index' style="padding-top:10px">
             <Card >
          <div style="padding-left:20px;display:inline-block">
            <div style="float:left">
            <img :src="'/src/static/'+i.linkid+'/main.jpg'" style="max-width:100px" alt srcset />
          </div>
          <div style="padding-left:120px">
            <h1>{{i.pname}}</h1>
            <p>{{i.peizhi}}</p>
            <p>单价：{{i.price}} 元</p>
            <p>数量：{{count[index]}}</p>
          </div>
        </div>
        </Card>
          </div>
         
        <div style="padding-left:60%;padding-top:10px">
          <p>订单金额：{{total}} 元</p>
          <p>订单状态：{{order.status}}</p>
          <p>创建时间：{{order.creatime}}</p>
          <div v-if='notpaid'>
          <Button type="primary" size='large' @click="topay">立即支付</Button>
          <Button type="error" size='large' @click="cancel">取消订单</Button>
          </div>
          <div v-else-if="order.status=='已支付'">
            
            <Button size='large' @click="receive" type="primary">确认收货</Button>
            <Button size='large' @click="cancel">取消订单</Button>
          </div>
          <div v-else>
             <Button type="error" size='large' @click="del">删除订单</Button>
          </div>
        </div>
      
          <div>
          <Button @click="goback">返回上一页</Button>
        </div>
        </div>
      </div>
    </Card>
    <!-- <pay v-bind:show="showpay" v-bind:oid="orid"></pay> -->
     <div>
        <Modal
        v-model="showpay"
        :title="modaltitle"
        
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
import { getOrder,delOrder,toPay,cancelOrder,Products,receiveProduct } from "../api/goods";
import {NothavePaypass,updatePaypass} from '../api/index'
// import pay from '../component/payorder.vue'
export default {
  data() {
    return {
      orid: this.$route.query.orid,
      showpay:false,
      payload:false,
      havapaypass:false,
      modaltitle:'',
      oktext:"",
      passwd:'',
      order:{},
      products:[],
      count:[]
    };
  },
  methods: {
    initorder() {
      getOrder({ orid: this.orid })
        .then(res => {
          if (res.status == 1) {
            this.order = res.data[0];
            let c=res.data[0].count.split(" ")
            let a=[]
            for(let i=0;i<c.length;i++){
              a.push(Number(c[i]))
            }
            this.count=a
            let pid=res.data[0].pid.split(" ")
            Products({pid:pid}).then(res=>{
              console.log(res.data[0])
              if(res.status==1){
                this.products=res.data[0]
              }else{
                this.$Message.info(res.msg)
              }
            }).catch(e=>{
              this.$Message.error("服务器不给力，等会再试吧")
            })
          } else {
            this.$Message.error(res.msg);
          }
        })
        .catch(e => {
          this.$Message.error("服务器不给力，等会再试吧");
        });
    },
    goback(){
      this.$router.go(-1)
    },
    topay(){
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
           this.order.status="已支付"
           this.showpay=false
           this.payload=false
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
        this.$Message.error("服务器不给力，等会再试吧")
           this.payload=false
      })
      }else{
        updatePaypass({npaypass:this.passwd}).then(res=>{
          console.log(res)
          if(res.status==1){
            this.payload=false
            this.$Message.success("设置成功")
            this.showpay=false
            this.passwd=""
            this.topay()
          }else{
            this.$Message.error(res.msg)
          }
        }).catch(e=>{
          this.$Message.error("服务器不给力，等会再试吧")
        })
      }
      this.payload=false
    },
    del(){
      delOrder({orid:this.orid}).then(res=>{
        if(res.status==1){
          this.$Message.success(res.msg)
          this.$router.go(-1)
        }else{
          this.$Message.info(res.msg)
        }
      }).catch(e=>{
        this.$Message.error("服务器不给力，等会再试吧")
      })
    },
    cancel(){
      console.log("cancel")
      cancelOrder({orid:this.orid}).then(res=>{
        if(res.status==1){
          this.$Message.success(res.msg)
          this.order.status="已取消"
        }else{
          this.$Message.error(res.msg)
        }
      }).catch(e=>{
        this.$Message.error("服务器不给力，等会再试吧")
      })

    },
    receive(){
      receiveProduct({orid:this.orid}).then(res=>{
        if(res.status==1){
          this.$Message.success(res.msg)
          this.order.status="已完成"
        }else{
          this.$Message.error(res.msg)
        }
      }).catch(e=>{
        this.$Message.error("服务器不给力，等会再试吧")
      })
    }
  },
  mounted() {
     this.initorder();
     
     },
  computed:{
    total(){
      let t=0
      for(let i=0;i<this.products.length;i++){
        t+=this.products[i].price*this.count[i]
      }
      return t
    },
    notpaid(){
      if(this.order.status=="未支付"){
        return true
      }
      return false
    },
    checkok(){
      if(this.passwd.length==6){
        return false
      }else{
        return true
      }
    },
    currentstep(){
      switch (this.order.status){
        case "未支付": return 0;break;
        case "已支付" : return 1;break;
        case "已完成" : return 2; break
        case "已取消": return 3;break;
      }
    }
  },
};
</script>