<template>
  <div >
    <top/>
      <router-view :key="$route.fullPath" v-if="isAlive"></router-view>
      <BackTop></BackTop>
  </div>
</template>

<script>
import top from './component/top.vue'
export default {
     name:'app',
     data(){
       return{
         isAlive:true
       }
     },
     components:{
         top,
     },
     methods: {
       reload(){
         this.isAlive=false
         this.$nextTick(function(){
           this.isAlive=true
         })
       }
     },
}
</script>

