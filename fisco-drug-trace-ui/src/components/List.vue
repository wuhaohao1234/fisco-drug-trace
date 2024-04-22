<template>
    <div class="p-10 bg-white min-h-full rounded">
        <el-card class="content-query">
            <slot name="header"></slot>
        </el-card>
        <div class="flex-end">
            <el-button @click="reset">重置</el-button>
            <el-button type="primary" @click="find">查询</el-button>
            <el-button type="danger" @click="openAddForm" v-if="showAdd">添加</el-button>
            <slot name="btns"></slot>
        </div>
        <div class="content">
            <slot name="table">
                <el-table :data="tableDate" stripe :border="true" style="width: 100%" >
                    <template v-for="(itemProps, index) in tableColProps" :key="index">
                        <el-table-column v-bind="itemProps" align="center" :width="itemProps.width ? itemProps.width : 'auto' ">
                            <template #default="scope">
                                <slot :name="itemProps.slotName ? itemProps.slotName : itemProps.prop" :scope="scope"
                                    :row="scope.row" :index="scope.$index">
                                    <template v-if="!itemProps.type">
                                        {{
                scope.row[itemProps.prop], itemProps.prop
            }}
                                    </template>
                                    <template v-if="itemProps.type === 'time'">
                                        {{
                scope.row[itemProps.prop] != null ?
                    moment(scope.row[itemProps.prop]).format('YYYY-MM-DD HH:mm:ss') : ""
            }}
                                    </template>
                                    <template v-if="itemProps.type === 'map'">
                                        {{
                itemProps.map[scope.row[itemProps.prop]]
            }}
                                    </template>
                                    <template v-if="itemProps.type === 'tag'">
                                        <el-tag>
                                            {{
                scope.row[itemProps.prop]
            }}
                                        </el-tag>
                                    </template>
                                </slot>
                            </template>
                        </el-table-column>
                    </template>
                    <el-table-column label="操作" align="center" :width="operateWidth" v-if="showOperate">
                        <template #default="scope">
                            <el-button type="warning" @click="openUpdateForm(scope.row)"
                                v-if="showUpdate">编辑</el-button>
                            <el-popconfirm v-if="showDelete" title="请确定是否要删除这条数据?" @confirm="handleDelete(scope.row)">
                                <template #reference>
                                    <el-button type="danger">删除</el-button>
                                </template>
                            </el-popconfirm>
                            <slot name="handler" :scope="scope" :row="scope.row" :index="scope.$index"></slot>
                        </template>
                    </el-table-column>
                </el-table>
            </slot>
            <div class="flex justify-center mt-8">
                <el-pagination v-model:currentPage="page.page" :page-size="page.pageSize" layout="prev, pager, next"
                    :total="page.total" @current-change="find"></el-pagination>
            </div>
        </div>
        <el-dialog v-model="formVisiable" :title="title" width="680px">
            <slot name="form">
                <el-form :model="formData" ref="formRef" :rules="formRules" label-position="left" label-width="100px">
                    <template v-for="(itemProps, index) in formProps" :key="index">
                        <slot :name="itemProps.slotName ? itemProps.slotName : itemProps.prop" :index="index"
                            :formData="formData">
                            <el-form-item :label="itemProps.label" :prop="itemProps.prop">
                                <template v-if="!itemProps.type">
                                    <el-input v-model="formData[itemProps.prop]"></el-input>
                                </template>
                                <template v-if="itemProps.type === 'textarea'">
                                    <el-input v-model="formData[itemProps.prop]" type="textarea" autosize></el-input>
                                </template>
                                <template v-if="itemProps.type === 'time'">
                                    <el-date-picker v-model="formData[itemProps.prop]" :disabled-date="disabledDate"
                                        type="datetime"></el-date-picker>
                                </template>
                                <template v-if="itemProps.type === 'date'">
                                    <el-date-picker v-model="formData[itemProps.prop]" 
                                        type="date" ></el-date-picker>
                                </template>
                                <template v-if="itemProps.type === 'select'">
                                    <el-select v-model="formData[itemProps.prop]" placeholder="请选择">
                                        <el-option v-for="(item, index) in itemProps.options" :key="index"
                                            :label="item.label" :value="item.value">
                                        </el-option>
                                    </el-select>
                                </template>
                                <template v-if="itemProps.type === 'switch'">
                                    <el-switch v-model="formData[itemProps.prop]" inline-prompt
                                        :active-text="itemProps.activeText" :inactive-text="itemProps.inactiveText"
                                        :active-value="itemProps.activeValue" :inactive-value="itemProps.inactiveValue" />
                                </template>
                                <template v-if="itemProps.type === 'tag'">
                                    <div class="flex gap-2">
                                        <el-tag v-for="tag in dynamicTags" :key="tag" closable
                                            :disable-transitions="false" @close="handleClose(tag)">
                                            {{ tag }}
                                        </el-tag>
                                        <el-input v-if="inputVisible" ref="InputRef" v-model="inputValue" class="w-20"
                                            size="small" @keyup.enter="handleInputConfirm" @blur="handleInputConfirm" />
                                        <el-button v-else class="button-new-tag" size="small" @click="showInput">
                                            + 投票项
                                        </el-button>
                                    </div>
                                </template>

                            </el-form-item>
                        </slot>
                    </template>
                </el-form>
            </slot>
            <div class="flex flex-row-reverse">
                <el-button type="primary" @click="emitSubmit">确认</el-button>
            </div>
        </el-dialog>
    </div>
