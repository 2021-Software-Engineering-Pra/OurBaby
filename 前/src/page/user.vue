<template>
  <div>
        <h1 style="padding-left:20px">{{uname}},欢迎</h1>
        <keep-alive>
          <Col span='6'>
        <Menu :theme="theme3" :active-name="isact">
          <MenuGroup title="订单管理">
              <MenuItem  :to="{name:'orders',params:{act:1}}" name="1">
                <Icon type="md-document" />所有订单
              </MenuItem>
          </MenuGroup>
          <MenuGroup title="账户管理">

              <MenuItem :to="{name:'udetail',params:{act:3}}" name="3">
                <Icon type="md-heart" />账户详情
              </MenuItem>


              <MenuItem :to="{name:'address',params:{act:4}}" name="4">
                <Icon type="md-leaf" />收货地址
              </MenuItem>

          </MenuGroup>
        </Menu>
        </Col>
        </keep-alive>
        <Col span='10' >
            <router-view></router-view>
        </Col>
        
  </div>
</template>
<script>
export default {
  data() {
    return {
      theme3: "light",
      uname: this.$route.params.uname,
      isact: this.$route.params.act
    };
  },
  methods: {
    //   select(){
    //       console.log(this.isact)
    //   }
  },
  
};
</script>
