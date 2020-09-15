<template>
  <div>
      <div class="addnew"><Button type="primary" size='large' @click="showadd=true" >新增商品</Button></div>
       <Modal  fullscreen title="新增商品" v-model="showadd" :mask-closable="false" :draggable='true'>
           <Form :label-width="80">
        <FormItem label="商品名称">
            <Input v-model="pname" placeholder="请输入商品名称"></Input>
        </FormItem>
        <FormItem label="商品描述">
            <Input v-model="pdesp" placeholder="请输入商品描述"></Input>
        </FormItem>
        <FormItem label="商品配置">
            <Input v-model="peizhi" placeholder="请输入商品配置"></Input>
        </FormItem>
        <FormItem label="商品价格">
            <InputNumber style="width:100%" :active-change='false' v-model="price" placeholder="请输入商品价格"></InputNumber>
        </FormItem>
        <FormItem label="商品库存">
            <InputNumber style="width:100%" :active-change='false' v-model="kucun" placeholder="请输入商品库存"></InputNumber>
        </FormItem>
        <FormItem label="商品系列">
            <Input v-model="type" placeholder="请输入商品系列"></Input>
        </FormItem>
        <FormItem label="商品种类">
            <Select v-model="kind">
                <Option value="light">轻薄本</Option>
                <Option value="game">游戏本</Option>
                <Option value="ThinkPad">ThinkPad</Option>
            </Select>
        </FormItem>
        <FormItem>
            <Row>
                <Col span="8">
                <Upload type='drag' :action='uploadurl' :format='imageformat' :before-upload='mainfile'>
                <!-- <Button icon="ios-cloud-upload-outline">请上传商品主图</Button> -->
                <Icon type="ios-cloud-upload" size="52" style="color: #3399ff"></Icon>
                    <p>主图</p>
            </Upload>
              <a @click="removemain" v-if='main!==null'>{{main.name}}</a>
                </Col>
            <Col span="8">
                <Upload :action='uploadurl' type='drag'  :format='imageformat' :before-upload='bigfile'>
                <!-- <Button icon="ios-cloud-upload-outline">请上传商品大图</Button> -->
                <Icon type="ios-cloud-upload" size="52" style="color: #3399ff"></Icon>
                <p>大图</p>
                
            </Upload>
             <a @click="removebig" v-if='big!==null'>{{big.name}}</a>
            </Col>
            <Col span="8">
                <Upload :action='uploadurl' type='drag' :format='imageformat' multiple :before-upload='introfile'>
                <!-- <Button icon="ios-cloud-upload-outline">请上传商品介绍图</Button> -->
                <Icon type="ios-cloud-upload" size="52" style="color: #3399ff"></Icon>
                <p>介绍图</p>
               
            </Upload>
             <div v-if='intro!==null'>
                <div v-for='(i,index) in intro' :key=i> <a @click="removeintro(index)" v-if='i!==null'>{{i.name}}</a></div>
                </div>
            </Col>
            </Row>
        </FormItem>
            <!-- <Button type='primary' @click="done">上传</Button> -->
        </Form>
        <div slot="footer">
            <Button type="error" size="large"   @click="cancel">取消</Button>
            <Button type="primary" size="large" :loading="isload" @click="done">确定</Button>
        </div>
       </Modal>
       <div style="padding-top:20px;padding-left:20px">
            <table style="width:100%"  >
                <tr>
                    <td style="text-align: center;"><p>主图</p></td>
                    <td style="text-align: center;"><p>商品名称</p></td>
                    <td style="text-align: center;"><p>商品描述</p></td>
                    <td style="text-align: center;"><p>商品配置</p></td>
                    <td class="smalltd"><p>商品价格</p></td>
                    <td class="smalltd"><p>商品库存</p></td>
                    <td style="text-align: center;"><p>商品系列</p></td>
                    <td class="smalltd"><p>商品种类</p></td>
                    <td class="smalltd"><p>状态</p></td>
                    <td style="text-align: center;"><p>大图</p></td>
                    <td style="text-align: center;"><p>详情</p></td>
                    <td style="text-align: center;"><p>操作</p></td>
                    
                </tr>
                <tr v-for="(item,index) in data" :key="index">
                    <!-- <td><Input v-model="item.uid"></Input></td> -->
                     <td><a @click="toshowmainimage('/src/static/'+item.linkid+'/main.jpg')"><img :src="'/src/static/'+item.linkid+'/main.jpg'" class="smallimage" alt="" srcset=""></a></td>
                    <td><Input v-model="item.pname"></Input></td>
                    <!-- <td><Input v-model="item.uname"></Input></td> -->
                    <td><Input v-model="item.pdesp"></Input></td>
                    <td><Input v-model="item.peizhi"></Input></td>
                    <td><Input v-model="item.price"></Input></td>
                    <td><Input v-model="item.kucun"></Input></td>
                    <td><Input v-model="item.type"></Input></td>
                    <td><Select v-model="item.kind">
                        <Option value="light">轻薄本</Option>
                        <Option value="game">游戏本</Option>
                        <Option value="ThinkPad">ThinkPad</Option>
                        </Select></td>
                    <td><Input v-model="item.status"></Input></td>
                    <td><a @click="toshowbigimage('/src/static/'+item.linkid+'/big.jpg')"><img :src="'/src/static/'+item.linkid+'/big.jpg'" class="smallimage" alt="" srcset=""></a></td>
                    <td><a @click="toshowintroimage('/src/static/'+item.linkid+'/intro.jpg')"><img :src="'/src/static/'+item.linkid+'/intro.jpg'" class="smallimage" alt="" srcset=""></a></td>
                    <!-- <td><DatePicker v-model="item.updatime" format="yyyy-MM-dd" type="date" placeholder="Select date"></DatePicker></td> -->
                    <!-- <td><p format="yyyy-MM-dd">{{item.updatime}}</p></td> -->
                    <td><Button  type="primary" @click="updateitem(index)" :loading='processing'>更新</Button> </td>
                    <!-- <td><Button  type="error" @click="resetitem(index)">下架</Button> </td> -->
                </tr>
            </table>
        </div>
        <Modal v-model="showmainimage" fullscreen @on-cancle='removemain'>
            <div>
                <img :src="mainimage" alt="" srcset="">
            </div>
            <Upload type='drag' :format='imageformat'  :before-upload='mainfile'>
              <Icon type="ios-cloud-upload" size="52" style="color: #3399ff"></Icon>
              <p>上传新图片替换</p>
            </Upload>
                <a @click="removemain" v-if='main!==null'>{{main.name}}</a>
        </Modal>
        <Modal v-model="showbigimage" fullscreen @on-cancle='removebig'>
            <div>
                <img :src="bigimage" style="width:800px" alt="" srcset="">
            </div>
            <Upload type='drag' :format='imageformat'  :before-upload='bigfile'>
              <Icon type="ios-cloud-upload" size="52" style="color: #3399ff"></Icon>
              <p>上传新图片替换</p>
            </Upload>
            <a @click="removebig" v-if='big!==null'>{{big.name}}</a>

        </Modal>
        <Modal v-model="showintroimage" fullscreen @on-cancle='removeintro'>
            <div>
                <img :src="introimage" class="bigimage" alt="" srcset="">
            </div>
            <Upload type='drag' :format='imageformat' multiple :before-upload='introfile'>
              <Icon type="ios-cloud-upload" size="52" style="color: #3399ff"></Icon>
              <p>上传新图片替换</p>
            </Upload>
            <div v-if='intro!==null'>
                <div v-for='(i,index) in intro' :key=i> <a @click="removeintro(index)" v-if='i!==null'>{{i.name}}</a></div>
                </div>
        </Modal>
  </div>
