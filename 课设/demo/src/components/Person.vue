<template>
  <div v-if="user">
    <ul style="color:rgb(207, 234, 244);margin-left:150px;margin-top:80px;font-size:25px">
      <div style="display:flex;">
        <el-upload
          class="avatar-uploader"
          :show-file-list="false"
          :http-request="httpRequest"
          :before-upload="beforeImageUpload"
        >
        <img v-if="user.img" :src="toImg(user.img)" style="height:100px;width:100px;border-radius:50%"  />
          <el-icon v-else class="avatar-uploader-icon" style="height:70px;width:70px;border-radius:50%;background-color:white"><Plus />头像</el-icon>
        </el-upload>
      </div>
      <li style="margin-top:20px">账号：<input style="margin-left:50px;background-color:rgb(105, 105, 105);border:0px" type="text" v-model="user.account"/></li>
      <li style="margin-top:10px">用户名：<input style="margin-left:26px;background-color:rgb(105, 105, 105);border:0px" type="text" v-model="user.username"/></li>
      <li style="margin-top:10px">个性签名：<input style="background-color:rgb(105, 105, 105);border:0px" type="text" v-model="user.signed"/></li>
      <li style="margin-top:10px">邮箱：<input style="margin-left:50px;background-color:rgb(105, 105, 105);border:0px" type="text" v-model="user.email"/></li>
      <li style="margin-top:10px">生日：<input style="margin-left:50px;background-color:rgb(105, 105, 105);border:0px" type="text" v-model="user.birthday"/></li>
      <li style="margin-top:10px">居住地：<input style="margin-left:24px;background-color:rgb(105, 105, 105);border:0px" type="text" v-model="user.city"/></li>
    </ul>
    
    <div>
      <button @click="save()" style="margin-top:20px;margin-left:300px;width:110px;height:40px;font-size:20px;line-height:40px;border-radius:10px">保存</button>
    </div>
  </div>
</template>
<script setup>
import {ref,reactive,onMounted,inject} from 'vue'
import { ElNotification } from 'element-plus'
import service from '../axios-instance'
import axios from "axios";
import { useWsStore } from '../store/user';
const $MYGO = inject('$MYGO', '');
const wsStore=useWsStore()
// const user=reactive(JSON.parse(localStorage.getItem('user')))
var user = reactive({})
function toImg(url){
  return $MYGO+"/"+url
}
onMounted(()=>{
  wsStore.event=-1
  getuser()
  // service.post($MYGO+'/user/fidUser',{"username":localStorage.getItem("username")})
  // .then(tmp=>{
  //   user.username=tmp.data.username
  //   user.account=tmp.data.account
  //   user.signed=tmp.data.signed
  //   user.email=tmp.data.email
  //   user.pht=tmp.data.img
  //   user.birthday=tmp.data.birthday
  //   user.city=tmp.data.city
  // })
})
function getuser(){
  service.get($MYGO+'/user/info')
  .then(res=>{
    var item=res.data.data.item
    localStorage.setItem('user',JSON.stringify(item))
    user.username=item.username
    user.account=item.account
    user.email=item.email
    user.img=item.img
  }).catch(err=>{
    console.error(err)
  })
}
function save()
{
  console.log(user)
  service.post($MYGO+'/user/update',{
    'username':user.username
  })
  .then(res=>{
    localStorage.setItem('user',JSON.stringify(user))
    ElNotification({
      title: 'Success',
      message: '修改成功',
      type: 'success',
    }) 
  }).catch(err=>{
    ElNotification({
      title: 'Error',
      message: '修改失败',
      type: 'error',
    })
  })
}
const uploadUrl=ref($MYGO+'/user/upload')
function httpRequest(option){
  let dataForm = new FormData();
  dataForm.append('file',option.file)
  dataForm.append('uid',option.file.uid)
  console.log('uid:',option.file.uid)
  axios({
        method: 'POST',
        url: uploadUrl.value,
        data: dataForm,
//设置请求参数的规则
        headers: {
            "Content-Type": "multipart/form-data",
            "Authorization":localStorage.getItem('token')
        }
    }).then(response => {
        console.log(response.data)
        user.img=response.data.data.url
    }).catch(err=>{
      console.log(err)
    })

}

function beforeImageUpload(rawFile){
    if(rawFile.size / 1024 / 1024 > 10){
        ElMessage.error("单张图片大小不能超过10MB!");
        return false;
    }
    return true;
}
</script>
<style>
.disUoloadBtn .el-upload--picture-card{
    display:none;   /* 上传按钮隐藏 */
}
</style>