</template>
<script setup>
import { defineEmits, defineProps, getCurrentInstance, ref } from 'vue';
import moment from 'moment';
const { proxy } = getCurrentInstance()
const emit = defineEmits(['reset', 'find', 'add', 'update'])


const props = defineProps({
    page: {
        type: Object,
        default: () => {
            return {
                page: 1,
                pageSize: 10,
                total: 0
            }
        }
    },
    showAdd: {
        type: Boolean,
        default: false
    },
    showUpdate: {
        type: Boolean,
        default: false

    },
    showDelete: {
        type: Boolean,
        default: false
    },
    tableDate: [],
    tableColProps: {},
    operateWidth: {
        default: 100
    },
    formProps: {},
    formRules: undefined,
    showOperate: {
        default: true
    }
})

console.log("props", props.tableColProps);

const formVisiable = ref(false)

const title = ref("Add")

const formRef = ref(null)

const formData = ref({
})

const disabledDate = (time) => {
    // time > now + 1h, use moment.js
    return time.getTime() < moment().add(1, 'hours').valueOf();
}
const openAddForm = () => {
    if (title.value != "添加") {
        let tmp = {}
        for (const item of props.formProps) {
            if (item.type == "img" || item.type == 'file') {
                tmp[item.prop] = []
                continue
            }
            if (item.default != undefined) {
                tmp[item.prop] = item.default
                continue
            }
            tmp[item.prop] = ""
        }
        formData.value = tmp
    }
    title.value = "添加"
    formVisiable.value = true

}

const openUpdateForm = (data) => {
    title.value = "编辑"
    formVisiable.value = true
    let tmp = Object.assign({}, data)
    for (const item of props.formProps) {
        if (item.type == "img" || item.type == 'file') {
            tmp[item.prop] = data[item.prop]
            continue
        }
        tmp[item.prop] = data[item.prop]
    }
    console.log(tmp);
    formData.value = tmp
}

const handleDelete = (data) => {
    emit('delete', data)
}


const reset = () => {
    emit('reset')
    emit('find')
}

const find = () => {
    emit('find')
}

const emitSubmit = () => {
    console.log(formData);

    proxy.$refs.formRef.validate(async (valid) => {
        if (valid) {
            
            formVisiable.value = false
            if (title.value == "添加") {
                for (const item of props.formProps) {
                    if(item.type == "tag"){
                        formData.value = {
                            ...formData.value,
                            [item.prop]: dynamicTags.value
                        }
                    }
                }
                emit('add', formData.value)
                let tmp = {}
                for (const item of props.formProps) {
                    if (item.type == "img" || item.type == 'file') {
                        tmp[item.prop] = []
                        continue
                    }

                    tmp[item.prop] = ""
                }
                dynamicTags.value = []
                formData.value = tmp
            } else {
                emit('update', formData.value)
            }
        } else {

            return false;
        }
    });

}


const inputValue = ref('')
const dynamicTags = ref([])
const inputVisible = ref(false)
const InputRef = ref()

const handleClose = (tag) => {
    dynamicTags.value.splice(dynamicTags.value.indexOf(tag), 1)
}

const showInput = () => {
    inputVisible.value = true
    nextTick(() => {
        InputRef.value.input.focus()
    })
}

const handleInputConfirm = () => {
    if (inputValue.value) {
        dynamicTags.value.push(inputValue.value.trim())
    }
    inputVisible.value = false
    inputValue.value = ''
}
</script>
<style scoped>
.content-query {
    margin-bottom: 20px;
    display: flex;
}

.flex-end {
    display: flex;
    justify-content: flex-end;
}

.content {
    margin-top: 20px;
    padding: 10;
}
</style>