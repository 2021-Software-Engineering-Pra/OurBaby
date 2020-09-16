<template>
  <div>
    <h1>您的订单</h1>
    <h2>{{errmsg}}</h2>
    <div v-for="(item,index) in orders" :key="index">
        <div style="padding-top:15px;color:black">
            <Card :to="{path:'/orderdetail',query:{orid:item.oid}}" :title="'订单编号:'+item.oid">
            <Row>
            <Col span='5'>
            <img :src="'../src/static/'+item.linkid+'/main.jpg'" style="max-width:100px" alt srcset />
            </Col>
            <Col span="19">
            <h2>{{item.pname}}</h2>
            <p>单价：{{item.aprice}}</p>
            <p>数量：{{item.count}}</p>
            <P>总计：{{total[index]}}</P>
            <p>状态：<span style="color:red">{{item.status}}</span></p>
            <p>创建时间：{{item.creatime}}</p>
            <p v-if="isNaN(total[index])" style="color:red">该笔订单包含多个商品，请点击查看</p>
            </Col>
            </Row>
        
        <p></p>
      </Card>
        </div>
     
    </div>
  </div>
</template>
<script>
import {getOrder} from '../api/goods'
export default {
  data() {
    return {
      orders: [],
        total:[],
        errmsg:""
      
    }
  },
  methods:{
      getord(){
          getOrder().then(res=>{
            console.log(res)
              if(res.status==1){
                if(res.data[0]==null){
                  this.errmsg="还没有订单哦~"
                  return
                }
                  let t=res.data[0]
                  this.orders=t.reverse()
                   let w=[]
                for(let i=0;i<this.orders.length;i++){
                  let n=this.orders[i].aprice*this.orders[i].count
                  w.push(n)
                }
                this.total=w
                console.log(this.total)
              }else{
                  this.$Message.error(res.msg)
              }
          }).catch(e=>{
              this.$Message.error("服务器不给力，等会再试吧")
          })
      },
      
  },
  mounted() {
      this.getord()
      
  },
  computed:{
   
  }
};
</script>
<style lang="css" scoped>
    p{
        color: rgb(77,90,108)
    }
   
</style>