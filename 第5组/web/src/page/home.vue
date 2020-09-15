<template>
    <div >
        <Carousel autoplay style="" >
            <CarouselItem v-for="(i,index) in hot" :key='index' >
                 <router-link :to="{path:'detail/'+i.pid,query:{type:i.type}}">
                <img :src="'/src/static/'+i.linkid+'/big.jpg'"  alt="" srcset="">
                </router-link>
            </CarouselItem>
        </Carousel>
        <div style="padding-top:20px">
            <Card title="热卖推荐" style="max-width:90%;margin:auto;">
            <div  style="display:inline-block;margin:auto;" >
            <div  v-for="(item,index) in hot" style="max-width:370px;padding-left:30px;padding-top:20px;float:left;" :key='index' >
                <div style="border:1px  solid #E7E9EB;">
                    <router-link :to="{path:'detail/'+item.pid,query:{type:item.type}}">
                    <div style="max-width:300px;margin:auto">
                         <img :src="'/src/static/'+item.linkid+'/main.jpg'" style="max-width:300px" alt="" srcset="">
                    <h3 style="word-break:break-word;color:black">{{item.pname}}</h3>
                    <p style="color:red" >￥{{item.price}}</p>
                    <h4 style="word-break:break-word;color:black">{{item.pdesp}}</h4>
                    </div>
                    </router-link>
                </div>
            </div>
            </div>
        </Card>
        </div>
        <div style="padding-top:20px">
        <Card title="新品推荐" style="max-width:90%;margin:auto;padding-top:20px">
             <div  style="display:inline-block;margin:auto;">
            <div  v-for="(i,index) in newest" style="max-width:370px;padding-left:30px;padding-top:20px;float:left;" :key='index' >
                <div style="border:1px  solid #E7E9EB;" >
                    <router-link :to="{path:'detail/'+i.pid,query:{type:i.type}}">
                    <div style="max-width:300px;margin:auto">
                         <img :src="'/src/static/'+i.linkid+'/main.jpg'" style="max-width:300px" alt="" srcset="">
                    <h3 style="word-break:break-word;color:black;">{{i.pname}}</h3>
                    <p style="color:red" >￥{{i.price}}</p>
                    <h4 style="word-break:break-word;color:black">{{i.pdesp}}</h4>
                    </div>
                    </router-link>
                   
                </div>
            </div>
            </div>
        </Card>
        </div>
        <Button @click="show">点我</Button>
    </div>
</template>
<script>
import {getNewGoods,getHotGoods} from '../api/goods.js'
import {deluser} from '../api/index.js'
export default {
    data(){
        return{
            status:0,
            msg:'',
            hot:[],
            newest:[]
        }
    },
    methods:{
      show(){
        deluser({uid:"18"}).then(res=>{
            console.log(res)
        })
      },
      gethot(){
        getHotGoods().then(res=>{
            this.hot=res.data[0]
            console.log(this.hot[0])
        })
      },
      getnew(){
        getNewGoods().then(res1=>{
           this.newest=res1.data[0]
           console.log('new',this.newest[0])
          })
    },
    },
    mounted(){
     this.gethot()
     this.getnew()
    },
   
 
}
</script>