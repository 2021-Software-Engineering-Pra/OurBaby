<template>
  <div>
        <keep-alive>
        <Col span='4'>
        <Menu :theme="theme3" :active-name="isact">
          <MenuGroup title="订单管理">
              <MenuItem  :to="{name:'morders',params:{act:1}}" name="1">
                <Icon type="md-document" />订单列表
              </MenuItem>
          </MenuGroup>
          <MenuGroup title="用户管理">
              <MenuItem :to="{name:'umassage',params:{act:3}}" name="3">
                <Icon type="md-heart" />用户列表
              </MenuItem>
          </MenuGroup>
           <MenuGroup title="商品管理">

              <MenuItem :to="{name:'productadmin',params:{act:3}}" name="3">
                <Icon type="md-leaf" />商品列表
              </MenuItem>
          </MenuGroup>
        </Menu>
        </Col>
        </keep-alive>
        <Col span='20' >
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