</template>

<script>
import { addProduct,updateProduct,delProduct,getAllProduct } from "../api/admin";
export default {
    data(){
        return{
            uploadurl:"/api/addproduct",
            showadd:false,
            isload:false,
            processing:false,
            showmainimage:false,
            showbigimage:false,
            showintroimage:false,
            mainimage:'',
            bigimage:'',
            introimage:'',
            data:[],
            temp:{"status": 1,"msg": "ok","data": [{"pid": "1","pname": "小新Pro13.3英寸英特尔酷睿全面屏超轻薄笔记本电脑","pdesp": "全新10代i5，“真”全面屏","peizhi": "I5-10210U 16G 512G MX250 2.5K QHD 银","kucun": 3402,"price": 6699,"type": "小新Pro","kind": "light","sellcount": 0,"status": 1,"linkid": 1,"creatime": "2019-11-27 21:29:57","updatime": "2020-01-05 21:19:28"},{"pid": "2","pname": "拯救者 Y7000P 2019 英特尔酷睿i7 15.6英寸专业电竞本","pdesp": "为战而生，全新拯救者助你一臂之力","peizhi": "i7-9750H/Windows 10 家庭中文版/15.6英寸/16G/1T SSD/ GeForce RTX™ 2060 6G独显/红色","kucun": 3995,"price": 9898,"type": "拯救者Y7000P","kind": "game","sellcount": 0,"status": 1,"linkid": 2,"creatime": "2019-12-11 20:42:12","updatime": "2020-09-01 10:38:07"},{"pid": "3","pname": "拯救者 Y7000P 2019 英特尔酷睿i5 15.6英寸专业电竞本","pdesp": "为战而生，全新拯救者助你一臂之力","peizhi": "i5-9300H/Windows 10 家庭中文版/15.6英寸/8G/1T SSD/ GeForce RTX™ 1660Ti 6G独显/黑色","kucun": 3986,"price": 7899,"type": "拯救者Y7000P","kind": "game","sellcount": 0,"status": 1,"linkid": 2,"creatime": "2019-12-11 20:46:58","updatime": "2020-01-05 21:14:45"},{"pid": "4","pname": "ThinkPad X1 Carbon 2019 英特尔酷睿 笔记本电脑 20QD002BCD","pdesp": "超轻薄精英商务本","peizhi": "i7-8565U/Windows 10家庭中文版/8GB/512GB SSD/UHD 620显示芯片/14英寸WQHD IPS LED背光显示屏","kucun": 3983,"price": 12499,"type": "ThinkPad X1 Carbon","kind": "ThinkPad","sellcount": 0,"status": 1,"linkid": 3,"creatime": "2019-12-11 20:50:31","updatime": "2020-01-05 12:20:00"},{"pid": "5","pname": "ThinkPad X1 Carbon 2019 英特尔酷睿 笔记本电脑 20QD0020CD","pdesp": "超轻薄精英商务本","peizhi": "i5-8265U/Windows 10家庭中文版/8GB/512GB SSD/UHD 620显示芯片/14英寸WQHD IPS LED背光显示屏","kucun": 3997,"price": 9699,"type": "ThinkPad X1 Carbon","kind": "ThinkPad","sellcount": 0,"status": 1,"linkid": 3,"creatime": "2019-12-11 20:51:35","updatime": "2020-09-02 12:55:45"}]},
            imageformat:['jpg','png','bmp'],
                pname:"",
                pdesp:"",
                peizhi:'',
                price:null,
                kucun:null,
                type:"",
                kind:'',
                main:null,
                big:null,
                intro:[]
        }
    },
    methods:{
        initdata(){
            getAllProduct().then(res=>{
                if(res.status==1){
                    this.data=res.data
                }else{
                     this.$Message.info("获取数据失败："+res.msg)
                }
            }).catch(e=>{
                 this.$Message.error("连接不到服务器")
            })
            // this.data=this.temp.data
            
        },
        toshowmainimage(url){
            // console.log(url)
            this.mainimage=url
            this.showmainimage=true
        },
        mainfile(file){
            this.main=file
            return false
        },
        removemain(){
            this.main=null
        },
        bigfile(file){
            this.big=file
            return false
        },
         toshowbigimage(url){
            // console.log(url)
            this.bigimage=url
            this.showbigimage=true
        },
        removebig(){
            this.big=null
        },
        introfile(file){
          this.intro.push(file)
          return false
        },
         toshowintroimage(url){
            // console.log(url)
            this.introimage=url
            this.showintroimage=true
        },
        removeintro(index){
            this.intro.splice(index,1)
        },
        cancel(){
            this.showadd=false
            this.pname=''
            this.pdesp=''
            this.main=null
            this.big=null
            this.price=null
            this.kucun=null
            this.kind=''
            this.intro.splice(0,this.intro.length)
        },
        done(){
            this.isload=true
            let formdata=new FormData()
            formdata.append('pname',this.pname)
            formdata.append('pdesp',this.pdesp)
            formdata.append('peizhi',this.peizhi)
            formdata.append('price',this.price)
            formdata.append('kucun',this.kucun)
            formdata.append('type',this.type)
            formdata.append('kind',this.kind)
            formdata.append('main',this.main)
            formdata.append('big',this.big)
            this.intro.forEach(e => {
                formdata.append('intro',e)
            });
            addProduct(formdata).then(res=>{
                if(res.status==1){
                    this.$Message.success("添加成功")
                    this.cancel()
                }else{
                    this.$Message.info("添加失败："+res.msg)
                    this.showadd=true
                }
            }).catch(e=>{
                this.$Message.error("服务器不给力，等会再试吧")
                this.showadd=true
            })
            this.isload=false
        },
        updateitem(index){
            this.processing=true
             let formdata=new FormData()
             formdata.append('pid',this.data[index].pid)
            formdata.append('pname',this.data[index].pname)
            formdata.append('pdesp',this.data[index].pdesp)
            formdata.append('peizhi',this.data[index].peizhi)
            formdata.append('price',this.data[index].price)
            formdata.append('kucun',this.data[index].kucun)
            formdata.append('type',this.data[index].type)
            formdata.append('kind',this.data[index].kind)
            formdata.append('status',this.data[index].status)
            if(this.main!==null){
                formdata.append('main',this.main)
            }
            if(this.big!==null){
                formdata.append('big',this.big)
            }
            this.intro.forEach(e => {
                formdata.append('intro',e)
            });
            // console.log(formdata)
            updateProduct(formdata).then(res=>{
                if(res.status==1){
                    this.$Message.success("更新成功")
                    this.cancel()
                }else{
                    this.$Message.info("更新失败："+res.msg)
                }
            }).catch(e=>{
                this.$Message.error("服务器不给力，等会再试吧")
            })
            this.processing=false
        }
    },
    mounted(){
        this.initdata()
    }

}
</script>

<style>
    .addnew{
       padding: 10px;
    }
    .smallimage{
        width:50px;
        height: 50px;
    }
    .bigimage{
        width: 600px;
    }
    .smalltd{
        width:6%;
    }
</style>