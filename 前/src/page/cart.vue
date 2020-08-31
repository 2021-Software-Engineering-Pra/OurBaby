<template>
  <div style="min-width:1400px">
    <Card title="购物车">
        <h1>{{empty}}</h1>
        <CheckboxGroup v-model="chosenid">
      <div v-for="(i,index) in cartlist" :key="index" >
       <div style="display:inline-block" >
           <div style="float:left;padding-top:100px;padding-left:15px">
            <Checkbox size='large' :label='index'></Checkbox>
            </div>
            <div style="padding-top:15px;padding-left:100px;width:1300px;">
                <Card >
            <div style="display:inline-block">
                
            <router-link :to="{path:'/detail/'+i.pid,query:{type:i.type}}">
                <div style="float:left;padding-left:20px">
                <img
                    :src="'/src/static/'+cartlist[index].linkid+'/main.jpg'"
                    style="max-width:150px"
                    alt
                    srcset
                />
                </div>
                </router-link>
                <div style="padding-left:240px">
                    <router-link :to="{path:'/detail/'+i.pid,query:{type:i.type}}"><h1 style="color:rgb(63,63,99)">{{i.pname}}</h1></router-link>
                    <p>{{i.peizhi}}</p>
                    
                    <p style="color:red;padding-top:10px">单价：{{i.aprice}} 元</p>
                    <div style="padding-top:10px">
                        <label>数量</label>
                        <InputNumber :max="100" :min="1" v-model="cartlist[index].count" @on-blur="makechange(index)" ></InputNumber>
                    
                    </div>
                
                </div>
                    <div style="padding-left:900px">
                        <Button type="error" @click="remove(index)">删除</Button>
                    </div>
            </div>
            </Card>
        </div>
       </div>
       
      </div>
     
      </CheckboxGroup>
        <div style="padding-left:800px;padding-top:10px">
           <h1 style="display:inline;">总计：{{total}}</h1>
           <div style="padding-left:50px; display:inline">
               <Button type="primary" size='large' :disabled='isbuy' @click="tocheckout" >去结算</Button>
           </div>
           
       </div>
    </Card>
  </div>
</template>
<script>
import {getCart,delCart,setCartNum} from '../api/goods';
export default {
  data() {
    return {
      cartlist: [
        {
          cid: 11,
          uid: 1,
          pid: 3,
          pname: "拯救者 Y7000P 2019 英特尔酷睿i5 15.6英寸专业电竞本",
          peizhi:
            "i5-9300H/Windows 10 家庭中文版/15.6英寸/8G/1T SSD/ GeForce RTX™ 1660Ti 6G独显/黑色",
          linkid: 2,
          count: 2,
          aprice: 7899,
          creatime: "2019-12-21 15:09:32"
        },
        {
          cid: 12,
          uid: 1,
          pid: 2,
          pname: "拯救者 Y7000P 2019 英特尔酷睿i7 15.6英寸专业电竞本",
          peizhi:
            "i7-9750H/Windows 10 家庭中文版/15.6英寸/16G/1T SSD/ GeForce RTX™ 2060 6G独显/红色",
          linkid: 2,
          count: 1,
          aprice: 9898,
          creatime: "2019-12-21 15:11:58"
        },
        {
          cid: 13,
          uid: 1,
          pid: 1,
          pname: "小新Pro13.3英寸英特尔酷睿全面屏超轻薄笔记本电脑",
          peizhi: "I5-10210U 16G 512G MX250 2.5K QHD 银",
          linkid: 1,
          count: 1,
          aprice: 6699,
          creatime: "2019-12-21 15:12:01"
        }
      ],
      chosenid:[],
      empty:''
    };
  },
  methods:{
      init(){
          getCart().then(res=>{
              if(res.status==1){
                  this.cartlist=res.data
                  if(this.cartlist==null){
                      this.empty="什么也没有"
                  }
              }else{
                  this.$Message.error(res.msg)
              }
              
          }).catch(e=>{this.$Message.error("服务器不给力，等会再试吧")})
      },
      
      tocheckout(){
          this.$router.push({path:'/order',query:{pid:this.getpids(),count:this.getcounts()}})
      },
      makechange(index){
          
          setCartNum({pid:this.cartlist[index].pid,count:this.cartlist[index].count}).then(res=>{
              if(res.status==1){
                 
              }else{
                  this.$Message.error(res.msg)
              }
          }).catch(e=>{
                  this.$Message.error("服务器不给力，等会再试吧")
          })
      },
      getpids(){
          let temp=[]
          this.chosenid.forEach(i=>{
              temp.push(this.cartlist[i].pid)
          })
          return temp
      },
      remove(i){
          delCart({pid:this.cartlist[i].pid}).then(res=>{
              if(res.status==1){
                  
                  //首先删除选中
                  this.chosenid.forEach((item,index)=>{
                      if(item==i){
                          this.chosenid.splice(index,1)
                      }
                  })
                  this.$Message.success("移除成功")
                  this.cartlist.splice(i,1)
              }else{
                  this.$Message.error(res.msg)
              }
          }).catch(e=>{
              this.$Message.error("服务器不给力，等会再试吧")
          })
      },
      getcounts(){
          let tem=[]
           this.chosenid.forEach(i=>{
              tem.push(this.cartlist[i].count)
          })
          return tem
      }
  },
  mounted(){
      this.init()
  },
  computed:{
      total(){
          let temp=0
          this.chosenid.forEach(i=>{
              temp+=(this.cartlist[i].aprice)*(this.cartlist[i].count)
          })
          return temp
      },
      isbuy(){
          if(this.total>0){
              return false
          }
          return true
      }
  }
};
</script>
