<template>
  <div>
    <div style="padding-top:40px;min-width:1000px;display:inline-block">
      <div >
        <div style="float:left;padding-left:20px">
          <Card>
              <img
            :src="'/src/static/'+datas[chosenid].linkid+'/main.jpg'"
            style="max-width:500px"
            alt
            srcset
          />
          </Card>
        </div>
      </div>
      
      <div style="padding-left:600px">
        <h1 style="font-size:30px">{{datas[chosenid].pname}}</h1>
        <h3>{{datas[chosenid].pdesp}}</h3>
        <div style="padding-top:10px;max-width:200px">
          <Card :dis-hover='true' title="价格">
            <p style="color:red">￥ {{datas[chosenid].price}}</p>
          </Card>
        </div>
        <p style="padding-top:10px">选择配置</p>
        <div style="padding-top:10px; max-width:500px" v-for="(i,index) in datas" :key="index">
            <div>
                <a href="javascript:;" style="color:black" @click="changechosenid(index)">
                <Card :class="{'chosened':index==chosenid}">{{i.peizhi}}</Card>
                 </a>
            </div>
          
        </div>
        <div style="padding-top:10px">
          <label>购买数量</label>
          
          <InputNumber :max="100" :min="1" v-model="countnum" ></InputNumber>         
          </Input>
          
        </div>
        <div style="padding-top:20px">
          <Button size="large" type="primary" :disabled='isbuy' @click='buynow'>立即购买</Button>
          <Button size="large" :disabled='isbuy' @click="cart">加入购物车</Button>
        </div>
      </div>

    </div>
    <div style="padding-top:20px;max-width:1530px;margin:auto">
        <Card title="商品详情">
          <img :src="'/src/static/'+datas[chosenid].linkid+'/intro.jpg'" style="max-width:1480px" alt srcset />
        </Card>
      </div>
  </div>
</template>
<script>
import {getProduct,addCart} from '../api/goods'
export default {
  data() {
    return {
      id: this.$route.params.id,
      type: this.$route.query.type,
      status: 0,
      msg: "",
      chosenid: 0,
      countnum:1,
      datas: [
        
      ]
    };
  },
  methods: {
    changechosenid(index) {
      this.chosenid = index;
    },
   
    init(){
        getProduct({type:this.type}).then(res=>{
            if(res.status==1){
                this.datas=res.data
            }else{
                console.log(res.msg)
            }
        })
    },
    buynow(){
        let uname = document.cookie.split("=");
            if(uname.length>1){
                let tem1=[],tem2=[]
                tem1.push(this.datas[this.chosenid].pid)
                tem2.push(this.countnum)
                this.$router.push({path:'/order',query:{pid:tem1,count:tem2}})
            }else{
                this.$Message.info("请先登录")
            }
    },
    cart(){
      addCart({pid:this.datas[this.chosenid].pid,count:this.countnum}).then(res=>{
          if(res.status==1){
            this.$Message.success(res.msg)
          }else{
            this.$Message.error(res.msg)
          }
      }).catch(e=>{
        this.$Message.error("服务器不给力，等会再试吧")
      })
    }
  },
    computed:{
        isbuy(){
            if(this.countnum>0){
                return false
            }else{
                return true
            }
        }
    },
    mounted(){
        this.init()
    }
};
</script>
<style lang="css" scoped>
.chosened{
    border:1px red solid
}
</style>