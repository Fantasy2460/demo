<template>
      <div style="width:400px;height:400px;position:absolute;left:550px;top:150px">
      <form  style="color:rgba(220, 228, 253, 0.942);">
        <div class="form-group" >
          <label for="exampleInputEmail1" style="color:rgba(220, 228, 253, 0.942);">请输入账号：</label>
          <div style="display: flex;">
            <input type="text" class="form-control"  id="exampleInputEmail1" placeholder="账号" v-on:input='checkname($event)' v-bind:value='username'>
            <span v-if="ckname" style="color: red;">!!!!!</span>
          </div>
        </div>
        <div class="form-group">
          <label for="exampleInputPassword1" style="color:rgba(220, 228, 253, 0.942);">密码：</label>
          <input type="password" class="form-control" id="exampleInputPassword1" placeholder="密码" v-model="passwd">
          
        </div>
        <div class="form-group" style="display: flex;">
          <div>
            <label for="exampleInputPassword1" style="color:rgba(220, 228, 253, 0.942);">验证码</label>
            <input type="text" class="form-control" id="exampleInputPassword1" placeholder="验证码" v-model="yzm">
          </div>
          <div @click="checkyz" style="margin-top:28px;font-family:楷体;margin-left:30px;background-color:white;color:black;width:100px;height:30px;text-align: center;line-height: 30px;">
        {{ yz }}
          </div>
        </div>
        <!-- <div class="form-group">
          <label for="exampleInputFile" style="color:rgba(220, 228, 253, 0.942);">File input</label>
          <input type="file" id="exampleInputFile">
          <p class="help-block"  style="color:rgba(220, 228, 253, 0.942);">Example block-level help text here.</p>
        </div> -->
        <div class="checkbox">
          <label>
            <input type="checkbox" style="color:rgba(220, 228, 253, 0.942);"> 记住我
          </label>
        </div>
        <el-button plain @click="sub">登录</el-button>
        <div class="form-group" style="margin-top:10px">
          <label class="exampleInputEmail1">
            没有账号？
            <router-link to="/Register" class="xx1">
              点我注册
            </router-link>
          </label>
        </div>
      </form> 
    </div> 
</template>

<script lang="ts" setup>  
import { ElNotification } from 'element-plus'
import { ref, onMounted ,h,inject} from 'vue';  
import { defineExpose } from 'vue'; // 如果你需要在模板外的地方访问组件内部变量或方法，可以使用 defineExpose  
import SideBar from './SideBar.vue';  
import Demo from './Demo.vue';  
import axios from 'axios';  
import bus from "../EventBus/eventbus"; // 请注意确保你的 eventbus 正确工作并支持你想要的功能  
import { useRouter } from 'vue-router' 
import { useUserStore } from '../store/user';
import service from '../axios-instance';
const $MYGO = inject('$MYGO', '');
const userStore=useUserStore()
const router = useRouter()  

const yz=ref(localStorage.getItem('yz'))
const yzm=ref('')
function checkyz(){
  // service.get($MYGO+'/user/createYz')
  // .then(res=>{
  //   localStorage.setItem('yz',res.data.yz)
  //   yz.value=res.data.yz
  // })
}
// 响应式状态  
const username = ref('');  
const passwd = ref('');  
const ckname = ref(false);  
  
// 方法  
function sub() {  
  localStorage.setItem("username", username.value);  
  axios.post($MYGO+'/user/login', {  
    userName: username.value,  
    password: passwd.value
  })  
  .then(response => {
    
    localStorage.setItem('token',response.data.data.token)
    localStorage.setItem("name",username.value);
    ElNotification({
      title: 'Success',
      message: '登录成功',
      type: 'success',
    }) 
    router.push('person')
    
  })  
  .catch(err => {  
    checkyz()
    // 注意：在 <script setup> 中，你可能需要自定义通知函数或使用其他库/插件来显示通知  
    // 这里只是一个示例，你可能需要调整  
    ElNotification({
      title: 'Error',
      message: err.response.data.msg,
      type: 'error',
    })
  });  
}  
function checkname(event) {  
  username.value = event.target.value;  
  if (username.value !== '' && (username.value.length < 5 || username.value.length > 20)) {  
    ckname.value = true;  
  } else {  
    ckname.value = false;  
  }  
}  
  
// 组件挂载后执行的操作  
onMounted(() => {  
  username.value = localStorage.getItem('username') || '';
  checkyz()
});  
  
// 如果你需要在模板外的地方访问组件内部变量或方法，可以使用 defineExpose  
defineExpose({  
  // 可以选择暴露给外部的方法或变量  
  // 例如：sub, checkname, ...  
});  
</script> 


<style>
.xx1{
  color:rgba(220, 228, 253, 0.942);
}

.xx1:hover{
  color:rgb(98, 98, 211);
}
</style>