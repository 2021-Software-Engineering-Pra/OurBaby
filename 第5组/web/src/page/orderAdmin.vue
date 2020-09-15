<template>
    <div>
           <div style="padding-top:20px;padding-left:20px">
            <table style="width:100%"  >
                <tr>
                    <td style="text-align: center;"><p>订单编号</p></td>
                    <td style="text-align: center;"><p>商品名称</p></td>
                    <td class="smalltd"><p>商品单价</p></td>
                    <td class="smalltd"><p>商品数量</p></td>
                    <td style="text-align: center;"><p>下单用户</p></td>
                    <td class="smalltd"><p>订单状态</p></td>
                    <td class="smalltd"><p>收货人</p></td>
                    <td style="text-align: center;"><p>手机号</p></td>
                    <td style="text-align: center;"><p>详细地址</p></td>
                    <td style="text-align: center;"><p>创建时间</p></td>
                    <td style="text-align: center;"><p>操作</p></td>
                </tr>
                <tr v-for="(item,index) in data" :key="index">
                    <!-- <td><Input v-model="item.uid"></Input></td> -->
                    <td>{{item.oid}}</td>
                    <!-- <td><Input v-model="item.uname"></Input></td> -->
                    <td>{{item.pname}}</td>
                    <td>{{item.aprice}}</td>
                    <td>{{item.count}}</td>
                    <td>{{item.uname}}</td>
                    <td><Input v-model="item.status"></Input></td>
                    <td><Input v-model="item.recvname"></Input></td>
                    <td><Input v-model="item.phone"></Input></td>
                    <td><Input v-model="item.detailed"></Input></td>
                    <td>{{item.creatime}}</td>
                    <!-- <td><DatePicker v-model="item.updatime" format="yyyy-MM-dd" type="date" placeholder="Select date"></DatePicker></td> -->
                    <!-- <td><p format="yyyy-MM-dd">{{item.updatime}}</p></td> -->
                    <td><Button  type="primary" @click="updateitem(index)" :disabled='item.isdis' :loading='processing'>更新</Button> </td>
                    <!-- <td><Button  type="error" @click="resetitem(index)">下架</Button> </td> -->
                </tr>
            </table>
        </div>
        <!-- <h1 style="padding-left:20px">订单管理</h1>
        <div>
            <div v-for="(item,index) in data" :key='index'>
                <Col span='20' style="padding-top:20px">
                    <Card>
                        <div>
                            <p>{{item.oid}}</p>
                            <p>单价：{{item.aprice}}</p>
                            <p>数量：{{item.count}}</p>
                            <p>状态：<span style="color:red">{{item.status}}</span></p>
                            <p>创建时间：{{item.creatime}}</p>
    
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
            </Modal> -->
    </div>
</template>
<script>
import {getallOrder,updateOrder} from '../api/admin'
export default {
  data(){
    return{
        data:[],
        re:{"status": 1,"msg": "ok","data": [{"oid": "202001051214221","uid": 1,"count": "1 2","pid": "4 3","aprice": 12499,"status": "已取消","pname": "ThinkPad X1 Carbon 2019 英特尔酷睿 笔记本电脑 20QD002BCD","peizhi": "i7-8565U/Windows 10家庭中文版/8GB/512GB SSD/UHD 620显示芯片/14英寸WQHD IPS LED背光显示屏","uname": "test","linkid": 3,"recvname": "老王1","phone": "1234567892","detailed": "山东省青岛市崂山区中国海洋大学","creatime": "2020-01-05 12:14:22","updatime": "2020-01-05 12:20:00"},{"oid": "202001051216141","uid": 1,"count": "1","pid": "4","aprice": 12499,"status": "已支付","pname": "ThinkPad X1 Carbon 2019 英特尔酷睿 笔记本电脑 20QD002BCD","peizhi": "i7-8565U/Windows 10家庭中文版/8GB/512GB SSD/UHD 620显示芯片/14英寸WQHD IPS LED背光显示屏","uname": "test","linkid": 3,"recvname": "李晓明","phone": "17843215678","detailed": "山东青岛崂山","creatime": "2020-01-05 12:16:14","updatime": "2020-01-05 12:18:29"},{"oid": "2020010512250818","uid": 18,"count": "1","pid": "5","aprice": 9699,"status": "已完成","pname": "ThinkPad X1 Carbon 2019 英特尔酷睿 笔记本电脑 20QD0020CD","peizhi": "i5-8265U/Windows 10家庭中文版/8GB/512GB SSD/UHD 620显示芯片/14英寸WQHD IPS LED背光显示屏","uname": "nnn222","linkid": 3,"recvname": "小张","phone": "12345678911","detailed": "Shandong Qingdao","creatime": "2020-01-05 12:25:08","updatime": "2020-01-05 12:27:52"},{"oid": "202001061145501","uid": 1,"count": "1","pid": "2","aprice": 9898,"status": "已完成","pname": "拯救者 Y7000P 2019 英特尔酷睿i7 15.6英寸专业电竞本","peizhi": "i7-9750H/Windows 10 家庭中文版/15.6英寸/16G/1T SSD/ GeForce RTX™ 2060 6G独显/红色","uname": "test","linkid": 2,"recvname": "张晓明","phone": "12345678911","detailed": "山东省青岛市崂山区松岭路238号中国海洋大学东区","creatime": "2020-01-06 11:45:50","updatime": "2020-01-06 11:46:23"},{"oid": "202009011038071","uid": 1,"count": "1","pid": "2","aprice": 9898,"status": "已支付","pname": "拯救者 Y7000P 2019 英特尔酷睿i7 15.6英寸专业电竞本","peizhi": "i7-9750H/Windows 10 家庭中文版/15.6英寸/16G/1T SSD/ GeForce RTX™ 2060 6G独显/红色","uname": "test","linkid": 2,"recvname": "你好啊","phone": "12345678912","detailed": "山东青岛","creatime": "2020-09-01 10:38:07","updatime": "2020-09-02 14:51:00"}]}

    }
  },
  methods:{
        initdata(){
            getallOrder().then(res=>{
                if(res.status==1){
                    var t =new Array()
                    t=res.data
                    t.forEach((e,i)=>{
                        t[i].isdis=this.isdisable(e)
                    })
                    this.data= t.sort((a,b)=>a.isdis-b.isdis)
                    // this.data=res.data
                }else{
                    this.$Message.info(res.msg)
                }
            })
            // this.data=this.re.data
        },
        isdisable(e){
            if(e.status=="已完成"||e.status=="已取消"){
                return true
            }else{
                return false
            }
        },
        updateitem(i){
            updateOrder({orid:this.data[i].oid,recvname:this.data[i].recvname,phone:this.data[i].phone,detail:this.data[i].detailed,status:this.data[i].status}).then(res=>{
                if(res.status==1){
                     this.$Message.success("更新成功")
                }else{
                    this.$Message.info(res.msg)
                }
            }).catch(e=>{ this.$Message.error("暂时无法连接到服务器，请稍后再试")})
        },
  },
  mounted() {
      this.initdata()
      
  },
};   
</script>

