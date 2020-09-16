<template>
  <div>
    <h1>所有用户</h1>
    <h2>{{msg}}</h2>
       <div v-for='(i,index) in data' :key='index' style="padding-top:10px">
                <Col span='20' style="padding-top:20px">
                    <Card>
                    <div>
                        <h3>用户id：{{i.uid}}</h3>
                        <h3>用户名：{{i.uname}}</h3>
                        <h3>用户邮箱：{{i.email}}</h3>
                        <h3>用户手机号：{{i.phone}}</h3>
                        <h3>用户余额：{{i.yue}}</h3>
                        <h3>用户状态：{{i.status}}</h3>
                    </div>
                    </Card>
                    <!-- 添加收货地址面板 -->
                    <Modal
                    v-model="modal"
                    title="更新用户信息"
                    @on-ok="ok"
                    >
                    <Form>
                        <FormItem label="用户邮箱">
                        <Input v-model="email" placeholder="请输入用户邮箱"></Input>
                        </FormItem>
                        <FormItem label="手机">
                        <Input v-model="phone" type="tel" @on-blur='cphone' placeholder="请输入用户手机号"></Input>
                        <span style="color:red">{{phoneerror}}</span>
                        </FormItem>
                        <FormItem label="用户余额">
                        <Input v-model="yue" type="tel" placeholder="请输入用户余额"></Input>
                        </FormItem>
                        <FormItem label="用户状态">
                        <Input v-model="status" type="tel" placeholder="请输入用户状态"></Input>
                        </FormItem>
                    </Form>
                </Modal>
                </Col>
                <Col span="3" style="padding-left:20px;padding-top:25px">
                        <Button @click="changeadd(index)">更新</Button>
                    <div style="padding-top:20px">
                        <Button   type="error" @click="reseti(index)">重置</Button>
                    </div>
                     <div style="padding-top:20px">
                        <Button  type="error" @click="deli(index)">删除</Button>
                    </div>
                </Col>
            </div>
            </div>
</template>
<script>
import {deluser, updateuser,getalluser,resetuser  } from "../api/admin";
export default {
    data(){
        return{
            data:[],
            re:{ "status": 1,"msg": "ok","data": [{"uid": 1,"uname": "test","email": "ouc@stu.ouc.edu.cn","phone": "17685842001","yue": 124807,"status": 1,"creatime": "2019-11-27 21:11:09","updatime": "2020-09-01 10:38:03"},{"uid": 12,"uname": "xyzxx","email": "3512198443@qq.com","phone": "17860716615","yue": 0,"status": 1,"creatime": "2019-12-31 08:19:22","updatime": "2019-12-31 08:22:17"},{"uid": 14,"uname": "test111","email": "e@w.c","phone": "12345678911","yue": 100000,"status": 1,"creatime": "2020-01-04 21:52:51","updatime": "2020-01-04 21:52:51"},{"uid": 15,"uname": "test1112","email": "2W@e.c","phone": "11111111111","yue": 100000,"status": 1,"creatime": "2020-01-04 21:54:56","updatime": "2020-01-04 21:54:56"},{"uid": 16,"uname": "test1113","email": "e@qs.c","phone": "12345678911","yue": 87501,"status": 1,"creatime": "2020-01-04 21:55:29","updatime": "2020-01-04 21:57:29"},{"uid": 17,"uname": "nnn111","email": "ei@q.c1","phone": "12345678910","yue": 87501,"status": 1,"creatime": "2020-01-05 11:31:31","updatime": "2020-01-05 12:09:28"},{"uid": 18,"uname": "nnn222","email": "qwe@q.c","phone": "12345678911","yue": 90301,"status": 0,"creatime": "2020-01-05 12:24:24","updatime": "2020-09-01 09:17:19"}]},
            modal:false,
            title:'',
            email:"",
            phone:'',
            yue:'',
            status:'',
            chosenid:null,
            phoneerror:'',
        }
    },
    methods:{
        initdata(){
            this.data=this.re.data      
        },
        changeadd(index){
            console.log(index)
            this.title="更新用户信息"
            this.modal=true
            this.chosenid=index
            this.email=this.data[index].email
            this.phone=this.data[index].phone
            this.yue=this.data[index].yue
            this.status=this.data[index].status
        },
         cphone(){
             if (this.phone.length!=11){
                this.phoneerror="手机号长度不对"
                return false
            }
            else if(this.phone.match('^1[0-9]+')==null){
                this.phoneerror="手机号格式不对"
                return false
            }
            else{
                this.phoneerror=''
                return true
            }
        },
        ok(){
            if(this.chosenid!=null){
            updateuser({uid:this.data[this.chosenid].uid,email:this.email,phone:this.phone,yue:this.yue}).then(res=>{
                if(res.status==1){
                     this.$Message['success']({
                            background:true,
                            content:"更改成功"
                        })
                     this.data[this.chosenid].email=this.email
                     this.data[this.chosenid].phone=this.phone
                     this.data[this.chosenid].yuee=this.yue
                     this.data[this.chosenid].statue=this.status
                }else{
                    this.$Message['error']({
                            background:true,
                            content:res.msg
                        })
                }
            }).catch(e=>{ this.$Message.error("暂时无法连接到服务器，请稍后再试")})
        }},
        deli(i){
            deluser({uid:this.data[i].uid}).then(res=>{
                 if(res.status==1){
                     this.$Message.success("删除成功")
                     this.data[i].status=0
                }else{
                    this.$Message.info(res.msg)
                }
            }).catch(e=>{ this.$Message.error("暂时无法连接到服务器，请稍后再试")})
        },
        reseti(i){
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
        console.log(this.isact)
    }
}
</script>