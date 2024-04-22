<script setup>
import { ElMessage, ElMessageBox } from 'element-plus';
import _ from 'lodash';
import moment from 'moment';
import List from '/@/components/List.vue';
import api from '/@/api/api';
import { getCurrentInstance, ref } from 'vue';

const { proxy } = getCurrentInstance()

const defaultQuery = {
    nickname: "",

    page: 1,
    size: 10,
    total: 0

}

const query = ref(_.cloneDeep(defaultQuery))

const reset = () => {
    query.value = _.cloneDeep(defaultQuery)
}

// - id
// - 用户名
// - 密码
// - 名称
// - 角色 （生产商、经销商，用户）
// - 状态 （启用 禁用）
// - 注册时间
const tableColProps = [
    // {
    //   "prop": "id",
    //   "label": "ID"
    // },
    {
        "prop": "Username",
        "label": "用户名"
    },
    {
        "prop": "Nickname",
        "label": "名称"
    },
    {
        "prop": "Role",
        "label": "角色",
        "type": "map",
        "map": {
            "merchant": "生产商",
            "dealer": "经销商",
            "user": "用户"
        }
    },
    {
        "prop": "Status",
        "label": "状态",
        "type": "map",
        "map": {
            "enabled": "启用",
            "disabled": "禁用"
        }
    },
    {
        "prop": "CreatedAt",
        "label": "注册时间",
        "type": "time"
    }
]

const formProps = [
    // {
    //     "prop": "Role",
    //     "label": "角色",
    //     "type": "select",
    //     "default": "user",
    //     "options": [
    //         {
    //             "label": "生产商",
    //             "value": "manufacturer"
    //         },
    //         {
    //             "label": "经销商",
    //             "value": "dealer"
    //         },
    //         {
    //             "label": "用户",
    //             "value": "user"
    //         }
    //     ]
    // },
    {
        "prop": "Status",
        "label": "状态",
        "type": "switch",
        "default": "启用",
        "activeText": "启用",
        "inactiveText": "禁用",
        "activeValue": "enabled",
        "inactiveValue": "disabled"
    },
]

const formRules = {

}

const tableData = ref([
  
])

const find = () => {
    console.log("find query: ", query.value);
    api.user.list(query.value).then((res) => {
        console.log("find res: ", res);
        tableData.value = res.data
        query.value.total = res.total
    })
}

find()


const update = (item) => {
    console.log(item);
    api.user.update(item).then((res) => {
        if (res) {
            ElMessage.success("更新成功")
            find()
        } else {
            ElMessage.error("更新失败")
        }
    })
}

</script>
<template>
    <div class="min-h-full h-full">
        <List :page="query" :tableColProps="tableColProps" :tableDate="tableData" :formProps="formProps"
            :formRules="formRules" @reset="reset" @find="find" :showOperate="true" :showUpdate="true" @update="update">
            <template #header>
                <el-form :inline="true" :model="query">
                    <el-form-item label="名称: ">
                        <el-input v-model="query.nickname"></el-input>
                    </el-form-item>
                </el-form>
            </template>
        </List>
    </div>
</template>