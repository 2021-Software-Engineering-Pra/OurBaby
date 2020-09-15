<template>
    <div>
            <div style="padding-top:20px;padding-left:20px">
            <table style="width:100%">
                <tr>
                    <td style="text-align: center;"><p>用户名</p></td>
                    <td style="text-align: center;"><p>用户邮箱</p></td>
                    <td style="text-align: center;"><p>手机号</p></td>
                    <td style="text-align: center;"><p>余额</p></td>
                    <td style="text-align: center;"><p>用户状态</p></td>
                    <td style="text-align: center;"><p>操作</p></td>
                </tr>
                <tr v-for="(item,index) in data" :key="index">
                    <!-- <td><Input v-model="item.uid"></Input></td> -->
                    <!-- <td><p >{{item.uid}}</p></td> -->
                    <!-- <td><Input v-model="item.uname"></Input></td> -->
                    <td><p>{{item.uname}}</p></td>
                    <td><Input v-model="item.email"></Input></td>
                    <td><Input v-model="item.phone"></Input></td>
                    <td><Input v-model="item.yue"></Input></td>
                    <td><Input v-model="item.status"></Input></td>
                    <!-- <td><DatePicker v-model="item.updatime" format="yyyy-MM-dd" type="date" placeholder="Select date"></DatePicker></td> -->
                    <!-- <td><p format="yyyy-MM-dd">{{item.updatime}}</p></td> -->
                    <td><Button  type="primary" @click="updateitem(index)">更新</Button> 
                    <Button   @click="resetitem(index)">重置</Button> 
                    <Button  type="error" @click="delitem(index)">删除</Button> </td>
                </tr>
            </table>
                </div>
    </div>
</template>
<script>
import { deluser,updateuser,getalluser,resetuser } from "../api/admin";
export default {
    data(){
        return{
            data:[],
            // isact:this.$router.params.act,
            re:{ "status": 1,"msg": "ok","data": [{"uid": 1,"uname": "test","email": "ouc@stu.ouc.edu.cn","phone": "17685842001","yue": 124807,"status": 1,"creatime": "2019-11-27 21:11:09","updatime": "2020-09-01 10:38:03"},{"uid": 12,"uname": "xyzxx","email": "3512198443@qq.com","phone": "17860716615","yue": 0,"status": 1,"creatime": "2019-12-31 08:19:22","updatime": "2019-12-31 08:22:17"},{"uid": 14,"uname": "test111","email": "e@w.c","phone": "12345678911","yue": 100000,"status": 1,"creatime": "2020-01-04 21:52:51","updatime": "2020-01-04 21:52:51"},{"uid": 15,"uname": "test1112","email": "2W@e.c","phone": "11111111111","yue": 100000,"status": 1,"creatime": "2020-01-04 21:54:56","updatime": "2020-01-04 21:54:56"},{"uid": 16,"uname": "test1113","email": "e@qs.c","phone": "12345678911","yue": 87501,"status": 1,"creatime": "2020-01-04 21:55:29","updatime": "2020-01-04 21:57:29"},{"uid": 17,"uname": "nnn111","email": "ei@q.c1","phone": "12345678910","yue": 87501,"status": 1,"creatime": "2020-01-05 11:31:31","updatime": "2020-01-05 12:09:28"},{"uid": 18,"uname": "nnn222","email": "qwe@q.c","phone": "12345678911","yue": 90301,"status": 0,"creatime": "2020-01-05 12:24:24","updatime": "2020-09-01 09:17:19"}]},

        }
    },
    methods:{
        initdata(){
            getalluser().then(res=>{
                if(res.status==1){
                    this.data=res.data
                }else{
                    this.$Message.info(res.msg)
                }
            })
            // this.data=this.re.data
        },
        updateitem(i){
            updateuser({uid:this.data[i].uid,email:this.data[i].email,phone:this.data[i].phone,yue:this.data[i].yue}).then(res=>{
                if(res.status==1){
                     this.$Message.success("更新成功")
                }else{
                    this.$Message.info(res.msg)
                }
            }).catch(e=>{ this.$Message.error("暂时无法连接到服务器，请稍后再试")})
        },
        delitem(i){
            deluser({uid:this.data[i].uid}).then(res=>{
                 if(res.status==1){
                     this.$Message.success("删除成功")
                     this.data[i].status=0
                }else{
                    this.$Message.info(res.msg)
                }
            }).catch(e=>{ this.$Message.error("暂时无法连接到服务器，请稍后再试")})
        },
        resetitem(i){
           resetuser({uid:this.data[i].uid}).then(res=>{
                if(res.status==1){
                     this.$Message.success("重置成功")
                }else{
                    this.$Message.info(res.msg)
                }
           }).catch(e=>{ this.$Message.error("暂时无法连接到服务器，请稍后再试")})
        }
    },
    mounted(){
        this.initdata()
    }
}
</script>