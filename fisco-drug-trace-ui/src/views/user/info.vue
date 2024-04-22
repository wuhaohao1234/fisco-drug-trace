<script setup>
import { ref } from 'vue';
import { useInfoStore } from '/@/store/info';
import api from '/@/api/api';
import { ElMessage } from 'element-plus';
const infoStore = useInfoStore()
const infoRef = ref()
const infoForm = ref({
    nickname: infoStore.nickname,
})
const infoRules = ref({
    username: [
        { required: true, message: '请输入名称', trigger: 'blur' },
    ],

})
const save = async () => {
    if (infoRef.value.validate()) {
        api.user.updateMe({
            nickname: infoForm.value.nickname
        }).then((res) => {
            if (res) {
                ElMessage.success('保存成功');
                infoStore.setNickname(infoForm.value.nickname)
            } else {
                ElMessage.error('保存失败');
            }
        })
    }
}

const getRoleName = () => {
    if (infoStore.role == 'admin') {
        return '管理员'
    } else if (infoStore.role == 'merchant') {
        return '生产商'
    } else if (infoStore.role == 'dealer') {
        return '经销商'
    } else {
        return '用户'
    }

}

</script>
<template>
    <div class="w-3/5 " style="margin: 0 auto;">
        <el-card>
            <template #header>
                <div class="text-center">
                    <h1>个人信息</h1>
                    <p>{{ getRoleName() }}</p>
                </div>
            </template>
            <div class="flex justify-center">
                <el-form state-icon ref="infoRef" label-width="0px" class="demo-ruleForm" style="width:300px">
                    <el-form-item prop="username" :rules="infoRules" @keyup.enter.native="info">
                        <el-input :prefix-icon="User" v-model="infoForm.nickname" placeholder="请输入名称" autofocus="off"
                            autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <div class="flex justify-between w-full">
                            <el-button type="primary" class="w-full" size="large" @click="save">保存</el-button>
                        </div>
                    </el-form-item>
                </el-form>
            </div>
        </el-card>
    </div>
</template>