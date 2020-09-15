<template>
  <div>
      <!-- <Modal title="欢迎你，管理员" v-model="notice">
          <p>点击左边开始管理商城吧</p>
      </Modal> -->
        <Row>
        <keep-alive>
          <Col span='3'>
            <Menu :theme="theme3"  style="width:100%">
            <MenuGroup title="管理">
                <MenuItem  :to="{name:'productadmin'}" name='1' >
                    <Icon type="md-document" />商品管理
                </MenuItem>
                <MenuItem :to="{name:'orderadmin'}" name='2'>
                    <Icon type="md-heart" />订单管理
                </MenuItem>
                <MenuItem :to="{name:'useradmin'}" name='3'>
                    <Icon type="md-leaf" />用户管理
                </MenuItem>
                </MenuGroup>
            </Menu>
        </Col>
        </keep-alive>
        <Col span="21">
        <router-view></router-view>
        </Col>
        </Row>
  </div>
</template>

<script>
export default {
    data(){
        return{
            notice:false,
        }
    },
    methods:{
        killnotice(){
            document.cookie="shownotice="+'no'
        },
        isshow(){
            if(!document.cookie.includes("shownotice")){
                this.notice=true
            }
        }
    },
    mounted(){
        // this.isshow()
        //  setTimeout(() => {
        //             this.notice = false;
        //             this.killnotice()
        //         }, 5000);
    }
}
</script>

<style>

</style>