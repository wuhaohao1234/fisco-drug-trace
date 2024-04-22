<template>
  <el-container class="h-full min-h-screen bg-color-#fbffff">
    <div class="h-full min-h-screen flex flex-col bg-color-#EFFFFF color-white aside-wrapper w-180px"
      style="min-width: 180px;">
      <div class="h-15 flex justify-center">
        <img :src="logo" class="w-120px" />
      </div>
      <div class="menu-wrapper">
        <div :class="`menu-item ${currentRoute.name == 'UserList' ? 'menu-item-active' : ''}`"
          @click="go('/user/list')" v-if="infoStore.getRole == 'admin'">用户管理</div>
          <div :class="`menu-item ${currentRoute.name == 'Info' ? 'menu-item-active' : ''}`"
          @click="go('/info')">个人信息</div>
        <div :class="`menu-item ${currentRoute.name == 'DrugManufacturer' ? 'menu-item-active' : ''}`"
          @click="go('/drug/manufacturer')" v-if="infoStore.getRole == 'merchant'">药品列表</div>
        <div :class="`menu-item ${currentRoute.name == 'DrugDealer' ? 'menu-item-active' : ''}`"
          @click="go('/drug/dealer')" v-if="infoStore.getRole == 'dealer'">药品销售</div>
        <div :class="`menu-item ${currentRoute.name == 'DrugUser' ? 'menu-item-active' : ''}`"
          @click="go('/drug/user')" v-if="infoStore.getRole == 'user'">药品购买</div>
          <div :class="`menu-item ${currentRoute.name == 'DrugBuy' ? 'menu-item-active' : ''}`"
          @click="go('/drug/buy')" v-if="infoStore.getRole == 'user'">我购买得药品</div>
      </div>
    </div>
    <el-container>
      <el-header class="p-2" height="56px">
        <div class="flex h-full w-full justify-between items-center">
          <div style="font-size: 2rem;">基于区块链的药品溯源系统</div>
          <el-dropdown @command="handleCommand">
            <span class="el-dropdown-link flex items-center">
              <el-icon>
                <user-filled />
              </el-icon>
              {{ infoStore.getNickname }}
              <el-icon class="el-icon--right">
                <arrow-down />
              </el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="changePassword">修改密码</el-dropdown-item>
                <el-dropdown-item command="logout">登出</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>
      <el-main class="w-full bg-color-#ebf3f5" style="padding: 0;">
        <div class="content-wrapper p-5">
          <router-view v-slot="{ Component, route }">
            <transition name="fade-transform" mode="out-in">
              <component :is="Component" :key="route.path" />
            </transition>
          </router-view>
        </div>

      </el-main>
    </el-container>
    <el-dialog v-model="changePasswordVisiable" title="修改密码" width="400px">
      <el-form :model="form" ref="formRef" label-width="80px" label-position="left">
        <el-form-item label="旧密码: " prop="oldPwd">
          <el-input v-model="form.oldPwd" autocomplete="off" type="password" />
        </el-form-item>
        <el-form-item label="新密码: " prop="pwd">
          <el-input v-model="form.pwd" autocomplete="off" type="password" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button type="primary" @click="changePassword">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </el-container>
</template>

<script setup>
import { getCurrentInstance, ref } from 'vue';
import logo from '/@/assets/logo.svg';
import api from '/@/api/api';
import { useInfoStore } from '/@/store/info';
const infoStore = useInfoStore()
const { proxy } = getCurrentInstance()

import { useRouter } from 'vue-router';
const { currentRoute, push } = useRouter()


const go = (url) => {
  push(url)
}

const username = ref(infoStore.username)

const changePasswordVisiable = ref(false)

const form = ref({
  oldPwd: "",
  pwd: ""
})

const changePassword = () => {
  if (form.value.pwd == "") {
    ElMessage.error("密码不能为空")
    return
  }
  api.user.updateMe({
    oldPassword: form.value.oldPwd,
    password: form.value.pwd
  }).then((res) => {
    if (res) {
      ElMessage.success("修改成功")
      changePasswordVisiable.value = false
    } else {
      ElMessage.error("修改失败")
    }
  })
}

const handleCommand = (command) => {
  if (command == 'logout') {
    infoStore.clearAll()
    push("/login")
  }
  if (command == 'changePassword') {
    changePasswordVisiable.value = true
    form.value = {
      oldPassword: "",
      newPassword: "",
      newPasswordBackup: "",
    }
  }
}
</script>

<style lang="scss" scoped>
.aside-wrapper {
  box-sizing: border-box;
  border-top-right-radius: 32px;
  box-shadow: 0 1px 2px #00000045 !important;
  border-right: solid 1px rgba(0, 0, 0, .12);
  overflow: visible;
  /* 或者其他合适的值 */
  z-index: 2;
}

.menu-wrapper {
  color: #0b163e;
  height: calc(100vh - 2.75rem);
  padding: 0 .5rem;
}

.menu-item {
  padding: .7rem;
  border-radius: 4px;
  margin-bottom: .25rem;
  cursor: pointer;
}

.menu-item-active {
  background-color: #fbffff;
}

.menu-item:hover {
  background-color: #fbffff;

}

.nav-item {
  display: flex;
  align-items: center;
  height: 100%;
  margin: 0 .5rem;
}

::v-deep .el-progress__text {
  min-width: 30px;
}

.content-wrapper {
  min-height: calc(100vh - 80px);
}
</style>