<template>
    <div style="color:rgba(220, 228, 253, 0.942);">
      <div style="display:flex">
        <div style="border:1px;height:600px;width:1px;float:left"></div>
        <div>
          <div class="Message" >
            <div>
              <div style="margin-left:20px;line-height:20px">{{ group.groupName?group.groupName:'' }}</div>
              <button style="float:right;margin-right:20px;margin-top:-10px;background-color:rgb(105, 105, 105);border:0px;" @click="drawer1=true">···</button>
            </div>
            <hr>
            <div class="Top" style="width:auto">
              <el-scrollbar style="width:607px;height:345px;margin-top:-10px" ref="scrollbarRef" always>
                <div ref="innerRef">
                  <p v-for="(message,i) in groupMessages" 
                  :key="i" 
                   :class="getMessageClass(message.receiverUserId==message.senderUserId,message.isWithdrawn)">
                   <div v-if="message.isWithdrawn" style="display: flex;">
                    {{ message.receiverUserId==message.senderUserId?'您已撤回一条消息':message.senderUserName+'已撤回一条消息' }}
                    <div v-if="message.standby" class="mb-4">
                      <el-button
                        type="danger"
                        text
                        @click="checkstandby(message.standby)"
                      >
                       重新编辑
                      </el-button>
                    </div>
                    <div @click="messagedelete(message.id,i)">
                      <img :src="$MYGO + '/static/images/close.png'"/>
                    </div>
                   </div>
                   <div v-else-if="message.receiverUserId==message.senderUserId" style="display: flex;">
                        <div style="width:300px;height:1px"></div>
                      <div class="bubble">
                        <div class="message">
                          <el-popover :visible="message.visible?message.visible:false" placement="top" :width="185">
                            <div style="text-align: right; margin: 0;">
                              <div style="display:flex">
                              <el-button size="small" type="primary" style="background-color:rgb(207, 230, 244)" text @click="message.visible = false">取消</el-button>
                              <el-button size="small" type="primary" @click="revocation(message.id)">
                                撤回
                              </el-button>
                              <el-button size="small" type="primary" @click="messagedelete(message.id,i)">
                                删除
                              </el-button>
                              </div>
                            </div>
                            <template #reference>
                              <el-button @click="message.visible = true">{{ message.message }}</el-button>
                            </template>
                          </el-popover>
                        </div>
                      </div>
                      <div class="avatar">
                        {{ message.receiverUserName }}
                        <img :src="message.receiverUserImg" class="avatar-image" style="margin-right:20px" @click="checknwgroupuser(message.senderUser)"/>
                      </div>
                    </div>
                    <div v-else  style="display: flex;">
                      <div class="avatar">
                        <img :src="message.receiverUserImg" class="avatar-image" style="margin-left:20px" @click="checknwgroupuser(message.senderUser)"/>
                        {{ message.receiverUserName }}
                      </div>
                      <div class="bubble">
                        <div class="message">
                          <el-popover :visible="message.visible?message.visible:false" placement="top" :width="160">
                            <div style="text-align: right; margin: 0;">
                              <el-button size="small" type="primary" style="background-color:rgb(207, 230, 244)" text @click="message.visible = false">取消</el-button>
                              <el-button size="small" type="primary" @click="messagedelete(message.id,i)">
                                删除
                              </el-button>
                            </div>
                            <template #reference>
                              <el-button @click="message.visible = true">{{ message.message }}</el-button>
                            </template>
                          </el-popover>
                        </div>
                      </div>
                        <div style="width:300px;height:1px"></div>
                    </div>
                  </p>
                </div>
              </el-scrollbar>
            </div>
          </div>
          <div style="width:1px; background-color: black;"></div>
          <div class="Chat">
            <div style="height:10px;width:1px"></div>
            <textarea style="width:590px;height:120px;background-color:rgb(141, 141, 141);border:1px double black;margin-left:10px" v-model="message" @input="handleInput"></textarea>
            <button style="color:rgba(220, 228, 253, 0.942);background-color:#82838372;width:60px;height:30px;margin-left:500px" @click="send">发送</button>
          </div>
        </div>
      </div>
    </div>
    <div>
      <el-drawer v-model="drawer1" direction="rtl">
        <template #header>
          <div style="display: flex;">
            <div style="font-size: 20px;margin-right: 50px;">头像: </div> 
            <el-upload
              class="avatar-uploader"
              :show-file-list="false"
              :http-request="httpRequest"
              :before-upload="beforeImageUpload"
            >
              <img v-if="group.groupAvatar" :src="group.groupAvatar" style="height:100px;width:100px;border-radius:50%"   />
              <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
            </el-upload>
            <div style="font-size: 20px;margin-right: 50px;">{{ group.groupName }} </div> 
          </div>
        </template>
        <template #default>
          <div>群号：{{ group.groupNumber }}</div>
          <div>群名：{{ group.groupName }}</div>
          <div>群公告：{{ group.groupAnnouncement }}</div>
          <!-- <div>我的群备注：{{ usertoUsers[index].remarks }}</div> -->
          <div style="display: flex;">
            <div v-for="(item,i) in users" :key="i" >
              <img :src="item.user.img" @click="checknwgroupuser(item)" style="width: 50px;"/>{{ item.text }}
            </div>
          </div>
          <div>其他：</div>
          <el-button @click="toSpace(nwgroupuser.userId)" type="primary" style="background-color:rgb(207, 230, 244);margin-top:20px;margin-left:10px" text >访问空间</el-button>
        </template>
        <template #footer>
          <div style="flex: auto">
            <el-button @click="drawer1 = false">取消</el-button>
            <el-button type="primary" @click="confirmClick">保存</el-button>
          </div>
        </template>
      </el-drawer>
      <el-drawer v-model="drawer2" direction="ltr" v-if="nwgroupuser.userId==user.userId">
        <template #header>
          <h4><img :src="nwgroupuser.img" style="width: 50px;border-radius:25px"/> {{ nwgroupuser.text }} </h4>
        </template>
        <template #default>
          <div>账号：{{ nwgroupuser.remark }}</div>
          <div>账号名：{{ nwgroupuser.account }}</div>
          <div>群备注：可更改</div>
          <div>权限：
            <span v-if="nwgroupuser.permissionLevel==0">群用户</span>
            <span v-if="nwgroupuser.permissionLevel==1">群管理</span>
            <span v-if="nwgroupuser.permissionLevel==2">群主</span>
          </div>
          <div>其他：</div>
          <el-button @click="toSpace(nwgroupuser.userId)" type="primary" style="background-color:rgb(207, 230, 244);margin-top:20px;margin-left:10px" text >访问空间</el-button>
        </template>
        <template #footer>
          <div style="flex: auto">
            <el-button @click="drawer2 = false">取消</el-button>
            <el-button type="primary" @click="confirmClick">保存</el-button>
          </div>
        </template>
      </el-drawer>
      <el-drawer v-model="drawer2" direction="ltr" v-if="nwgroupuser.userId!=user.userId">
        <template #header>
          <h4><img :src="nwgroupuser.user.img" style="width: 50px;border-radius:25px"/> {{ nwgroupuser.text }} </h4>
        </template>
        <template #default>
          <div>账号：{{ nwgroupuser.remark }}</div>
          <div>账号名：{{ nwgroupuser.account }}</div>
          <div>群备注：不更改</div>
          <div>权限：
            <span v-if="nwgroupuser.permissionLevel==0">群用户</span>
            <span v-if="nwgroupuser.permissionLevel==1">群管理</span>
            <span v-if="nwgroupuser.permissionLevel==2">群主</span>
          </div>
          <div>其他：</div>
           <el-button @click="toSpace(nwgroupuser.userId)" type="primary" style="background-color:rgb(207, 230, 244);margin-top:20px;margin-left:10px" text >访问空间</el-button>
        </template>
        <template #footer>
          <div style="flex: auto">
            <el-button @click="drawer2 = false">取消</el-button>
            <el-button type="primary" @click="confirmClick">保存</el-button>
          </div>
        </template>
      </el-drawer>
    </div>
    
  </template>
  <script lang="ts" setup>
  import ImageViewer from "@luohc92/vue3-image-viewer";
  import '@luohc92/vue3-image-viewer/dist/style.css';
  import { ref, onMounted ,h,reactive,nextTick,inject,watch } from 'vue'; 
  import { ElNotification,ElScrollbar } from 'element-plus'
  import service from '../axios-instance'
  import axios from "axios";
  import { useRouter } from 'vue-router' 
  import { useWsStore } from '../store/user';
  import { da } from "element-plus/es/locale";
  const $MYGO = inject('$MYGO', '');
  const wsStore=useWsStore()
  const router = useRouter()
  var user = reactive(JSON.parse(localStorage.getItem('user')))
  const $Ws: ((data) => string) | undefined = inject('$Ws')
  
  let message=ref('')
  function handleInput() {  
      // 检查inputValue是否包含换行符  
      if (message.value.includes('\n')) {  
        // 如果包含换行符，则执行你的函数  
        send();  
      }  
  }
  let nwgroupuser=reactive({
    user:{},
    groupId: 4,
    id: 0,
    isAdmin:0,
    isGag: false,
    text:'',
    userId:0,
  })
  function toSpace(e)
  {
    localStorage.setItem("toId",e)
    router.push("/toSpace")
    console.log(e)
  }
  let drawer1=ref(false)
  let drawer2=ref(false)
  function checknwgroupuser(item){
    nwgroupuser.user=item.user
    nwgroupuser.groupId=item.groupId
    nwgroupuser.id=item.id
    nwgroupuser.isAdmin=item.isAdmin
    nwgroupuser.isGag=item.isGag
    nwgroupuser.text=item.text
    nwgroupuser.userId=item.userId
  
    drawer2.value=true
  }
  const uploadUrl=ref($MYGO+'/group/updateImg')
  const imageUrl = ref('')
  function httpRequest(option){
    let dataForm = new FormData();
    dataForm.append('file',option.file)
    dataForm.append('uid',option.file.uid)
    dataForm.append('id',groups[index.value].id)
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
          groups[index.value].img=response.data.img
          console.log(response.data)
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
  
  function send(){
    $Ws && $Ws({
              groupId: groups[index.value].id-0,
              message:message.value,
              event:'/group/sendMessage',
              token:localStorage.getItem('token')
          })
    message.value=''
  }
  function messagedelete(id,j){
    groupMessages[j].visible=false
    service.post($MYGO+'/group/deleteMessage',{
      'id': id-0,
    }).then(res=>{
      groupMessages.splice(j,1)
    }).catch(err=>{
      console.error(err)
      ElNotification({
          title: 'Error',
          message: err.response.data.errorMessage,
          type: 'error',
        })
    })
  }
  function revocation(id){
    service.post($MYGO+'/group/revocationMessage',{
      'id': id-0,
    }).then(res=>{
      
    }).catch(err=>{
      console.error(err)
      ElNotification({
          title: 'Error',
          message: err.response.data.errorMessage,
          type: 'error',
        })
    })
  }
  
  //消息框样式动态选择
  const getMessageClass = (isSent,isDeleted) => {
    if(isDeleted){
      return 'message-container-centre'
    }
    return isSent ? 'message-container-right' : 'message-container-left';
  };
  
  
  var groups=reactive([{groupName:''}])
  
  const innerRef = ref<HTMLDivElement>()
  const scrollbarRef = ref<InstanceType<typeof ElScrollbar>>()
  function gobottom(){//抵达最底部
    nextTick(() => {  
      scrollbarRef.value!.setScrollTop(20000)
    })
  }
  function getgroups(){
    console.log('发送请求')
     service.get($MYGO+'/group/fidMessage/'+localStorage.getItem('groupId')+'/1/100')
     .then(res=>{
      console.log(res.data)
      messages.pop()
      let i=0
      res.data.forEach(element => {
        messages.push(element)
        getcount(i)
        i++
      });
      gobottom()
     }).catch(err=>{
        console.error(err)
        let data=err.response.data
        if(data.type=='token'){
          localStorage.removeItem('token')
        }
        ElNotification({
          title: 'Error',
          message: err,
          type: 'error',
        })
     })
  }
  onMounted(() => {
    wsStore.event=1
    getgroups()
  })
  </script>
  <style scoped>
  .scrollbar-demo-item {
    display: flex;
    align-items: center;
    height: 50px;
    margin: 25px;
    text-align: center;
    border-radius: 4px;
    background: var(--el-color-primary-light-9);
    color: var(--el-color-primary);
    width: 100%;
  }
  </style>
  <style>
  
  .avatar {
    margin-left: 10px; /* 修改这里将头像放在消息框的右边 */
  }
   
  .avatar-image {
    width: 40px;
    height: 40px;
    border-radius: 50%;
    object-fit: cover;
  }
   
  .bubble {
    color: #000;
    padding: 10px;
    border-radius: 5px;
  }
  .message {
    text-align: left;
    margin: 0;
  }
  .message-container {
    display: flex;
    align-items: center;
    margin-bottom: 10px;
  }
  .message-container-centre{
    float: left;
    margin-left: 40%;
  }
  .message-container-right {
    float:right;
    /*justify-content: flex-end;*/
  }
   
  .message-container-left {
    float:left;
    /*justify-content: flex-start;*/
  }
  .List{
    width:310px;
    height:596px;
    border-top-left-radius: 18px;
    border-bottom-left-radius: 18px;
    background-color:#666464;
  }
  
  .Message{
    width:607px;
    height:420px;
    border-top-right-radius: 18px;
    background-color: #82838372;
  }
  .friend{
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    width: 200px;
  }
  .friend:hover{
    background-color:#cdcdcda2;
  }
  
  .Chat
  {
    height:176px;
    width: 607px;;
    border-bottom-right-radius: 18px;;
    background-color: #8e8f8f;
  }
  
  .find{
      height: 32px;
      width: 200px;
      background-color: #525257a2;
    
  }
  
  .bTn{
    margin-top:14px;
    color: rgb(207, 230, 244);
    font-size:17px;
    background-color: #2d9cf8;
    text-align: center;
    border-left:1px solid rgba(0, 0, 0, 0.45);
    width:50px;
    height: 34px;
    line-height: 34px;
    border-top-right-radius: 18px;
    border-bottom-right-radius: 18px
  }
  
  .inputT{
      margin-left: 210px;
      margin-top:15px;
    width:150px;
    height: 32px;
    text-indent: 1em;
     line-height: 34px;
    border: 1px solid black;
    border-top-left-radius: 17px;
    border-bottom-left-radius: 17px;
    font-size: 12px;
  }
  .unread-indicator {  
    /* 红点的样式 */  
    position: relative;  
    display: inline-block;  
    margin-left: 10px;
    /* 其他样式... */  
  }  
    
  .unread-count {  
    /* 数字的样式 */  
    position: absolute;  
    background-color: red;
    border-radius:100%;
    width: 20px;
    height: 20px;
    top: -30px; /* 假设你想要将数字放在红点的上方 */  
    right: -15px; /* 假设你想要将数字放在红点的右侧 */  
    text-align: center;
    line-height: 20px;
    /* 其他样式... */  
  } 
  </style>
  