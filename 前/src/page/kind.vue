<template>
    <div id='home'>
        <h1>{{errmsg}}</h1>
        <div v-if="errmsg.length==0" style="padding-left:40px;padding-top:20px">
             <ButtonGroup type="button">
               <Button  @click="paixu(1)">价格<Icon :type="pricekind" /></Button>
                <Button @click="paixu(2)">时间<Icon :type="timekind" /></Button>
            </ButtonGroup>
        </div>
        <div style="padding-left:20px;">
                <Col span="7" v-for="(item,index) in datas" :key='index' style="width:350px;padding-left:20px;padding-top:20px">
                <Card :title="item.msg" :to="{path:'/detail/'+item.pid,query:{type:item.type}}">
                    <img :src="'/src/static/'+item.linkid+'/main.jpg'" style="max-width:300px" alt="图片" srcset="">
                <h3 style="word-break:break-word;color:black">{{item.pname}}</h3>
                <p style="color:red" >￥{{item.price}}</p>
                <h4 style="word-break:break-word;color:black">{{item.pdesp}}</h4>
                </Card>
            </Col>
            
        </div>
        <BackTop></BackTop>
    </div>
</template>
<script>
import {getKind,searchProduct} from '../api/goods'
export default {
    data(){
        return{
            status: 1,
            kind:this.$route.params.kind,
            key:this.$route.query.key,
            msg: "",
            datas:[],
            timekind:"ios-arrow-down",
            pricekind:'',
            button1:'',
            errmsg:''
        }
    },
    methods:{
      search(){
          searchProduct({key:this.key}).then(
              res=>{
                  if (res.status==1&&res.data!=null){
                       this.datas=res.data 
                  }else if(res.data==null){
                      this.errmsg="什么也没有搜到呀~"
                  }else{
                      this.$Message.error(res.msg)
                  }
              }
          ).catch(e=>{this.$Message.error("服务器不给力，等会再试吧")})
      },
      gkind(){
          getKind({kind:this.kind}).then(
              res=>{
                  if(res.status==1) {
                      this.datas=res.data
                  }else{
                      this.errmsg=res.msg
                  }

              }
          ).catch(e=>{this.$Message.error("服务器不给力，等会再试吧")})
      },
      paixu(kind){
          let t=this.datas
          if(kind==1){
              this.timekind=""
              if(this.status%2==1){
              t.sort((a,b)=>{ return a.price-b.price})
              this.pricekind="ios-arrow-up"
          } 
          else{
              this.pricekind="ios-arrow-down"
              t.reverse()
          }
          }else{
              this.pricekind=''
              if(this.status%2==1){
              t.sort((a,b)=>{ return a.pid-b.pid})
              
              this.timekind="ios-arrow-up"
          } 
          else{
              this.timekind="ios-arrow-down"
              t.reverse()
          }
          }
          
         this.status++
      }
    },
    mounted(){
        // console.log(this.kind,this.key)
        if (this.kind!=undefined){
            this.gkind()
        }else{
            this.search()
        }
        
    }

}
</script>