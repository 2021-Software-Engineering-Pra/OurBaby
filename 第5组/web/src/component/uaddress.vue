<template>
    <div>
        <Button type="primary" size='large' @click='add'>新建地址</Button>
        <div>
            <div v-for="(item,index) in addres" :key='index'>
                     <Col span='20' style="padding-top:20px">
                    <Card>
                    <div>
                        <h2>{{item.recvname}}</h2>
                        <h3>{{item.phone}}</h3>
                        <p>{{item.detailed}}</p>
                    </div>
                    </Card>
                </Col>
                <Col span="3" style="padding-left:20px;padding-top:25px">
                    <Button @click="changeadd(index)">更改</Button>
                    <div style="padding-top:20px">
                        <Button  type="error" @click="del(index)">删除</Button>
                    </div>
                </Col>
            </div>
        </div>
            <Modal
                v-model="modal"
                :title="mtitle"
                @on-ok="ok"
                >
                <Form>
                    <FormItem label="收货人">
                    <Input v-model="recv" placeholder="请输入收货人"></Input>
                    </FormItem>
                    <FormItem label="手机">
                    <Input v-model="phone" type="tel" @on-blur='cphone' placeholder="请输入手机号"></Input>
                    <span style="color:red">{{phoneerror}}</span>
                    </FormItem>
                    <FormItem label="详细地址">
                    <Input v-model="detail" type="textarea" :autosize="{minRows: 2,maxRows: 5}" placeholder="请输入详细地址"/>
                </FormItem>
                </Form>
            </Modal>
    </div>
</template>
<script>
import {getAddress,addAddress,updateAddress,delAddress} from '../api/index'
export default {
    data(){
        return{
            addres:[],
            modal:false,
            recv:"",
            phone:'',
            detail:'',
            mtitle:'',
            chosenid:null,
            phoneerror:'',
        }
    },
    methods:{
        getadd(){
            getAddress().then(res1=>{
                this.$set(this,'addres',res1.data)
            }).catch(e=>{
                this.$Message.error("服务器繁忙请稍后再试")
            })
        },
        changeadd(index){
            console.log(index)
            this.mtitle="更改地址"
            this.modal=true
            this.chosenid=index
            this.recv=this.addres[index].recvname
            this.phone=this.addres[index].phone
            this.detail=this.addres[index].detailed
            
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
        del(index){
            console.log(index)
            delAddress({aid:this.addres[index].aid}).then(res=>{
                if(res.status==1){
                        this.$Message['success']({
                            background:true,
                            content:"删除成功"
                        })
                        this.addres.splice(index,1)
                    }else{
                        this.$Message['error']({
                            background:true,
                            content:res.msg
                        })
                    }
            })
        },
        add(){
            this.mtitle="添加地址"
            this.chosenid=null
            this.recv=''
            this.phone=''
            this.detail=''
            this.modal=true
        },
        ok(){
            if(this.chosenid==null){
                //新建地址
                if(this.cphone()&&this.recvname!=''&&this.detail!=""){
                    addAddress({recvname:this.recv,phone:this.phone,detailed:this.detail}).then(res=>{
                    console.log(res)
                    if(res.status==1){
                        this.$Message['success']({
                            background:true,
                            content:"添加成功"
                        })
                        this.addres.push({recvname:this.recv,phone:this.phone,detailed:this.detail})
                    }else{
                        this.$Message['error']({
                            background:true,
                            content:res.msg
                        })
                    }
                    
                }).catch(e=>{
                    console.log("服务器不给力，等会再试吧")
                })
                }else{
                    this.$Message.error("数据存在错误，未能成功添加")
                }
                
            }else{
                updateAddress({aid:this.addres[this.chosenid].aid,recvname:this.recv,phone:this.phone,detailed:this.detail}).then(res=>{
                    console.log(res)
                    if(res.status==1){
                        this.$Message['success']({
                            background:true,
                            content:"更改成功"
                        })
                      
                        this.addres[this.chosenid].recvname=this.recv
                        this.addres[this.chosenid].phone=this.phone
                        this.addres[this.chosenid].detailed=this.detail
                    }else{
                        this.$Message['error']({
                            background:true,
                            content:res.msg
                        })
                    }
                })
                
            }
        }
    },
    mounted(){
        this.getadd()
    }
}
</script>