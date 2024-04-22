<script setup>
import { reactive, ref } from 'vue';
import { ElMessage, ElMessageBox } from 'element-plus';
import { useRouter } from "vue-router";
import api from "/@/api/api";
const router = useRouter()
import { useInfoStore } from '/@/store/info';
const infoStore = useInfoStore()
const formData = ref({
    username: "",
    password: "",
    role: "",
})
const registFlag = ref(false)
const formRef = ref(null)
const formRules = ref({
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    // { min: 5, max: 30, message: '长度在 5 到 20 个字符', trigger: 'blur' }
  ],
  password: [
  { required: true, message: '请输入密码', trigger: 'blur' },
    // { min: 8, max: 20, message: '长度在 8 到 20 个字符', trigger: 'blur' }
  ],
  nickname: [
    { required: true, message: '请输入名称', trigger: 'blur' },
    // { min: 2, max: 20, message: '长度在 2 到 20 个字符', trigger: 'blur' }
  ]
})



const login = async () => {
    formRef.value.validate((valid) => {
        if (valid) {
            api.user.login(formData.value).then((res) => {
                if(res){
                    ElMessage.success('登录成功');
                    infoStore.setNickname(res.nickname)
                    infoStore.setRole(res.role)
                    if(res.role == 'admin'){
                        router.push({path: '/user/list'})
                    }else if (res.role == 'merchant'){
                        router.push("/drug/manufacturer")
                    }else if (res.role == 'dealer'){
                        router.push("/drug/dealer")
                    }else {
                        router.push("/drug/user")
                    }
                }
            })
        } else {
            console.log('error submit!!');
            return false;
        }
    });
}

const register = async () => {
    if(registFlag.value){
        formRef.value.validate((valid) => {
            if (valid) {
                api.user.register(formData.value).then((res) => {
                    if(res){
                        ElMessage.success('注册成功');
                        registFlag.value = false
                    }
                })
            } else {
                console.log('error submit!!');
                return false;
            }
        });
    }else{
        registFlag.value = true
    }

}

</script>
<template>
    <div class="wrapper">
        <div class="content">
            <div class="title">
                <h1>基于区块链的药品溯源系统</h1>
            </div>
            <el-form ref="formRef" :model="formData" :rule="formRules" label-width="80px" >
                <el-form-item label="用户名: " prop="username" class="item" >
                    <el-input v-model="formData.username" style="width:400px"/>
                </el-form-item>
                <el-form-item label="密码: " prop="password" class="item">
                    <el-input v-model="formData.password" type="password" style="width:400px"/>
                </el-form-item>
                <el-form-item label="名称: " prop="nickname" class="item" v-if="registFlag" >
                    <el-input v-model="formData.nickname" style="width:400px"/>
                </el-form-item>
                <el-form-item label="角色: " prop="role" class="item" v-if="registFlag">
                    <el-radio-group v-model="formData.role">
                        <el-radio label="manufacturer" style="color: black;">生产商</el-radio>
                        <el-radio label="dealer" style="color: black;">经销商</el-radio>
                        <el-radio label="user" style="color: black;">用户</el-radio>
                    </el-radio-group>
                </el-form-item>

            </el-form>
            <div class="flex justify-center mt-2">
                <el-button class="w-40" type="primary" @click="register">注册</el-button>
                <el-button class="w-40" type="primary" @click="login">登录</el-button>
            </div>
        </div>

    </div>
</template>
<style lang="scss" >
.wrapper {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    background-image: url('/@/assets/bg.jpg');
    background-size: cover;

    .content {
        background-color: rgba(200,
                200,
                200,
                0.3);
        border-color: rgba(200, 200, 200, 0.5);
        border-radius: 18px;
        padding: 40px;
    }

    .title {
        text-align: center;
        color: black;
        margin-bottom: 1em;
    }


}
.item .el-form-item__label{
    color: black; 
}
</style>